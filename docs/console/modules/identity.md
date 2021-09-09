<!--
order: 6
-->

# 身份

## create

创建一个新的身份。

```bash
irita tx identity create [flags]
```

**标志：**

| 名称，速记     | 类型   | 必须 | 默认 | 描述                                                   |
| -------------- | ------ | ---- | ---- | ------------------------------------------------------ |
| --id  | string | 否   |      | 身份唯一 ID，长度为32的 Hex 字符串；如不提供，则自动生成                   |
| --pubkey       | string | 否   |      | 身份主体的公钥, Hex 字符串；"rsa" 与 "dsa" 公钥采用 DER 编码，"ecdsa"、"ed25519" 与 "sm2" 公钥采用压缩形式                              |
| --pubkey-algo | string | 否   |      | 公钥算法，支持 "rsa"、"dsa"、"ecdsa"、"ed25519"、"sm2" |
| --cert-file    | string | 否   |      | 公钥证书文件路径；仅支持 PEM 编码的 X.509 证书         |
| --credentials  | string | 否   |      | 身份主体凭证信息的 URI，最大长度为140                    |

### 创建身份示例

```bash
irita tx identity create --pubkey=03281ce4ba0b8c97e5b1434f8f298b064f03d4c1d21aae9276065e170fc90a5d51 --pubkey-algo=sm2 --credentials=https://security.com/kyc/10001/ --from=node0 --chain-id=irita-test -b=block -o=json --indent -y --home=testnet/node0/iritacli
```

结果

```json
{
  "height": "4145",
  "txhash": "9C0FF80B52318245B4EC5DDA14414E88DE1BCEC2E8C0578243E1F1127267C0FA",
  "raw_log": "<raw-log>",
  "logs": [
    {
      "msg_index": 0,
      "log": "",
      "events": [
        {
          "type": "create_identity",
          "attributes": [
            {
              "key": "id",
              "value": "f0933a5745ba489495b9da982c5c60c1"
            },
            {
              "key": "owner",
              "value": "iaa12v3r0unp6nprp9zur0fn446n832dfag5l3w38e"
            }
          ]
        },
        {
          "type": "message",
          "attributes": [
            {
              "key": "action",
              "value": "create_identity"
            },
            {
              "key": "module",
              "value": "identity"
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
  "gas_used": "49688"
}
```

**_注_：**`id` 即为所创建身份的唯一 ID

## update

更新指定的身份。

```bash
irita tx identity update [id] [flags]
```

**标志：**

| 名称，速记     | 类型   | 必须 | 默认 | 描述                                                   |
| -------------- | ------ | ---- | ---- | ------------------------------------------------------ |
| --id  | string | 是   |      | 身份唯一 ID，长度为32的 Hex 字符串                   |
| --pubkey       | string | 否   |      | 身份主体的公钥，Hex 字符串；"rsa" 与 "dsa" 公钥采用 DER 编码，"ecdsa"、"ed25519" 与 "sm2" 公钥采用压缩形式                              |
| --pubkey-algo | string | 否   |      | 公钥算法，支持 "rsa"、"dsa"、"ecdsa"、"ed25519"、"sm2" |
| --cert-file    | string | 否   |      | 公钥证书文件路径；仅支持 PEM 编码的 X.509 证书         |
| --credentials  | string | 否   |      | 身份主体凭证信息的 URI，最大长度为140                    |

### 更新身份示例

```bash
irita tx identity update f0933a5745ba489495b9da982c5c60c1 --pubkey=03576aac14e47fc165789df8f86268faaa8f012a9bfbdef3ae18c22a63cfe5eac0 --pubkey-algo=ecdsa --credentials=https://security.com/aml/10001/ --from=node0 --chain-id=irita-test -b=block -o=json --indent -y --home=testnet/node0/iritacli
```

结果

```json
{
  "height": "4353",
  "txhash": "537271A9C2F4C147F3E784FB1AE77693D716CBC8009D752292A48BA88A9E1AA5",
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
              "value": "update_identity"
            },
            {
              "key": "module",
              "value": "identity"
            },
            {
              "key": "sender",
              "value": "iaa12v3r0unp6nprp9zur0fn446n832dfag5l3w38e"
            }
          ]
        },
        {
          "type": "update_identity",
          "attributes": [
            {
              "key": "id",
              "value": "f0933a5745ba489495b9da982c5c60c1"
            },
            {
              "key": "owner",
              "value": "iaa12v3r0unp6nprp9zur0fn446n832dfag5l3w38e"
            }
          ]
        }
      ]
    }
  ],
  "gas_wanted": "200000",
  "gas_used": "47148"
}
```

## identity

查询指定的身份。

```bash
irita query identity identity [id] [flags]
```

**参数：**

| 名称      | 类型   | 必须 | 默认 | 描述          |
| --------- | ------ | ---- | ---- | ------------- |
| id | string | 是   |      | 身份的唯一 ID |

### 查询身份示例

```bash
irita query identity identity f0933a5745ba489495b9da982c5c60c1 -o=json --indent --chain-id=irita-test
```

结果

```json
{
  "id": "F0933A5745BA489495B9DA982C5C60C1",
  "pubkeys": [
    {
      "pubkey": "03576AAC14E47FC165789DF8F86268FAAA8F012A9BFBDEF3AE18C22A63CFE5EAC0",
      "algorithm": "ECDSA"
    },
    {
      "pubkey": "03281CE4BA0B8C97E5B1434F8F298B064F03D4C1D21AAE9276065E170FC90A5D51",
      "algorithm": "SM2"
    }
  ],
  "credentials": "https://security.com/aml/10001/",
  "owner": "iaa12v3r0unp6nprp9zur0fn446n832dfag5l3w38e"
}
```
