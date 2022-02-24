<!--
order: 7
-->

# 资产批量数字化建模

## issue

发行资产。

```bash
irita tx mt issue [flags]
```

**标志：**

| 名称，速记       | 类型    | 必须 | 默认          | 描述                                                                     |
| ---------------- | ------- | ---- | ------------- | ------------------------------------------------------------------------ |
| --name        | string | 是 |  | 资产类别的名称；长度为3到64，字母数字字符，以字母开始 |
| --from        | string | 否 |  | 资产类别的发行者 |

### 发行资产示例

```bash
irita tx mt issue --name=validator-denom --from=validator --chain-id=irita-tesnet -b=block -o=json -y
```

结果

```json
{
  "height":"6079",
  "txhash":"843A99C892265243B070E5E164A366D69159734EFE2FBD82DAB6B78756EAC0B3",
  "codespace":"",
  "code":0,
  "data":"0A1B0A192F697269736D6F642E6D742E4D7367497373756544656E6F6D",
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
              "value":"ee084d58e26993894026d2a026f984b7bc07ea0ac813c56753ea293c112e5217"
            },
            {
              "key":"denom_name",
              "value":"validator-denom"
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
              "value":"/irismod.mt.MsgIssueDenom"
            },
            {
              "key":"module",
              "value":"mt"
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
  "gas_used":"58384"
}
```

## mint

创建指定类别的具体资产；可指定发行数量。

```bash
irita tx mt mint [denom-id] [flags]
```

**参数：**

| 名称      | 类型    | 必须 | 默认          | 描述                                                                     |
| ---------------- | ------- | ---- | ------------- | ------------------------------------------------------------------------ |
| denom-id  | string  | 是   |             | 资产类别的唯一 ID |

**标志：**

| 名称，速记   | 类型   | 必须 | 默认  | 描述                                                  |
| ------------ | ------ | ---- | ----- | ----------------------------------------------------- |
| --amount  | string | 否 | | 发行资产数量 |
| --data    | string | 否 | | 资产元数据 |
| --mt-id   | string | 否 | | 资产的唯一 ID；若不填写则是发行，若填写则是增发 |
| --from    | string | 是 | | 资产发起者地址；必须是 denom 的拥有者 |
| --recipient | string | 否 | | 资产接收者地址，默认为交易发起者地址 |

### 创建资产示例

```bash
irita tx mt mint ee084d58e26993894026d2a026f984b7bc07ea0ac813c56753ea293c112e5217 --amount=100 --data='{"name":"test security"}' --from=validator --chain-id=irita-tesnet -b=block -o=json -y
```

结果

```json
{
  "height":"6232",
  "txhash":"265B93B8F110868BCB7E6AD273418E2CD3498FD67D3353F97DFE5D8453D6C0B9",
  "codespace":"",
  "code":0,
  "data":"0A170A152F697269736D6F642E6D742E4D73674D696E744D54",
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
              "value":"/irismod.mt.MsgMintMT"
            },
            {
              "key":"module",
              "value":"mt"
            },
            {
              "key":"sender",
              "value":"iaa17y3qs2zuanr93nk844x0t7e6ktchwygnc8fr0g"
            }
          ]
        },
        {
          "type":"mint_mt",
          "attributes":[
            {
              "key":"mt_id",
              "value":"dc1d1f5a54cfdc4ed5b9ca90ad09f5b8b9bfa3b78a94b64ce51d4d77c6c212f3"
            },
            {
              "key":"denom_id",
              "value":"ee084d58e26993894026d2a026f984b7bc07ea0ac813c56753ea293c112e5217"
            },
            {
              "key":"amount",
              "value":"100"
            },
            {
              "key":"supply",
              "value":"100"
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
  "gas_used":"71443"
}
```

## edit

编辑指定的资产。可更新的属性包括：资产元数据。

```bash
irita tx mt edit [denom-id] [mt-id] [flags]
```

**参数：**

| 名称      | 类型    | 必须 | 默认          | 描述                                                                     |
| ---------------- | ------- | ---- | ------------- | ------------------------------------------------------------------------ |
| denom-id  | string  | 是   |             | 资产类别的唯一 ID  |
| mt-id     | string  | 是   |             | 资产的唯一 ID |

**标志：**

| 名称，速记 | 类型   | 必须 | 默认 | 描述                                       |
| ---------- | ------ | ---- | ---- | ------------------------------------------ |
| --data | string | 否 | | 资产的元数据 |
| --from | string | 否 | | 资产拥有者地址 |

### 编辑资产示例

```bash
irita tx mt edit ee084d58e26993894026d2a026f984b7bc07ea0ac813c56753ea293c112e5217 dc1d1f5a54cfdc4ed5b9ca90ad09f5b8b9bfa3b78a94b64ce51d4d77c6c212f3 --data='{"name":"test edit"}' --from=validator --chain-id=irita-tesnet -b=block -o=json -y
```

结果

```json
{
  "height":"6400",
  "txhash":"F6CC3019F1E6096960705283E10C2B170A1D06381CC8563DE8B48728EB90A7CD",
  "codespace":"",
  "code":0,
  "data":"0A170A152F697269736D6F642E6D742E4D7367456469744D54",
  "raw_log":"<raw-log>",
  "logs":[
    {
      "msg_index":0,
      "log":"",
      "events":[
        {
          "type":"edit_mt",
          "attributes":[
            {
              "key":"mt_id",
              "value":"dc1d1f5a54cfdc4ed5b9ca90ad09f5b8b9bfa3b78a94b64ce51d4d77c6c212f3"
            },
            {
              "key":"denom_id",
              "value":"ee084d58e26993894026d2a026f984b7bc07ea0ac813c56753ea293c112e5217"
            }
          ]
        },
        {
          "type":"message",
          "attributes":[
            {
              "key":"action",
              "value":"/irismod.mt.MsgEditMT"
            },
            {
              "key":"module",
              "value":"mt"
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
  "gas_used":"59039"
}
```

## transfer

转移指定资产；可指定转移数量。

```bash
irita tx mt transfer [from_key_or_address] [recipient] [denom-id] [mt-id] [amount] [flags]
```

**参数：**

| 名称      | 类型    | 必须 | 默认          | 描述                                                                     |
| ---------------- | ------- | ---- | ------------- | ------------------------------------------------------------------------ |
| from_key_or_address   | string | 是 |             | 资产发起者地址；必须拥有此资产 |
| recipient             | string | 是 |             | 资产接收者地址 |
| denom-id              | string | 是 |             | 资产类别的唯一 ID |
| mt-id                 | string | 是 |             | 资产的唯一 ID |
| amount                | string | 是 |             | 转移资产数量 |

### 转移资产示例

```bash
irita tx mt transfer validator iaa1pjprrg6xy0gkck94msu04j4q36m9wku70v6kfm ee084d58e26993894026d2a026f984b7bc07ea0ac813c56753ea293c112e5217 dc1d1f5a54cfdc4ed5b9ca90ad09f5b8b9bfa3b78a94b64ce51d4d77c6c212f3 50 --chain-id=irita-tesnet -b=block -o=json -y
```

结果

```json
{
  "height":"6496",
  "txhash":"20132FD609E3E4ECD27E33B36C26B0456D8C6850FF85DF5DC097D7B1DDC18ACA",
  "codespace":"",
  "code":0,
  "data":"0A1B0A192F697269736D6F642E6D742E4D73675472616E736665724D54",
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
              "value":"/irismod.mt.MsgTransferMT"
            },
            {
              "key":"module",
              "value":"mt"
            },
            {
              "key":"sender",
              "value":"iaa17y3qs2zuanr93nk844x0t7e6ktchwygnc8fr0g"
            }
          ]
        },
        {
          "type":"transfer_mt",
          "attributes":[
            {
              "key":"mt_id",
              "value":"dc1d1f5a54cfdc4ed5b9ca90ad09f5b8b9bfa3b78a94b64ce51d4d77c6c212f3"
            },
            {
              "key":"denom_id",
              "value":"ee084d58e26993894026d2a026f984b7bc07ea0ac813c56753ea293c112e5217"
            },
            {
              "key":"amount",
              "value":"50"
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
  "gas_used":"58082"
}
```

## burn

销毁指定资产；可指定销毁数量。

```bash
irita tx mt burn [denom-id] [mt-id] [amount] [flags]
```

**参数：**

| 名称      | 类型    | 必须 | 默认          | 描述                                                                     |
| ---------------- | ------- | ---- | ------------- | ------------------------------------------------------------------------ |
| denom-id  | string  | 是  |             | 资产类别的唯一 ID |
| mt-id     | string  | 是  |             | 资产的唯一 ID |
| amount    | string  | 是  |             | 销毁资产的数量 |

**标志：**

| 名称，速记 | 类型   | 必须 | 默认 | 描述                                       |
| ---------- | ------ | ---- | ---- | ------------------------------------------ |
| --from | string | 否 | | 资产拥有者地址 |

### 销毁资产示例

```bash
irita tx mt burn ee084d58e26993894026d2a026f984b7bc07ea0ac813c56753ea293c112e5217 dc1d1f5a54cfdc4ed5b9ca90ad09f5b8b9bfa3b78a94b64ce51d4d77c6c212f3 20 --from=iaa17y3qs2zuanr93nk844x0t7e6ktchwygnc8fr0g --chain-id=irita-tesnet -b=block -o=json -y
```

结果

```json
{
  "height":"6563",
  "txhash":"0EA1F45987AC5B88785EC4D94C8D6269C92F6A0C5A239465966CD667263335B1",
  "codespace":"",
  "code":0,
  "data":"0A170A152F697269736D6F642E6D742E4D73674275726E4D54",
  "raw_log":"<raw-log>",
  "logs":[
    {
      "msg_index":0,
      "log":"",
      "events":[
        {
          "type":"burn_mt",
          "attributes":[
            {
              "key":"mt_id",
              "value":"dc1d1f5a54cfdc4ed5b9ca90ad09f5b8b9bfa3b78a94b64ce51d4d77c6c212f3"
            },
            {
              "key":"denom_id",
              "value":"ee084d58e26993894026d2a026f984b7bc07ea0ac813c56753ea293c112e5217"
            },
            {
              "key":"amount",
              "value":"20"
            }
          ]
        },
        {
          "type":"message",
          "attributes":[
            {
              "key":"action",
              "value":"/irismod.mt.MsgBurnMT"
            },
            {
              "key":"module",
              "value":"mt"
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
  "gas_used":"57608"
}
```

## denom

根据 `DenomID` 查询资产类别信息。

```bash
irita query mt denom [denom-id] [flags]
```

**参数：**

| 名称      | 类型    | 必须 | 默认          | 描述                                                                     |
| ---------------- | ------- | ---- | ------------- | ------------------------------------------------------------------------ |
| denom-id  | string  | 是  |             | 资产类别的唯一 ID |

### 查询指定资产示例

```bash
irita query mt denom ee084d58e26993894026d2a026f984b7bc07ea0ac813c56753ea293c112e5217 -o=json
```

结果

```json
{
  "id":"ee084d58e26993894026d2a026f984b7bc07ea0ac813c56753ea293c112e5217",
  "name":"validator-denom",
  "data":null,
  "owner":"iaa17y3qs2zuanr93nk844x0t7e6ktchwygnc8fr0g"
}
```

## denoms

查询所有类别的资产信息。

```bash
irita query mt denoms [flags]
```

### 查询所有类别的资产示例

```bash
irita query mt denoms -o=json
```

结果

```json
{
  "denoms":[
    {
      "id":"068b196f26c91d2d17f9c31440572686309cf9558358c91a3c81fbf801e3aa29",
      "name":"denom3",
      "data":null,
      "owner":"iaa17y3qs2zuanr93nk844x0t7e6ktchwygnc8fr0g"
    },
    {
      "id":"ee084d58e26993894026d2a026f984b7bc07ea0ac813c56753ea293c112e5217",
      "name":"validator-denom",
      "data":null,
      "owner":"iaa17y3qs2zuanr93nk844x0t7e6ktchwygnc8fr0g"
    }
  ]
}
```

## supply

根据 `DenomID` 和 `MtID` 查询资产总量。

```bash
irita query mt supply [denom-id] [mt-id] [flags]
```

**参数：**

| 名称      | 类型    | 必须 | 默认          | 描述                                                                     |
| ---------------- | ------- | ---- | ------------- | ------------------------------------------------------------------------ |
| denom-id  | string  | 是  |             | 资产类别的唯一 ID |
| mt-id     | string  | 是  |             | 资产的唯一 ID |

### 查询指定类别的资产总量示例

```bash
irita query mt supply ee084d58e26993894026d2a026f984b7bc07ea0ac813c56753ea293c112e5217 dc1d1f5a54cfdc4ed5b9ca90ad09f5b8b9bfa3b78a94b64ce51d4d77c6c212f3 -o=json
```

结果

```text
{"amount":"80"}
```

## balances

查询指定账户某类别资产的总量。

```bash
irita query mt balances [owner] [denom-id] [flags]
```

**参数：**

| 名称      | 类型    | 必须 | 默认          | 描述                                                                     |
| ---------------- | ------- | ---- | ------------- | ------------------------------------------------------------------------ |
| owner     | string  | 是 |             | 资产拥有者地址 |
| denom-id  | string  | 是 |             | 资产类别的唯一 ID |

### 查询指定账户某类别资产的总量示例

```bash
irita q mt balances iaa17y3qs2zuanr93nk844x0t7e6ktchwygnc8fr0g ee084d58e26993894026d2a026f984b7bc07ea0ac813c56753ea293c112e5217 -o=json
```

结果

```json
{
  "balance":[
    {
      "mt_id":"dc1d1f5a54cfdc4ed5b9ca90ad09f5b8b9bfa3b78a94b64ce51d4d77c6c212f3",
      "amount":"30"
    }
  ]
}
```

## token

根据 `DenomID` 以及 `MtID` 查询具体资产信息。

```bash
irita query mt token [denom-id] [mt-id] [flags]
```

**参数：**

| 名称      | 类型    | 必须 | 默认          | 描述                                                                     |
| ---------------- | ------- | ---- | ------------- | ------------------------------------------------------------------------ |
| denom-id  | string  | 是  |             | 资产类别的唯一 ID |
| mt-id     | string  | 是  |             | 资产的唯一 ID |

### 查询指定资产示例

```bash
irita query mt token ee084d58e26993894026d2a026f984b7bc07ea0ac813c56753ea293c112e5217 dc1d1f5a54cfdc4ed5b9ca90ad09f5b8b9bfa3b78a94b64ce51d4d77c6c212f3 -o=json
```

结果

```json
{
  "id":"dc1d1f5a54cfdc4ed5b9ca90ad09f5b8b9bfa3b78a94b64ce51d4d77c6c212f3",
  "supply":"80",
  "data":"eyJuYW1lIjoidGVzdCBlZGl0In0="
}
```

## tokens

根据 `DenomID` 查询所有资产信息。

```bash
irita query mt tokens [denom-id] [flags]
```

**参数：**

| 名称      | 类型    | 必须 | 默认          | 描述                                                                     |
| ---------------- | ------- | ---- | ------------- | ------------------------------------------------------------------------ |
| denom-id  | string  | 是  |             | 资产类别的唯一 ID |

### 查询所有资产示例

```bash
irita query mt tokens ee084d58e26993894026d2a026f984b7bc07ea0ac813c56753ea293c112e5217 -o=json
```

结果

```json
{
  "mts":[
    {
      "id":"dc1d1f5a54cfdc4ed5b9ca90ad09f5b8b9bfa3b78a94b64ce51d4d77c6c212f3",
      "supply":"80",
      "data":"eyJuYW1lIjoidGVzdCBlZGl0In0="
    }
  ]
}
```
