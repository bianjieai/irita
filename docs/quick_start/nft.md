<!--
order: 4
-->

# 创建数字资产

本教程将开发一个简单的应用来演示资产数字化建模相关的功能，包括：

- 发行资产
- 创建资产
- 转移资产
- 查询资产类别信息
- 查询资产
- 查询账户资产

有关`资产数字化建模`的介绍请参考[这里](../core_modules/nft.md)。

>**_需求：_** 开发前请完成[准备工作](prepare.md)。

## 开发步骤

### 初始化 SDK

参考[初始化 SDK](sdk_init.md)

### 定义资产变量

```go
// 定义资产类别属性
denom := "security"
schema := `{"type":"object","properties":{"name":{"type":"string"}}}`

// 定义资产数据
id := "e269969972be451aa44ca12bccd88a4c"
uri := "https://metadata.io/id"
metadata := `{"name":"security001"}`
```

### 发行资产

调用 `NFT` 模块的 `IssueDenom` 方法发行资产。

```go
// 构造 BaseTx
baseTx := types.BaseTx{
    From:     accountName,
    Gas:      uint64(gas),
    Fee:      fee,
    Memo:     "",
    Mode:     mode,
    Password: password,
}

// 构造发行资产请求
issueDenomReq := nft.IssueDenomRequest(
    Denom: denom,
    Schema: schema,
}

// 发行资产
_, err := client.NFT.IssueNFT(issueDenomReq, baseTx)
```

### 查询资产类别信息

根据 `Denom` 查询资产类别信息。

```go
// 查询资产类别信息
res, err := client.NFT.QueryDenom(denom)
```

### 创建资产

调用 `NFT` 模块的 `MintNFT` 方法创建资产。

```go
// 构造创建资产请求
mintNFTReq := nft.MintNFTRequest{
    Denom: denom,
    TokenID: id,
    TokenURI: uri,
    TokenData: metadata,
}

// 创建资产
_, err := client.NFT.MintNFT(mintNFTReq, baseTx)
```

### 查询资产

根据 `Denom` 和 `ID` 查询具体的资产。

```go
// 通过 Denom 和 ID 查询资产信息
res, err := client.NFT.QueryNFT(denom, id)
```

### 转移资产

调用 `NFT` 模块的 `TransferNFT` 完成资产转移。

```go
// 构造转移资产请求
recipient := "iaa18up8anyjpal8rncm8rd4ukp5f7etga795gp33q"
transferNFTReq := nft.TransferNFTRequest(
    Recipient: recipient,
    Denom: denom,
    TokenID: id,
}

// 转移资产
_, err := client.NFT.TransferNFT(transferNFTReq, baseTx)
```

### 查询账户资产

查询 `recipient` 的全部资产。

```go
res, err := client.NFT.QueryOwner(recipient)
```

## 完整示例代码

此数字资产应用示例完整代码如下：

```go
TODO
```
