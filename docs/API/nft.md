<!--
order: 4
-->

# 数字资产建模

## 查询指定的资产

根据资产类别和 ID 查询资产。

**API：**

```bash
GET /nft/nfts/{denom}/{token-id}
```

**参数：**

- denom：string，资产类别

- token-id：string，资产 ID

**返回值：**

- 资产查询结果：object

**请求示例：**

```bash
curl -X GET "http://localhost:1317/nft/nfts/security/a4c74c4203af41619d00bb3e2f462c10" -H "accept: application/json" | jq
```

```json
{
  "height": "1978",
  "result": {
    "type": "irismod/nft/BaseNFT",
    "value": {
      "id": "a4c74c4203af41619d00bb3e2f462c10",
      "owner": "iaa12v3r0unp6nprp9zur0fn446n832dfag5l3w38e",
      "token_uri": "http://metadata.io/a4c74c4203af41619d00bb3e2f462c10",
      "token_data": "{\"name\":\"test security\"}"
    }
  }
}
```

## 查询指定的资产类别

查询指定的资产类别定义。

**API：**

```bash
GET /nft/nfts/denoms/{denom}
```

**参数：**

- denom：string，资产类别

**返回值：**

- 资产类别查询结果：object

**请求示例：**

```bash
curl -X GET "http://localhost:1317/nft/nfts/denom/security" -H "accept: application/json" | jq
```

```json
{
  "height": "2007",
  "result": {
    "name": "security",
    "schema": "{\"type\":\"object\",\"properties\":{\"name\":{\"type\":\"string\"}}}",
    "creator": "iaa12v3r0unp6nprp9zur0fn446n832dfag5l3w38e"
  }
}
```

## 查询所有资产类别

查询所有的资产类别。

**API：**

```bash
GET /nft/nfts/denoms
```

**参数：**

无

**返回值：**

- 资产类别列表查询结果：object

**请求示例：**

```bash
curl -X GET "http://localhost:1317/nft/nfts/denoms" -H "accept: application/json" | jq
```

```json
{
  "height": "2020",
  "result": [
    {
      "name": "security",
      "schema": "{\"type\":\"object\",\"properties\":{\"name\":{\"type\":\"string\"}}}",
      "creator": "iaa12v3r0unp6nprp9zur0fn446n832dfag5l3w38e"
    }
  ]
}
```

## 查询资产供应量

查询指定类别资产的总量。如 `owner` 被指定，则查询此 `owner` 所拥有的该类别资产的总量。

**API：**

```bash
GET /nft/nfts/supplies/{denom}?owner=<owner>
```

**参数：**

- denom：string，资产类别

- owner：string，所有者账户地址，可选

**返回值：**

- 资产总量查询结果：object

**请求示例：**

```bash
curl -X GET "http://localhost:1317/nft/nfts/supplies/security" -H "accept: application/json" | jq
```

```json
{
  "height": "2027",
  "result": "1"
}
```

## 查询指定账户的资产

查询指定账户的资产集合。如提供 `denom`，则查询该账户指定 `denom` 的资产列表。

**API：**

```bash
GET /nft/nfts/owners/{owner}?denom=<denom>
```

**参数：**

- owner：string，所有者账户地址

- denom：string，资产类别，可选

**返回值：**

- 账户资产查询结果：object

**请求示例：**

```bash
curl -X GET "http://localhost:1317/nft/nfts/owners/iaa12v3r0unp6nprp9zur0fn446n832dfag5l3w38e" -H "accept: application/json" | jq
```

```json
{
  "height": "2040",
  "result": {
    "address": "iaa12v3r0unp6nprp9zur0fn446n832dfag5l3w38e",
    "id_collections": [
      {
        "denom": "security",
        "ids": [
          "a4c74c4203af41619d00bb3e2f462c10"
        ]
      }
    ]
  }
}
```

## 查询指定类别的资产集合

查询指定类别的所有资产。

**API：**

```bash
GET /nft/nfts/collections/{denom}
```

**参数：**

- denom：string，资产类别

**返回值：**

- 资产集合查询结果：object

**请求示例：**

```bash
curl -X GET "http://localhost:1317/nft/nfts/collections/security" -H "accept: application/json" | jq
```

```json
{
  "height": "2049",
  "result": {
    "denom": {
      "name": "security",
      "schema": "{\"type\":\"object\",\"properties\":{\"name\":{\"type\":\"string\"}}}",
      "creator": "iaa12v3r0unp6nprp9zur0fn446n832dfag5l3w38e"
    },
    "nfts": [
      {
        "type": "irismod/nft/BaseNFT",
        "value": {
          "id": "a4c74c4203af41619d00bb3e2f462c10",
          "owner": "iaa12v3r0unp6nprp9zur0fn446n832dfag5l3w38e",
          "token_uri": "http://metadata.io/a4c74c4203af41619d00bb3e2f462c10",
          "token_data": "{\"name\":\"test security\"}"
        }
      }
    ]
  }
}
```
