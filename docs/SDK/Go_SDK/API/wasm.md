<!--
order: 6
-->

# 智能合约

Go SDK [智能合约](../../../core_modules/wams.md)模块实现了合约的上传、初始化、执行、升级和查询功能。

## 接口

### 上传合约

**接口：**

```go
Store(request StoreRequest, config sdk.BaseTx) (string, error)
```

- request：创建上传合约对象。

    ```go
    type StoreRequest struct {
        // WASMByteCode can be raw or gzip compressed
        WASMByteCode []byte
        // WASMFile can be raw or gzip file
        WASMFile string
        // Source is a valid absolute HTTPS URI to the contract's source code, optional
        Source string
        // Builder is a valid docker image name with tag, optional
        Builder string
        // InstantiatePermission access control to apply on contract creation, optional
        permission *AccessConfig
    }
    ```

### 初始化合约

**接口：**

```go
Instantiate(request InstantiateRequest, config sdk.BaseTx) (string, error)
```

- request：创建上传合约对象。

    ```go
    type InstantiateRequest struct {
        // Admin is an optional address that can execute migrations
        Admin string
        // CodeID is the reference to the stored WASM code
        CodeID string
        // Label is optional metadata to be stored with a contract instance.
        Label string
        // InitMsg json encoded message to be passed to the contract on instantiation
        InitMsg Args
        // InitFunds coins that are transferred to the contract on instantiation
        InitFunds sdk.Coins
    }
    ```

### 执行合约

**接口：**

```go
Execute(contractAddress string,
    abi *ContractABI,
    sentFunds sdk.Coins,
    config sdk.BaseTx) (sdk.ResultTx, error)
```

- abi：创建合约的方法签名，主要包括方法名和参数列表。

    ```go
    type Args map[string]interface{}
    type ContractABI struct {
        Method string
        Args   Args
    }
    ```

- sentFunds：转账给合约的代币数量

### 升级智能合约

**接口：**

```go
Migrate(contractAddress string,
    newCodeID string,
    msgByte []byte,
    config sdk.BaseTx) (sdk.ResultTx, error)
```

- contractAddress：智能合约地址
- newCodeID：新的智能合约`code_id`(执行`Store`方法时的返回值)
- msgByte：执行合约升级的参数信息。

### 根据合约地址查询智能合约信息

**接口：**

```go
QueryContractInfo(contractAddress string) (*ContractInfo, error)
```

- contractAddress：智能合约地址

### 执行智能合约的查询方法

**接口：**

```go
QueryContract(contractAddress string, abi *ContractABI) ([]byte, error)
```

- contractAddress：智能合约地址
- abi：创建合约的方法签名，主要包括方法名和参数列表。

### 导出合约中的所有状态信息

**接口：**

```go
ExportAllContractState(contractAddress string) (map[string][]byte, error)
```

- contractAddress：智能合约地址。
