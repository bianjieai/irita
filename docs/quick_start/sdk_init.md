<!--
order: 2
-->

# SDK 初始化

## 客户端配置

通过实例化 `ClientConfig` 来完成 `SDK` 客户端的配置。各参数可按实际开发需求设置。本文档教程所用 Key 通过 `KeyStore` 导入。

```go
// 定义客户端配置变量
nodeURI := "tcp://localhost:26657"
network := types.Mainnet
chainID := "irita-test"
gas := 200000
fee := types.NewDecCoins(types.NewDecCoin("point", types.NewInt(5)))
mode := types.Commit
algo := "sm2"
timeout := 10
level := "info"
dbPath := os.Getwd()

// 生成客户端配置对象
clientConfig := types.ClientConfig{
    NodeURI:   nodeURI,
    Network:   network,
    ChainID:   chainID,
    Gas:       uint64(gas),
    Fee:       fee,
    Mode:      mode,
    Algo:      algo,
    KeyDAO:    store.NewFileDAO(dbPath),
    Timeout:   uint(timeout),
    Level:     level,
}
```

## 构建客户端

用生成的 `clientConfig` 构建客户端。

```go
client := sdk.NewIRITAClient(clientConfig)
```

## Key 设置

教程使用 `KeyStore` 导入签名所需 Key。开发者可以通过实现 `KeyDAO` 接口满足更复杂的业务需求。详见 [SDK](../SDK/Go_SDK/overview.md)。

```go
// keystore
keyName := "test"
password := "1234567890"
keyStore := `-----BEGIN TENDERMINT PRIVATE KEY-----
kdf: bcrypt
salt: 0DFC160024DC06F28878EC59F1D86C64
type: secp256k1

WPewPlMKHVXRoLnBZOT1IC9hNh6vjb8RVAbIHD97uYNR2lf+SnDCp4WgD9a4UOi/
qC5uVKGBKf6jNGqAx30vkBfIX2pmwa5gKh7Wqhs=
=C5OM
-----END TENDERMINT PRIVATE KEY-----`

// 导入 keystore
_, err := client.Key.Import(keyName, password, keyStore)
if err != nil {
    panic(err)
}
```
