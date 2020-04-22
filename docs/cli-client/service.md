# iritacli service

Service模块允许在IRIS Hub中定义、绑定、调用服务。[了解更多iService内容](../features/service.md)。

## 可用命令

| 名称                                                    | 描述                                         |
| ------------------------------------------------------- | -------------------------------------------- |
| [define](#iritacli-service-define)                       | 定义一个新的服务                             |
| [definition](#iritacli-service-definition)               | 查询服务定义                                 |
| [bind](#iritacli-service-bind)                           | 绑定一个服务                                 |
| [binding](#iritacli-service-binding)                     | 查询服务绑定                                 |
| [bindings](#iritacli-service-bindings)                   | 查询服务绑定列表                             |
| [set-withdraw-addr](#iritacli-service-set-withdraw-addr) | 设置服务提供者的提取地址                     |
| [withdraw-addr](#iritacli-service-withdraw-addr)         | 查询服务提供者的提取地址                     |
| [update-binding](#iritacli-service-update-binding)       | 更新一个存在的服务绑定                       |
| [disable](#iritacli-service-disable)                     | 禁用一个可用的服务绑定                       |
| [enable](#iritacli-service-enable)                       | 启用一个不可用的服务绑定                     |
| [refund-deposit](#iritacli-service-refund-deposit)       | 退还一个服务绑定的所有押金                   |
| [call](#iritacli-service-call)                           | 发起服务调用                               |
| [request](#iritacli-service-request)                     | 通过请求ID查询服务请求                       |
| [requests](#iritacli-service-requests)                   | 通过服务绑定或请求上下文查询服务请求列表     |
| [respond](#iritacli-service-respond)                     | 响应服务请求                                 |
| [response](#iritacli-service-response)                   | 通过请求ID查询服务响应                       |
| [responses](#iritacli-service-responses)                 | 通过请求上下文ID和批次计数器查询服务响应列表 |
| [request-context](#iritacli-service-request-context)     | 查询请求上下文                               |
| [update](#iritacli-service-update)                       | 更新请求上下文                               |
| [pause](#iritacli-service-pause)                         | 暂停一个正在进行的请求上下文                 |
| [start](#iritacli-service-start)                         | 启动一个暂停的请求上下文                     |
| [kill](#iritacli-service-kill)                           | 终止请求上下文                               |
| [fees](#iritacli-service-fees)                           | 查询服务提供者的收益                         |
| [withdraw-fees](#iritacli-service-withdraw-fees)         | 提取服务提供者的收益                         |

## iritacli service define

定义一个新的服务。

```bash
iritacli tx service define [flags]
```

**标志：**

| 名称，速记           | 默认 | 描述                        | 必须 |
| -------------------- | ---- | --------------------------- | ---- |
| --name               |      | 服务名称                    | 是   |
| --description        |      | 服务的描述                  |      |
| --author-description |      | 服务创建者的描述            |      |
| --tags               |      | 服务的标签列表              |      |
| --schemas            |      | 服务接口schemas的内容或文件路径 | 是   |

### 定义一个新的服务

```bash
iritacli tx service define --chain-id=irishub --from=<key-name> --fees=0.3iris 
--name=<service name> --description=<service description> --author-description=<author description>
--tags=tag1,tag2 --schemas=<schemas content or path/to/schemas.json>
```

### Schemas内容示例

```json
{"input":{"$schema":"http://json-schema.org/draft-04/schema#","title":"BioIdentify service input","description":"BioIdentify service input specification","type":"object","properties":{"id":{"description":"id","type":"string"},"name":{"description":"name","type":"string"},"data":{"description":"data","type":"string"}},"required":["id","data"]},"output":{"$schema":"http://json-schema.org/draft-04/schema#","title":"BioIdentify service output","description":"BioIdentify service output specification","type":"object","properties":{"data":{"description":"result data","type":"string"}},"required":["data"]}}
```

## iritacli service definition

查询服务定义。

```bash
iritacli tx service definition [service-name] [flags]
```

**标志：**

| 名称，速记     | 默认 | 描述     | 必须 |
| -------------- | ---- | -------- | ---- |
| --service-name |      | 服务名称 | 是   |

### 查询一个服务定义

查询指定服务定义的详细信息。

```bash
iritacli q service definition <service name>
```

## iritacli service bind

绑定一个服务。

```bash
iritacli tx service bind [flags]
```

**标志：**

| 名称，速记     | 默认 | 描述                                                  | 必须 |
| -------------- | ---- | ----------------------------------------------------- | ---- |
| --service-name |      | 服务名称                                              | 是   |
| --deposit      |      | 服务绑定的押金                                        | 是   |
| --pricing      |      | 服务定价内容或文件路径，是一个[Irishub Service Pricing JSON Schema](../features/service-pricing.md)实例 |   是   |
| --min-resp-time |     | 最小响应时间 | 是 |

### 绑定一个存在的服务定义

抵押`deposit`应该满足最小抵押数量需求，最小抵押数量为`price` * `MinDepositMultiple` 和 `MinDeposit`中的最大值（`MinDepositMultiple`以及`MinDeposit`是可治理参数）。

```bash
iritacli tx service bind --chain-id=irishub --from=<key-name> --fees=0.3iris
--service-name=<service name> --deposit=10000iris --pricing=<pricing content or path/to/pricing.json> --min-resp-time=50
```

### Pricing内容示例

```json
{
    "price": "1iris"
}
```

## iritacli service binding

查询服务绑定。

```bash
iritacli q service binding [service-name] [provider] [flags]
```

### 查询一个服务绑定

```bash
iritacli q service binding <service name> <provider>
```

## iritacli service bindings

查询服务绑定列表。

```bash
iritacli q service bindings [service-name] [flags]
```

### 查询服务绑定列表

```bash
iritacli q service bindings <service name>
```

## iritacli service update-binding

更新服务绑定。

```bash
iritacli tx service update-binding [service-name] [flags]
```

**标志：**
| 名称，速记     | 默认 | 描述                                                  | 必须 |
| -------------- | ---- | ----------------------------------------------------- | ---- |
| --deposit      |      | 增加的绑定押金，为空则不更新                                      |      |
| --pricing      |      | 服务定价内容或文件路径，是一个[Irishub Service Pricing JSON Schema](../features/service-pricing.md)实例，为空则不更新 |      |
| --min-resp-time |     | 最小响应时间，为0则不更新 | |

### 更新一个存在的服务绑定

更新服务绑定，追加 10 IRIS 的抵押。

```bash
iritacli tx service update-binding <service-name> --chain-id=irishub --from=<key-name> --fees=0.3iris --deposit=10iris
```

## iritacli service set-withdraw-addr

设置服务提供者的提取地址。

```bash
iritacli tx service set-withdraw-addr [withdrawal-address] [flags]
```

### 设置一个提取地址

```bash
iritacli tx service set-withdraw-addr <withdrawal address> --chain-id=irishub --from=<key-name> --fees=0.3iris
```

## iritacli service withdraw-addr

查询服务提供者的提取地址。

```bash
iritacli q service withdraw-addr [provider] [flags]
```

### 查询一个服务提供者的提取地址

```bash
iritacli q service withdraw-addr <provider>
```

## iritacli service disable

禁用一个可用的服务绑定。

```bash
iritacli tx service disable [service-name] [flags]
```

### 禁用一个可用的服务绑定

```bash
iritacli tx service disable <service name> --chain-id=irishub --from=<key-name> --fees=0.3iris
```

## iritacli service enable

启用一个不可用的服务绑定。

```bash
iritacli tx service enable [service-name] [flags]
```

**标志：**

| 名称，速记 | 默认 | 描述               | 必须 |
| ---------- | ---- | ------------------ | ---- |
| --deposit  |      | 启用绑定增加的押金 |      |

### 启用一个不可用的服务绑定

启用一个不可用的服务绑定，追加 10 IRIS 的抵押。

```bash
iritacli tx service enable <service name> --chain-id=irishub --from=<key-name> --fees=0.3iris --deposit=10iris
```

## iritacli service refund-deposit

从一个服务绑定中退还所有的押金。

```bash
iritacli tx service refund-deposit [service-name] [flags]
```

### 取回所有押金

取回抵押之前，必须先[禁用](#iritacli-service-disable)服务绑定。

```bash
iritacli tx service refund-deposit <service name> --chain-id=irishub --from=<key-name> --fees=0.3iris
```

## iritacli service call

发起服务调用。

```bash
iritacli tx service call [flags]
```

**标志：**

| 名称，速记        | 默认  | 描述                                    | 必须 |
| ----------------- | ----- | --------------------------------------- | ---- |
| --name            |       | 服务名称                                | 是   |
| --providers       |       | 服务提供者列表                          | 是   |
| --service-fee-cap |       | 愿意为单个请求支付的最大服务费用        | 是   |
| --data            |       | 请求输入的内容或文件路径，是一个Input JSON Schema实例 | 是   |
| --timeout         |       | 请求超时                                |   是   |
| --super-mode      | false | 签名者是否为超级用户                    |
| --repeated        | false | 请求是否为重复性的                      |      |
| --frequency       |       | 重复性请求的请求频率；默认为`timeout`值 |      |
| --total           |       | 重复性请求的请求总数，-1表示无限制      |      |

### 发起一个服务调用请求

```bash
iritacli tx service call --chain-id=irishub --from=<key name> --fees=0.3iris --service-name=<service name>
--providers=<provider list> --service-fee-cap=1iris --data=<request input or path/to/input.json> --timeout=100 --repeated --frequency=150 --total=100
```

### 请求输入示例

```json
{
    "id": "1",
    "name": "irisnet",
    "data": "facedata"
}
```

## iritacli service request

通过请求ID查询服务请求。

```bash
iritacli q service request [request-id] [flags]
```

### 查询一个服务请求

```bash
iritacli q service request <request-id>
```

:::tip
你可以从[按高度获取区块信息](./tendermint.md#iritacli-tendermint-block)的结果中获取`request-id`。
:::

## iritacli service requests

通过服务绑定或请请求上下文ID查询服务请求列表。

```bash
iritacli q service requests [service-name] [provider] | [request-context-id] [batch-counter] [flags]
```

### 查询服务绑定的活跃请求

```bash
iritacli q service requests <service name> <provider>
```

### 通过请求上下文ID和批次计数器查询服务请求列表

```bash
iritacli q service requests <request-context-id> <batch-counter>
```

## iritacli service respond

响应服务请求。

```bash
iritacli tx service respond [flags]
```

**标志：**

| 名称，速记   | 默认 | 描述                                         | 必须 |
| ------------ | ---- | -------------------------------------------- | ---- |
| --request-id |      | 欲响应请求的ID                               | 是   |
| --result     |      | 响应结果的内容或文件路径, 是一个[Irishub Service Result JSON Schema](../features/service-result.md)实例 | 是   |
| --data       |      | 响应输出的内容或文件路径, 是一个Output JSON Schema实例 |      |

### 响应一个服务请求

```bash
iritacli tx service respond --chain-id=irishub --from=<key-name> --fees=0.3iris
--request-id=<request-id> --result=<response result or path/to/result.json> --data=<response output or path/to/output.json>
```

:::tip
你可以从[按高度获取区块信息](./tendermint.md#iritacli-tendermint-block)的结果中获取`request-id`。
:::

### 响应结果示例

```json
{
    "code": 200,
    "message": ""
}
```

### 响应输出示例

```json
{
    "data": "userdata"
}
```

## iritacli service response

通过请求ID查询服务响应。

```bash
iritacli q service response [request-id] [flags]
```

### 查询一个服务响应

```bash
iritacli q service response <request-id>
```

:::tip
你可以从[按高度获取区块信息](./tendermint.md#iritacli-tendermint-block)的结果中获取`request-id`。
:::

## iritacli service responses

通过请求上下文ID以及批次计数器查询服务响应列表。

```bash
iritacli q service responses [request-context-id] [batch-counter] [flags]
```

### 根据指定的请求上下文ID以及批次计数器查询服务响应

```bash
iritacli q service responses <request-context-id> <batch-counter>
```

## iritacli service request-context

查询请求上下文。

```bash
iritacli q service request-context [request-context-id] [flags]
```

### 查询一个请求上下文

```bash
iritacli q service request-context <request-context-id>
```

:::tip
你可以从[调用服务](#iritacli-service-call)的结果中获取`request-context-id`
:::

## iritacli service update

更新请求上下文。

```bash
iritacli tx service update [request-context-id] [flags]
```

**标志：**

| 名称，速记        | 默认 | 描述                                           | 必须 |
| ----------------- | ---- | ---------------------------------------------- | ---- |
| --providers       |      | 服务提供者列表，为空则不更新                   |      |
| --service-fee-cap |      | 愿意为单个请求支付的最大服务费用，为空则不更新 |      |
| --timeout         |      | 请求超时，为0则不更新                          |      |
| --frequency       |      | 请求频率，为0则不更新                          |      |
| --total           |      | 请求总数，为0则不更新                          |      |

### 更新一个请求上下文

```bash
iritacli tx service update <request-context-id> --chain-id=irishub --from=<key name> --fees=0.3iris
--providers=<provider list> --service-fee-cap=1iris --timeout=0 --frequency=150 --total=100
```

## iritacli service pause

暂停一个正在进行的请求上下文。

```bash
iritacli tx service pause [request-context-id] [flags]
```

### 暂停一个正在进行的请求上下文

```bash
iritacli tx service pause <request-context-id>
```

## iritacli service start

启动一个暂停的请求上下文。

```bash
iritacli tx service start [request-context-id] [flags]
```

### 启动一个暂停的请求上下文

```bash
iritacli tx service start <request-context-id>
```

## iritacli service kill

终止请求上下文。

```bash
iritacli tx service kill [request-context-id] [flags]
```

### 终止一个请求上下文

```bash
iritacli tx service kill <request-context-id>
```

## iritacli service fees

查询服务提供者的收益。

```bash
iritacli q service fees [provider] [flags]
```

### 查询服务提供者的收益

```bash
iritacli q service fees <provider>
```

## iritacli service withdraw-fees

提取服务提供者的收益。

```bash
iritacli q service withdraw-fees [flags]
```

### 提取服务提供者的收益

```bash
iritacli q service withdraw-fees --chain-id=irishub --from=<key-name> --fee=0.3iris
```
