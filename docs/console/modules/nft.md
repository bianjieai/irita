<!--
order: 2
-->

# 资产数字化建模

## issue

发行资产。

```bash
irita tx nft issue [denom] [flags]
```

**参数：**

| 名称      | 类型    | 必须 | 默认          | 描述                                                                     |
| ---------------- | ------- | ---- | ------------- | ------------------------------------------------------------------------ |
| denom  | string  | 是   |             | 资产的类别，全局唯一；长度为3到64，字母数字字符，以字母开始 |

**标志：**

| 名称，速记       | 类型    | 必须 | 默认          | 描述                                                                     |
| ---------------- | ------- | ---- | ------------- | ------------------------------------------------------------------------ |
| --schema       | string  | 否   |               | 资产元数据 [JSON Schema](https://JSON-Schema.org/) 规范 |

### 发行资产示例

```bash
irita tx nft issue security --schema='{"type":"object","properties":{"name":{"type":"string"}}}' --from=node0 --chain-id=irita-test -b=block -o=json --indent -y --home=testnet/node0/iritacli
```

结果

```json
{
  "height": "87",
  "txhash": "833DFD23566B67DFE9F81FFFB1C6F58173F3027CA1BC84D3AAFC6C51E9B34AC8",
  "raw_log": "<raw-log>",
  "logs": [
    {
      "msg_index": 0,
      "log": "",
      "events": [
        {
          "type": "issue_denom",
          "attributes": [
            {
              "key": "denom",
              "value": "security"
            }
          ]
        },
        {
          "type": "message",
          "attributes": [
            {
              "key": "action",
              "value": "issue_denom"
            },
            {
              "key": "module",
              "value": "nft"
            },
            {
              "key": "sender",
              "value": "iaa1w9g6g2692y973597x5euw9dfwm53w8tya4zkyn"
            }
          ]
        }
      ]
    }
  ],
  "gas_wanted": "200000",
  "gas_used": "43198"
}
```

## mint

创建指定类别的具体资产。

```bash
irita tx nft mint [denom] [tokenID] [flags]
```

**参数：**

| 名称      | 类型    | 必须 | 默认          | 描述                                                                     |
| ---------------- | ------- | ---- | ------------- | ------------------------------------------------------------------------ |
| denom  | string  | 是   |             | 资产的类别；长度为3到64，字母数字字符，以字母开始 |
| tokenID  | string  | 是   |             | 资产的唯一 ID，如 `UUID` |

**标志：**

| 名称，速记   | 类型   | 必须 | 默认  | 描述                                                  |
| ------------ | ------ | ---- | ----- | ----------------------------------------------------- |
| --token-uri | string | 否   | | 资产元数据的 `URI` |
| --token-data | string | 否 | | 资产元数据 |
| --recipient | string | 否   | | 资产接收者地址，默认为交易发起者地址 |

### 创建资产示例

```bash
irita tx nft mint security a4c74c4203af41619d00bb3e2f462c10 --token-uri=http://metadata.io/a4c74c4203af41619d00bb3e2f462c10 --token-data='{"name":"test security"}' --from=node0 --chain-id=irita-test -b=block -o=json --indent -y --home=testnet/node0/iritacli
```

结果

```json
{
  "height": "105",
  "txhash": "91713F006A50E66036B82AAEA7109244C8D67D0B9BCB475DB24CE2444B1E1445",
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
              "value": "mint_nft"
            },
            {
              "key": "module",
              "value": "nft"
            },
            {
              "key": "sender",
              "value": "iaa1w9g6g2692y973597x5euw9dfwm53w8tya4zkyn"
            }
          ]
        },
        {
          "type": "mint_nft",
          "attributes": [
            {
              "key": "recipient",
              "value": "iaa1w9g6g2692y973597x5euw9dfwm53w8tya4zkyn"
            },
            {
              "key": "denom",
              "value": "security"
            },
            {
              "key": "token-id",
              "value": "a4c74c4203af41619d00bb3e2f462c10"
            },
            {
              "key": "token-uri",
              "value": "http://metadata.io/a4c74c4203af41619d00bb3e2f462c10"
            }
          ]
        }
      ]
    }
  ],
  "gas_wanted": "200000",
  "gas_used": "52368"
}
```

## edit

编辑指定的资产。可更新的属性包括：资产元数据、元数据 `URI`

```bash
irita tx nft edit [denom] [tokenID] [flags]
```

**参数：**

| 名称      | 类型    | 必须 | 默认          | 描述                                                                     |
| ---------------- | ------- | ---- | ------------- | ------------------------------------------------------------------------ |
| denom  | string  | 是   |             | 资产的类别；长度为3到64，字母数字字符，以字母开始 |
| tokenID  | string  | 是   |             | 资产的唯一 ID |

**标志：**

| 名称，速记 | 类型   | 必须 | 默认 | 描述                                       |
| ---------- | ------ | ---- | ---- | ------------------------------------------ |
| --token-uri | string | 否   | | 资产元数据的 `URI` |
| --token-data | string | 否 | | 资产元数据 |

### 编辑资产示例

```bash
irita tx nft edit security a4c74c4203af41619d00bb3e2f462c10 --token-data='{"name":"new test security"}' --from=node0 --chain-id=irita-test -b=block -o=json --indent -y --home=testnet/node0/iritacli
```

结果

```json
{
  "height": "132",
  "txhash": "84AF2212C4651007F3838096032BA96C44C1CEBA245D3E8425D326417D0AB719",
  "raw_log": "<raw-log>",
  "logs": [
    {
      "msg_index": 0,
      "log": "",
      "events": [
        {
          "type": "edit_nft",
          "attributes": [
            {
              "key": "denom",
              "value": "security"
            },
            {
              "key": "token-id",
              "value": "a4c74c4203af41619d00bb3e2f462c10"
            },
            {
              "key": "token-uri",
              "value": "[do-not-modify]"
            }
          ]
        },
        {
          "type": "message",
          "attributes": [
            {
              "key": "action",
              "value": "edit_nft"
            },
            {
              "key": "module",
              "value": "nft"
            },
            {
              "key": "sender",
              "value": "iaa1w9g6g2692y973597x5euw9dfwm53w8tya4zkyn"
            }
          ]
        }
      ]
    }
  ],
  "gas_wanted": "200000",
  "gas_used": "45183"
}
```

## transfer

转移指定资产。

```bash
irita tx nft transfer [recipient] [denom] [tokenID] [flags]
```

**参数：**

| 名称      | 类型    | 必须 | 默认          | 描述                                                                     |
| ---------------- | ------- | ---- | ------------- | ------------------------------------------------------------------------ |
| recipient  | string  | 是   |             | 积分的唯一标识符 |
| denom  | string  | 是   |             | 资产的类别；长度为3到64，字母数字字符，以字母开始 |
| tokenID  | string  | 是   |             | 资产的唯一 ID |

### 转移资产示例

```bash
irita tx nft transfer iaa1gjmj3r0h9krjm9sg4hjkkv5wnsy52xck80g2sf security a4c74c4203af41619d00bb3e2f462c10 --from=node0 --chain-id=irita-test -b=block -o=json --indent -y --home=testnet/node0/iritacli
```

结果

```json
{
  "height": "174",
  "txhash": "2881520417D4C9A83837D271F856E6C94086629D6EDE8EB397CD23A579D6AD0E",
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
              "value": "transfer_nft"
            },
            {
              "key": "module",
              "value": "nft"
            },
            {
              "key": "sender",
              "value": "iaa1w9g6g2692y973597x5euw9dfwm53w8tya4zkyn"
            }
          ]
        },
        {
          "type": "transfer_nft",
          "attributes": [
            {
              "key": "recipient",
              "value": "iaa1gjmj3r0h9krjm9sg4hjkkv5wnsy52xck80g2sf"
            },
            {
              "key": "denom",
              "value": "security"
            },
            {
              "key": "token-id",
              "value": "a4c74c4203af41619d00bb3e2f462c10"
            }
          ]
        }
      ]
    }
  ],
  "gas_wanted": "200000",
  "gas_used": "49197"
}
```

## burn

销毁指定资产。

```bash
irita tx nft burn [denom] [tokenID] [flags]
```

**参数：**

| 名称      | 类型    | 必须 | 默认          | 描述                                                                     |
| ---------------- | ------- | ---- | ------------- | ------------------------------------------------------------------------ |
| denom  | string  | 是   |             | 资产的类别；长度为3到64，字母数字字符，以字母开始 |
| tokenID  | string  | 是   |             | 资产的唯一 ID |

### 销毁资产示例

```bash
irita tx nft burn security a4c74c4203af41619d00bb3e2f462c10 --from=iaa1gjmj3r0h9krjm9sg4hjkkv5wnsy52xck80g2sf --chain-id=irita-test -b=block -o=json --indent -y --home=testnet/node0/iritacli
```

结果

```json
{
  "height": "490",
  "txhash": "84D5CF42DF73A8D72E687FE47C26D22B67084D269B7B825940C45A62B7138CC3",
  "raw_log": "<raw-log>",
  "logs": [
    {
      "msg_index": 0,
      "log": "",
      "events": [
        {
          "type": "burn_nft",
          "attributes": [
            {
              "key": "denom",
              "value": "security"
            },
            {
              "key": "token-id",
              "value": "a4c74c4203af41619d00bb3e2f462c10"
            }
          ]
        },
        {
          "type": "message",
          "attributes": [
            {
              "key": "action",
              "value": "burn_nft"
            },
            {
              "key": "module",
              "value": "nft"
            },
            {
              "key": "sender",
              "value": "iaa1gjmj3r0h9krjm9sg4hjkkv5wnsy52xck80g2sf"
            }
          ]
        }
      ]
    }
  ],
  "gas_wanted": "200000",
  "gas_used": "47453"
}
```

## token

查询指定类别和 `ID` 的资产。

```bash
irita query nft token [denom] [tokenID] [flags]
```

**参数：**

| 名称      | 类型    | 必须 | 默认          | 描述                                                                     |
| ---------------- | ------- | ---- | ------------- | ------------------------------------------------------------------------ |
| denom  | string  | 是   |             | 资产的类别；长度为3到64，字母数字字符，以字母开始 |
| tokenID  | string  | 是   |             | 资产的唯一 ID |

### 查询指定资产示例

```bash
irita query nft token security a4c74c4203af41619d00bb3e2f462c10 -o=json --indent --chain-id=irita-test
```

结果

```json
{
  "type": "irismod/nft/BaseNFT",
  "value": {
    "ID": "a4c74c4203af41619d00bb3e2f462c10",
    "owner": "iaa1pf0r9rhfyzdyw3ed2hk0kyzjfwz4tehwsynxvt",
    "tokenURI": "http://metadata.io/a4c74c4203af41619d00bb3e2f462c10",
    "token_data": "{\"name\":\"test security\"}"
  }
}
```

## denom

查询指定类别的资产信息。

```bash
irita query nft denom [denom] [flags]
```

**参数：**

| 名称      | 类型    | 必须 | 默认          | 描述                                                                     |
| ---------------- | ------- | ---- | ------------- | ------------------------------------------------------------------------ |
| denom  | string  | 是   |             | 资产的类别；长度为3到64，字母数字字符，以字母开始 |

### 查询指定类别的资产信息示例

```bash
irita query nft denom security -o=json --indent --chain-id=irita-test
```

结果

```json
{
  "name": "security",
  "schema": "{\"type\":\"object\",\"properties\":{\"name\":{\"type\":\"string\"}}}",
  "creator": "iaa1pf0r9rhfyzdyw3ed2hk0kyzjfwz4tehwsynxvt"
}
```

## denoms

查询所有类别的资产信息。

```bash
irita query nft denoms [flags]
```

### 查询所有类别的资产信息示例

```bash
irita query nft denoms -o=json --indent --chain-id=irita-test
```

结果

```json
[
  {
    "name": "security",
    "schema": "{\"type\":\"object\",\"properties\":{\"name\":{\"type\":\"string\"}}}",
    "creator": "iaa1pf0r9rhfyzdyw3ed2hk0kyzjfwz4tehwsynxvt"
  }
]
```

## supply

查询指定类别资产的总量。如 `owner` 被指定，则查询此 `owner` 所拥有的该类别资产的总量。 

```bash
irita query nft supply [denom] [flags]
```

**参数：**

| 名称      | 类型    | 必须 | 默认          | 描述                                                                     |
| ---------------- | ------- | ---- | ------------- | ------------------------------------------------------------------------ |
| denom  | string  | 是   |             | 资产的类别；长度为3到64，字母数字字符，以字母开始 |

**标志：**

| 名称，速记 | 类型   | 必须 | 默认 | 描述                                       |
| ---------- | ------ | ---- | ---- | ------------------------------------------ |
| --owner | string | 否   | | 资产所有者地址 |

### 查询指定类别的资产总量示例

```bash
irita query nft supply security --chain-id=irita-test
```

结果

```text
1
```

### 查询指定账户某类别资产的总量

```bash
irita query nft supply security --owner=iaa1pf0r9rhfyzdyw3ed2hk0kyzjfwz4tehwsynxvt --chain-id=irita-test
```

结果

```text
1
```

## owner

查询指定账户的资产列表。如提供 `denom`，则查询该账户指定 `denom` 的资产列表。

```bash
irita query nft owner [address] [flags]
```

**参数：**

| 名称      | 类型    | 必须 | 默认          | 描述                                                                     |
| ---------------- | ------- | ---- | ------------- | ------------------------------------------------------------------------ |
| address  | string  | 是   |             | 目标账户地址 |

**标志：**

| 名称，速记 | 类型   | 必须 | 默认 | 描述                                       |
| ---------- | ------ | ---- | ---- | ------------------------------------------ |
| --denom  | string  | 是   |             | 资产的类别；长度为3到64，字母数字字符，以字母开始 |

### 查询账户所有资产示例

```bash
irita query nft owner iaa1pf0r9rhfyzdyw3ed2hk0kyzjfwz4tehwsynxvt -o=json --indent --chain-id=irita-test
```

结果

```json
{
  "address": "iaa1pf0r9rhfyzdyw3ed2hk0kyzjfwz4tehwsynxvt",
  "id_collections": [
    {
      "denom": "security",
      "ids": [
        "a4c74c4203af41619d00bb3e2f462c10"
      ]
    }
  ]
}
```

### 查询账户指定类别的所有资产示例

```bash
irita query nft owner iaa1pf0r9rhfyzdyw3ed2hk0kyzjfwz4tehwsynxvt --denom=security -o=json --indent --chain-id=irita-test
```

结果

```json
{
  "address": "iaa1pf0r9rhfyzdyw3ed2hk0kyzjfwz4tehwsynxvt",
  "id_collections": [
    {
      "denom": "security",
      "ids": [
        "a4c74c4203af41619d00bb3e2f462c10"
      ]
    }
  ]
}
```

## collection

查询指定类别的所有资产。

```bash
irita query nft collection [denom] [flags]
```

**参数：**

| 名称      | 类型    | 必须 | 默认          | 描述                                                                     |
| ---------------- | ------- | ---- | ------------- | ------------------------------------------------------------------------ |
| denom  | string  | 是   |             | 资产的类别；长度为3到64，字母数字字符，以字母开始 |

### 查询指定类别的所有资产示例

```bash
irita query nft collection security -o=json --indent --chain-id=irita-test
```

结果

```json
{
  "denom": {
    "name": "security",
    "schema": "{\"type\":\"object\",\"properties\":{\"name\":{\"type\":\"string\"}}}",
    "creator": "iaa1pf0r9rhfyzdyw3ed2hk0kyzjfwz4tehwsynxvt"
  },
  "nfts": [
    {
      "type": "irismod/nft/BaseNFT",
      "value": {
        "ID": "a4c74c4203af41619d00bb3e2f462c10",
        "owner": "iaa1pf0r9rhfyzdyw3ed2hk0kyzjfwz4tehwsynxvt",
        "tokenURI": "http://metadata.io/a4c74c4203af41619d00bb3e2f462c10",
        "token_data": "{\"name\":\"test security\"}"
      }
    }
  ]
}
```
