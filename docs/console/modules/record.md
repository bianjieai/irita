<!--
order: 3
-->

# 存证

## create

创建存证记录。

```bash
irita tx record create [digest] [digest-algo] [flags]
```

**参数：**

| 名称        | 类型   | 必须 | 默认 | 描述                                                       |
| ----------- | ------ | ---- | ---- | ---------------------------------------------------------- |
| digest      | string | 是   |      | 存证元数据的摘要                                           |
| digest-algo | string | 是   |      | 存证元数据摘要的生成算法，如 `sha256`，`sha512`，`sha3` 等 |

**标志：**

| 名称，速记 | 类型   | 必须 | 默认 | 描述                           |
| ---------- | ------ | ---- | ---- | ------------------------------ |
| --uri        | string | 否   |      | 存证元数据的 URI，如 IPFS 链接 |
| --meta       | string | 否   |      | 存证的元数据                   |

### 创建存证示例

```bash
irita tx record create c7a147baa5fb8da269da8dc565bb8522e23a7664f523370f8b8957efbdf8052b sha256 --uri=http://metadata.io/c7a147baa5fb8da269da8dc565bb8522e23a7664f523370f8b8957efbdf8052b --meta="test record" --from=node0 --chain-id=irita-test -b=block -o=json --indent -y --home=testnet/node0/iritacli
```

结果

```json
{
  "height": "1338",
  "txhash": "8ED89CDB2D5A06B3000DDD3D7C57CFBEE5D1B02B7D614D5BE393D5DC50E24542",
  "raw_log": "<raw-log>",
  "logs": [
    {
      "msg_index": 0,
      "log": "",
      "events": [
        {
          "type": "create_record",
          "attributes": [
            {
              "key": "creator",
              "value": "iaa19x57qemq4f3nret0hltehlgkv9vl0lxr0x552c"
            },
            {
              "key": "record_id",
              "value": "335db1d496e782e93f2339b8e2c99fd1be8d08410b303b0f462385d4a826913c"
            }
          ]
        },
        {
          "type": "message",
          "attributes": [
            {
              "key": "action",
              "value": "create_record"
            },
            {
              "key": "module",
              "value": "record"
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
  "gas_used": "50528"
}
```

**_注_：**`record-id` 即为所创建存证的唯一 ID

## record

查询指定的存证记录。

```bash
irita query record record [record-id] [flags]
```

**参数：**

| 名称      | 类型   | 必须 | 默认 | 描述          |
| --------- | ------ | ---- | ---- | ------------- |
| record-id | string | 是   |      | 存证的唯一 ID |

### 查询存证示例

```bash
irita query record record 335db1d496e782e93f2339b8e2c99fd1be8d08410b303b0f462385d4a826913c -o=json --indent --chain-id=irita-test
```

结果

```json
{
  "tx_hash": "8ED89CDB2D5A06B3000DDD3D7C57CFBEE5D1B02B7D614D5BE393D5DC50E24542",
  "contents": [
    {
      "digest": "c7a147baa5fb8da269da8dc565bb8522e23a7664f523370f8b8957efbdf8052b",
      "digest_algo": "sha256",
      "URI": "http://metadata.io/c7a147baa5fb8da269da8dc565bb8522e23a7664f523370f8b8957efbdf8052b",
      "meta": "test record"
    }
  ],
  "creator": "iaa19x57qemq4f3nret0hltehlgkv9vl0lxr0x552c"
}
```
