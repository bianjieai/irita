<!--
order: 4
-->

# iService

## define

创建一个新的服务定义。

```bash
 irita tx service define [flags]
```

**标志：**

| 名称，速记           | 默认 | 描述                            | 必须 |
| -------------------- | ---- | ------------------------------- | ---- |
| --name               |      | 服务名称                        | 是   |
| --description        |      | 服务的描述                      |      |
| --author-description |      | 服务创建者的描述                |      |
| --tags               |      | 服务的标签列表                  |      |
| --schemas            |      | 服务接口schemas的内容或文件路径 | 是   |

### 服务定义示例

```bash
irita tx service define --name=test --description=test --author-description=test --tags=test --schemas='{"input":{"type":"object","properties":{"id":{"type":"string"}}},"output":{"type":"object","properties":{"name":{"type":"string"}}}}' --from=node0 --chain-id=irita-test -b=block  -y --home=testnet/node0/iritacli
```

结果

```json
{"height":"1525","txhash":"C7BE33CB20F05D41B31FE0A072159B95048B8134A4EF43BEFA6EA5AB52788BF4","codespace":"","code":0,"data":"0A100A0E646566696E655F73657276696365","raw_log":"[{\"events\":[{\"type\":\"create_definition\",\"attributes\":[{\"key\":\"service_name\",\"value\":\"test\"},{\"key\":\"author\",\"value\":\"iaa17rc02z9tfec74pq8avjg5uj6kj8d57992q7ys0\"}]},{\"type\":\"message\",\"attributes\":[{\"key\":\"action\",\"value\":\"define_service\"},{\"key\":\"module\",\"value\":\"service\"},{\"key\":\"sender\",\"value\":\"iaa17rc02z9tfec74pq8avjg5uj6kj8d57992q7ys0\"}]}]}]","logs":[{"msg_index":0,"log":"","events":[{"type":"create_definition","attributes":[{"key":"service_name","value":"test"},{"key":"author","value":"iaa17rc02z9tfec74pq8avjg5uj6kj8d57992q7ys0"}]},{"type":"message","attributes":[{"key":"action","value":"define_service"},{"key":"module","value":"service"},{"key":"sender","value":"iaa17rc02z9tfec74pq8avjg5uj6kj8d57992q7ys0"}]}]}],"info":"","gas_wanted":"400000","gas_used":"59930","tx":null,"timestamp":""}
```

## bind

绑定一个服务。

```bash
irita tx service bind [flags]
```

**标志：**

| 名称，速记     | 默认 | 描述                                                         | 必须 |
| -------------- | ---- | ------------------------------------------------------------ | ---- |
| --service-name |      | 服务名称                                                     | 是   |
| --provider     |      | 服务提供者地址，默认为签名账户地址                           | 否   |
| --deposit      |      | 服务绑定的押金                                               | 是   |
| --pricing      |      | 服务定价内容或文件路径；服务定价需符合 [iService Pricing Schema](../../core_modules/schemas/iservice-pricing.md) | 是   |
| --qos          |      | 服务质量，即最小响应时间                                     | 是   |
| --options      |      | 非功能性需求的选择，没有额外选择请设置为空（'{}'）           | 是   |

### 服务绑定示例

```bash
irita tx service bind --service-name=test --deposit=10000upoint --pricing='{"price":"1upoint"}' --qos=50 --from=node0 --chain-id=irita-test -b=block -y --options='{}' --home=testnet/node0/iritacli
```

结果

```json
{"height":"1652","txhash":"4EC54BE07048E711CA72EFC48E5D60F5F016B2698EE3E966A03194FE29B3648B","codespace":"","code":0,"data":"0A0E0A0C62696E645F73657276696365","raw_log":"[{\"events\":[{\"type\":\"create_binding\",\"attributes\":[{\"key\":\"service_name\",\"value\":\"test\"},{\"key\":\"provider\",\"value\":\"iaa17rc02z9tfec74pq8avjg5uj6kj8d57992q7ys0\"},{\"key\":\"owner\",\"value\":\"iaa17rc02z9tfec74pq8avjg5uj6kj8d57992q7ys0\"}]},{\"type\":\"message\",\"attributes\":[{\"key\":\"action\",\"value\":\"bind_service\"},{\"key\":\"sender\",\"value\":\"iaa17rc02z9tfec74pq8avjg5uj6kj8d57992q7ys0\"},{\"key\":\"module\",\"value\":\"service\"},{\"key\":\"sender\",\"value\":\"iaa17rc02z9tfec74pq8avjg5uj6kj8d57992q7ys0\"}]},{\"type\":\"transfer\",\"attributes\":[{\"key\":\"recipient\",\"value\":\"iaa1c8t8npfed4xc29755wwwvw2x834q36828duh55\"},{\"key\":\"sender\",\"value\":\"iaa17rc02z9tfec74pq8avjg5uj6kj8d57992q7ys0\"},{\"key\":\"amount\",\"value\":\"10000upoint\"}]}]}]","logs":[{"msg_index":0,"log":"","events":[{"type":"create_binding","attributes":[{"key":"service_name","value":"test"},{"key":"provider","value":"iaa17rc02z9tfec74pq8avjg5uj6kj8d57992q7ys0"},{"key":"owner","value":"iaa17rc02z9tfec74pq8avjg5uj6kj8d57992q7ys0"}]},{"type":"message","attributes":[{"key":"action","value":"bind_service"},{"key":"sender","value":"iaa17rc02z9tfec74pq8avjg5uj6kj8d57992q7ys0"},{"key":"module","value":"service"},{"key":"sender","value":"iaa17rc02z9tfec74pq8avjg5uj6kj8d57992q7ys0"}]},{"type":"transfer","attributes":[{"key":"recipient","value":"iaa1c8t8npfed4xc29755wwwvw2x834q36828duh55"},{"key":"sender","value":"iaa17rc02z9tfec74pq8avjg5uj6kj8d57992q7ys0"},{"key":"amount","value":"10000upoint"}]}]}],"info":"","gas_wanted":"400000","gas_used":"96551","tx":null,"timestamp":""}
```

## update-binding

更新已存在的服务绑定。

```bash
irita tx service update-binding [service-name] [provider-address] [flags]
```

**参数：**

| 名称             | 类型   | 必须 | 默认 | 描述                               |
| ---------------- | ------ | ---- | ---- | ---------------------------------- |
| service-name     | string | 是   |      | 服务名称                           |
| provider-address | string | 否   |      | 服务提供者地址，默认为签名账户地址 |

**标志：**

| 名称，速记 | 默认 | 描述                                                         | 必须 |
| ---------- | ---- | ------------------------------------------------------------ | ---- |
| --deposit  |      | 增加的绑定押金，为空则不更新                                 |      |
| --pricing  |      | 服务定价内容或文件路径；服务定价需符合 [iService Pricing Schema](../../core_modules/schemas/iservice-pricing.md)；为空则不更新 |      |
| --qos      |      | 服务质量，为0则不更新                                        |      |

### 更新服务绑定示例

```bash
irita tx service update-binding test --deposit=100upoint --qos=30 --from=node0 --chain-id=irita-test -b=block  -y --home=testnet/node0/iritacli
```

结果

```json
{"height":"1740","txhash":"4252EAEA1DF2451E4A4FA52F3FEAE744BEBD9E1A0306C889237C1D8DC9A19CA3","codespace":"","code":0,"data":"0A180A167570646174655F736572766963655F62696E64696E67","raw_log":"[{\"events\":[{\"type\":\"message\",\"attributes\":[{\"key\":\"action\",\"value\":\"update_service_binding\"},{\"key\":\"sender\",\"value\":\"iaa17rc02z9tfec74pq8avjg5uj6kj8d57992q7ys0\"},{\"key\":\"module\",\"value\":\"service\"},{\"key\":\"sender\",\"value\":\"iaa17rc02z9tfec74pq8avjg5uj6kj8d57992q7ys0\"}]},{\"type\":\"transfer\",\"attributes\":[{\"key\":\"recipient\",\"value\":\"iaa1c8t8npfed4xc29755wwwvw2x834q36828duh55\"},{\"key\":\"sender\",\"value\":\"iaa17rc02z9tfec74pq8avjg5uj6kj8d57992q7ys0\"},{\"key\":\"amount\",\"value\":\"100upoint\"}]},{\"type\":\"update_binding\",\"attributes\":[{\"key\":\"service_name\",\"value\":\"test\"},{\"key\":\"provider\",\"value\":\"iaa17rc02z9tfec74pq8avjg5uj6kj8d57992q7ys0\"},{\"key\":\"owner\",\"value\":\"iaa17rc02z9tfec74pq8avjg5uj6kj8d57992q7ys0\"}]}]}]","logs":[{"msg_index":0,"log":"","events":[{"type":"message","attributes":[{"key":"action","value":"update_service_binding"},{"key":"sender","value":"iaa17rc02z9tfec74pq8avjg5uj6kj8d57992q7ys0"},{"key":"module","value":"service"},{"key":"sender","value":"iaa17rc02z9tfec74pq8avjg5uj6kj8d57992q7ys0"}]},{"type":"transfer","attributes":[{"key":"recipient","value":"iaa1c8t8npfed4xc29755wwwvw2x834q36828duh55"},{"key":"sender","value":"iaa17rc02z9tfec74pq8avjg5uj6kj8d57992q7ys0"},{"key":"amount","value":"100upoint"}]},{"type":"update_binding","attributes":[{"key":"service_name","value":"test"},{"key":"provider","value":"iaa17rc02z9tfec74pq8avjg5uj6kj8d57992q7ys0"},{"key":"owner","value":"iaa17rc02z9tfec74pq8avjg5uj6kj8d57992q7ys0"}]}]}],"info":"","gas_wanted":"400000","gas_used":"75555","tx":null,"timestamp":""}
```

## disable

禁用一个可用的服务绑定。

```bash
irita tx service disable [service-name] [provider-address] [flags]
```

**参数：**

| 名称             | 类型   | 必须 | 默认 | 描述                               |
| ---------------- | ------ | ---- | ---- | ---------------------------------- |
| service-name     | string | 是   |      | 服务名称                           |
| provider-address | string | 否   |      | 服务提供者地址，默认为签名账户地址 |

### 禁用服务绑定示例

```bash
irita tx service disable test --from=node0 --chain-id=irita-test -b=block  -y --home=testnet/node0/iritacli
```

结果

```json
{"height":"1752","txhash":"E806327E4E44BF552BE0ECF8F4B8D0A93C80653B7FBB7E4E83C49E99A6508115","codespace":"","code":0,"data":"0A190A1764697361626C655F736572766963655F62696E64696E67","raw_log":"[{\"events\":[{\"type\":\"disable_binding\",\"attributes\":[{\"key\":\"service_name\",\"value\":\"test\"},{\"key\":\"provider\",\"value\":\"iaa17rc02z9tfec74pq8avjg5uj6kj8d57992q7ys0\"},{\"key\":\"owner\",\"value\":\"iaa17rc02z9tfec74pq8avjg5uj6kj8d57992q7ys0\"}]},{\"type\":\"message\",\"attributes\":[{\"key\":\"action\",\"value\":\"disable_service_binding\"},{\"key\":\"module\",\"value\":\"service\"},{\"key\":\"sender\",\"value\":\"iaa17rc02z9tfec74pq8avjg5uj6kj8d57992q7ys0\"}]}]}]","logs":[{"msg_index":0,"log":"","events":[{"type":"disable_binding","attributes":[{"key":"service_name","value":"test"},{"key":"provider","value":"iaa17rc02z9tfec74pq8avjg5uj6kj8d57992q7ys0"},{"key":"owner","value":"iaa17rc02z9tfec74pq8avjg5uj6kj8d57992q7ys0"}]},{"type":"message","attributes":[{"key":"action","value":"disable_service_binding"},{"key":"module","value":"service"},{"key":"sender","value":"iaa17rc02z9tfec74pq8avjg5uj6kj8d57992q7ys0"}]}]}],"info":"","gas_wanted":"400000","gas_used":"57809","tx":null,"timestamp":""}
```

## enable

启用一个不可用的服务绑定。

```bash
irita tx service enable [service-name] [provider-address] [flags]
```

**参数：**

| 名称             | 类型   | 必须 | 默认 | 描述                               |
| ---------------- | ------ | ---- | ---- | ---------------------------------- |
| service-name     | string | 是   |      | 服务名称                           |
| provider-address | string | 否   |      | 服务提供者地址，默认为签名账户地址 |

**标志：**

| 名称，速记 | 默认 | 描述               | 必须 |
| ---------- | ---- | ------------------ | ---- |
| --deposit  |      | 启用绑定增加的押金 |      |

### 启用服务绑定示例

```bash
irita tx service enable test --deposit=10point --from=node0 --chain-id=irita-test -b=block  -y --home=testnet/node0/iritacli
```

结果

```json
{"height":"1760","txhash":"787B319D0667636AEC297AE776A20714F32315E51F6A9A6370364781285C2FFD","codespace":"","code":0,"data":"0A180A16656E61626C655F736572766963655F62696E64696E67","raw_log":"[{\"events\":[{\"type\":\"enable_binding\",\"attributes\":[{\"key\":\"service_name\",\"value\":\"test\"},{\"key\":\"provider\",\"value\":\"iaa17rc02z9tfec74pq8avjg5uj6kj8d57992q7ys0\"},{\"key\":\"owner\",\"value\":\"iaa17rc02z9tfec74pq8avjg5uj6kj8d57992q7ys0\"}]},{\"type\":\"message\",\"attributes\":[{\"key\":\"action\",\"value\":\"enable_service_binding\"},{\"key\":\"sender\",\"value\":\"iaa17rc02z9tfec74pq8avjg5uj6kj8d57992q7ys0\"},{\"key\":\"module\",\"value\":\"service\"},{\"key\":\"sender\",\"value\":\"iaa17rc02z9tfec74pq8avjg5uj6kj8d57992q7ys0\"}]},{\"type\":\"transfer\",\"attributes\":[{\"key\":\"recipient\",\"value\":\"iaa1c8t8npfed4xc29755wwwvw2x834q36828duh55\"},{\"key\":\"sender\",\"value\":\"iaa17rc02z9tfec74pq8avjg5uj6kj8d57992q7ys0\"},{\"key\":\"amount\",\"value\":\"10000000upoint\"}]}]}]","logs":[{"msg_index":0,"log":"","events":[{"type":"enable_binding","attributes":[{"key":"service_name","value":"test"},{"key":"provider","value":"iaa17rc02z9tfec74pq8avjg5uj6kj8d57992q7ys0"},{"key":"owner","value":"iaa17rc02z9tfec74pq8avjg5uj6kj8d57992q7ys0"}]},{"type":"message","attributes":[{"key":"action","value":"enable_service_binding"},{"key":"sender","value":"iaa17rc02z9tfec74pq8avjg5uj6kj8d57992q7ys0"},{"key":"module","value":"service"},{"key":"sender","value":"iaa17rc02z9tfec74pq8avjg5uj6kj8d57992q7ys0"}]},{"type":"transfer","attributes":[{"key":"recipient","value":"iaa1c8t8npfed4xc29755wwwvw2x834q36828duh55"},{"key":"sender","value":"iaa17rc02z9tfec74pq8avjg5uj6kj8d57992q7ys0"},{"key":"amount","value":"10000000upoint"}]}]}],"info":"","gas_wanted":"400000","gas_used":"74744","tx":null,"timestamp":""}
```

## refund-deposit

取回服务绑定的押金。服务绑定必须处于不可用状态。

```bash
irita tx service refund-deposit [service-name] [provider-address] [flags]
```

**参数：**

| 名称             | 类型   | 必须 | 默认 | 描述                               |
| ---------------- | ------ | ---- | ---- | ---------------------------------- |
| service-name     | string | 是   |      | 服务名称                           |
| provider-address | string | 否   |      | 服务提供者地址，默认为签名账户地址 |

### 取回押金示例

```bash
irita tx service refund-deposit test --from=node0 --chain-id=irita-test    -y --home=testnet/node0/iritacli
```

结果

```json
{
  "height": "523",
  "txhash": "3D0F9CBFEE8127981967AED670EC818236406C0A6C7A9C9E002D803ADF0FBABA",
  "raw_log": "<raw-log>",
  "logs": [
    {
      "msg_index": 0,
      "log": "",
      "events": [
        {
          "type": "message",
          "attributes": [
            {
              "key": "action",
              "value": "refund_service_deposit"
            },
            {
              "key": "sender",
              "value": "iaa1gjhme86fhwnv32974cnn70zayvgy8zxqk6huj3"
            },
            {
              "key": "module",
              "value": "service"
            },
            {
              "key": "sender",
              "value": "iaa1xnzrqm8kfvyyw8tpr2lcjenemd2u00pl3pwenv"
            }
          ]
        },
        {
          "type": "transfer",
          "attributes": [
            {
              "key": "recipient",
              "value": "iaa1xnzrqm8kfvyyw8tpr2lcjenemd2u00pl3pwenv"
            },
            {
              "key": "amount",
              "value": "10100point"
            }
          ]
        }
      ]
    }
  ],
  "gas_wanted": "200000",
  "gas_used": "54366"
}
```

## call

发起服务调用。

```bash
irita tx service call [flags]
```

**标志：**

| 名称，速记        | 默认  | 描述                                                       | 必须 |
| ----------------- | ----- | ---------------------------------------------------------- | ---- |
| --name            |       | 服务名称                                                   | 是   |
| --providers       |       | 服务提供者列表                                             | 是   |
| --service-fee-cap |       | 愿意为单个请求支付的最大服务费用                           | 是   |
| --data            |       | 请求输入的内容或文件路径；请求输入需符合 Input JSON Schema | 是   |
| --timeout         |       | 请求超时                                                   | 是   |
| --super-mode      | false | 签名者是否为超级用户                                       |      |
| --repeated        | false | 请求是否为重复性的                                         |      |
| --frequency       |       | 重复性请求的请求频率；默认为 `--timeout`                   |      |
| --total           |       | 重复性请求的请求总数，-1表示无限制                         |      |

### 服务调用示例

```bash
irita tx service call --service-name=test --providers=iaa17rc02z9tfec74pq8avjg5uj6kj8d57992q7ys0 --service-fee-cap=1upoint --data='{"header":{},"body":{"id":"001"}}'  --repeated --frequency=150 --total=100   --timeout=100  --from=node0 --chain-id=irita-test -b=block  -y --home=testnet/node0/iritacli 
```

结果

```json
{"height":"3231","txhash":"39E0CB69D2A7C4AFD2D637F98C0679E1BF273258E178462C60A8590BE69FDD3A","codespace":"","code":0,"data":"0A620A0C63616C6C5F7365727669636512520A503339453043423639443241374334414644324436333746393843303637394531424632373332353845313738343632433630413835393042453639464444334130303030303030303030303030303030","raw_log":"[{\"events\":[{\"type\":\"create_context\",\"attributes\":[{\"key\":\"request_context_id\",\"value\":\"39E0CB69D2A7C4AFD2D637F98C0679E1BF273258E178462C60A8590BE69FDD3A0000000000000000\"},{\"key\":\"service_name\",\"value\":\"test\"},{\"key\":\"consumer\",\"value\":\"iaa17rc02z9tfec74pq8avjg5uj6kj8d57992q7ys0\"}]},{\"type\":\"message\",\"attributes\":[{\"key\":\"action\",\"value\":\"call_service\"},{\"key\":\"module\",\"value\":\"service\"},{\"key\":\"sender\",\"value\":\"iaa17rc02z9tfec74pq8avjg5uj6kj8d57992q7ys0\"}]}]}]","logs":[{"msg_index":0,"log":"","events":[{"type":"create_context","attributes":[{"key":"request_context_id","value":"39E0CB69D2A7C4AFD2D637F98C0679E1BF273258E178462C60A8590BE69FDD3A0000000000000000"},{"key":"service_name","value":"test"},{"key":"consumer","value":"iaa17rc02z9tfec74pq8avjg5uj6kj8d57992q7ys0"}]},{"type":"message","attributes":[{"key":"action","value":"call_service"},{"key":"module","value":"service"},{"key":"sender","value":"iaa17rc02z9tfec74pq8avjg5uj6kj8d57992q7ys0"}]}]}],"info":"","gas_wanted":"400000","gas_used":"68978","tx":null,"timestamp":""}
```

**_注_：**`request_context_id` 是服务调用所创建的请求上下文的 ID 。

## respond

响应指定的服务请求。

```bash
irita tx service respond [flags]
```

**标志：**

| 名称，速记   | 默认 | 描述                                                         | 必须 |
| ------------ | ---- | ------------------------------------------------------------ | ---- |
| --request-id |      | 请求 ID                                                      | 是   |
| --result     |      | 响应结果的内容或文件路径；响应结果需符合 [iService Result Schema](.../../core_modules/schemas/iservice-result.md) | 是   |
| --data       |      | 响应输出的内容或文件路径；响应输出需符合 Output JSON Schema  |      |

### 响应服务请求示例

```bash
irita tx service respond --request-id=5C4324F9B7C1A0A20A7C272DBE75138C531D1BF4000D862A9866B0DA87EC8725000000000000000000000000000000010000000000000D020000 --result='{"code":200,"message":"ok"}' --data='{"header":{},"body": { "name":"test"}}' --from=node0 --chain-id=irita-test -b=block  -y --home=testnet/node0/iritacli
```

结果

```json
{"height":"3253","txhash":"4F795E3D851A85DF119AC10CC450693A49A8274B78D47312B7B507DB8EB588B1","codespace":"","code":0,"data":"0A110A0F726573706F6E645F73657276696365","raw_log":"[{\"events\":[{\"type\":\"complete_batch\",\"attributes\":[{\"key\":\"request_context_id\",\"value\":\"39E0CB69D2A7C4AFD2D637F98C0679E1BF273258E178462C60A8590BE69FDD3A0000000000000000\"},{\"key\":\"request_context_state\",\"value\":\"{\\\"batch_counter\\\":1,\\\"state\\\":1,\\\"batch_response_threshold\\\":0,\\\"batch_request_count\\\":1,\\\"batch_response_count\\\":1}\"}]},{\"type\":\"message\",\"attributes\":[{\"key\":\"action\",\"value\":\"respond_service\"},{\"key\":\"sender\",\"value\":\"iaa1wnrllnlwm67jvqs963x8dhvqz4kdaaswmawrqy\"},{\"key\":\"module\",\"value\":\"service\"},{\"key\":\"sender\",\"value\":\"iaa17rc02z9tfec74pq8avjg5uj6kj8d57992q7ys0\"}]},{\"type\":\"respond_service\",\"attributes\":[{\"key\":\"request_context_id\",\"value\":\"39E0CB69D2A7C4AFD2D637F98C0679E1BF273258E178462C60A8590BE69FDD3A0000000000000000\"},{\"key\":\"request_id\",\"value\":\"39E0CB69D2A7C4AFD2D637F98C0679E1BF273258E178462C60A8590BE69FDD3A000000000000000000000000000000010000000000000C9F0000\"},{\"key\":\"service_name\",\"value\":\"test\"},{\"key\":\"provider\",\"value\":\"iaa17rc02z9tfec74pq8avjg5uj6kj8d57992q7ys0\"},{\"key\":\"consumer\",\"value\":\"iaa17rc02z9tfec74pq8avjg5uj6kj8d57992q7ys0\"}]},{\"type\":\"transfer\",\"attributes\":[{\"key\":\"recipient\",\"value\":\"iaa1k83ewmsh9t5ra60urmcj5jc8ev2agmfez0jawf\"},{\"key\":\"sender\",\"value\":\"iaa1wnrllnlwm67jvqs963x8dhvqz4kdaaswmawrqy\"},{\"key\":\"amount\"}]}]}]","logs":[{"msg_index":0,"log":"","events":[{"type":"complete_batch","attributes":[{"key":"request_context_id","value":"39E0CB69D2A7C4AFD2D637F98C0679E1BF273258E178462C60A8590BE69FDD3A0000000000000000"},{"key":"request_context_state","value":"{\"batch_counter\":1,\"state\":1,\"batch_response_threshold\":0,\"batch_request_count\":1,\"batch_response_count\":1}"}]},{"type":"message","attributes":[{"key":"action","value":"respond_service"},{"key":"sender","value":"iaa1wnrllnlwm67jvqs963x8dhvqz4kdaaswmawrqy"},{"key":"module","value":"service"},{"key":"sender","value":"iaa17rc02z9tfec74pq8avjg5uj6kj8d57992q7ys0"}]},{"type":"respond_service","attributes":[{"key":"request_context_id","value":"39E0CB69D2A7C4AFD2D637F98C0679E1BF273258E178462C60A8590BE69FDD3A0000000000000000"},{"key":"request_id","value":"39E0CB69D2A7C4AFD2D637F98C0679E1BF273258E178462C60A8590BE69FDD3A000000000000000000000000000000010000000000000C9F0000"},{"key":"service_name","value":"test"},{"key":"provider","value":"iaa17rc02z9tfec74pq8avjg5uj6kj8d57992q7ys0"},{"key":"consumer","value":"iaa17rc02z9tfec74pq8avjg5uj6kj8d57992q7ys0"}]},{"type":"transfer","attributes":[{"key":"recipient","value":"iaa1k83ewmsh9t5ra60urmcj5jc8ev2agmfez0jawf"},{"key":"sender","value":"iaa1wnrllnlwm67jvqs963x8dhvqz4kdaaswmawrqy"},{"key":"amount","value":""}]}]}],"info":"","gas_wanted":"400000","gas_used":"88704","tx":null,"timestamp":""}
```

## update

更新指定的请求上下文。

```bash
irita tx service update [request-context-id] [flags]
```

**参数：**

| 名称               | 类型   | 必须 | 默认 | 描述          |
| ------------------ | ------ | ---- | ---- | ------------- |
| request-context-id | string | 是   |      | 请求上下文 ID |

**标志：**

| 名称，速记        | 默认 | 描述                                           | 必须 |
| ----------------- | ---- | ---------------------------------------------- | ---- |
| --providers       |      | 服务提供者列表，为空则不更新                   |      |
| --service-fee-cap |      | 愿意为单个请求支付的最大服务费用，为空则不更新 |      |
| --timeout         |      | 请求超时，为0则不更新                          |      |
| --frequency       |      | 请求频率，为0则不更新                          |      |
| --total           |      | 请求总数，为0则不更新                          |      |

### 更新请求上下文示例

```bash
irita tx service update 39E0CB69D2A7C4AFD2D637F98C0679E1BF273258E178462C60A8590BE69FDD3A0000000000000000 --service-fee-cap=1point --timeout=0 --frequency=100 --total=50 --from=node0 --chain-id=irita-test -b=block  -y --home=testnet/node0/iritacli
```

结果

```json
{"height":"3114","txhash":"3939EF00D1B045227997BB0700C50FC0906B7BC53CB243AC48643649C4CA159E","codespace":"","code":0,"data":"0A180A167570646174655F726571756573745F636F6E74657874","raw_log":"[{\"events\":[{\"type\":\"message\",\"attributes\":[{\"key\":\"action\",\"value\":\"update_request_context\"},{\"key\":\"module\",\"value\":\"service\"},{\"key\":\"sender\",\"value\":\"iaa17rc02z9tfec74pq8avjg5uj6kj8d57992q7ys0\"}]},{\"type\":\"update_context\",\"attributes\":[{\"key\":\"request_context_id\",\"value\":\"8F97355330F8C41DDF4863B48B502C868480B8E47D42DA4B838884BB2175ECB50000000000000000\"},{\"key\":\"consumer\",\"value\":\"iaa17rc02z9tfec74pq8avjg5uj6kj8d57992q7ys0\"}]}]}]","logs":[{"msg_index":0,"log":"","events":[{"type":"message","attributes":[{"key":"action","value":"update_request_context"},{"key":"module","value":"service"},{"key":"sender","value":"iaa17rc02z9tfec74pq8avjg5uj6kj8d57992q7ys0"}]},{"type":"update_context","attributes":[{"key":"request_context_id","value":"8F97355330F8C41DDF4863B48B502C868480B8E47D42DA4B838884BB2175ECB50000000000000000"},{"key":"consumer","value":"iaa17rc02z9tfec74pq8avjg5uj6kj8d57992q7ys0"}]}]}],"info":"","gas_wanted":"400000","gas_used":"62037","tx":null,"timestamp":""}
```

## pause

暂停一个正在进行的请求上下文。

```bash
irita tx service pause [request-context-id] [flags]
```

**参数：**

| 名称               | 类型   | 必须 | 默认 | 描述          |
| ------------------ | ------ | ---- | ---- | ------------- |
| request-context-id | string | 是   |      | 请求上下文 ID |

### 暂停请求上下文示例

```bash
irita tx service pause 39E0CB69D2A7C4AFD2D637F98C0679E1BF273258E178462C60A8590BE69FDD3A0000000000000000 --from=node0 --chain-id=irita-test -b=block -y --home=testnet/node0/iritacli
```

结果

```json
{"height":"3236","txhash":"CF32DEE3A3FA08A2B68E0705684071A3C16BFCBEACEFFC1C35E41FD8FBB186C4","codespace":"","code":0,"data":"0A170A1570617573655F726571756573745F636F6E74657874","raw_log":"[{\"events\":[{\"type\":\"message\",\"attributes\":[{\"key\":\"action\",\"value\":\"pause_request_context\"},{\"key\":\"module\",\"value\":\"service\"},{\"key\":\"sender\",\"value\":\"iaa17rc02z9tfec74pq8avjg5uj6kj8d57992q7ys0\"}]},{\"type\":\"pause_context\",\"attributes\":[{\"key\":\"request_context_id\",\"value\":\"39E0CB69D2A7C4AFD2D637F98C0679E1BF273258E178462C60A8590BE69FDD3A0000000000000000\"},{\"key\":\"consumer\",\"value\":\"iaa17rc02z9tfec74pq8avjg5uj6kj8d57992q7ys0\"}]}]}]","logs":[{"msg_index":0,"log":"","events":[{"type":"message","attributes":[{"key":"action","value":"pause_request_context"},{"key":"module","value":"service"},{"key":"sender","value":"iaa17rc02z9tfec74pq8avjg5uj6kj8d57992q7ys0"}]},{"type":"pause_context","attributes":[{"key":"request_context_id","value":"39E0CB69D2A7C4AFD2D637F98C0679E1BF273258E178462C60A8590BE69FDD3A0000000000000000"},{"key":"consumer","value":"iaa17rc02z9tfec74pq8avjg5uj6kj8d57992q7ys0"}]}]}],"info":"","gas_wanted":"400000","gas_used":"59790","tx":null,"timestamp":""}
```

## start

启动一个暂停的请求上下文。

```bash
irita tx service start [request-context-id] [flags]
```

**参数：**

| 名称               | 类型   | 必须 | 默认 | 描述          |
| ------------------ | ------ | ---- | ---- | ------------- |
| request-context-id | string | 是   |      | 请求上下文 ID |

### 启动请求上下文示例

```bash
irita tx service start 39E0CB69D2A7C4AFD2D637F98C0679E1BF273258E178462C60A8590BE69FDD3A0000000000000000 --from=node0 --chain-id=irita-test -b=block  -y --home=testnet/node0/iritacli
```

结果

```json
{"height":"3264","txhash":"D6D76F7AD0D7C1CA6851951A8EF41763E999696BB7B478832A94D883371CE59E","codespace":"","code":0,"data":"0A170A1573746172745F726571756573745F636F6E74657874","raw_log":"[{\"events\":[{\"type\":\"message\",\"attributes\":[{\"key\":\"action\",\"value\":\"start_request_context\"},{\"key\":\"module\",\"value\":\"service\"},{\"key\":\"sender\",\"value\":\"iaa17rc02z9tfec74pq8avjg5uj6kj8d57992q7ys0\"}]},{\"type\":\"start_context\",\"attributes\":[{\"key\":\"request_context_id\",\"value\":\"39E0CB69D2A7C4AFD2D637F98C0679E1BF273258E178462C60A8590BE69FDD3A0000000000000000\"},{\"key\":\"consumer\",\"value\":\"iaa17rc02z9tfec74pq8avjg5uj6kj8d57992q7ys0\"}]}]}]","logs":[{"msg_index":0,"log":"","events":[{"type":"message","attributes":[{"key":"action","value":"start_request_context"},{"key":"module","value":"service"},{"key":"sender","value":"iaa17rc02z9tfec74pq8avjg5uj6kj8d57992q7ys0"}]},{"type":"start_context","attributes":[{"key":"request_context_id","value":"39E0CB69D2A7C4AFD2D637F98C0679E1BF273258E178462C60A8590BE69FDD3A0000000000000000"},{"key":"consumer","value":"iaa17rc02z9tfec74pq8avjg5uj6kj8d57992q7ys0"}]}]}],"info":"","gas_wanted":"400000","gas_used":"61078","tx":null,"timestamp":""}
```

## kill

永久终止一个请求上下文。

```bash
irita tx service kill [request-context-id] [flags]
```

**参数：**

| 名称               | 类型   | 必须 | 默认 | 描述          |
| ------------------ | ------ | ---- | ---- | ------------- |
| request-context-id | string | 是   |      | 请求上下文 ID |

### 终止请求上下文示例

```bash
irita tx service kill 39E0CB69D2A7C4AFD2D637F98C0679E1BF273258E178462C60A8590BE69FDD3A0000000000000000 --from=node0 --chain-id=irita-test -b=block  -y --home=testnet/node0/iritacli
```

结果

```json
{"height":"3268","txhash":"900905D10D21AED7837B3FED228F38BD99192FC11B66491A43F1FF8758B24750","codespace":"","code":0,"data":"0A160A146B696C6C5F726571756573745F636F6E74657874","raw_log":"[{\"events\":[{\"type\":\"kill_context\",\"attributes\":[{\"key\":\"request_context_id\",\"value\":\"39E0CB69D2A7C4AFD2D637F98C0679E1BF273258E178462C60A8590BE69FDD3A0000000000000000\"},{\"key\":\"consumer\",\"value\":\"iaa17rc02z9tfec74pq8avjg5uj6kj8d57992q7ys0\"}]},{\"type\":\"message\",\"attributes\":[{\"key\":\"action\",\"value\":\"kill_request_context\"},{\"key\":\"module\",\"value\":\"service\"},{\"key\":\"sender\",\"value\":\"iaa17rc02z9tfec74pq8avjg5uj6kj8d57992q7ys0\"}]}]}]","logs":[{"msg_index":0,"log":"","events":[{"type":"kill_context","attributes":[{"key":"request_context_id","value":"39E0CB69D2A7C4AFD2D637F98C0679E1BF273258E178462C60A8590BE69FDD3A0000000000000000"},{"key":"consumer","value":"iaa17rc02z9tfec74pq8avjg5uj6kj8d57992q7ys0"}]},{"type":"message","attributes":[{"key":"action","value":"kill_request_context"},{"key":"module","value":"service"},{"key":"sender","value":"iaa17rc02z9tfec74pq8avjg5uj6kj8d57992q7ys0"}]}]}],"info":"","gas_wanted":"400000","gas_used":"60140","tx":null,"timestamp":""}
```

## set-withdraw-addr

设置所有者的服务费提取地址。

```bash
irita tx service set-withdraw-addr [withdrawal-address] [flags]
```

**参数：**

| 名称               | 类型   | 必须 | 默认 | 描述     |
| ------------------ | ------ | ---- | ---- | -------- |
| withdrawal-address | string | 是   |      | 提取地址 |

### 设置提取地址示例

```bash
irita tx service set-withdraw-addr yjctestaddress--from=node0 --chain-id=irita-test -b=block -y --home=testnet/node0/iritacli
```

结果

```json
{"height":"3285","txhash":"A87DF1B1F828AA087133390B982DA89B8D99C9D56374E8E3D633A0D06CB76DD5","codespace":"","code":0,"data":"0A160A147365745F77697468647261775F61646472657373","raw_log":"[{\"events\":[{\"type\":\"message\",\"attributes\":[{\"key\":\"action\",\"value\":\"set_withdraw_address\"},{\"key\":\"module\",\"value\":\"service\"},{\"key\":\"sender\",\"value\":\"iaa17rc02z9tfec74pq8avjg5uj6kj8d57992q7ys0\"}]},{\"type\":\"set_withdraw_address\",\"attributes\":[{\"key\":\"owner\",\"value\":\"iaa17rc02z9tfec74pq8avjg5uj6kj8d57992q7ys0\"},{\"key\":\"withdraw_address\",\"value\":\"iaa17rc02z9tfec74pq8avjg5uj6kj8d57992q7ys0\"}]}]}]","logs":[{"msg_index":0,"log":"","events":[{"type":"message","attributes":[{"key":"action","value":"set_withdraw_address"},{"key":"module","value":"service"},{"key":"sender","value":"iaa17rc02z9tfec74pq8avjg5uj6kj8d57992q7ys0"}]},{"type":"set_withdraw_address","attributes":[{"key":"owner","value":"iaa17rc02z9tfec74pq8avjg5uj6kj8d57992q7ys0"},{"key":"withdraw_address","value":"iaa17rc02z9tfec74pq8avjg5uj6kj8d57992q7ys0"}]}]}],"info":"","gas_wanted":"400000","gas_used":"52330","tx":null,"timestamp":""}
```

## withdraw-fees

提取服务提供者赚取的服务费。如未指定服务提供者，则提取该所有者全部服务提供者的服务费。

```bash
irita tx service withdraw-fees [provider] [flags]
```

**参数：**

| 名称     | 类型   | 必须 | 默认 | 描述           |
| -------- | ------ | ---- | ---- | -------------- |
| provider | string | 否   |      | 服务提供者地址 |

### 提取服务费示例

```bash
irita tx service withdraw-fees iaa17rc02z9tfec74pq8avjg5uj6kj8d57992q7ys0 --from=node0 --chain-id=irita-test -b=block -y --home=testnet/node0/iritacli
```

结果

```json
{"height":"3295","txhash":"2CABF824DEBEE1B230A288F9C3E616A7BA68793E44A81CE7A51439C00080ECBA","codespace":"","code":0,"data":"0A160A1477697468647261775F6561726E65645F66656573","raw_log":"[{\"events\":[{\"type\":\"message\",\"attributes\":[{\"key\":\"action\",\"value\":\"withdraw_earned_fees\"},{\"key\":\"sender\",\"value\":\"iaa1wnrllnlwm67jvqs963x8dhvqz4kdaaswmawrqy\"},{\"key\":\"module\",\"value\":\"service\"},{\"key\":\"sender\",\"value\":\"iaa17rc02z9tfec74pq8avjg5uj6kj8d57992q7ys0\"}]},{\"type\":\"transfer\",\"attributes\":[{\"key\":\"recipient\",\"value\":\"iaa17rc02z9tfec74pq8avjg5uj6kj8d57992q7ys0\"},{\"key\":\"sender\",\"value\":\"iaa1wnrllnlwm67jvqs963x8dhvqz4kdaaswmawrqy\"},{\"key\":\"amount\",\"value\":\"2upoint\"}]},{\"type\":\"withdraw_earned_fees\",\"attributes\":[{\"key\":\"provider\",\"value\":\"iaa17rc02z9tfec74pq8avjg5uj6kj8d57992q7ys0\"},{\"key\":\"owner\",\"value\":\"iaa17rc02z9tfec74pq8avjg5uj6kj8d57992q7ys0\"}]}]}]","logs":[{"msg_index":0,"log":"","events":[{"type":"message","attributes":[{"key":"action","value":"withdraw_earned_fees"},{"key":"sender","value":"iaa1wnrllnlwm67jvqs963x8dhvqz4kdaaswmawrqy"},{"key":"module","value":"service"},{"key":"sender","value":"iaa17rc02z9tfec74pq8avjg5uj6kj8d57992q7ys0"}]},{"type":"transfer","attributes":[{"key":"recipient","value":"iaa17rc02z9tfec74pq8avjg5uj6kj8d57992q7ys0"},{"key":"sender","value":"iaa1wnrllnlwm67jvqs963x8dhvqz4kdaaswmawrqy"},{"key":"amount","value":"2upoint"}]},{"type":"withdraw_earned_fees","attributes":[{"key":"provider","value":"iaa17rc02z9tfec74pq8avjg5uj6kj8d57992q7ys0"},{"key":"owner","value":"iaa17rc02z9tfec74pq8avjg5uj6kj8d57992q7ys0"}]}]}],"info":"","gas_wanted":"400000","gas_used":"64193","tx":null,"timestamp":""}
```

## definition

查询服务定义。

```bash
irita query service definition [service-name] [flags]
```

**参数：**

| 名称         | 类型   | 必须 | 默认 | 描述     |
| ------------ | ------ | ---- | ---- | -------- |
| service-name | string | 是   |      | 服务名称 |

### 查询服务定义示例

```bash
irita query service definition test --chain-id=irita-test
```

结果

```json
author: iaa17rc02z9tfec74pq8avjg5uj6kj8d57992q7ys0
author_description: test
description: test
name: test
schemas: '{"input":{"type":"object","properties":{"id":{"type":"string"}}},"output":{"type":"object","properties":{"name":{"type":"string"}}}}'
tags:
- test
```

## binding

查询指定的服务绑定。

```bash
irita q service binding [service-name] [provider] [flags]
```

**参数：**

| 名称         | 类型   | 必须 | 默认 | 描述           |
| ------------ | ------ | ---- | ---- | -------------- |
| service-name | string | 是   |      | 服务名称       |
| provider     | string | 是   |      | 服务提供者地址 |

### 查询服务绑定示例

```bash
irita query service binding test iaa17rc02z9tfec74pq8avjg5uj6kj8d57992q7ys0 --chain-id=irita-test
```

结果

```json
available: true
deposit:
- amount: "2012948138"
  denom: upoint
disabled_time: "0001-01-01T00:00:00Z"
options: '{}'
owner: iaa17rc02z9tfec74pq8avjg5uj6kj8d57992q7ys0
pricing: '{"price":"1upoint"}'
provider: iaa17rc02z9tfec74pq8avjg5uj6kj8d57992q7ys0
qos: "30"
service_name: test
```

## bindings

查询指定服务的绑定列表，可指定 `owner` 。

```bash
irita q service bindings [service-name] [flags]
```

**参数：**

| 名称         | 类型   | 必须 | 默认 | 描述     |
| ------------ | ------ | ---- | ---- | -------- |
| service-name | string | 是   |      | 服务名称 |

**标志：**

| 名称，速记 | 默认 | 描述           | 必须 |
| ---------- | ---- | -------------- | ---- |
| --owner    |      | 所有者账户地址 | 否   |

### 查询服务绑定列表示例

```bash
irita q service bindings test  --chain-id=irita-test
```

结果

```json
pagination:
  next_key: null
  total: "0"
service_bindings:
- available: true
  deposit:
  - amount: "2012948138"
    denom: upoint
  disabled_time: "0001-01-01T00:00:00Z"
  options: '{}'
  owner: iaa17rc02z9tfec74pq8avjg5uj6kj8d57992q7ys0
  pricing: '{"price":"1upoint"}'
  provider: iaa17rc02z9tfec74pq8avjg5uj6kj8d57992q7ys0
  qos: "30"
  service_name: test
```

## requests

通过服务绑定或请求上下文 ID 查询当前活跃的服务请求列表。

```bash
irita q service requests [service-name] [provider] | [request-context-id] [batch-counter] [flags]
```

**参数：**

| 名称               | 类型   | 必须 | 默认 | 描述                                                 |
| ------------------ | ------ | ---- | ---- | ---------------------------------------------------- |
| service-name       | string | 否   |      | 服务名称，如通过绑定查询时必须指定                   |
| provider           | string | 否   |      | 服务提供者地址，如通过绑定查询时必须指定             |
| request-context-id | string | 否   |      | 请求上下文 ID，如通过请求上下文查询时必须指定        |
| batch-counter      | uint64 | 否   |      | 请求上下文批次计数器，如通过请求上下文查询时必须指定 |

### 查询服务绑定的请求列表示例

```bash
irita query service requests test iaa17rc02z9tfec74pq8avjg5uj6kj8d57992q7ys0 --chain-id=irita-test
```

结果

```json
pagination:
  next_key: null
  total: "0"
requests:
- consumer: iaa17rc02z9tfec74pq8avjg5uj6kj8d57992q7ys0
  expiration_height: "3329"
  id: 62735964FE3EA957C395E828FA31977EC5A7A2A92A829D2BAC5AAFF7BA43ED0C000000000000000000000000000000040000000000000C9D0000
  input: '{"header":{},"body":{"id":"001"}}'
  provider: iaa17rc02z9tfec74pq8avjg5uj6kj8d57992q7ys0
  request_context_batch_counter: "4"
  request_context_id: 62735964FE3EA957C395E828FA31977EC5A7A2A92A829D2BAC5AAFF7BA43ED0C0000000000000000
  request_height: "3229"
  service_fee:
  - amount: "1"
    denom: upoint
  service_name: test
- consumer: iaa17rc02z9tfec74pq8avjg5uj6kj8d57992q7ys0
  expiration_height: "3338"
  id: 479D2E87AC48C877A1530F055D5AFFC9AC7033248EEE16AD884600552AD5ADF4000000000000000000000000000000080000000000000CA60000
  input: '{"header":{},"body":{"id":"001"}}'
  provider: iaa17rc02z9tfec74pq8avjg5uj6kj8d57992q7ys0
  request_context_batch_counter: "8"
  request_context_id: 479D2E87AC48C877A1530F055D5AFFC9AC7033248EEE16AD884600552AD5ADF40000000000000000
  request_height: "3238"
  service_fee:
  - amount: "1"
    denom: upoint
  service_name: test
```

### 查询请求上下文的服务请求示例

```bash
irita query service requests 5C4324F9B7C1A0A20A7C272DBE75138C531D1BF4000D862A9866B0DA87EC87250000000000000000 1 --chain-id=irita-test
```

结果

```json
pagination:
  next_key: null
  total: "0"
requests:
- consumer: iaa17rc02z9tfec74pq8avjg5uj6kj8d57992q7ys0
  expiration_height: "3430"
  id: 5C4324F9B7C1A0A20A7C272DBE75138C531D1BF4000D862A9866B0DA87EC8725000000000000000000000000000000010000000000000D020000
  input: '{"header":{},"body":{"id":"001"}}'
  provider: iaa17rc02z9tfec74pq8avjg5uj6kj8d57992q7ys0
  request_context_batch_counter: "1"
  request_context_id: 5C4324F9B7C1A0A20A7C272DBE75138C531D1BF4000D862A9866B0DA87EC87250000000000000000
  request_height: "3330"
  service_fee:
  - amount: "1"
    denom: upoint
  service_name: test
```

## request

通过请求 ID 查询服务请求。

```bash
irita query service request [request-id] [flags]
```

**参数：**

| 名称       | 类型   | 必须 | 默认 | 描述        |
| ---------- | ------ | ---- | ---- | ----------- |
| request-id | string | 是   |      | 服务请求 ID |

### 查询服务请求示例

```bash
irita q service request 5C4324F9B7C1A0A20A7C272DBE75138C531D1BF4000D862A9866B0DA87EC8725000000000000000000000000000000010000000000000D020000 --chain-id=irita-test
```

结果

```json
consumer: iaa17rc02z9tfec74pq8avjg5uj6kj8d57992q7ys0
expiration_height: "3430"
id: 5C4324F9B7C1A0A20A7C272DBE75138C531D1BF4000D862A9866B0DA87EC8725000000000000000000000000000000010000000000000D020000
input: '{"header":{},"body":{"id":"001"}}'
provider: iaa17rc02z9tfec74pq8avjg5uj6kj8d57992q7ys0
request_context_batch_counter: "1"
request_context_id: 5C4324F9B7C1A0A20A7C272DBE75138C531D1BF4000D862A9866B0DA87EC87250000000000000000
request_height: "3330"
service_fee:
- amount: "1"
  denom: upoint
service_name: test
```

## response

查询指定服务请求的服务响应。

```bash
irita query service response [request-id] [flags]
```

**参数：**

| 名称       | 类型   | 必须 | 默认 | 描述    |
| ---------- | ------ | ---- | ---- | ------- |
| request-id | string | 是   |      | 请求 ID |

### 查询服务响应示例

```bash
irita query service response 5C4324F9B7C1A0A20A7C272DBE75138C531D1BF4000D862A9866B0DA87EC8725000000000000000000000000000000010000000000000D020000  --chain-id=irita-test
```

结果

```json
consumer: iaa17rc02z9tfec74pq8avjg5uj6kj8d57992q7ys0
output: '{"header":{},"body":{"name":"test"}}'
provider: iaa17rc02z9tfec74pq8avjg5uj6kj8d57992q7ys0
request_context_batch_counter: "1"
request_context_id: 5C4324F9B7C1A0A20A7C272DBE75138C531D1BF4000D862A9866B0DA87EC87250000000000000000
result: '{"code":200,"message":"ok"}'
```

## responses

通过请求上下文 ID 以及批次计数器查询活跃的服务响应。

```bash
irita query service responses [request-context-id] [batch-counter] [flags]
```

**参数：**

| 名称               | 类型   | 必须 | 默认 | 描述                 |
| ------------------ | ------ | ---- | ---- | -------------------- |
| request-context-id | string | 是   |      | 请求上下文 ID        |
| batch-counter      | uint64 | 是   |      | 请求上下文批次计数器 |

### 查询请求上下文服务响应示例

```bash
irita query service responses 5C4324F9B7C1A0A20A7C272DBE75138C531D1BF4000D862A9866B0DA87EC87250000000000000000 1 --chain-id=irita-test
```

结果

```json
pagination:
  next_key: null
  total: "0"
responses:
- consumer: iaa17rc02z9tfec74pq8avjg5uj6kj8d57992q7ys0
  output: '{"header":{},"body":{"name":"test"}}'
  provider: iaa17rc02z9tfec74pq8avjg5uj6kj8d57992q7ys0
  request_context_batch_counter: "1"
  request_context_id: 5C4324F9B7C1A0A20A7C272DBE75138C531D1BF4000D862A9866B0DA87EC87250000000000000000
  result: '{"code":200,"message":"ok"}'
```

## request-context

查询指定的请求上下文。

```bash
irita query service request-context [request-context-id] [flags]
```

**参数：**

| 名称               | 类型   | 必须 | 默认 | 描述          |
| ------------------ | ------ | ---- | ---- | ------------- |
| request-context-id | string | 是   |      | 请求上下文 ID |

### 查询请求上下文示例

```bash
irita query service request-context 5C4324F9B7C1A0A20A7C272DBE75138C531D1BF4000D862A9866B0DA87EC87250000000000000000  --chain-id=irita-test
```

结果

```json
batch_counter: "1"
batch_request_count: 1
batch_response_count: 1
batch_response_threshold: 0
batch_state: BATCH_COMPLETED
consumer: iaa17rc02z9tfec74pq8avjg5uj6kj8d57992q7ys0
input: '{"header":{},"body":{"id":"001"}}'
module_name: ""
providers:
- iaa17rc02z9tfec74pq8avjg5uj6kj8d57992q7ys0
repeated: true
repeated_frequency: "150"
repeated_total: "100"
response_threshold: 0
service_fee_cap:
- amount: "1"
  denom: upoint
service_name: test
state: RUNNING
timeout: "100"
```

## withdraw-addr

查询所有者的服务费提取地址。

```bash
irita query service withdraw-addr [address] [flags]
```

**参数：**

| 名称    | 类型   | 必须 | 默认 | 描述           |
| ------- | ------ | ---- | ---- | -------------- |
| address | string | 是   |      | 所有者账户地址 |

### 查询提取地址示例

```bash
irita query service withdraw-addr iaa17rc02z9tfec74pq8avjg5uj6kj8d57992q7ys0 --chain-id=irita-test
```

结果

```text
withdraw_address: iaa17rc02z9tfec74pq8avjg5uj6kj8d57992q7ys0
```

## fees

查询指定服务提供者赚取的服务费。

```bash
irita query service fees [provider] [flags]
```

**参数：**

| 名称     | 类型   | 必须 | 默认 | 描述           |
| -------- | ------ | ---- | ---- | -------------- |
| provider | string | 是   |      | 服务提供者地址 |

### 查询服务费示例

```bash
irita query service fees iaa17rc02z9tfec74pq8avjg5uj6kj8d57992q7ys0  --chain-id=irita-test
```

结果

```json
fees:
- amount: "1"
  denom: upoint
```

