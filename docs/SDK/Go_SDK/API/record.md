<!--
order: 3
-->

# 存证

Go SDK 支持[存证](../../../core_modules/token.md)的创建和查询。

## 导入

导入 Go SDK 存证模块：

```go
import (
  "github.com/bianjieai/irita-sdk-go/modules/record"
)
```

## 接口

### 创建存证

**接口：**

```go
client.Record.CreateRecord(request CreateRecordRequest, baseTx sdk.BaseTx) (string, sdk.Error)
```

**参数：**

- request：创建存证请求对象

  ```go
  type CreateRecordRequest struct {
    Contents []Content
  }

  type Content struct {
    Digest     string
    DigestAlgo string
    URI        string
    Meta       string
  }
  ```

**返回值：**

- 存证 ID

### 查询存证

**接口：**

```go
client.Record.QueryRecord(recordID string) (QueryRecordResponse, sdk.Error)
```

**参数：**

- recordID：存证 ID

**返回值：**

- 查询存证响应

  ```go
  type QueryRecordResponse struct {
    TxHash   string
    Contents []Content
    Creator  string
  }
  ```
