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

| 名称，速记           | 默认 | 描述                        | 必须 |
| -------------------- | ---- | --------------------------- | ---- |
| --name               |      | 服务名称                    | 是   |
| --description        |      | 服务的描述                  |      |
| --author-description |      | 服务创建者的描述            |      |
| --tags               |      | 服务的标签列表              |      |
| --schemas            |      | 服务接口schemas的内容或文件路径 | 是   |

### 服务定义示例

```bash
irita tx service define --name=test --description=test --author-description=test --tags=test --schemas='{"input":{"type":"object","properties":{"id":{"type":"string"}}},"output":{"type":"object","properties":{"name":{"type":"string"}}}}' --from=node0 --chain-id=irita-test -b=block -o=json --indent -y --home=testnet/node0/iritacli
```

结果

```json
{
  "height": "56",
  "txhash": "726688B0E031B4CF328F4BEB63F277569459A7A53FBDF4382DE426BEF3E0293C",
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
              "value": "define_service"
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
        }
      ]
    }
  ],
  "gas_wanted": "200000",
  "gas_used": "46808"
}
```

## bind

绑定一个服务。

```bash
irita tx service bind [flags]
```

**标志：**

| 名称，速记     | 默认 | 描述                                                  | 必须 |
| -------------- | ---- | ----------------------------------------------------- | ---- |
| --service-name |      | 服务名称                                              | 是   |
| --provider | | 服务提供者地址，默认为签名账户地址 | 否 |
| --deposit      |      | 服务绑定的押金                                        | 是   |
| --pricing      |      | 服务定价内容或文件路径；服务定价需符合 [iService Pricing Schema](../../core_modules/schemas/iservice-pricing.md) |   是   |
| --qos |     | 服务质量，即最小响应时间 | 是 |

### 服务绑定示例

```bash
irita tx service bind --service-name=test --deposit=10000point --pricing='{"price":"1point"}' --qos=50 --from=node0 --chain-id=irita-test -b=block -o=json --indent -y --home=testnet/node0/iritacli
```

结果

```json
{
  "height": "75",
  "txhash": "3E617B9EF93F2078CC750F5ED04BD86ABC3EF8C115F1D4AB14A627A9F9A862DC",
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
              "value": "bind_service"
            },
            {
              "key": "sender",
              "value": "iaa1xnzrqm8kfvyyw8tpr2lcjenemd2u00pl3pwenv"
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
              "value": "iaa1gjhme86fhwnv32974cnn70zayvgy8zxqk6huj3"
            },
            {
              "key": "amount",
              "value": "10000point"
            }
          ]
        }
      ]
    }
  ],
  "gas_wanted": "200000",
  "gas_used": "78781"
}
```

## update-binding

更新已存在的服务绑定。

```bash
irita tx service update-binding [service-name] [provider] [flags]
```

**参数：**

| 名称      | 类型    | 必须 | 默认          | 描述                                                                     |
| ---------------- | ------- | ---- | ------------- | ------------------------------------------------------------------------ |
| service-name  | string  | 是   |             | 服务名称 |
| provider  | string  | 否   |             | 服务提供者地址，默认为签名账户地址 |

**标志：**
| 名称，速记     | 默认 | 描述                                                  | 必须 |
| -------------- | ---- | ----------------------------------------------------- | ---- |
| --deposit      |      | 增加的绑定押金，为空则不更新                                      |      |
| --pricing      |      | 服务定价内容或文件路径；服务定价需符合 [iService Pricing Schema](../../core_modules/schemas/iservice-pricing.md)；为空则不更新 |      |
| --qos |     | 服务质量，为0则不更新 | |

### 更新服务绑定示例

```bash
irita tx service update-binding test --deposit=100point --qos=30 --from=node0 --chain-id=irita-test -b=block -o=json --indent -y --home=testnet/node0/iritacli
```

结果

```json
{
  "height": "119",
  "txhash": "6EA1C8D1FBAC9B030FAFC8091C85288A1217A2D70C4BAB7EAFA32FE0DD8AABD2",
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
              "value": "update_service_binding"
            },
            {
              "key": "sender",
              "value": "iaa1xnzrqm8kfvyyw8tpr2lcjenemd2u00pl3pwenv"
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
              "value": "iaa1gjhme86fhwnv32974cnn70zayvgy8zxqk6huj3"
            },
            {
              "key": "amount",
              "value": "100point"
            }
          ]
        }
      ]
    }
  ],
  "gas_wanted": "200000",
  "gas_used": "60450"
}
```

## disable

禁用一个可用的服务绑定。

```bash
irita tx service disable [service-name] [provider] [flags]
```

**参数：**

| 名称      | 类型    | 必须 | 默认          | 描述                                                                     |
| ---------------- | ------- | ---- | ------------- | ------------------------------------------------------------------------ |
| service-name  | string  | 是   |             | 服务名称 |
| provider  | string  | 否   |             | 服务提供者地址，默认为签名账户地址 |

### 禁用服务绑定示例

```bash
irita tx service disable test --from=node0 --chain-id=irita-test -b=block -o=json --indent -y --home=testnet/node0/iritacli
```

结果

```json
{
  "height": "158",
  "txhash": "9AA4E1020486C0C2AC7F8A3318B3AD95B27C69C3C496B6F7754780F986C4FE72",
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
              "value": "disable_service_binding"
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
        }
      ]
    }
  ],
  "gas_wanted": "200000",
  "gas_used": "43427"
}
```

## enable

启用一个不可用的服务绑定。

```bash
irita tx service enable [service-name] [provider] [flags]
```

**参数：**

| 名称      | 类型    | 必须 | 默认          | 描述                                                                     |
| ---------------- | ------- | ---- | ------------- | ------------------------------------------------------------------------ |
| service-name  | string  | 是   |             | 服务名称 |
| provider  | string  | 否   |             | 服务提供者地址，默认为签名账户地址 |

**标志：**

| 名称，速记 | 默认 | 描述               | 必须 |
| ---------- | ---- | ------------------ | ---- |
| --deposit  |      | 启用绑定增加的押金 |      |

### 启用服务绑定示例

```bash
irita tx service enable test --deposit=10point --from=node0 --chain-id=irita-test -b=block -o=json --indent -y --home=testnet/node0/iritacli
```

结果

```json
{
  "height": "168",
  "txhash": "94899A9B25C1F0500DC830AC346FF846C50C309C9CAB60978407DC1019F51EFE",
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
              "value": "enable_service_binding"
            },
            {
              "key": "sender",
              "value": "iaa1xnzrqm8kfvyyw8tpr2lcjenemd2u00pl3pwenv"
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
              "value": "iaa1gjhme86fhwnv32974cnn70zayvgy8zxqk6huj3"
            },
            {
              "key": "amount",
              "value": "10point"
            }
          ]
        }
      ]
    }
  ],
  "gas_wanted": "200000",
  "gas_used": "59402"
}
```

## refund-deposit

取回服务绑定的押金。服务绑定必须处于不可用状态。

```bash
irita tx service refund-deposit [service-name] [provider] [flags]
```

**参数：**

| 名称      | 类型    | 必须 | 默认          | 描述                                                                     |
| ---------------- | ------- | ---- | ------------- | ------------------------------------------------------------------------ |
| service-name  | string  | 是   |             | 服务名称 |
| provider  | string  | 否   |             | 服务提供者地址，默认为签名账户地址 |

### 取回押金示例

```bash
irita tx service refund-deposit test --from=node0 --chain-id=irita-test -b=block -o=json --indent -y --home=testnet/node0/iritacli
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

| 名称，速记        | 默认  | 描述                                    | 必须 |
| ----------------- | ----- | --------------------------------------- | ---- |
| --name            |       | 服务名称                                | 是   |
| --providers       |       | 服务提供者列表                          | 是   |
| --service-fee-cap |       | 愿意为单个请求支付的最大服务费用        | 是   |
| --data            |       | 请求输入的内容或文件路径；请求输入需符合 Input JSON Schema | 是   |
| --timeout         |       | 请求超时                                |   是   |
| --super-mode      | false | 签名者是否为超级用户                    |
| --repeated        | false | 请求是否为重复性的                      |      |
| --frequency       |       | 重复性请求的请求频率；默认为 `--timeout` |      |
| --total           |       | 重复性请求的请求总数，-1表示无限制      |      |

### 服务调用示例

```bash
irita tx service call --service-name=test --providers=iaa1xnzrqm8kfvyyw8tpr2lcjenemd2u00pl3pwenv --service-fee-cap=1point --data='{"id":"001"}' --timeout=100 --repeated --frequency=150 --total=100 --from=node0 --chain-id=irita-test -b=block -o=json --indent -y --home=testnet/node0/irita
```

结果

```json
{
  "height": "233",
  "txhash": "B63FD99BB9DC26724B65519FA5184FC518C2F646D0C5A4B5D20DE6A85D534740",
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
              "value": "call_service"
            },
            {
              "key": "module",
              "value": "service"
            },
            {
              "key": "sender",
              "value": "iaa1xnzrqm8kfvyyw8tpr2lcjenemd2u00pl3pwenv"
            },
            {
              "key": "request_context_id",
              "value": "B63FD99BB9DC26724B65519FA5184FC518C2F646D0C5A4B5D20DE6A85D5347400000000000000000"
            }
          ]
        }
      ]
    }
  ],
  "gas_wanted": "200000",
  "gas_used": "50977"
}
```

**_注_：**`request_context_id` 是服务调用所创建的请求上下文的 ID 。

## respond

响应指定的服务请求。

```bash
irita tx service respond [flags]
```

**标志：**

| 名称，速记   | 默认 | 描述                                         | 必须 |
| ------------ | ---- | -------------------------------------------- | ---- |
| --request-id |      | 请求 ID                               | 是   |
| --result     |      | 响应结果的内容或文件路径；响应结果需符合 [iService Result Schema](.../../core_modules/schemas/iservice-result.md) | 是   |
| --data       |      | 响应输出的内容或文件路径；响应输出需符合 Output JSON Schema |      |

### 响应服务请求示例

```bash
irita tx service respond --request-id=B63FD99BB9DC26724B65519FA5184FC518C2F646D0C5A4B5D20DE6A85D5347400000000000000000000000000000000100000000000000E90000 --result='{"code":200,"message":"ok"}' --data='{"name":"test"}' --from=node0 --chain-id=irita-test -b=block -o=json --indent -y --home=testnet/node0/iritacli
```

结果

```json
{
  "height": "308",
  "txhash": "6CEC447C53386E0C7301B96E37A456E447D43058DB5F9A8E8E47493009C131F5",
  "raw_log": "<raw-log>",
  "logs": [
    {
      "msg_index": 0,
      "log": "",
      "events": [
        {
          "type": "complete_batch",
          "attributes": [
            {
              "key": "request_context_id",
              "value": "B63FD99BB9DC26724B65519FA5184FC518C2F646D0C5A4B5D20DE6A85D5347400000000000000000"
            },
            {
              "key": "request_context_state",
              "value": "{\"batch_counter\":1,\"state\":\"completed\",\"batch_response_threshold\":0,\"batch_request_count\":1,\"batch_response_count\":1}"
            }
          ]
        },
        {
          "type": "message",
          "attributes": [
            {
              "key": "action",
              "value": "respond_service"
            },
            {
              "key": "sender",
              "value": "iaa1adzua8vf72s6jvmr3kehxuuhvl2yzvths54nph"
            },
            {
              "key": "module",
              "value": "service"
            },
            {
              "key": "sender",
              "value": "iaa1xnzrqm8kfvyyw8tpr2lcjenemd2u00pl3pwenv"
            },
            {
              "key": "request_context_id",
              "value": "B63FD99BB9DC26724B65519FA5184FC518C2F646D0C5A4B5D20DE6A85D5347400000000000000000"
            },
            {
              "key": "request_id",
              "value": "B63FD99BB9DC26724B65519FA5184FC518C2F646D0C5A4B5D20DE6A85D5347400000000000000000000000000000000100000000000000E90000"
            }
          ]
        },
        {
          "type": "transfer",
          "attributes": [
            {
              "key": "recipient",
              "value": "iaa1h0srvfqqv2336aasp223seps59m6smrf2ccjna"
            },
            {
              "key": "amount"
            }
          ]
        }
      ]
    }
  ],
  "gas_wanted": "200000",
  "gas_used": "70217"
}
```

## update

更新指定的请求上下文。

```bash
irita tx service update [request-context-id] [flags]
```

**参数：**

| 名称      | 类型    | 必须 | 默认          | 描述                                                                     |
| ---------------- | ------- | ---- | ------------- | ------------------------------------------------------------------------ |
| request-context-id  | string  | 是   |             | 请求上下文 ID |

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
irita tx service update B63FD99BB9DC26724B65519FA5184FC518C2F646D0C5A4B5D20DE6A85D5347400000000000000000 --service-fee-cap=1point --timeout=0 --frequency=0 --total=50 --from=node0 --chain-id=irita-test -b=block -o=json --indent -y --home=testnet/node0/iritacli
```

结果

```json
{
  "height": "418",
  "txhash": "C7C4310C086C7B360A0A6E6A5EA3C12FFCFB095E4A941A30231F35CF7BC9C88F",
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
              "value": "update_request_context"
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
        }
      ]
    }
  ],
  "gas_wanted": "200000",
  "gas_used": "46578"
}
```

## pause

暂停一个正在进行的请求上下文。

```bash
irita tx service pause [request-context-id] [flags]
```

**参数：**

| 名称      | 类型    | 必须 | 默认          | 描述                                                                     |
| ---------------- | ------- | ---- | ------------- | ------------------------------------------------------------------------ |
| request-context-id  | string  | 是   |             | 请求上下文 ID |

### 暂停请求上下文示例

```bash
irita tx service pause B63FD99BB9DC26724B65519FA5184FC518C2F646D0C5A4B5D20DE6A85D5347400000000000000000 --from=node0 --chain-id=irita-test -b=block -o=json --indent -y --home=testnet/node0/iritacli
```

结果

```json
{
  "height": "437",
  "txhash": "8F7231B1EBA18A9577F247B26820B084641E476AF6478021548E0C1AC63A498C",
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
              "value": "pause_request_context"
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
        }
      ]
    }
  ],
  "gas_wanted": "200000",
  "gas_used": "44492"
}
```

## start

启动一个暂停的请求上下文。

```bash
irita tx service start [request-context-id] [flags]
```

**参数：**

| 名称      | 类型    | 必须 | 默认          | 描述                                                                     |
| ---------------- | ------- | ---- | ------------- | ------------------------------------------------------------------------ |
| request-context-id  | string  | 是   |             | 请求上下文 ID |

### 启动请求上下文示例

```bash
irita tx service start B63FD99BB9DC26724B65519FA5184FC518C2F646D0C5A4B5D20DE6A85D5347400000000000000000 --from=node0 --chain-id=irita-test -b=block -o=json --indent -y --home=testnet/node0/iritacli
```

结果

```json
{
  "height": "448",
  "txhash": "9A3873CBB68A9291F0102FDFAD076BEB1FD80235840E5CD4EE40CC27452A1E02",
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
              "value": "start_request_context"
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
        }
      ]
    }
  ],
  "gas_wanted": "200000",
  "gas_used": "45420"
}
```

## kill

永久终止一个请求上下文。

```bash
irita tx service kill [request-context-id] [flags]
```

**参数：**

| 名称      | 类型    | 必须 | 默认          | 描述                                                                     |
| ---------------- | ------- | ---- | ------------- | ------------------------------------------------------------------------ |
| request-context-id  | string  | 是   |             | 请求上下文 ID |

### 终止请求上下文示例

```bash
irita tx service kill B63FD99BB9DC26724B65519FA5184FC518C2F646D0C5A4B5D20DE6A85D5347400000000000000000 --from=node0 --chain-id=irita-test -b=block -o=json --indent -y --home=testnet/node0/iritacli
```

结果

```json
{
  "height": "461",
  "txhash": "93C1798DF1B79E1924764075157D8E3795697AF5AB21E4DAF1BC6C90633D9544",
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
              "value": "kill_request_context"
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
        }
      ]
    }
  ],
  "gas_wanted": "200000",
  "gas_used": "44492"
}
```

## set-withdraw-addr

设置所有者的服务费提取地址。

```bash
irita tx service set-withdraw-addr [withdrawal-address] [flags]
```

**参数：**

| 名称      | 类型    | 必须 | 默认          | 描述                                                                     |
| ---------------- | ------- | ---- | ------------- | ------------------------------------------------------------------------ |
| withdrawal-address  | string  | 是   |             | 提取地址 |

### 设置提取地址示例

```bash
irita tx service set-withdraw-addr c3gxeeqztx69uqaagcwfl7d0e9w0fggcs78yszv --from=node0 --chain-id=irita-test -b=block -o=json --indent -y --home=testnet/node0/iritacli
```

结果

```json
{
  "height": "133",
  "txhash": "4CAFD8D1D059BF431D3436F5A3DBB77A196E0BFE01C0CAF9C87EB02991F9A3BF",
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
              "value": "set_withdraw_address"
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
        }
      ]
    }
  ],
  "gas_wanted": "200000",
  "gas_used": "39598"
}
```

## withdraw-fees

提取服务提供者赚取的服务费。如未指定服务提供者，则提取该所有者全部服务提供者的服务费。

```bash
irita tx service withdraw-fees [provider] [flags]
```

**参数：**

| 名称      | 类型    | 必须 | 默认          | 描述                                                                     |
| ---------------- | ------- | ---- | ------------- | ------------------------------------------------------------------------ |
| provider  | string  | 否   |             | 服务提供者地址 |

### 提取服务费示例

```bash
irita tx service withdraw-fees iaa1xnzrqm8kfvyyw8tpr2lcjenemd2u00pl3pwenv --from=node0 --chain-id=irita-test -b=block -o=json --indent -y --home=testnet/node0/iritacli
```

结果

```json
{
  "height": "478",
  "txhash": "22BB2B2336A6D5A5E964D02B0AAA838D561493720FEB31F04CA30128D72940D4",
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
              "value": "withdraw_earned_fees"
            },
            {
              "key": "sender",
              "value": "iaa18up8anyjpal8rncm8rd4ukp5f7etga795gp33q"
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
              "value": "iaa1adzua8vf72s6jvmr3kehxuuhvl2yzvths54nph"
            },
            {
              "key": "amount",
              "value": "1point"
            }
          ]
        }
      ]
    }
  ],
  "gas_wanted": "200000",
  "gas_used": "54053"
}
```

## definition

查询服务定义。

```bash
irita query service definition [service-name] [flags]
```

**参数：**

| 名称      | 类型    | 必须 | 默认          | 描述                                                                     |
| ---------------- | ------- | ---- | ------------- | ------------------------------------------------------------------------ |
| service-name  | string  | 是   |             | 服务名称 |

### 查询服务定义示例

```bash
irita query service definition test -o=json --indent --chain-id=irita-test
```

结果

```json
{
  "type": "irismod/service/ServiceDefinition",
  "value": {
    "name": "test",
    "description": "test",
    "tags": [
      "test"
    ],
    "author": "iaa1xnzrqm8kfvyyw8tpr2lcjenemd2u00pl3pwenv",
    "author_description": "test",
    "schemas": "{\"input\":{\"type\":\"object\",\"properties\":{\"id\":{\"type\":\"string\"}}},\"output\":{\"type\":\"object\",\"properties\":{\"name\":{\"type\":\"string\"}}}}"
  }
}
```

## binding

查询指定的服务绑定。

```bash
irita q service binding [service-name] [provider] [flags]
```

**参数：**

| 名称      | 类型    | 必须 | 默认          | 描述                                                                     |
| ---------------- | ------- | ---- | ------------- | ------------------------------------------------------------------------ |
| service-name  | string  | 是   |             | 服务名称 |
| provider  | string  | 是   |             | 服务提供者地址 |

### 查询服务绑定示例

```bash
irita query service binding test iaa1xnzrqm8kfvyyw8tpr2lcjenemd2u00pl3pwenv -o json --indent --chain-id=irita-test
```

结果

```json
{
  "type": "irismod/service/ServiceBinding",
  "value": {
    "service_name": "test",
    "provider": "iaa1xnzrqm8kfvyyw8tpr2lcjenemd2u00pl3pwenv",
    "deposit": [
      {
        "denom": "point",
        "amount": "10000"
      }
    ],
    "pricing": "{\"price\":\"1point\"}",
    "qos": "50",
    "available": true,
    "disabled_time": "0001-01-01T00:00:00Z",
    "owner": "iaa1xnzrqm8kfvyyw8tpr2lcjenemd2u00pl3pwenv"
  }
}
```

## bindings

查询指定服务的绑定列表，可指定 `owner` 。

```bash
irita q service bindings [service-name] [flags]
```

**参数：**

| 名称      | 类型    | 必须 | 默认          | 描述                                                                     |
| ---------------- | ------- | ---- | ------------- | ------------------------------------------------------------------------ |
| service-name  | string  | 是   |             | 服务名称 |

**标志：**

| 名称，速记        | 默认 | 描述                                           | 必须 |
| ----------------- | ---- | ---------------------------------------------- | ---- |
| --owner       |      | 所有者账户地址                   |  否    |

### 查询服务绑定列表示例

```bash
irita q service bindings test -o json --indent --chain-id=irita-test
```

结果

```json
[
  {
    "service_name": "test",
    "provider": "iaa1xnzrqm8kfvyyw8tpr2lcjenemd2u00pl3pwenv",
    "deposit": [
      {
        "denom": "point",
        "amount": "10000"
      }
    ],
    "pricing": "{\"price\":\"1point\"}",
    "qos": "50",
    "available": true,
    "disabled_time": "0001-01-01T00:00:00Z",
    "owner": "iaa1xnzrqm8kfvyyw8tpr2lcjenemd2u00pl3pwenv"
  }
]
```

## requests

通过服务绑定或请求上下文 ID 查询当前活跃的服务请求列表。

```bash
irita q service requests [service-name] [provider] | [request-context-id] [batch-counter] [flags]
```

**参数：**

| 名称      | 类型    | 必须 | 默认          | 描述                                                                     |
| ---------------- | ------- | ---- | ------------- | ------------------------------------------------------------------------ |
| service-name  | string  | 否   |             | 服务名称，如通过绑定查询时必须指定 |
| provider  | string  | 否   |             | 服务提供者地址，如通过绑定查询时必须指定 |
| request-context-id  | string  | 否   |             | 请求上下文 ID，如通过请求上下文查询时必须指定 |
| batch-counter  | uint64  | 否   |             | 请求上下文批次计数器，如通过请求上下文查询时必须指定 |

### 查询服务绑定的请求列表示例

```bash
irita query service requests test iaa1xnzrqm8kfvyyw8tpr2lcjenemd2u00pl3pwenv -o=json --indent --chain-id=irita-test
```

结果

```json
[
  {
    "id": "B63FD99BB9DC26724B65519FA5184FC518C2F646D0C5A4B5D20DE6A85D5347400000000000000000000000000000000100000000000000E90000",
    "service_name": "test",
    "provider": "iaa1xnzrqm8kfvyyw8tpr2lcjenemd2u00pl3pwenv",
    "consumer": "iaa1xnzrqm8kfvyyw8tpr2lcjenemd2u00pl3pwenv",
    "input": "{\"id\":\"001\"}",
    "service_fee": [
      {
        "denom": "point",
        "amount": "1"
      }
    ],
    "request_height": "233",
    "expiration_height": "333",
    "request_context_id": "B63FD99BB9DC26724B65519FA5184FC518C2F646D0C5A4B5D20DE6A85D5347400000000000000000",
    "request_context_batch_counter": "1"
  }
]
```

### 查询请求上下文的服务请求示例

```bash
irita query service requests B63FD99BB9DC26724B65519FA5184FC518C2F646D0C5A4B5D20DE6A85D5347400000000000000000 1 -o=json --indent --chain-id=irita-test
```

结果

```json
[
  {
    "id": "B63FD99BB9DC26724B65519FA5184FC518C2F646D0C5A4B5D20DE6A85D5347400000000000000000000000000000000100000000000000E90000",
    "service_name": "test",
    "provider": "iaa1xnzrqm8kfvyyw8tpr2lcjenemd2u00pl3pwenv",
    "consumer": "iaa1xnzrqm8kfvyyw8tpr2lcjenemd2u00pl3pwenv",
    "input": "{\"id\":\"001\"}",
    "service_fee": [
      {
        "denom": "point",
        "amount": "1"
      }
    ],
    "request_height": "233",
    "expiration_height": "333",
    "request_context_id": "B63FD99BB9DC26724B65519FA5184FC518C2F646D0C5A4B5D20DE6A85D5347400000000000000000",
    "request_context_batch_counter": "1"
  }
]
```

## request

通过请求 ID 查询服务请求。

```bash
irita query service request [request-id] [flags]
```

**参数：**

| 名称      | 类型    | 必须 | 默认          | 描述                                                                     |
| ---------------- | ------- | ---- | ------------- | ------------------------------------------------------------------------ |
| request-id  | string  | 是   |             | 服务请求 ID |

### 查询服务请求示例

```bash
irita q service request B63FD99BB9DC26724B65519FA5184FC518C2F646D0C5A4B5D20DE6A85D5347400000000000000000000000000000000100000000000000E90000 -o=json --indent --chain-id=irita-test
```

结果

```json
{
  "type": "irismod/service/Request",
  "value": {
    "id": "B63FD99BB9DC26724B65519FA5184FC518C2F646D0C5A4B5D20DE6A85D5347400000000000000000000000000000000100000000000000E90000",
    "service_name": "test",
    "provider": "iaa1xnzrqm8kfvyyw8tpr2lcjenemd2u00pl3pwenv",
    "consumer": "iaa1xnzrqm8kfvyyw8tpr2lcjenemd2u00pl3pwenv",
    "input": "{\"id\":\"001\"}",
    "service_fee": [
      {
        "denom": "point",
        "amount": "1"
      }
    ],
    "request_height": "233",
    "expiration_height": "333",
    "request_context_id": "B63FD99BB9DC26724B65519FA5184FC518C2F646D0C5A4B5D20DE6A85D5347400000000000000000",
    "request_context_batch_counter": "1"
  }
}
```

## response

查询指定服务请求的服务响应。

```bash
irita query service response [request-id] [flags]
```

**参数：**

| 名称      | 类型    | 必须 | 默认          | 描述                                                                     |
| ---------------- | ------- | ---- | ------------- | ------------------------------------------------------------------------ |
| request-id  | string  | 是   |             | 请求 ID |

### 查询服务响应示例

```bash
irita query service response B63FD99BB9DC26724B65519FA5184FC518C2F646D0C5A4B5D20DE6A85D5347400000000000000000000000000000000100000000000000E90000 -o=json --indent --chain-id=irita-test
```

结果

```json
{
  "type": "irismod/service/Response",
  "value": {
    "provider": "iaa1xnzrqm8kfvyyw8tpr2lcjenemd2u00pl3pwenv",
    "consumer": "iaa1xnzrqm8kfvyyw8tpr2lcjenemd2u00pl3pwenv",
    "result": "{\"code\":200,\"message\":\"ok\"}",
    "output": "{\"name\":\"test\"}",
    "request_context_id": "B63FD99BB9DC26724B65519FA5184FC518C2F646D0C5A4B5D20DE6A85D5347400000000000000000",
    "request_context_batch_counter": "1"
  }
}
```

## responses

通过请求上下文 ID 以及批次计数器查询活跃的服务响应。

```bash
irita query service responses [request-context-id] [batch-counter] [flags]
```

**参数：**

| 名称      | 类型    | 必须 | 默认          | 描述                                                                     |
| ---------------- | ------- | ---- | ------------- | ------------------------------------------------------------------------ |
| request-context-id  | string  | 是   |             | 请求上下文 ID |
| batch-counter  | uint64  | 是   |             | 请求上下文批次计数器 |

### 查询请求上下文服务响应示例

```bash
irita query service responses B63FD99BB9DC26724B65519FA5184FC518C2F646D0C5A4B5D20DE6A85D5347400000000000000000 1 -o=json --indent --chain-id=irita-test
```

结果

```json
[
  {
    "type": "irismod/service/Response",
    "value": {
      "provider": "iaa1xnzrqm8kfvyyw8tpr2lcjenemd2u00pl3pwenv",
      "consumer": "iaa1xnzrqm8kfvyyw8tpr2lcjenemd2u00pl3pwenv",
      "result": "{\"code\":200,\"message\":\"ok\"}",
      "output": "{\"name\":\"test\"}",
      "request_context_id": "B63FD99BB9DC26724B65519FA5184FC518C2F646D0C5A4B5D20DE6A85D5347400000000000000000",
      "request_context_batch_counter": "1"
    }
  }
]
```

## request-context

查询指定的请求上下文。

```bash
irita query service request-context [request-context-id] [flags]
```

**参数：**

| 名称      | 类型    | 必须 | 默认          | 描述                                                                     |
| ---------------- | ------- | ---- | ------------- | ------------------------------------------------------------------------ |
| request-context-id  | string  | 是   |             | 请求上下文 ID |

### 查询请求上下文示例

```bash
irita query service request-context B63FD99BB9DC26724B65519FA5184FC518C2F646D0C5A4B5D20DE6A85D5347400000000000000000 -o=json --indent --chain-id=irita-test
```

结果

```json
{
  "type": "irismod/service/RequestContext",
  "value": {
    "service_name": "test",
    "providers": [
      "iaa1xnzrqm8kfvyyw8tpr2lcjenemd2u00pl3pwenv"
    ],
    "consumer": "iaa1xnzrqm8kfvyyw8tpr2lcjenemd2u00pl3pwenv",
    "input": "{\"id\":\"001\"}",
    "service_fee_cap": [
      {
        "denom": "point",
        "amount": "1"
      }
    ],
    "timeout": "100",
    "repeated": true,
    "repeated_frequency": "150",
    "repeated_total": "100",
    "batch_counter": "2",
    "batch_request_count": 1
  }
}
```

## withdraw-addr

查询所有者的服务费提取地址。

```bash
irita query service withdraw-addr [address] [flags]
```

**参数：**

| 名称      | 类型    | 必须 | 默认          | 描述                                                                     |
| ---------------- | ------- | ---- | ------------- | ------------------------------------------------------------------------ |
| address  | string  | 是   |             | 所有者账户地址 |

### 查询提取地址示例

```bash
irita query service withdraw-addr iaa1xnzrqm8kfvyyw8tpr2lcjenemd2u00pl3pwenv --chain-id=irita-test
```

结果

```text
c3gxeeqztx69uqaagcwfl7d0e9w0fggcs78yszv
```

## fees

查询指定服务提供者赚取的服务费。

```bash
irita query service fees [provider] [flags]
```

**参数：**

| 名称      | 类型    | 必须 | 默认          | 描述                                                                     |
| ---------------- | ------- | ---- | ------------- | ------------------------------------------------------------------------ |
| provider  | string  | 是   |             | 服务提供者地址 |

### 查询服务费示例

```bash
irita query service fees iaa1xnzrqm8kfvyyw8tpr2lcjenemd2u00pl3pwenv -o=json --indent --chain-id=irita-test
```

结果

```json
[
  {
    "denom": "point",
    "amount": "1"
  }
]
```
