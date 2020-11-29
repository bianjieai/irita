<!--
order: 1
-->

# 积分

## issue

发行积分。

```bash
irita tx token issue [flags]
```

**标志：**

| 名称，速记       | 类型    | 必须 | 默认          | 描述                                                                     |
| ---------------- | ------- | ---- | ------------- | ------------------------------------------------------------------------ |
| --name           | string  | 是   |               | 积分的名称；unicode字符，最大长度为32字节，例如 "IRITA Credit"             |
| --symbol         | string  | 是   |               | 积分的唯一标识符；长度在3到8之间，字母数字字符，以字符开始，不区分大小写 |
| --initial-supply | uint64  | 是   |               | 积分的初始供应；增发前的数量不应超过1000亿                               |
| --max-supply     | uint64  |      | 1000000000000 | 积分的最大供应，总供应不能超过最大供应；增发前的数量不应超过1万亿        |
| --scale          | uint8   | 是   |               | 积分的精度，最多可以有18位小数；为0将默认到18位小数                      |
| --min-unit       | string  | 是   |               | 最小单位；长度在3到10之间，字母数字字符，以字符开始，不区分大小写        |
| --mintable       | boolean |      | false         | 发行后是否可以增发                                                       |

### 发行积分示例

```bash
irita tx token issue --symbol=mycredit --name="my credit" --initial-supply=10000 --max-supply=100000 --scale=0 --min-unit=mycretdit --mintable=true --from=node0 --chain-id=irita-test -b=block -o=json --indent -y --home=testnet/node0/iritacli
```

结果

```json
{
  "height": "11",
  "txhash": "261F5664FF680CD6E74FEF3C2A27E022DE09D402214F58A14BB1380BBE27FEDB",
  "raw_log": "<raw-log>",
  "logs": [
    {
      "msg_index": 0,
      "log": "",
      "events": [
        {
          "type": "issue_token",
          "attributes": [
            {
              "key": "symbol",
              "value": "mycredit"
            }
          ]
        },
        {
          "type": "message",
          "attributes": [
            {
              "key": "action",
              "value": "issue_token"
            },
            {
              "key": "sender",
              "value": "iaa19x57qemq4f3nret0hltehlgkv9vl0lxr0x552c"
            },
            {
              "key": "sender",
              "value": "iaa17ythh6yydzy9ngh2jxkhp40vwnmqgnx3qt840w"
            },
            {
              "key": "sender",
              "value": "iaa17ythh6yydzy9ngh2jxkhp40vwnmqgnx3qt840w"
            },
            {
              "key": "module",
              "value": "token"
            },
            {
              "key": "sender",
              "value": "iaa19x57qemq4f3nret0hltehlgkv9vl0lxr0x552c"
            }
          ]
        },
        {
          "type": "transfer",
          "attributes": [
            {
              "key": "recipient",
              "value": "iaa17ythh6yydzy9ngh2jxkhp40vwnmqgnx3qt840w"
            },
            {
              "key": "amount",
              "value": "4672point"
            },
            {
              "key": "recipient",
              "value": "iaa1h0srvfqqv2336aasp223seps59m6smrf2ccjna"
            },
            {
              "key": "amount",
              "value": "1868point"
            },
            {
              "key": "recipient",
              "value": "iaa19x57qemq4f3nret0hltehlgkv9vl0lxr0x552c"
            },
            {
              "key": "amount",
              "value": "10000mycretdit"
            }
          ]
        }
      ]
    }
  ],
  "gas_wanted": "200000",
  "gas_used": "119510"
}
```

## edit

编辑存在的积分。可编辑的属性包括：名称、最大供应以及可增发性

```bash
irita tx token edit [symbol] [flags]
```

**标志：**

| 名称，速记   | 类型   | 必须 | 默认  | 描述                                                  |
| ------------ | ------ | ---- | ----- | ----------------------------------------------------- |
| --name       | string |      |       | 积分名称，为空将不更新                                |
| --max-supply | uint   |      | 0     | 积分的最大供应量，应不小于当前的总供应量，为0将不更新 |
| --mintable   | bool   |      | false | 积分是否可以增发，默认为false                         |

### 编辑积分示例

```bash
irita tx token edit mycredit --max-supply=1000000 --mintable=true --from=node0 --chain-id=irita-test -b=block -o=json --indent -y --home=testnet/node0/iritacli
```

结果

```json
{
  "height": "35",
  "txhash": "D21095FDDB4F8AF89335710F8BAF2873631E907241BC985C7F5AE469E60C3B6E",
  "raw_log": "<raw-log>",
  "logs": [
    {
      "msg_index": 0,
      "log": "",
      "events": [
        {
          "type": "edit_token",
          "attributes": [
            {
              "key": "symbol",
              "value": "mycredit"
            }
          ]
        },
        {
          "type": "message",
          "attributes": [
            {
              "key": "action",
              "value": "edit_token"
            },
            {
              "key": "module",
              "value": "token"
            },
            {
              "key": "sender",
              "value": "iaa19x57qemq4f3nret0hltehlgkv9vl0lxr0x552c"
            }
          ]
        }
      ]
    }
  ],
  "gas_wanted": "200000",
  "gas_used": "43431"
}
```

## mint

增发积分到指定地址。

```bash
irita tx token mint [symbol] [flags]
```

**标志：**

| 名称，速记 | 类型   | 必须 | 默认 | 描述                                       |
| ---------- | ------ | ---- | ---- | ------------------------------------------ |
| --to       | string |      |      | 积分的接收地址，默认为交易发起者的账户地址 |
| --amount   | uint64 | 是   | 0    | 增发的数量                                 |

### 增发积分示例

```bash
irita tx token mint mycredit --to=iaa1lq8ye9aksqtyg2mn46esz9825zuxt5zatm5uxm --amount=1000 --from=node0 --chain-id=irita-test -b=block -o=json --indent -y --home=testnet/node0/iritacli
```

结果

```json
{
  "height": "52",
  "txhash": "49FEEFCB257656061640BA14EE203BA63AE301F69C5944F52E8C5C29E9953F84",
  "raw_log": "raw_log",
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
              "value": "mint_token"
            },
            {
              "key": "sender",
              "value": "iaa19x57qemq4f3nret0hltehlgkv9vl0lxr0x552c"
            },
            {
              "key": "sender",
              "value": "iaa17ythh6yydzy9ngh2jxkhp40vwnmqgnx3qt840w"
            },
            {
              "key": "sender",
              "value": "iaa17ythh6yydzy9ngh2jxkhp40vwnmqgnx3qt840w"
            },
            {
              "key": "module",
              "value": "token"
            },
            {
              "key": "sender",
              "value": "iaa19x57qemq4f3nret0hltehlgkv9vl0lxr0x552c"
            }
          ]
        },
        {
          "type": "mint_token",
          "attributes": [
            {
              "key": "symbol",
              "value": "mycredit"
            },
            {
              "key": "amount",
              "value": "1000"
            }
          ]
        },
        {
          "type": "transfer",
          "attributes": [
            {
              "key": "recipient",
              "value": "iaa17ythh6yydzy9ngh2jxkhp40vwnmqgnx3qt840w"
            },
            {
              "key": "amount",
              "value": "467point"
            },
            {
              "key": "recipient",
              "value": "iaa1h0srvfqqv2336aasp223seps59m6smrf2ccjna"
            },
            {
              "key": "amount",
              "value": "186point"
            },
            {
              "key": "recipient",
              "value": "iaa1lq8ye9aksqtyg2mn46esz9825zuxt5zatm5uxm"
            },
            {
              "key": "amount",
              "value": "1000mycretdit"
            }
          ]
        }
      ]
    }
  ],
  "gas_wanted": "200000",
  "gas_used": "118133"
}
```

## transfer

转让积分所有权。

```bash
irita tx token transfer [symbol] [flags]
```

**标志：**

| 名称，速记 | 类型   | 必须 | 默认 | 描述           |
| ---------- | ------ | ---- | ---- | -------------- |
| --to       | string | 是   |      | 新的所有者地址 |

### 转让积分所有权示例

```bash
irita tx token transfer mycredit --to=iaa1lq8ye9aksqtyg2mn46esz9825zuxt5zatm5uxm --from=node0 --chain-id=irita-test -b=block -o=json --indent -y --home=testnet/node0/iritacli
```

结果

```json
{
  "height": "69",
  "txhash": "81FBCF77E17D2739E38472C84927259E6DA335A6299CB5DBEFFA94B3337653A5",
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
              "value": "transfer_token_owner"
            },
            {
              "key": "module",
              "value": "token"
            },
            {
              "key": "sender",
              "value": "iaa19x57qemq4f3nret0hltehlgkv9vl0lxr0x552c"
            }
          ]
        },
        {
          "type": "transfer_token_owner",
          "attributes": [
            {
              "key": "symbol",
              "value": "mycredit"
            }
          ]
        }
      ]
    }
  ],
  "gas_wanted": "200000",
  "gas_used": "45477"
}
```

## token

查询指定的积分。

```bash
irita query token token [symbol] [flags]
```

### 查询积分示例

```bash
irita query token token mycredit -o=json --indent --chain-id=irita-test
```

结果

```json
{
  "type": "irismod/token/Token",
  "value": {
    "symbol": "mycredit",
    "name": "my credit",
    "min_unit": "mycretdit",
    "initial_supply": "10000",
    "max_supply": "1000000",
    "mintable": true,
    "owner": "iaa1lq8ye9aksqtyg2mn46esz9825zuxt5zatm5uxm"
  }
}
```

## tokens

查询已发行的所有积分，包括系统原生积分。如指定 `owner` 参数，则查询该 `owner` 发行的积分列表

```bash
irita query token tokens [owner] [flags]
```

### 查询所有积分示例

```bash
irita query token tokens -o=json --indent --chain-id=irita-test
```

结果

```json
[
  {
    "type": "irismod/token/Token",
    "value": {
      "symbol": "point",
      "name": "Network staking token ",
      "min_unit": "point",
      "initial_supply": "2000000000",
      "max_supply": "10000000000",
      "mintable": true,
      "owner": "iaa17ythh6yydzy9ngh2jxkhp40vwnmqgnx3qt840w"
    }
  },
  {
    "type": "irismod/token/Token",
    "value": {
      "symbol": "mycredit",
      "name": "my credit",
      "min_unit": "mycretdit",
      "initial_supply": "10000",
      "max_supply": "1000000",
      "mintable": true,
      "owner": "iaa1lq8ye9aksqtyg2mn46esz9825zuxt5zatm5uxm"
    }
  }
]
```

### 查询指定所有者的积分列表示例

```bash
irita query token tokens iaa1lq8ye9aksqtyg2mn46esz9825zuxt5zatm5uxm -o=json --indent --chain-id=irita-test
```

结果

```json
[
  {
    "type": "irismod/token/Token",
    "value": {
      "symbol": "mycredit",
      "name": "my credit",
      "min_unit": "mycretdit",
      "initial_supply": "10000",
      "max_supply": "1000000",
      "mintable": true,
      "owner": "iaa1lq8ye9aksqtyg2mn46esz9825zuxt5zatm5uxm"
    }
  }
]
```

## fee

查询积分相关的费用，包括发行和增发。

```bash
irita query token fee [symbol] [flags]
```

### 查询发行和增发积分费用示例

```bash
irita query token fee credit -o=json --indent --chain-id=irita-test
```

结果

```json
{
  "exist": false,
  "issue_fee": {
    "denom": "point",
    "amount": "8474"
  },
  "mint_fee": {
    "denom": "point",
    "amount": "847"
  }
}
```

**_注：_**`exist` 指示此 `symbol` 是否已经存在。
