<!--
order: 5
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
irita tx identity create --pubkey=03281ce4ba0b8c97e5b1434f8f298b064f03d4c1d21aae9276065e170fc90a5d51 --pubkey-algo=sm2 --credentials=https://security.com/kyc/10001/ --from=node0 --chain-id=irita-test -y --home=testnet/node0/iritacli
```

结果

```json
{"height":"3731","txhash":"B95A2E5AC48A4776FBBC2F258F240E1DC2D34C8D3344CA32EB875FB8C2B330CE","codespace":"","code":0,"data":"0A110A0F6372656174655F6964656E74697479","raw_log":"[{\"events\":[{\"type\":\"create_identity\",\"attributes\":[{\"key\":\"id\",\"value\":\"61F3270E79044629896832B7925F89CB\"},{\"key\":\"owner\",\"value\":\"iaa17rc02z9tfec74pq8avjg5uj6kj8d57992q7ys0\"}]},{\"type\":\"message\",\"attributes\":[{\"key\":\"action\",\"value\":\"create_identity\"},{\"key\":\"module\",\"value\":\"identity\"},{\"key\":\"sender\",\"value\":\"iaa17rc02z9tfec74pq8avjg5uj6kj8d57992q7ys0\"}]}]}]","logs":[{"msg_index":0,"log":"","events":[{"type":"create_identity","attributes":[{"key":"id","value":"61F3270E79044629896832B7925F89CB"},{"key":"owner","value":"iaa17rc02z9tfec74pq8avjg5uj6kj8d57992q7ys0"}]},{"type":"message","attributes":[{"key":"action","value":"create_identity"},{"key":"module","value":"identity"},{"key":"sender","value":"iaa17rc02z9tfec74pq8avjg5uj6kj8d57992q7ys0"}]}]}],"info":"","gas_wanted":"400000","gas_used":"62680","tx":null,"timestamp":""}
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
irita tx identity update 61F3270E79044629896832B7925F89CB --pubkey=03576aac14e47fc165789df8f86268faaa8f012a9bfbdef3ae18c22a63cfe5eac0 --pubkey-algo=ecdsa --credentials=https://security.com/aml/10001/ --from=node0 --chain-id=irita-test   -y --home=testnet/node0/iritacli
```

结果

```json
{"height":"3763","txhash":"D32B58D0524E6BF0C8AA80CF69160D55BD43BA8585B5FAEBB45A97ABE98872C9","codespace":"","code":0,"data":"0A110A0F7570646174655F6964656E74697479","raw_log":"[{\"events\":[{\"type\":\"message\",\"attributes\":[{\"key\":\"action\",\"value\":\"update_identity\"},{\"key\":\"module\",\"value\":\"identity\"},{\"key\":\"sender\",\"value\":\"iaa17rc02z9tfec74pq8avjg5uj6kj8d57992q7ys0\"}]},{\"type\":\"update_identity\",\"attributes\":[{\"key\":\"id\",\"value\":\"61F3270E79044629896832B7925F89CB\"},{\"key\":\"owner\",\"value\":\"iaa17rc02z9tfec74pq8avjg5uj6kj8d57992q7ys0\"}]}]}]","logs":[{"msg_index":0,"log":"","events":[{"type":"message","attributes":[{"key":"action","value":"update_identity"},{"key":"module","value":"identity"},{"key":"sender","value":"iaa17rc02z9tfec74pq8avjg5uj6kj8d57992q7ys0"}]},{"type":"update_identity","attributes":[{"key":"id","value":"61F3270E79044629896832B7925F89CB"},{"key":"owner","value":"iaa17rc02z9tfec74pq8avjg5uj6kj8d57992q7ys0"}]}]}],"info":"","gas_wanted":"400000","gas_used":"60140","tx":null,"timestamp":""}
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
irita query identity identity 61F3270E79044629896832B7925F89CB  --chain-id=irita-test
```

结果

```json
certificates: []
credentials: https://security.com/aml/10001/
id: 61F3270E79044629896832B7925F89CB
owner: iaa17rc02z9tfec74pq8avjg5uj6kj8d57992q7ys0
pub_keys:
- algorithm: ECDSA
  pub_key: 03576AAC14E47FC165789DF8F86268FAAA8F012A9BFBDEF3AE18C22A63CFE5EAC0
- algorithm: SM2
  pub_key: 03281CE4BA0B8C97E5B1434F8F298B064F03D4C1D21AAE9276065E170FC90A5D51
```
