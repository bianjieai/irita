<!--
order: 1
-->

# 权限管理

IRITA 管理员可以通过控制台进行链的管理操作，主要管理功能有：

- [权限管理](#权限管理)
  - [add-roles](#add-roles)
    - [增加权限示例](#增加权限示例)
  - [remove-roles](#remove-roles)
    - [移除权限示例](#移除权限示例)
  - [block-account](#block-account)
    - [加入黑名单示例](#加入黑名单示例)
  - [unblock-account](#unblock-account)
    - [移出黑名单示例](#移出黑名单示例)

## add-roles

IRITA 管理员可以为指定账户增加相应的操作权限。

```bash
irita tx admin add-roles [address] [roles] [flags]
```

**参数：**

| 名称      | 类型    | 必须 | 默认          | 描述                                                                     |
| ---------------- | ------- | ---- | ------------- | ------------------------------------------------------------------------ |
| address  | string  | 是   |             | 账户地址 |
| roles  | string  | 是   |             | 权限值，可用值包括：PermAdmin，BlacklistAdmin，NodeAdmin，ParamAdmin，PowerUser |

### 增加权限示例

```bash
irita tx admin add-roles iaa18up8anyjpal8rncm8rd4ukp5f7etga795gp33q PermAdmin --from=node0 --chain-id=irita-test -b=block -o=json --indent -y --home=testnet/node0/iritacli
```

结果

```json
{
  "height": "329",
  "txhash": "8104DB91B083EB18F64CA4CC418FC54D4453AEA691C578EE64A22128FF724EC3",
  "raw_log": "<raw-log>",
  "logs": [
    {
      "msg_index": 0,
      "log": "",
      "events": [
        {
          "type": "add_roles",
          "attributes": [
            {
              "key": "account",
              "value": "iaa18up8anyjpal8rncm8rd4ukp5f7etga795gp33q"
            }
          ]
        },
        {
          "type": "message",
          "attributes": [
            {
              "key": "action",
              "value": "add_roles"
            },
            {
              "key": "module",
              "value": "admin"
            },
            {
              "key": "sender",
              "value": "iaa12v3r0unp6nprp9zur0fn446n832dfag5l3w38e"
            }
          ]
        }
      ]
    }
  ],
  "gas_wanted": "200000",
  "gas_used": "42088"
}
```

## remove-roles

IRITA 管理员可以移除指定账户的操作权限。

```bash
irita tx admin remove-roles [address] [roles] [flags]
```

**参数：**

| 名称      | 类型    | 必须 | 默认          | 描述                                                                     |
| ---------------- | ------- | ---- | ------------- | ------------------------------------------------------------------------ |
| address  | string  | 是   |             | 账户地址 |
| roles  | string  | 是   |             | 权限值，可用值包括：PermAdmin，BlacklistAdmin，NodeAdmin，ParamAdmin，PowerUser |

### 移除权限示例

```bash
irita tx admin add-roles iaa18up8anyjpal8rncm8rd4ukp5f7etga795gp33q PermAdmin --from=node0 --chain-id=irita-test -b=block -o=json --indent -y --home=testnet/node0/iritacli
```

结果

```json
{
  "height": "365",
  "txhash": "F298C81F1760503E1E7699EBBE57F5B8230DA3093BDF481FD7BF146A2240D416",
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
              "value": "remove_roles"
            },
            {
              "key": "module",
              "value": "admin"
            },
            {
              "key": "sender",
              "value": "iaa12v3r0unp6nprp9zur0fn446n832dfag5l3w38e"
            }
          ]
        },
        {
          "type": "remove_roles",
          "attributes": [
            {
              "key": "account",
              "value": "iaa18up8anyjpal8rncm8rd4ukp5f7etga795gp33q"
            }
          ]
        }
      ]
    }
  ],
  "gas_wanted": "200000",
  "gas_used": "42052"
}
```

## block-account

IRITA 管理员可以将指定账户加入黑名单。

```bash
irita tx admin block-account [address] [flags]
```

**参数：**

| 名称      | 类型    | 必须 | 默认          | 描述                                                                     |
| ---------------- | ------- | ---- | ------------- | ------------------------------------------------------------------------ |
| address  | string  | 是   |             | 账户地址 |

### 加入黑名单示例

```bash
irita tx admin block-account iaa18up8anyjpal8rncm8rd4ukp5f7etga795gp33q --from=node0 --chain-id=irita-test -b=block -o=json --indent -y --home=testnet/node0/iritacli
```

结果

```json
{
  "height": "388",
  "txhash": "F266F1734DDE56A8C4A5676BD929C414F0E9C874131856AF362762E4474B489D",
  "raw_log": "<raw-log>",
  "logs": [
    {
      "msg_index": 0,
      "log": "",
      "events": [
        {
          "type": "block_account",
          "attributes": [
            {
              "key": "account",
              "value": "iaa18up8anyjpal8rncm8rd4ukp5f7etga795gp33q"
            }
          ]
        },
        {
          "type": "message",
          "attributes": [
            {
              "key": "action",
              "value": "block_account"
            },
            {
              "key": "module",
              "value": "admin"
            },
            {
              "key": "sender",
              "value": "iaa12v3r0unp6nprp9zur0fn446n832dfag5l3w38e"
            }
          ]
        }
      ]
    }
  ],
  "gas_wanted": "200000",
  "gas_used": "41058"
}
```

## unblock-account

IRITA 管理员可以将指定账户从黑名单移出。

```bash
irita tx admin unblock-account [address] [flags]
```

**参数：**

| 名称      | 类型    | 必须 | 默认          | 描述                                                                     |
| ---------------- | ------- | ---- | ------------- | ------------------------------------------------------------------------ |
| address  | string  | 是   |             | 账户地址 |

### 移出黑名单示例

```bash
irita tx admin unblock-account iaa18up8anyjpal8rncm8rd4ukp5f7etga795gp33q --from=node0 --chain-id=irita-test -b=block -o=json --indent -y --home=testnet/node0/iritacli
```

结果

```json
{
  "height": "403",
  "txhash": "92BE6F2E8867288208D6DC68F0F99F179A8B6AA93253CEC64142B96EA724AA0F",
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
              "value": "unblock_account"
            },
            {
              "key": "module",
              "value": "admin"
            },
            {
              "key": "sender",
              "value": "iaa12v3r0unp6nprp9zur0fn446n832dfag5l3w38e"
            }
          ]
        },
        {
          "type": "unblock_account",
          "attributes": [
            {
              "key": "account",
              "value": "iaa18up8anyjpal8rncm8rd4ukp5f7etga795gp33q"
            }
          ]
        }
      ]
    }
  ],
  "gas_wanted": "200000",
  "gas_used": "39004"
}
```
