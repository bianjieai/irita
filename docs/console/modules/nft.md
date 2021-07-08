<!--
order: 2
-->

# 数字资产建模

## issue

发行资产。

```bash
irita tx nft issue [denom] [flags]
```

**参数：**

| 名称      | 类型    | 必须 | 默认          | 描述                                                                     |
| ---------------- | ------- | ---- | ------------- | ------------------------------------------------------------------------ |
| denom  | string  | 是   |             | 资产的类别，全局唯一；长度为3到64，字母数字字符，以字母开始 |

**标志：**

| 名称，速记       | 类型    | 必须 | 默认          | 描述                                                                     |
| ---------------- | ------- | ---- | ------------- | ------------------------------------------------------------------------ |
| --schema       | string  | 否   |               | 资产元数据 [JSON Schema](https://JSON-Schema.org/) 规范 |

### 发行资产示例

```bash
irita tx nft issue security --name security --schema='{"type":"object","properties":{"name":{"type":"string"}}}' --from=node0 --chain-id=test   -y --home=node0 -b=block
```

结果

```json
{"height":"5","txhash":"4093B022BD7676E4EE52C3C47712162617A795491A1A8D56296D4BDA11FF7F5C","codespace":"","code":0,"data":"0A0D0A0B69737375655F64656E6F6D","raw_log":"[{\"events\":[{\"type\":\"issue_denom\",\"attributes\":[{\"key\":\"denom_id\",\"value\":\"security\"},{\"key\":\"denom_name\",\"value\":\"security\"},{\"key\":\"creator\",\"value\":\"iaa1t07s27vgvgczpsvu5z75703azmmc9wcmje452z\"}]},{\"type\":\"message\",\"attributes\":[{\"key\":\"action\",\"value\":\"issue_denom\"},{\"key\":\"module\",\"value\":\"nft\"},{\"key\":\"sender\",\"value\":\"iaa1t07s27vgvgczpsvu5z75703azmmc9wcmje452z\"}]}]}]","logs":[{"msg_index":0,"log":"","events":[{"type":"issue_denom","attributes":[{"key":"denom_id","value":"security"},{"key":"denom_name","value":"security"},{"key":"creator","value":"iaa1t07s27vgvgczpsvu5z75703azmmc9wcmje452z"}]},{"type":"message","attributes":[{"key":"action","value":"issue_denom"},{"key":"module","value":"nft"},{"key":"sender","value":"iaa1t07s27vgvgczpsvu5z75703azmmc9wcmje452z"}]}]}],"info":"","gas_wanted":"400000","gas_used":"66022","tx":null,"timestamp":""}
```

## mint

创建指定类别的具体资产。

```bash
irita tx nft mint [denom] [tokenID] [flags]
```

**参数：**

| 名称      | 类型    | 必须 | 默认          | 描述                                                                     |
| ---------------- | ------- | ---- | ------------- | ------------------------------------------------------------------------ |
| denom  | string  | 是   |             | 资产的类别；长度为3到64，字母数字字符，以字母开始 |
| tokenID  | string  | 是   |             | 资产的唯一 ID，如 `UUID` |

**标志：**

| 名称，速记   | 类型   | 必须 | 默认  | 描述                                                  |
| ------------ | ------ | ---- | ----- | ----------------------------------------------------- |
| --uri | string | 否   | | 资产元数据的 `URI` |
| --data | string | 否 | | 资产元数据 |
| --recipient | string | 否   | | 资产接收者地址，默认为交易发起者地址 |

### 创建资产示例

```bash
irita tx nft mint security test --uri=https://test.com --data='{"name":"test security"}' --from=node0 --chain-id=test -y --home=node0 -b=block
```

结果

```json
{"height":"22","txhash":"EB0E454E13EFBA541525F8538944611CE92F49BBE7C0DA3D05E3953CB9080CD9","codespace":"","code":0,"data":"0A0A0A086D696E745F6E6674","raw_log":"[{\"events\":[{\"type\":\"message\",\"attributes\":[{\"key\":\"action\",\"value\":\"mint_nft\"},{\"key\":\"module\",\"value\":\"nft\"},{\"key\":\"sender\",\"value\":\"iaa1t07s27vgvgczpsvu5z75703azmmc9wcmje452z\"}]},{\"type\":\"mint_nft\",\"attributes\":[{\"key\":\"token_id\",\"value\":\"test\"},{\"key\":\"denom_id\",\"value\":\"security\"},{\"key\":\"token_uri\",\"value\":\"https://test.com\"},{\"key\":\"recipient\",\"value\":\"iaa1t07s27vgvgczpsvu5z75703azmmc9wcmje452z\"}]}]}]","logs":[{"msg_index":0,"log":"","events":[{"type":"message","attributes":[{"key":"action","value":"mint_nft"},{"key":"module","value":"nft"},{"key":"sender","value":"iaa1t07s27vgvgczpsvu5z75703azmmc9wcmje452z"}]},{"type":"mint_nft","attributes":[{"key":"token_id","value":"test"},{"key":"denom_id","value":"security"},{"key":"token_uri","value":"https://test.com"},{"key":"recipient","value":"iaa1t07s27vgvgczpsvu5z75703azmmc9wcmje452z"}]}]}],"info":"","gas_wanted":"400000","gas_used":"62250","tx":null,"timestamp":""}
```

## edit

编辑指定的资产。可更新的属性包括：资产元数据、元数据 `URI`

```bash
irita tx nft edit [denom] [tokenID] [flags]
```

**参数：**

| 名称      | 类型    | 必须 | 默认          | 描述                                                                     |
| ---------------- | ------- | ---- | ------------- | ------------------------------------------------------------------------ |
| denom  | string  | 是   |             | 资产的类别；长度为3到64，字母数字字符，以字母开始 |
| tokenID  | string  | 是   |             | 资产的唯一 ID |

**标志：**

| 名称，速记 | 类型   | 必须 | 默认 | 描述                                       |
| ---------- | ------ | ---- | ---- | ------------------------------------------ |
| --token-uri | string | 否   | | 资产元数据的 `URI` |
| --token-data | string | 否 | | 资产元数据 |

### 编辑资产示例

```bash
irita tx nft edit security test --data='{"name":"new test security"}' --from=node0 --chain-id=test  -y --home=node0  -b=block
```

结果

```json
{"height":"33","txhash":"715E19828C451E6500B61240E304F4758C9A2844914CDE67C8D7302F02DFECCC","codespace":"","code":0,"data":"0A0A0A08656469745F6E6674","raw_log":"[{\"events\":[{\"type\":\"edit_nft\",\"attributes\":[{\"key\":\"token_id\",\"value\":\"test\"},{\"key\":\"denom_id\",\"value\":\"security\"},{\"key\":\"token_uri\",\"value\":\"[do-not-modify]\"},{\"key\":\"owner\",\"value\":\"iaa1t07s27vgvgczpsvu5z75703azmmc9wcmje452z\"}]},{\"type\":\"message\",\"attributes\":[{\"key\":\"action\",\"value\":\"edit_nft\"},{\"key\":\"module\",\"value\":\"nft\"},{\"key\":\"sender\",\"value\":\"iaa1t07s27vgvgczpsvu5z75703azmmc9wcmje452z\"}]}]}]","logs":[{"msg_index":0,"log":"","events":[{"type":"edit_nft","attributes":[{"key":"token_id","value":"test"},{"key":"denom_id","value":"security"},{"key":"token_uri","value":"[do-not-modify]"},{"key":"owner","value":"iaa1t07s27vgvgczpsvu5z75703azmmc9wcmje452z"}]},{"type":"message","attributes":[{"key":"action","value":"edit_nft"},{"key":"module","value":"nft"},{"key":"sender","value":"iaa1t07s27vgvgczpsvu5z75703azmmc9wcmje452z"}]}]}],"info":"","gas_wanted":"400000","gas_used":"57162","tx":null,"timestamp":""}
```

## transfer

转移指定资产。

```bash
irita tx nft transfer [recipient] [denom] [tokenID] [flags]
```

**参数：**

| 名称      | 类型    | 必须 | 默认          | 描述                                                                     |
| ---------------- | ------- | ---- | ------------- | ------------------------------------------------------------------------ |
| recipient  | string  | 是   |             | 积分的唯一标识符 |
| denom  | string  | 是   |             | 资产的类别；长度为3到64，字母数字字符，以字母开始 |
| tokenID  | string  | 是   |             | 资产的唯一 ID |

### 转移资产示例

```bash
irita tx nft transfer  iaa177w2evwnx3uje646k78zxlp82mc9eatuwkdwlh<node1>   security test --from=node0 --chain-id=test  -y --home=node0 -b=block
```

结果

```json
{"height":"105","txhash":"17BDE3AA8A11C473EE286E420CE2ECA811DF1A5D7712D7FE1A416A3C1121DD64","codespace":"","code":0,"data":"0A0E0A0C7472616E736665725F6E6674","raw_log":"[{\"events\":[{\"type\":\"message\",\"attributes\":[{\"key\":\"action\",\"value\":\"transfer_nft\"},{\"key\":\"module\",\"value\":\"nft\"},{\"key\":\"sender\",\"value\":\"iaa1t07s27vgvgczpsvu5z75703azmmc9wcmje452z\"}]},{\"type\":\"transfer_nft\",\"attributes\":[{\"key\":\"token_id\",\"value\":\"test\"},{\"key\":\"denom_id\",\"value\":\"security\"},{\"key\":\"sender\",\"value\":\"iaa1t07s27vgvgczpsvu5z75703azmmc9wcmje452z\"},{\"key\":\"recipient\",\"value\":\"iaa177w2evwnx3uje646k78zxlp82mc9eatuwkdwlh\"}]}]}]","logs":[{"msg_index":0,"log":"","events":[{"type":"message","attributes":[{"key":"action","value":"transfer_nft"},{"key":"module","value":"nft"},{"key":"sender","value":"iaa1t07s27vgvgczpsvu5z75703azmmc9wcmje452z"}]},{"type":"transfer_nft","attributes":[{"key":"token_id","value":"test"},{"key":"denom_id","value":"security"},{"key":"sender","value":"iaa1t07s27vgvgczpsvu5z75703azmmc9wcmje452z"},{"key":"recipient","value":"iaa177w2evwnx3uje646k78zxlp82mc9eatuwkdwlh"}]}]}],"info":"","gas_wanted":"400000","gas_used":"60714","tx":null,"timestamp":""}
```

## burn

销毁指定资产。

```bash
irita tx nft burn [denom] [tokenID] [flags]
```

**参数：**

| 名称      | 类型    | 必须 | 默认          | 描述                                                                     |
| ---------------- | ------- | ---- | ------------- | ------------------------------------------------------------------------ |
| denom  | string  | 是   |             | 资产的类别；长度为3到64，字母数字字符，以字母开始 |
| tokenID  | string  | 是   |             | 资产的唯一 ID |

### 销毁资产示例

```bash
irita tx nft burn security test --from=iaa1t07s27vgvgczpsvu5z75703azmmc9wcmje452z<node1>  --chain-id=test -b=block  -y --home=node0 
```

结果

```json
{"height":"119","txhash":"DA7EDB31F795F39D37456BCDA9647193720B4FDC588AF42D99E261423BC3E9E4","codespace":"","code":0,"data":"0A0A0A086275726E5F6E6674","raw_log":"[{\"events\":[{\"type\":\"burn_nft\",\"attributes\":[{\"key\":\"denom_id\",\"value\":\"security\"},{\"key\":\"token_id\",\"value\":\"test\"},{\"key\":\"owner\",\"value\":\"iaa177w2evwnx3uje646k78zxlp82mc9eatuwkdwlh\"}]},{\"type\":\"message\",\"attributes\":[{\"key\":\"action\",\"value\":\"burn_nft\"},{\"key\":\"module\",\"value\":\"nft\"},{\"key\":\"sender\",\"value\":\"iaa177w2evwnx3uje646k78zxlp82mc9eatuwkdwlh\"}]}]}]","logs":[{"msg_index":0,"log":"","events":[{"type":"burn_nft","attributes":[{"key":"denom_id","value":"security"},{"key":"token_id","value":"test"},{"key":"owner","value":"iaa177w2evwnx3uje646k78zxlp82mc9eatuwkdwlh"}]},{"type":"message","attributes":[{"key":"action","value":"burn_nft"},{"key":"module","value":"nft"},{"key":"sender","value":"iaa177w2evwnx3uje646k78zxlp82mc9eatuwkdwlh"}]}]}],"info":"","gas_wanted":"400000","gas_used":"61862","tx":null,"timestamp":""}
```



## token

查询指定类别和 `ID` 的资产。

```bash
irita query nft token [denom] [tokenID] [flags]
```

**参数：**

| 名称      | 类型    | 必须 | 默认          | 描述                                                                     |
| ---------------- | ------- | ---- | ------------- | ------------------------------------------------------------------------ |
| denom  | string  | 是   |             | 资产的类别；长度为3到64，字母数字字符，以字母开始 |
| tokenID  | string  | 是   |             | 资产的唯一 ID |

### 查询指定资产示例

```bash
irita query nft token security test<tokenid>  --chain-id=test
```

结果

```json
data: '{"name":"new test security"}'
id: test
name: ""
owner: iaa1t07s27vgvgczpsvu5z75703azmmc9wcmje452z
uri: https://test.com
```

## denom

查询指定类别的资产信息。

```bash
irita query nft denom [denom] [flags]
```

**参数：**

| 名称      | 类型    | 必须 | 默认          | 描述                                                                     |
| ---------------- | ------- | ---- | ------------- | ------------------------------------------------------------------------ |
| denom  | string  | 是   |             | 资产的类别；长度为3到64，字母数字字符，以字母开始 |

### 查询指定类别的资产信息示例

```bash
irita query nft denom security  --chain-id=test 
```

结果

```json
creator: iaa1t07s27vgvgczpsvu5z75703azmmc9wcmje452z
id: security
name: security
schema: '{"type":"object","properties":{"name":{"type":"string"}}}'
```

## denoms

查询所有类别的资产信息。

```bash
irita query nft denoms [flags]
```

### 查询所有类别的资产信息示例

```bash
irita query nft denoms  --chain-id=irita-test
```

结果

```json
  creator: iaa1t07s27vgvgczpsvu5z75703azmmc9wcmje452z
  id: security
  name: security
  schema: '{"type":"object","properties":{"name":{"type":"string"}}}'
```

## supply

查询指定类别资产的总量。如 `owner` 被指定，则查询此 `owner` 所拥有的该类别资产的总量。 

```bash
irita query nft supply [denom] [flags]
```

**参数：**

| 名称      | 类型    | 必须 | 默认          | 描述                                                                     |
| ---------------- | ------- | ---- | ------------- | ------------------------------------------------------------------------ |
| denom  | string  | 是   |             | 资产的类别；长度为3到64，字母数字字符，以字母开始 |

**标志：**

| 名称，速记 | 类型   | 必须 | 默认 | 描述                                       |
| ---------- | ------ | ---- | ---- | ------------------------------------------ |
| --owner | string | 否   | | 资产所有者地址 |

### 查询指定类别的资产总量示例

```bash
irita query nft supply security --chain-id=test
```

结果

```text
1
```

### 查询指定账户某类别资产的总量

```bash
irita query nft supply security --owner=iaa1pf0r9rhfyzdyw3ed2hk0kyzjfwz4tehwsynxvt --chain-id=irita-test
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
irita query nft owner iaa1t07s27vgvgczpsvu5z75703azmmc9wcmje452z  --chain-id=test
```

结果

```json
owner:
  address: iaa1t07s27vgvgczpsvu5z75703azmmc9wcmje452z
  id_collections:
  - denom_id: security
    token_ids:
    - test
```

### 查询账户指定类别的所有资产示例

```bash
 irita query nft owner iaa1t07s27vgvgczpsvu5z75703azmmc9wcmje452z --denom-id=security  --chain-id=test
```

结果

```json
owner:
  address: iaa1t07s27vgvgczpsvu5z75703azmmc9wcmje452z
  id_collections:
  - denom_id: security
    token_ids:
    - test
```

## collection

查询指定类别的所有资产。

```bash
irita query nft collection [denom] [flags]
```

**参数：**

| 名称      | 类型    | 必须 | 默认          | 描述                                                                     |
| ---------------- | ------- | ---- | ------------- | ------------------------------------------------------------------------ |
| denom  | string  | 是   |             | 资产的类别；长度为3到64，字母数字字符，以字母开始 |

### 查询指定类别的所有资产示例

```bash
irita query nft collection security  --chain-id=irita-test
```

结果

```json
  denom:
    creator: iaa1t07s27vgvgczpsvu5z75703azmmc9wcmje452z
    id: security
    name: security
    schema: '{"type":"object","properties":{"name":{"type":"string"}}}'
  nfts:
  - data: '{"name":"new test security"}'
    id: test
    name: ""
    owner: iaa1t07s27vgvgczpsvu5z75703azmmc9wcmje452z
    uri: https://test.com
```
