<!--
order: 1
-->

# 概览

## 需求

Go SDK 需要 `1.13.5` 或以上的 [Golang](https://golang.org/doc/install) 版本。

## 导入 Go SDK

```go
import (
    sdk "github.com/bianjieai/bsnhub-sdk-go"
    "github.com/bianjieai/bsnhub-sdk-go/types"
)
```

## 初始化 SDK

在调用 SDK 接口之前需初始化 SDK，即构建 IRITA 客户端。

```go
options := []types.Option{
  types.KeyDAOOption(store.NewMemory(nil)),
  types.TimeoutOption(10),
}

cfg, err := types.NewClientConfig(nodeURI, chainID, options...)
if err != nil {
  panic(err)
}

client = sdk.NewIRITAClient(cfg)
```

`ClientConfig` 结构代表客户端的配置，其属性和意义如下表所示：

| 属性名     | 类型     | 描述                                                         |
| ---------- | -------- | ------------------------------------------------------------ |
| NodeURI    | string   | 节点的 RPC API 地址, 例如 localhost: 26657                   |
| ChainID    | string   | 链 ID, 例如 `irita`                                           |
| Gas        | uint64   | 交易消耗的 gas 上限, 如 `20000`                              |
| Fee        | DecCoins | 交易费用                                                     |
| KeyDAO     | KeyDAO   | Key 存储接口, 如不提供, 则默认使用 `LevelDB`                 |
| Mode       | enum     | 交易广播模式, 可用值为：`Sync`，`Async`，`Commit`            |
| Algo       | string   | Key 生成算法，支持`sm2`,`secp256k1`, 默认：`sm2`             |
| Timeout    | uint     | 交易广播超时，单位为秒，例如 `5`                             |
| Level      | string   | 日志 level, 如 `info`                                        |
| MaxTxBytes | uint64   | 当前连接的节点允许一笔交易最大的字节数，默认：`1073741824`(5M) |

## KeyDAO

### KeyDAO 接口定义

`KeyDAO` 是 Key 的存储接口，封装了 Key 的存取操作，定义如下：

```go
type KeyDAO interface {
	// Write will use user password to encrypt data and save to file, the file name is user name
	Write(name, password string, store KeyInfo) error

	// Read will read encrypted data from file and decrypt with user password
	Read(name, password string) (KeyInfo, error)

	// Delete will delete user data and use user password to verify permissions
	Delete(name, password string) error

	// Has returns whether the specified user name exists
	Has(name string) bool
}

type Crypto interface {
    Encrypt(data string, password string) (string, error)
    Decrypt(data string, password string) (string, error)
}
```

### Key 存储结构

其中 `KeyInfo` 为 Key 的存储方式

```go
// KeyInfo saves the basic information of the key
type KeyInfo struct {
	Name         string `json:"name"`
	PubKey       []byte `json:"pubkey"`
	PrivKeyArmor string `json:"priv_key_armor"`
	Algo         string `json:"algo"`
}
```

### KeyDAO 实现

开发者可以实现 KeyDAO 接口来满足灵活和复杂的业务需求。

**_提示_：** 如果开发者使用私钥存储方式，且未实现 Crypto 接口，则默认使用 `AES` 来加密存储私钥。

Go SDK 提供了 KeyDAO 的一个基于内存的简单实现，**开发者不应在产品环境中直接使用**，仅用于参考或测试目的。如下所示：

```go
// Use memory as storage, use with caution in build environment
type MemoryDAO struct {
	store map[string]KeyInfo
	Crypto
}

func NewMemory(crypto Crypto) MemoryDAO {
	if crypto == nil {
		crypto = AES{}
	}
	return MemoryDAO{
		store:  make(map[string]KeyInfo),
		Crypto: crypto,
	}
}
func (m MemoryDAO) Write(name, password string, store KeyInfo) error {
	m.store[name] = store
	return nil
}

func (m MemoryDAO) Read(name, password string) (KeyInfo, error) {
	return m.store[name], nil
}

// ReadMetadata read a key information from the local store
func (m MemoryDAO) ReadMetadata(name string) (store KeyInfo, err error) {
	return m.store[name], nil
}

func (m MemoryDAO) Delete(name, password string) error {
	delete(m.store, name)
	return nil
}

func (m MemoryDAO) Has(name string) bool {
	_, ok := m.store[name]
	return ok
}
```

## BaseTx

Go SDK 各接口中均包含 `baseTx` 参数，表示基本交易数据，定义如下：

```go
type BaseTx struct {
  From     string
  Password string
  Gas      uint64
  Fee      DecCoins
  Memo     string
  Mode     BroadcastMode
  Simulate bool
}
```

| 属性名   | 类型     | 描述                                              |
| -------- | -------- | ------------------------------------------------- |
| From     | string   | 交易发起者的 Key 名称或账户地址                   |
| Password | string   | Key 的解密密码                                    |
| Gas      | uint64   | 交易的 gas 上限                                   |
| Fee      | DecCoins | 交易费用                                          |
| Memo     | string   | 交易的备注信息                                    |
| Mode     | Enum     | 交易广播模式，可用值为：`Sync`，`Async`，`Commit` |
| Simulate | bool     | 是否为模拟执行                                    |

## 交易结果

如无特别说明，Go SDK 的交易接口将返回一个交易结果对象。

```go
type ResultTx struct{
    GasWanted int64
    GasUsed   int64
    Events    []Event
    Hash      string
    Height    int64
}
```

| 属性名    | 类型    | 描述                                      |
| --------- | ------- | ----------------------------------------- |
| GasWanted | int64   | 交易的 gas 上限                           |
| GasUsed   | int64   | 交易实际消耗的 gas                        |
| Events    | []Event | 交易触发的事件列表，参考[事件](#交易事件) |
| Hash      | string  | 交易 Hash                                 |
| Height    | int64   | 交易所在的区块高度                        |

## 事件

交易在执行过程中可能触发相关的事件，可通过交易结果的 `Events` 属性获取事件。事件定义如下：

```go
type Event struct {
    Type       string
    Attributes []Attribute
}
```

| 属性名     | 类型        | 描述           |
| ---------- | ----------- | -------------- |
| Type       | string      | 事件的名称     |
| Attributes | []Attribute | 事件的属性列表 |

```go
type Attribute struct {
    Key   string
    Value string
}
```

| 属性名 | 类型   | 描述         |
| ------ | ------ | ------------ |
| Key    | string | 事件属性的键 |
| Value  | string | 事件属性的值 |

## 错误处理

Go SDK 扩展了 `Golang` 标准错误接口 `error`，提供了更丰富的错误处理方法。所有接口均返回一个 `Error` 接口实例，用于客户端错误处理。`Error` 接口定义如下：

```go
type Error interface {
    Error() string
    Code() uint32
    Codespace() string
}
```

- Error 方法：获取错误信息
  - 返回值：错误信息

- Code 方法：获取错误码
  - 返回值：错误码

- Codespace 方法：获取错误码名字空间
  - 返回值：错误码名字空间

错误码和名字空间详见[这里](./errcodes.md)

## 核心模块

Go SDK 实现的核心模块包括：工分、资产数字化建模、存证以及 iService。

- [工分](./API/token.md)
- [资产数字化建模](./API/nft.md)
- [存证](./API/record.md)
- [iService](./API/iservice.md)
