<!--
order: 4
-->

# iService

为简化基于 IRITA [iService](../../../core_modules/iservice.md) 构建分布式服务应用的复杂性，Go SDK 高度封装了 iService 核心接口，为服务消费者和提供者提供了人性化的事件监听机制。

## 导入

导入 Go SDK iService 模块：

```go
import (
  "github.com/bianjieai/irita-sdk-go/modules/service"
)
```

## 接口

### 服务定义

**接口：**

```go
client.Service.DefineService(request DefineServiceRequest, sdk.BaseTx) (sdk.ResultTx, sdk.Error)
```

**参数：**

- request: 服务定义请求对象

  ```go
  type DefineServiceRequest struct {
    ServiceName       string
    Description       string
    Tags              []string
    AuthorDescription string
    Schemas           string
  }
  ```

### 服务绑定

**接口：**

```go
client.Service.BindService(request BindServiceRequest, baseTx sdk.BaseTx) (sdk.ResultTx, sdk.Error)
```

**参数：**

- request: 服务绑定请求对象

  ```go
  type BindServiceRequest struct {
    ServiceName string
    Deposit     sdk.DecCoins
    Pricing     string
    MinRespTime uint64
  }
  ```

### 更新服务绑定

**接口：**

```go
client.Service.UpdateServiceBinding(request UpdateServiceBindingRequest, baseTx sdk.BaseTx) (sdk.ResultTx, sdk.Error)
```

**参数：**

- request：更新服务绑定请求对象

  ```go
  type UpdateServiceBindingRequest struct {
    ServiceName string
    Deposit     sdk.DecCoins
    Pricing     string
    MinRespTime uint64
  }
  ```

### 禁用服务绑定

**接口：**

```go
client.Token.DisableServiceBinding(serviceName string, baseTx sdk.BaseTx) (sdk.ResultTx, sdk.Error)
```

**参数：**

- serviceName：服务名称

### 启用服务绑定

**接口：**

```go
client.Token.EnableServiceBinding(serviceName string, baseTx sdk.BaseTx) (sdk.ResultTx, sdk.Error)
```

**参数：**

- serviceName：服务名称

### 取回服务绑定押金

**接口：**

```go
client.Service.RefundServiceDeposit(serviceName string, baseTx sdk.BaseTx) (sdk.ResultTx, sdk.Error)
```

**参数：**

- serviceName：服务名称

### 发起服务调用

服务消费者发起服务调用。

**接口：**

```go
client.Service.InvokeService(request InvokeServiceRequest, baseTx sdk.BaseTx) (string, sdk.Error)
```

**参数：**

- request：调用服务请求对象

  ```go
  type InvokeServiceRequest struct {
    ServiceName       string
    Providers         []string
    Input             string
    ServiceFeeCap     sdk.DecCoins
    Timeout           int64
    SuperMode         bool
    Repeated          bool
    RepeatedFrequency uint64
    RepeatedTotal     int64
    Callback          InvokeCallback
  }

  // 服务请求的回调接口，服务消费者可提供回调方法，用于响应自动处理。
  type InvokeCallback func(reqCtxID, reqID, responses string)
  ```

**返回值：**

- 请求上下文 ID（关于`请求上下文`请参考 [iService](../../../core_modules/iservice.md#请求上下文)）

### 订阅服务响应

服务消费者在发起服务调用时，如未提供回调方法，在服务调用成功后可以主动订阅服务响应。订阅时提供的回调方法将持续执行，直到请求上下文结束。

**接口：**

```go
client.Service.SubscribeServiceResponse(reqCtxID string, callback InvokeCallback) (subscription sdk.Subscription, err sdk.Error)
```

**参数：**

- reqCtxID：请求上下文 ID

- callback：实现 `InvokeCallback` 的回调方法

**返回值：**

- subscription：订阅对象，可以通过如下方法取消该订阅：

  ```go
  client.Unsubscribe(subscription)
  ```

### 订阅服务请求

服务提供者订阅向自己发起的服务请求。

#### 订阅单个服务的请求

**接口：**

```go
client.Service.SubscribeSingleServiceRequest(serviceName string, callback RespondCallback,
baseTx sdk.BaseTx) (subscription sdk.Subscription, err sdk.Error)
```

**参数：**

- serviceName：服务名称
  
- callback：实现 `RespondCallback` 的回调方法

  ```go
  type RespondCallback func(reqCtxID, reqID, input string) (output string, result string)
  ```

  该回调方法用于服务请求的自动处理。

**返回值：**

- subscription：订阅对象，可以进行取消操作。

#### 订阅多个服务的请求

**接口：**

```go
client.Service.SubscribeServiceRequest(serviceRegistry Registry, baseTx sdk.BaseTx) (subscription sdk.Subscription, err sdk.Error)
```

**参数：**

- serviceRegistry：服务映射表，用于服务名到回调方法的映射
  
  ```go
  type Registry map[string]RespondCallback
  ```

**返回值：**

- subscription：订阅对象，可以进行取消操作。

### 更新请求上下文

服务消费者可以更新创建的请求上下文。

**接口：**

```go
client.Service.UpdateRequestContext(request UpdateRequestContextRequest, baseTx sdk.BaseTx) (sdk.ResultTx, sdk.Error)
```

**参数：**

- request：更新请求上下文请求对象

  ```go
  type UpdateRequestContextRequest struct {
    RequestContextID  string
    Providers         []string
    ServiceFeeCap     sdk.DecCoins
    Timeout           int64
    RepeatedFrequency uint64
    RepeatedTotal     int64
  }
  ```

### 暂停请求上下文

暂停正在运行的请求上下文。

**接口：**

```go
client.Service.PauseRequestContext(requestContextID string, baseTx sdk.BaseTx) (sdk.ResultTx, sdk.Error)
```

**参数：**

- requestContextID：请求上下文 ID

### 启动请求上下文

启动暂停的请求上下文。

**接口：**

```go
client.Service.StartRequestContext(requestContextID string, baseTx sdk.BaseTx) (sdk.ResultTx, sdk.Error)
```

**参数：**

- requestContextID：请求上下文 ID

### 终止请求上下文

永久终止请求上下文。

**接口：**

```go
client.Service.KillRequestContext(requestContextID string, baseTx sdk.BaseTx) (sdk.ResultTx, sdk.Error)
```

**参数：**

- requestContextID：请求上下文 ID

### 设置提取地址

服务提供者的`所有者`设置服务费提取地址。

**接口：**

```go
client.Service.SetWithdrawAddress(withdrawAddress string, baseTx sdk.BaseTx) (sdk.ResultTx, sdk.Error)
```

**参数：**

- withdrawAddress：提取地址

### 提取服务费

所有者提取赚取的服务费；如未指定服务提供者，则提取该账户绑定的所有服务提供者的服务费。

**接口：**

```go
client.Service.WithdrawEarnedFees(baseTx sdk.BaseTx) (sdk.ResultTx, sdk.Error)
```

### 查询服务定义

**接口：**

```go
client.Service.QueryServiceDefinition(serviceName string) (QueryServiceDefinitionResponse, sdk.Error)
```

**参数：**

- serviceName：服务名称

**返回值：**

- 查询服务定义响应

  ```go
  type QueryServiceDefinitionResponse struct {
    Name              string
    Description       string
    Tags              []string
    Author            sdk.AccAddress
    AuthorDescription string
    Schemas           string
  }
  ```

### 查询指定的服务绑定

**接口：**

```go
client.Service.QueryServiceBinding(serviceName string, provider sdk.AccAddress) (QueryServiceBindingResponse, sdk.Error)
```

**参数：**

- serviceName：服务名称

- provider：服务提供者地址

**返回值：**

- 查询服务绑定响应

  ```go
  type QueryServiceBindingResponse struct {
    ServiceName     string
    Provider        sdk.AccAddress
    Deposit         sdk.Coins
    Pricing         string
    WithdrawAddress sdk.AccAddress
    Available       bool
    DisabledTime    time.Time
  }
  ```

### 查询指定服务的所有绑定

**接口：**

```go
client.Service.QueryServiceBindings(serviceName string) ([]QueryServiceBindingResponse, sdk.Error)
```

**参数：**

- serviceName：服务名称

**返回值：**

- 服务绑定查询响应数组

### 查询服务请求

根据请求 ID 查询服务请求。

**接口：**

```go
client.Service.QueryServiceRequest(requestID string) (QueryServiceRequestResponse, sdk.Error)
```

**参数：**

- requestID：请求 ID

**返回值：**

- 服务请求查询响应

  ```go
  type QueryServiceRequestResponse struct {
    ID                         string
    ServiceName                string
    Provider                   sdk.AccAddress
    Consumer                   sdk.AccAddress
    Input                      string
    ServiceFee                 sdk.Coins
    SuperMode                  bool
    RequestHeight              int64
    ExpirationHeight           int64
    RequestContextID           string
    RequestContextBatchCounter uint64
  }
  ```

### 查询服务绑定的请求列表

查询服务绑定的当前活跃请求。

**接口：**

```go
client.Service.QueryServiceRequests(serviceName string, provider sdk.AccAddress) ([]QueryServiceRequestResponse, sdk.Error)
```

**参数：**

- serviceName：服务名称

- provider：服务提供者地址

**返回值：**

- 服务请求查询响应数组

### 查询服务响应

根据请求 ID 查询服务响应。

**接口：**

```go
client.Service.QueryServiceResponse(requestID string) (QueryServiceResponseResponse, sdk.Error)
```

**参数：**

- requestID：请求 ID

**返回值：**

- 服务响应查询结果

  ```go
  type QueryServiceResponseResponse struct {
    Provider                   sdk.AccAddress
    Consumer                   sdk.AccAddress
    Output                     string
    Result                     string
    RequestContextID           string
    RequestContextBatchCounter uint64
  }
  ```

### 查询请求上下文

**接口：**

```go
client.Service.QueryRequestContext(reqCtxID string) (QueryRequestContextResponse, sdk.Error)
```

**参数：**

- reqCtxID：请求上下文 ID

**返回值：**

- 查询请求上下文响应

  ```go
  type QueryRequestContextResponse struct {
    ServiceName        string
    Providers          []sdk.AccAddress
    Consumer           sdk.AccAddress
    Input              string
    ServiceFeeCap      sdk.Coins
    Timeout            int64
    SuperMode          bool
    Repeated           bool
    RepeatedFrequency  uint64
    RepeatedTotal      int64
    BatchCounter       uint64
    BatchRequestCount  uint16
    BatchResponseCount uint16
    BatchState         string
    State              string
    ResponseThreshold  uint16
    ModuleName         string
  }
  ```

### 查询请求上下文的请求列表

查询请求上下文指定批次的活跃请求。

**接口：**

```go
client.Service.QueryRequestsByReqCtx(reqCtxID string, batchCounter uint64) ([]QueryServiceRequestResponse, sdk.Error)
```

**参数：**

- reqCtxID：请求上下文 ID

- batchCounter：批次计数器

**返回值：**

- 服务请求查询响应数组

### 查询请求上下文的响应列表

查询请求上下文指定批次的活跃响应。

**接口：**

```go
client.Service.QueryServiceResponses(reqCtxID string, batchCounter uint64) ([]QueryServiceResponseResponse, sdk.Error)
```

**参数：**

- reqCtxID：请求上下文 ID

- batchCounter：批次计数器

**返回值：**

- 服务响应查询结果数组

### 查询提取地址

查询所有者的服务费提取地址。

**接口：**

```go
client.Service.QueryWithdrawAddress(owner string) (string， sdk.Error)
```

**参数：**

- owner：所有者的账户地址

**返回值：**

- 提取地址

### 查询服务费

查询指定服务提供者赚取的服务费。

**接口：**

```go
client.Service.QueryFees(provider string) (sdk.DecCoins, sdk.Error)
```

**参数：**

- provider：服务提供者地址

**返回值：**

- 服务费
