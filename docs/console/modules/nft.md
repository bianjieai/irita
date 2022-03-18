<!--
order: 2
-->

# 资产数字化建模

## issue

发行资产。

```bash
irita tx nft issue [denom-id] [flags]
```

**参数：**

| 名称      | 类型    | 必须 | 默认          | 描述                                                                     |
| ---------------- | ------- | ---- | ------------- | ------------------------------------------------------------------------ |
| denom-id  | string  | 是   |             | 资产的类别，全局唯一；长度为3到64，字母数字字符，以字母开始 |

**标志：**

| 名称，速记       | 类型    | 必须 | 默认          | 描述                                                                     |
| ---------------- | ------- | ---- | ------------- | ------------------------------------------------------------------------ |
| --schema              | string  | 否   |               | 资产元数据 [JSON Schema](https://JSON-Schema.org/) 规范 |
| --mint-restricted     | bool    | 是   |               | 发行受限 |
| --update-restricted   | bool    | 是   |               | 更新受限 |

### 发行资产示例

```bash
irita tx nft issue nftdenom --schema='{"type":"object","properties":{"name":{"type":"string"}}}' --mint-restricted=false --update-restricted=false --from=validator --chain-id=irita-test -b=block -o=json -y
```

结果

```json
{
  "height":"11892",
  "txhash":"9AAA1057A8439ECC2B6E0D47CF353CA9AC8296E88AFEA55A51638140AF317115",
  "codespace":"",
  "code":0,
  "data":"0A1C0A1A2F697269736D6F642E6E66742E4D7367497373756544656E6F6D",
  "raw_log":"<raw-log>",
  "logs":[
    {
      "msg_index":0,
      "log":"",
      "events":[
        {
          "type":"issue_denom",
          "attributes":[
            {
              "key":"denom_id",
              "value":"nftdenom"
            },
            {
              "key":"denom_name",
              "value":""
            },
            {
              "key":"creator",
              "value":"iaa17y3qs2zuanr93nk844x0t7e6ktchwygnc8fr0g"
            }
          ]
        },
        {
          "type":"message",
          "attributes":[
            {
              "key":"action",
              "value":"/irismod.nft.MsgIssueDenom"
            },
            {
              "key":"module",
              "value":"nft"
            },
            {
              "key":"sender",
              "value":"iaa17y3qs2zuanr93nk844x0t7e6ktchwygnc8fr0g"
            }
          ]
        }
      ]
    }
  ],
  "info":"",
  "gas_wanted":"200000",
  "gas_used":"58650"
}
```

## mint

创建指定类别的具体资产。

```bash
irita tx nft mint [denom-id] [nft-id] [flags]
```

**参数：**

| 名称      | 类型    | 必须 | 默认          | 描述                                                                     |
| ---------------- | ------- | ---- | ------------- | ------------------------------------------------------------------------ |
| denom-id  | string  | 是   |             | 资产的类别；长度为3到64，字母数字字符，以字母开始 |
| nft-id    | string  | 是   |             | 资产的唯一 ID，如 `UUID` |

**标志：**

| 名称，速记   | 类型   | 必须 | 默认  | 描述                                                  |
| ------------ | ------ | ---- | ----- | ----------------------------------------------------- |
| --uri       | string | 否   | | 资产元数据的 `URI` |
| --uri-hash  | string | 否   | | 资产 `URI` 的哈希 |
| --data      | string | 否   | | 资产元数据 |
| --recipient | string | 否   | | 资产接收者地址，默认为交易发起者地址 |

### 创建资产示例

```bash
irita tx nft mint nftdenom nft1 --uri=https://metadata.io/a4c74c4203af41619d00bb3e2f462c10 --data='{"name":"test nftdenom"}' --from=validator --chain-id=irita-test -b=block -o=json -y
```

结果

```json
{
  "height":"12879",
  "txhash":"6C4986952ADC3E6F02EC7AFF8F9550A08B00BBEF837FEC6F5EB94BE443F413E7",
  "codespace":"",
  "code":0,
  "data":"0A190A172F697269736D6F642E6E66742E4D73674D696E744E4654",
  "raw_log":"<raw-log>",
  "logs":[
    {
      "msg_index":0,
      "log":"",
      "events":[
        {
          "type":"message",
          "attributes":[
            {
              "key":"action",
              "value":"/irismod.nft.MsgMintNFT"
            },
            {
              "key":"module",
              "value":"nft"
            },
            {
              "key":"sender",
              "value":"iaa17y3qs2zuanr93nk844x0t7e6ktchwygnc8fr0g"
            }
          ]
        },
        {
          "type":"mint_nft",
          "attributes":[
            {
              "key":"token_id",
              "value":"nft1"
            },
            {
              "key":"denom_id",
              "value":"nftdenom"
            },
            {
              "key":"token_uri",
              "value":"http://metadata.io/a4c74c4203af41619d00bb3e2f462c10"
            },
            {
              "key":"recipient",
              "value":"iaa17y3qs2zuanr93nk844x0t7e6ktchwygnc8fr0g"
            }
          ]
        }
      ]
    }
  ],
  "info":"",
  "gas_wanted":"200000",
  "gas_used":"64001"
}
```

## edit

编辑指定的资产。可更新的属性包括：资产元数据、元数据 `URI`、`URI` 的哈希

```bash
irita tx nft edit [denom-id] [nft-id] [flags]
```

**参数：**

| 名称      | 类型    | 必须 | 默认          | 描述                                                                     |
| ---------------- | ------- | ---- | ------------- | ------------------------------------------------------------------------ |
| denom-id  | string  | 是   |             | 资产的类别；长度为3到64，字母数字字符，以字母开始 |
| nft-id    | string  | 是   |             | 资产的唯一 ID |

**标志：**

| 名称，速记 | 类型   | 必须 | 默认 | 描述                                       |
| ---------- | ------ | ---- | ---- | ------------------------------------------ |
| --uri       | string | 否   | | 资产元数据的 `URI` |
| --uri-hash  | string | 否   | | 资产 `URI` 的哈希 |
| --data      | string | 否   | | 资产元数据 |

### 编辑资产示例

```bash
irita tx nft edit nftdenom nft1 --uri=https://metadata.io/nft1 --data='{"name":"new test nftdenom"}' --from=validator --chain-id=irita-test -b=block -o=json -y
```

结果

```json
{
  "height":"13242",
  "txhash":"82155D946C9616CB1C3DB2E6CD24BE82DBEFF413B6B3F778FD0327A843A5ACBD",
  "codespace":"",
  "code":0,
  "data":"0A190A172F697269736D6F642E6E66742E4D7367456469744E4654",
  "raw_log":"<raw-log>",
  "logs":[
    {
      "msg_index":0,
      "log":"",
      "events":[
        {
          "type":"edit_nft",
          "attributes":[
            {
              "key":"token_id",
              "value":"nft1"
            },
            {
              "key":"denom_id",
              "value":"nftdenom"
            },
            {
              "key":"token_uri",
              "value":"https://metadata.io/nft1"
            },
            {
              "key":"owner",
              "value":"iaa17y3qs2zuanr93nk844x0t7e6ktchwygnc8fr0g"
            }
          ]
        },
        {
          "type":"message",
          "attributes":[
            {
              "key":"action",
              "value":"/irismod.nft.MsgEditNFT"
            },
            {
              "key":"module",
              "value":"nft"
            },
            {
              "key":"sender",
              "value":"iaa17y3qs2zuanr93nk844x0t7e6ktchwygnc8fr0g"
            }
          ]
        }
      ]
    }
  ],
  "info":"",
  "gas_wanted":"200000",
  "gas_used":"58132"
}
```

## transfer

转移指定资产。

```bash
irita tx nft transfer [recipient] [denom-id] [nft-id] [flags]
```

**参数：**

| 名称      | 类型    | 必须 | 默认          | 描述                                                                     |
| ---------------- | ------- | ---- | ------------- | ------------------------------------------------------------------------ |
| recipient  | string  | 是   |             | 积分的唯一标识符 |
| denom-id   | string  | 是   |             | 资产的类别；长度为3到64，字母数字字符，以字母开始 |
| nft-id     | string  | 是   |             | 资产的唯一 ID |

### 转移资产示例

```bash
irita tx nft transfer iaa1pjprrg6xy0gkck94msu04j4q36m9wku70v6kfm nftdenom nft1 --from=validator --chain-id=irita-test -b=block -o=json -y
```

结果

```json
{
  "height":"13276",
  "txhash":"42070121B1BC2ACD627E233F274172E07FFC820690D9393281CD506291EFAE29",
  "codespace":"",
  "code":0,
  "data":"0A1D0A1B2F697269736D6F642E6E66742E4D73675472616E736665724E4654",
  "raw_log":"<raw-log>",
  "logs":[
    {
      "msg_index":0,
      "log":"",
      "events":[
        {
          "type":"message",
          "attributes":[
            {
              "key":"action",
              "value":"/irismod.nft.MsgTransferNFT"
            },
            {
              "key":"module",
              "value":"nft"
            },
            {
              "key":"sender",
              "value":"iaa17y3qs2zuanr93nk844x0t7e6ktchwygnc8fr0g"
            }
          ]
        },
        {
          "type":"transfer_nft",
          "attributes":[
            {
              "key":"token_id",
              "value":"nft1"
            },
            {
              "key":"denom_id",
              "value":"nftdenom"
            },
            {
              "key":"sender",
              "value":"iaa17y3qs2zuanr93nk844x0t7e6ktchwygnc8fr0g"
            },
            {
              "key":"recipient",
              "value":"iaa1pjprrg6xy0gkck94msu04j4q36m9wku70v6kfm"
            }
          ]
        }
      ]
    }
  ],
  "info":"",
  "gas_wanted":"200000",
  "gas_used":"61500"
}
```

## burn

销毁指定资产。

```bash
irita tx nft burn [denom-id] [nft-id] [flags]
```

**参数：**

| 名称      | 类型    | 必须 | 默认          | 描述                                                                     |
| ---------------- | ------- | ---- | ------------- | ------------------------------------------------------------------------ |
| denom-id  | string  | 是   |             | 资产的类别；长度为3到64，字母数字字符，以字母开始 |
| nft-id    | string  | 是   |             | 资产的唯一 ID |

### 销毁资产示例

```bash
irita tx nft burn nftdenom nft1 --from=iaa17y3qs2zuanr93nk844x0t7e6ktchwygnc8fr0g --chain-id=irita-test -b=block -o=json --indent -y
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
              "value": "nftdenom"
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
irita query nft token [denom-id] [nft-id] [flags]
```

**参数：**

| 名称      | 类型    | 必须 | 默认          | 描述                                                                     |
| ---------------- | ------- | ---- | ------------- | ------------------------------------------------------------------------ |
| denom-id  | string  | 是   |             | 资产的类别；长度为3到64，字母数字字符，以字母开始 |
| nft-id    | string  | 是   |             | 资产的唯一 ID |

### 查询指定资产示例

```bash
irita query nft token nftdenom nft1 --chain-id=irita-test -o=json
```

结果

```json
{
  "type": "irismod/nft/BaseNFT",
  "value": {
    "id":"nft1",
    "name":"",
    "uri":"https://metadata.io/nft1",
    "data":"{\"name\":\"new test nftdenom\"}",
    "owner":"iaa1pjprrg6xy0gkck94msu04j4q36m9wku70v6kfm",
    "uri_hash":""
  }
}
```

## denom

查询指定类别的资产信息。

```bash
irita query nft denom [denom-id] [flags]
```

**参数：**

| 名称      | 类型    | 必须 | 默认          | 描述                                                                     |
| ---------------- | ------- | ---- | ------------- | ------------------------------------------------------------------------ |
| denom-id  | string  | 是   |             | 资产的类别；长度为3到64，字母数字字符，以字母开始 |

### 查询指定类别的资产信息示例

```bash
irita query nft denom nftdenom --chain-id=irita-test -o=json
```

结果

```json
{
  "id":"nftdenom",
  "name":"",
  "schema":"{\"type\":\"object\",\"properties\":{\"name\":{\"type\":\"string\"}}}",
  "creator":"iaa17y3qs2zuanr93nk844x0t7e6ktchwygnc8fr0g",
  "symbol":"",
  "mint_restricted":false,
  "update_restricted":false,
  "description":"",
  "uri":"",
  "uri_hash":"",
  "data":""
}
```

## denoms

查询所有类别的资产信息。

```bash
irita query nft denoms [flags]
```

### 查询所有类别的资产信息示例

```bash
irita query nft denoms --chain-id=irita-test -o=json
```

结果

```json
[
  {
    "id":"nftdenom1",
    "name":"",
    "schema":"",
    "creator":"iaa17y3qs2zuanr93nk844x0t7e6ktchwygnc8fr0g",
    "symbol":"",
    "mint_restricted":true,
    "update_restricted":true,
    "description":"",
    "uri":"",
    "uri_hash":"",
    "data":""
  },
  {
    "id":"nftdenom",
    "name":"",
    "schema":"{\"type\":\"object\",\"properties\":{\"name\":{\"type\":\"string\"}}}",
    "creator":"iaa17y3qs2zuanr93nk844x0t7e6ktchwygnc8fr0g",
    "symbol":"",
    "mint_restricted":false,
    "update_restricted":false,
    "description":"",
    "uri":"",
    "uri_hash":"",
    "data":""
  }
]
```

## supply

查询指定类别资产的总量。如 `owner` 被指定，则查询此 `owner` 所拥有的该类别资产的总量。 

```bash
irita query nft supply [denom-id] [flags]
```

**参数：**

| 名称      | 类型    | 必须 | 默认          | 描述                                                                     |
| ---------------- | ------- | ---- | ------------- | ------------------------------------------------------------------------ |
| denom-id  | string  | 是   |             | 资产的类别；长度为3到64，字母数字字符，以字母开始 |

**标志：**

| 名称，速记 | 类型   | 必须 | 默认 | 描述                                       |
| ---------- | ------ | ---- | ---- | ------------------------------------------ |
| --owner | string | 否   | | 资产所有者地址 |

### 查询指定类别的资产总量示例

```bash
irita query nft supply nftdenom
```

结果

```text
1
```

### 查询指定账户某类别资产的总量

```bash
irita query nft supply nftdenom --owner=iaa17y3qs2zuanr93nk844x0t7e6ktchwygnc8fr0g
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
      "denom": "nftdenom",
      "ids": [
        "a4c74c4203af41619d00bb3e2f462c10"
      ]
    }
  ]
}
```

### 查询账户指定类别的所有资产示例

```bash
irita query nft owner iaa1pf0r9rhfyzdyw3ed2hk0kyzjfwz4tehwsynxvt --denom=nftdenom -o=json --indent --chain-id=irita-test
```

结果

```json
{
  "address": "iaa1pf0r9rhfyzdyw3ed2hk0kyzjfwz4tehwsynxvt",
  "id_collections": [
    {
      "denom": "nftdenom",
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
irita query nft collection [denom-id] [flags]
```

**参数：**

| 名称      | 类型    | 必须 | 默认          | 描述                                                                     |
| ---------------- | ------- | ---- | ------------- | ------------------------------------------------------------------------ |
| denom-id  | string  | 是   |             | 资产的类别；长度为3到64，字母数字字符，以字母开始 |

### 查询指定类别的所有资产示例

```bash
irita query nft collection nftdenom -o=json --indent --chain-id=irita-test
```

结果

```json
{
  "denom": {
    "name": "nftdenom",
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
        "token_data": "{\"name\":\"test nftdenom\"}"
      }
    }
  ]
}
```
