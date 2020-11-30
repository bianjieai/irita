<!--
order: 3
-->

# 积分

## 查询积分

查询指定的积分。

**API：**

```bash
GET /token/tokens/{symbol}
```

**参数：**

- symbol：string，积分标识符

**返回值：**

- 积分查询结果：object

**请求示例：**

```bash
curl -X GET "http://localhost:1317/token/tokens/mycredit" -H "accept: application/json" | jq
```

```json
{
  "height": "1792",
  "result": {
    "type": "irismod/token/Token",
    "value": {
      "symbol": "mycredit",
      "name": "my credit",
      "min_unit": "mycretdit",
      "initial_supply": "10000",
      "max_supply": "100000",
      "mintable": true,
      "owner": "iaa12v3r0unp6nprp9zur0fn446n832dfag5l3w38e"
    }
  }
}
```

## 查询所有积分

查询已发行的所有积分，或者查询指定所有者发行的积分。

**API：**

```bash
GET /token/tokens?owner=<owner>
```

**参数：**

- owner：string，所有者账户地址，可选。

**返回值：**

- 积分查询结果：object

**请求示例：**

```bash
curl -X GET "http://localhost:1317/token/tokens?owner=iaa12v3r0unp6nprp9zur0fn446n832dfag5l3w38e" -H "accept: application/json" | jq
```

```json
{
  "height": "1842",
  "result": [
    {
      "type": "irismod/token/Token",
      "value": {
        "symbol": "mycredit",
        "name": "my credit",
        "min_unit": "mycretdit",
        "initial_supply": "10000",
        "max_supply": "100000",
        "mintable": true,
        "owner": "iaa12v3r0unp6nprp9zur0fn446n832dfag5l3w38e"
      }
    }
  ]
}
```

## 查询费用

查询积分发行和增发的费用。

**API：**

```bash
GET /token/tokens/{symbol}/fee
```

**参数：**

- symbol：string，积分标识符

**返回值：**

- 费用查询结果：object

**请求示例：**

```bash
curl -X GET "http://localhost:1317/token/tokens/mycredit/fee" -H "accept: application/json" | jq
```

```json
{
  "height": "1861",
  "result": {
    "exist": true,
    "issue_fee": {
      "denom": "point",
      "amount": "4672"
    },
    "mint_fee": {
      "denom": "point",
      "amount": "467"
    }
  }
}
```
