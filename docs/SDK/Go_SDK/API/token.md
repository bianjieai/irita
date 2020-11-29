<!--
order: 1
-->

# 积分

Go SDK 实现了 IRITA [积分模块](../../../core_modules/token.md) 的主要操作。

## 导入

导入 Go SDK 积分模块：

```go
import (
  "github.com/bianjieai/irita-sdk-go/modules/token"
)
```

## 接口

### 发行积分

**接口：**

```go
client.Token.IssueToken(request IssueTokenRequest, sdk.BaseTx) (sdk.ResultTx, sdk.Error)
```

**参数：**

- request: 发行积分请求对象

  ```go
  type IssueTokenRequest struct {
    Symbol        string
    Name          string
    Scale         uint8
    MinUnit       string
    InitialSupply uint64
    MaxSupply     uint64
    Mintable      bool
  }
  ```

### 编辑积分

**接口：**

```go
client.Token.EditToken(request EditTokenRequest, baseTx sdk.BaseTx) (sdk.ResultTx, sdk.Error)
```

**参数：**

- request: 编辑积分请求对象

  ```go
  type EditTokenRequest struct {
    Symbol    string
    Name      uint64
    MaxSupply string
    Mintable  bool
  }
  ```

### 增发积分

**接口：**

```go
client.Token.MintToken(request MintTokenRequest, baseTx sdk.BaseTx) (sdk.ResultTx, sdk.Error)
```

**参数：**

- request：增发积分请求对象

  ```go
  type MintTokenRequest struct {
    Symbol    string
    Amount    uint64
    Recipient string
  }
  ```

### 转让积分

**接口：**

```go
client.Token.TransferToken(request TransferTokenRequest, baseTx sdk.BaseTx) (sdk.ResultTx, sdk.Error)
```

**参数：**

- request：转让积分请求对象

  ```go
  type TransferTokenRequest struct {
    Symbol    string
    Recipient string
  }
  ```

### 查询指定积分

**接口：**

```go
client.Token.QueryToken(denom string, baseTx sdk.BaseTx) (sdk.Token, sdk.Error)
```

**参数：**

- denom: 积分的唯一标识符

**返回值：**

- 积分对象

  ```go
  type Token struct {
    Symbol        string
    Name          string
    Scale         uint8
    MinUnit       string
    InitialSupply uint64
    MaxSupply     uint64
    Mintable      bool
    Owner         string
  }
  ```

### 查询账户积分

**接口：**

```go
client.Token.QueryTokens(owner string) (sdk.Tokens, sdk.Error)
```

**参数：**

- owner: 所有者账户地址

**返回值：**

- 积分对象数组
  
  ```go
  type Tokens []Token
  ```
