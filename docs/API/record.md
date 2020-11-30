<!--
order: 5
-->

# 存证

## 查询存证

根据存证 ID 查询存证。

**API：**

```bash
GET /record/records/{record-id}
```

**参数：**

- record-id：string，存证 ID

**返回值：**

- 存证查询结果：object

**请求示例：**

```bash
curl -X GET "http://localhost:1317/record/records/62b258724829f5165d6a5b5befcc889156fc0c81ff6a198f47b50037649201a8" -H "accept: application/json" | jq
```

```json
{
  "height": "1754",
  "result": {
    "tx_hash": "F8196D970CC5DEA89CA06AA269F2D74D46E52D6CEE37E6BB4171483BEC953145",
    "contents": [
      {
        "digest": "c7a147baa5fb8da269da8dc565bb8522e23a7664f523370f8b8957efbdf8052b",
        "digest_algo": "sha256",
        "uri": "http://metadata.io/c7a147baa5fb8da269da8dc565bb8522e23a7664f523370f8b8957efbdf8052b",
        "meta": "test record"
      }
    ],
    "creator": "iaa12v3r0unp6nprp9zur0fn446n832dfag5l3w38e"
  }
}
```
