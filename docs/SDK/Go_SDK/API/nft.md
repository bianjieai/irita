<!--
order: 2
-->

# 资产数字化建模

Go SDK 封装了 IRITA [资产数字化建模](../../../core_modules/nft.md) 的核心功能。

## 导入

导入 Go SDK 资产数字化建模模块：

```go
import (
  "github.com/bianjieai/irita-sdk-go/modules/nft"
)
```

## 接口

### 发行资产

**接口：**

```go
client.NFT.IssueDenom(request IssueDenomRequest, baseTx sdk.BaseTx) (sdk.ResultTx, sdk.Error)
```

**参数：**

- request：发行资产请求对象

  ```go
  type IssueDenomRequest struct {
    Name    string
    Schema  string
  }
  ```

### 创建资产

**接口：**

```go
client.NFT.MintNFT(request MintNFTRequest, baseTx sdk.BaseTx) (sdk.ResultTx, sdk.Error)
```

**参数：**

- request：创建资产请求对象

  ```go
  type MintNFTRequest struct {
    Recipient string
    Denom     string
    TokenID   string
    TokenURI  string
    TokenData string
  }
  ```

### 编辑资产

**接口：**

```go
client.NFT.EditNFT(request EditNFTRequest, baseTx sdk.BaseTx) (sdk.ResultTx, sdk.Error)
```

**参数：**

- request：编辑资产请求对象

  ```go
  type EditNFTRequest struct {
    Denom     string
    TokenID   string
    TokenURI  string
    TokenData string
  }
  ```

### 转移资产

**接口：**

```go
client.NFT.TransferNFT(request TransferNFTRequest, baseTx sdk.BaseTx) (sdk.ResultTx, sdk.Error)
```

**参数：**

- request：转移资产请求对象

  ```go
  type TransferNFTRequest struct {
    Recipient string
    Denom     string
    TokenID   string
  }
  ```

### 销毁资产

**接口：**

```go
client.NFT.BurnNFT(request BurnNFTRequest, baseTx sdk.BaseTx) (sdk.ResultTx, sdk.Error)
```

**参数：**

- request：销毁资产请求对象

  ```go
  type BurnNFTRequest struct {
    Denom   string
    TokenID string
  }
  ```

### 查询资产类别信息

**接口：**

```go
client.NFT.QueryDenom(denom string) (QueryDenomResponse, sdk.Error)
```

**参数：**

- denom：资产类别标识符

**返回值：**

- 查询资产类别响应

  ```go
  type QueryDenomResponse struct {
    Name    string
    Schema  string
    Creator string
  }
  ```

### 查询所有资产类别

**接口：**

```go
client.NFT.QueryDenoms() ([]QueryDenomResponse, sdk.Error)
```

**返回值：**

- 查询资产类别响应数组

### 查询指定资产

**接口：**

```go
client.NFT.QueryNFT(denom string, tokenID string) (QueryNFTResponse, sdk.Error)
```

**参数：**

- denom：资产类别
- tokenID：资产 ID

**返回值：**

- 查询资产响应

  ```go
  type QueryNFTResponse struct {
    TokenID   string
    TokenURI  string
    TokenData string
    Creator   string
  }
  ```

### 查询账户资产列表

**接口：**

```go
client.NFT.QueryOwner(owner string, denom string) (QueryOwnerResponse, sdk.Error)
```

**参数：**

- owner：所有者账户地址
- denom：资产类别

**返回值：**

- 查询账户资产列表响应

  ```go
  type QueryOwnerResponse struct {
    Address string
    IDCs    []IDC
  }
  
  type IDC struct {
    Denom    string
    TokenIDs []string
  }
  ```

### 查询指定类别的所有资产

**接口：**

```go
client.NFT.QueryCollection(denom string) (QueryCollectionResponse, sdk.Error)
```

**参数：**

- denom：资产类别

**返回值：**

- 查询资产集合响应

  ```go
  type QueryCollectionResponse struct {
    Denom QueryDenomResponse
    NFTs  []QueryNFTResponse
  }
  ```

### 查询账户指定资产类别的总量

**接口：**

```go
client.NFT.QuerySupply(denom string, owner string) (uint64, sdk.Error)
```

**参数：**

- denom：资产类别
- owner：所有者账户地址

**返回值：**

- 资产总量
