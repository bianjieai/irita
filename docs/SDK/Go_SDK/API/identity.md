<!--
order: 5
-->

# 身份

Go SDK [身份] 模块(../../../core_modules/identity.md)实现了身份的创建、更新和查询。

## 导入

导入 Go SDK 身份模块：

```go
import (
  "github.com/bianjieai/irita-sdk-go/modules/identity"
)
```

## 接口

### 创建身份

**接口：**

```go
client.Identity.CreateIdentity(request CreateIdentityRequest, baseTx sdk.BaseTx) (sdk.ResultTx, sdk.Error)
```

**参数：**

- request：创建身份请求对象

  ```go
  type CreateIdentityRequest struct {
    ID           string
    PubkeyInfo   *PubkeyInfo
    Certificate  string
    Credentials  string
  }

  type PubkeyInfo struct {
    PubKey     string
    PubKeyAlgo PubKeyAlgorithm
  }

  type PubKeyAlgorithm int32
  ```

### 更新身份

**接口：**

```go
client.Identity.UpdateIdentity(request UpdateIdentityRequest, baseTx sdk.BaseTx) (sdk.ResultTx, sdk.Error)
```

**参数：**

- request：更新身份请求对象

  ```go
  type UpdateIdentityRequest struct {
    ID           string
    PubkeyInfo   *PubkeyInfo
    Certificate  string
    Credentials  *string
  }
  ```

### 查询身份

**接口：**

```go
client.Identity.QueryIdentity(id string) (QueryIdentityResponse, sdk.Error)
```

**参数：**

- id：身份 ID

**返回值：**

- 查询身份响应

  ```go
  type QueryIdentityResponse struct {
    ID           string
    PubkeyInfos  []PubkeyInfo
    Certificates []string
    Credentials  string
    Owner        string
  }
  ```
