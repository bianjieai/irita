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
 irita tx record create test  --uri= http://test.com  --meta="test record" --from=node0 --chain-id=test -b=block  -y --home=node0
```

结果

```json
{"height":"4","txhash":"1781AC381D41AF90B970E5ADABB30096004E5BCD68EA432736749CF1A9E34844","codespace":"","code":0,"data":"0A530A0D6372656174655F7265636F726412420A4034663166346238613134623134333939353933653531313063643837653639386339333533643566306232333433316436303836323964613834333862326264","raw_log":"[{\"events\":[{\"type\":\"create_record\",\"attributes\":[{\"key\":\"creator\",\"value\":\"iaa1t07s27vgvgczpsvu5z75703azmmc9wcmje452z\"},{\"key\":\"record_id\",\"value\":\"4f1f4b8a14b14399593e5110cd87e698c9353d5f0b23431d608629da8438b2bd\"}]},{\"type\":\"message\",\"attributes\":[{\"key\":\"action\",\"value\":\"create_record\"},{\"key\":\"module\",\"value\":\"record\"},{\"key\":\"sender\",\"value\":\"iaa1t07s27vgvgczpsvu5z75703azmmc9wcmje452z\"}]}]}]","logs":[{"msg_index":0,"log":"","events":[{"type":"create_record","attributes":[{"key":"creator","value":"iaa1t07s27vgvgczpsvu5z75703azmmc9wcmje452z"},{"key":"record_id","value":"4f1f4b8a14b14399593e5110cd87e698c9353d5f0b23431d608629da8438b2bd"}]},{"type":"message","attributes":[{"key":"action","value":"create_record"},{"key":"module","value":"record"},{"key":"sender","value":"iaa1t07s27vgvgczpsvu5z75703azmmc9wcmje452z"}]}]}],"info":"","gas_wanted":"400000","gas_used":"65212","tx":null,"timestamp":""}

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
irita query record record 4f1f4b8a14b14399593e5110cd87e698c9353d5f0b23431d608629da8438b2bd --chain-id=test
```

结果

```json
contents:
- digest: test
  digest_algo: http://test.com
  meta: test record
  uri: ""
creator: iaa1t07s27vgvgczpsvu5z75703azmmc9wcmje452z
tx_hash: 1781AC381D41AF90B970E5ADABB30096004E5BCD68EA432736749CF1A9E34844
```
