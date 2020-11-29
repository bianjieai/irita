<!--
order: 7
-->

# 身份

## 查询身份

根据身份 ID 查询身份。

**API：**

```bash
GET /identity/identities/{id}
```

**参数：**

- id：string，身份 ID

**返回值：**

- 身份查询结果：object

**请求示例：**

```bash
curl -X GET "http://localhost:1317/identity/identities/f0933a5745ba489495b9da982c5c60c1" -H "accept: application/json" | jq
```

```json
{
  "height": "1754",
  "result": {
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
}
```
