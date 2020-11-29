<!--
order: 6
-->

# 分布式服务调用

本教程将通过开发一个简单的 `iService` 应用来演示分布式服务调用，功能包括：

- 服务定义
- 服务绑定
- 服务调用
- 服务响应
- 查询请求
- 查询响应

有关 `iService` 的介绍请参考[这里](../core_modules/iservice.md)。

>**_需求：_** 开发前请完成[准备工作](prepare.md)。

> **_说明：_** 为演示各功能，教程采用 _非事件监听_ 模式，即服务消费者和提供者主动查询请求或响应。实际应用中，开发者可以使用更高效的 _事件监听_ 模式。具体参见 [SDK](../SDK/Go_SDK/API/iService.md)

## 开发步骤

### 初始化 SDK

参考[初始化 SDK](sdk_init.md)

### 声明服务属性

首先声明服务定义的属性，包括服务名称、标签、创建者描述、接口 Schemas 。

```go
// 声明服务属性
serviceName := "helloworld"
serviceDescription := "sample service"
serviceTags := []string{"hello","world"}
authorDescription := "irita admin"
schemas := `{"input":{"type":"object","properties":{"ping":{"type":"string"}},"output":{"type":"object","properties":{"pong":{"type":"string"}}}`
```

### 服务定义

调用 `Service` 模块的 `DefineService` 方法实现服务定义。

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

// 构造服务定义请求
defineServiceReq := service.DefineServiceRequest(
    Name: serviceName,
    Description: serviceDescription,
    Tags: serviceTags,
    AuthorDescription: authorDescription,
    Schemas: schemas,
}

// 服务定义
_, err := client.Service.DefineService(defineServiceReq, baseTx)
```

### 服务绑定

定义绑定相关的属性：押金、服务定价、服务质量。押金必须满足最低押金需求，详见 [iService](../core_modules/iservice.md) 。

```go
// 声明服务绑定属性
deposit := types.NewDecCoins(types.NewDecCoin("point", types.NewInt(10000)))
pricing := `{"price":"1point"}`
qos := uint64(50)
```

调用 `Service` 模块的 `BindService` 方法绑定服务。

```go
// 构造服务绑定请求
bindServiceReq := service.BindServiceRequest{
    ServiceName: serviceName,
    Deposit: deposit,
    Pricing: pricing,
    QoS: qos,
}

// 服务绑定
_, err := client.Service.BindService(bindServiceReq, baseTx)
```

### 服务调用

服务绑定之后，服务消费者就可以发起服务调用。调用所需的参数包括：服务名称、服务提供者列表、服务费用上限、请求输入数据、请求超时以及重复性相关参数（是否可重复、重复频率和重复总数）。

```go
// 定义服务调用参数
providers := []string{"iaa1n2jahnzw2llayw68jnw8uvkvqrx34rxq07ymwz"}
serviceFeeCap := types.NewDecCoins(types.NewDecCoin("point", types.NewInt(10)))
input := `{"ping":"hello"}`
timeout := uint64(100)
repeated := ture
frequency: = uint64(100)
total := int64(1)
```

```go
// 构造服务调用请求
invokeServiceReq := service.InvokeServiceRequest{
    ServiceName: serviceName,
    Providers: providers,
    ServiceFeeCap: serviceFeeCap,
    Input: input,
    Timeout: timeout,
    Repeated: repeated,
    RepeatedFrequency: frequency,
    RepeatedTotal: total,
}

// 服务调用
requestContextID, err := client.Service.InvokeService(invokeServiceReq, baseTx)
```

注：`InvokeService` 方法会返回 `请求上下文 ID`，可以通过此 ID 进行请求上下文的相关操作。具体见 [iService](../core_modules/iservice.md) 。

### 查询请求

在消费者发起服务调用之后，相应的服务提供者就可以监听或者查询到活跃的服务请求。查询方式如下：

```go
// 查询指定服务绑定的活跃请求
provider := "iaa1n2jahnzw2llayw68jnw8uvkvqrx34rxq07ymwz"
res, err := service.QueryServiceRequests(serviceName, provider)
}
```

### 服务响应

当服务提供者查询到活跃请求时，就可以对请求作出响应。响应时需提供响应结果，用于指示请求是否成功处理。如成功，则另需提供响应输出。

```go
// 定义响应参数
requestID := res[0].ID
result := `{"code":200,"message":"ok"}`
output := `{"pong":"world"}`

// 构造服务响应请求
respondServiceReq := service.RespondServiceRequest{
    RequestID: requestID,
    Result: result,
    Output: output,
}

// 服务响应
_, err := client.Service.RespondService(respondServiceReq, baseTx)
```

### 查询响应

服务消费者可以通过 `requestID` 查询对应的响应。

```go
// 查询响应
res, err := service.QueryServiceResponse(requestID)
}
```

## 完整示例代码

iService 应用示例的完整代码如下：

```go
TODO
```
