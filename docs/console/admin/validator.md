<!--
order: 3
-->

# 验证人

IRITA 管理员可以通过控制台进行链节点的管理操作，主要管理功能有：

- [验证人](#验证人)
  - [create](#create)
  - [update](#update)
  - [remove](#remove)
  - [validator](#validator)
  - [list](#list)
    - [params](#params)

**创建、更新、删除三种操作只有拥有`RoleRootAdmin`或者`RoleNodeAdmin`权限的地址才可以操作。**

## create

创建一个新的验证人节点。

1. 初始化验证人节点

   ```bash
   irita init node1 --chain-id=irita-test --home=testnet
   ```

2. 导出验证节点node1(`步骤1`)私钥为pem格式，方便用于申请节点证书

   ```bash
   irita genkey --home=testnet --out-file priv.pem
   ```

3. 使用`步骤2`中的私钥文件生成证书请求并申请[签发证书](../../node_identity_management/cert.md)

4. 使用`步骤3`生成的证书创建validator交易

   ```bash
   irita tx validator create --name=v1 --cert=<cert.crt> --power=<100> --from=node0 --chain-id=irita-test -b=block -o=json --indent -y --home=testnet/node0/iritacli
   ```

   **参数：**

   | 名称    | 类型   | 必须 | 默认 | 描述           |
   | ------- | ------ | ---- | ---- | -------------- |
   | name    | string | 是   |      | 验证人名称     |
   | cert    | string | 是   |      | 证书           |
   | power   | uint64 | 是   |      | 验证人权重     |
   | details | string | 否   | ""   | 验证人描述信息 |
   | node-id | string | 否   | ""   | 节点ID         |

   ```json
   {
     "height": "643",
     "txhash": "9BEDB90F6018D3B1395B884DFA010469906E3B173595DCB7FF94EE4D1490A17C",
     "raw_log": "<raw-log>",
     "logs": [
       {
         "msg_index": 0,
         "log": "",
         "events": [
           {
             "type": "create_validator",
             "attributes": [
               {
                 "key": "validator",
                 "value": "v1"
               }
             ]
           },
           {
             "type": "message",
             "attributes": [
               {
                 "key": "action",
                 "value": "create_validator"
               },
               {
                 "key": "module",
                 "value": "validator"
               },
               {
                 "key": "sender",
                 "value": "iaa1mjwj7h8cln4m5aw7uuu4d4pkh9xwqjdvs7u94r"
               }
             ]
           }
         ]
       }
     ],
     "gas_wanted": "200000",
     "gas_used": "91383"
   }
   ```

## update

更新验证人信息。

```bash
irita tx validator update [name] --cert=<cert.crt> --power=<100> --details=<details>
```

**参数：**

| 名称    | 类型   | 必须 | 默认 | 描述           |
| ------- | ------ | ---- | ---- | -------------- |
| cert    | string | 是   |      | 证书           |
| power   | uint64 | 是   |      | 验证人权重     |
| details | string | 否   | ""   | 验证人描述信息 |

示例：

```bash
irita tx validator update v1 --details="hahhah" --from=node0 --chain-id=test -b=block -o=json --indent -y --home=testnet/node0/iritacli
```

结果：

```json
{
  "height": "681",
  "txhash": "8EDC132688630ACB3BD9E0C52BB9F5B44EDC851257F7AC9F84CEB56EA6294FB6",
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
              "value": "update_validator"
            },
            {
              "key": "module",
              "value": "validator"
            },
            {
              "key": "sender",
              "value": "iaa1mjwj7h8cln4m5aw7uuu4d4pkh9xwqjdvs7u94r"
            }
          ]
        },
        {
          "type": "update_validator",
          "attributes": [
            {
              "key": "validator",
              "value": "v1"
            }
          ]
        }
      ]
    }
  ],
  "gas_wanted": "200000",
  "gas_used": "66109"
}
```

## remove

删除一个已经存在的验证节点。

```bash
irita tx validator remove [name]
```

示例：

```bash
irita tx validator remove v1 --from=node0 --chain-id=test -b=block -o=json --indent -y --home=testnet/node0/iritacli
```

结果：

```json
{
  "height": "774",
  "txhash": "A847DE4DBD87F9FF253B36C798767B4D47159E6703C6AAFF090C4EB32A9E46F8",
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
              "value": "remove_validator"
            },
            {
              "key": "module",
              "value": "validator"
            },
            {
              "key": "sender",
              "value": "iaa1mjwj7h8cln4m5aw7uuu4d4pkh9xwqjdvs7u94r"
            }
          ]
        },
        {
          "type": "remove_validator",
          "attributes": [
            {
              "key": "validator",
              "value": "v1"
            }
          ]
        }
      ]
    }
  ],
  "gas_wanted": "200000",
  "gas_used": "45203"
}
```

## validator

通过名称查询验证人信息

```bash
irita q validator validator <name>
```

示例：

```bash
irita q validator validator node0 --trust-node
```

输出结果：

```text
name: node0
pubkey: ivalconspub1ulx45dfpq0hd53sy0cxaggmlj8ae4eesczxs7r76fh0ealf368t8kjyrz0t2v5nrwxw
certificate: |
  -----BEGIN CERTIFICATE-----
  MIIBrDCCAVICFECIvVOIFXDFUeQrb3XicjX7Bu4+MAoGCCqBHM9VAYN1MFgxCzAJ
  BgNVBAYTAkNOMQ0wCwYDVQQIDARyb290MQ0wCwYDVQQHDARyb290MQ0wCwYDVQQK
  DARyb290MQ0wCwYDVQQLDARyb290MQ0wCwYDVQQDDARyb290MB4XDTIwMDYyNDA1
  NTIwNloXDTIxMDYyNDA1NTIwNlowWTELMAkGA1UEBhMCQ04xDTALBgNVBAgMBHRl
  c3QxDTALBgNVBAcMBHRlc3QxDTALBgNVBAoMBHRlc3QxDTALBgNVBAsMBHRlc3Qx
  DjAMBgNVBAMMBW5vZGUwMFkwEwYHKoZIzj0CAQYIKoEcz1UBgi0DQgAE7tpGBH4N
  1CN/kfua5zDAjQ8P2k3fnv0x0dZ7SIMT1qYp4pOAJFWK6TVmZTQe2wVH+Z814Gar
  fXJwZ0jG4ps0ZzAKBggqgRzPVQGDdQNIADBFAiBwkv+vIiG7DRs4NimggDIsLTlG
  C54uaAiJh0M3kLBp+AIhAOuxjULhIXmhoCll7k3UrhRh98onBz5FOcv35iBC6n78
  -----END CERTIFICATE-----
power: 100
details: node0
jailed: false
operator: iaa1mjwj7h8cln4m5aw7uuu4d4pkh9xwqjdvs7u94r
```

## list

查询验证人列表

```bash
irita q validator list
```

示例：

```bash
irita q validator list --trust-node
```

输出结果：

```bash
- name: node0
  pubkey: ivalconspub1ulx45dfpq0hd53sy0cxaggmlj8ae4eesczxs7r76fh0ealf368t8kjyrz0t2v5nrwxw
  certificate: |
    -----BEGIN CERTIFICATE-----
    MIIBrDCCAVICFECIvVOIFXDFUeQrb3XicjX7Bu4+MAoGCCqBHM9VAYN1MFgxCzAJ
    BgNVBAYTAkNOMQ0wCwYDVQQIDARyb290MQ0wCwYDVQQHDARyb290MQ0wCwYDVQQK
    DARyb290MQ0wCwYDVQQLDARyb290MQ0wCwYDVQQDDARyb290MB4XDTIwMDYyNDA1
    NTIwNloXDTIxMDYyNDA1NTIwNlowWTELMAkGA1UEBhMCQ04xDTALBgNVBAgMBHRl
    c3QxDTALBgNVBAcMBHRlc3QxDTALBgNVBAoMBHRlc3QxDTALBgNVBAsMBHRlc3Qx
    DjAMBgNVBAMMBW5vZGUwMFkwEwYHKoZIzj0CAQYIKoEcz1UBgi0DQgAE7tpGBH4N
    1CN/kfua5zDAjQ8P2k3fnv0x0dZ7SIMT1qYp4pOAJFWK6TVmZTQe2wVH+Z814Gar
    fXJwZ0jG4ps0ZzAKBggqgRzPVQGDdQNIADBFAiBwkv+vIiG7DRs4NimggDIsLTlG
    C54uaAiJh0M3kLBp+AIhAOuxjULhIXmhoCll7k3UrhRh98onBz5FOcv35iBC6n78
    -----END CERTIFICATE-----
  power: 100
  details: node0
  jailed: false
  operator: iaa1mjwj7h8cln4m5aw7uuu4d4pkh9xwqjdvs7u94r
```

### params

查询验证人模块的参数信息

```bash
irita q validator params
```

示例：

```bash
irita q validator params --trust-node
```

输出结果：

```text
historical_entries: 100
```
