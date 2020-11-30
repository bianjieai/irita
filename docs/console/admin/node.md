<!--
order: 4
-->

# 节点

IRITA 管理员可以通过控制台进行节点的管理操作，主要管理功能有：

- [节点](#节点)
  - [grant](#grant)
  - [revoke](#revoke)
  - [node](#node)
  - [nodes](#nodes)

**授权及撤销授权只有拥有 `RoleRootAdmin` 或者 `RoleNodeAdmin` 权限的账户才可以操作。**

## grant

授权节点加入到网络。

```bash
irita tx node grant --name=<name> --cert=<certificate-file> --from=node0 --chain-id=irita-test -b=block -y --home=testnet/node0/iritacli
```

::tip
节点证书的生成请参考[这里](../../node_identity_management/cert.md)
::

**标志：**

| 名称    | 类型   | 必须 | 默认 | 描述           |
| ------- | ------ | ---- | ---- | -------------- |
| name    | string | 是   |      | 节点名称     |
| cert    | string | 是   |      | 节点证书路径           |

示例：

```bash
irita tx node grant --name node1 cert=node1.crt --from=node0 --chain-id=test -b=block -y --home=testnet/node0/iritacli
```

结果：

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
         "type": "grant_node",
         "attributes": [
           {
             "key": "id",
             "value": "b0fc236d4b3c39031640d203329e8b9cbb277ac1"
           }
         ]
       },
       {
         "type": "message",
         "attributes": [
           {
             "key": "action",
             "value": "grant_node"
           },
           {
             "key": "module",
             "value": "node"
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

## revoke

撤销节点授权。

```bash
irita tx node revoke [id] --from=node0 --chain-id=irita-test -b=block -y --home=testnet/node0/iritacli
```

**参数：**

| 名称    | 类型   | 必须 | 默认 | 描述           |
| ------- | ------ | ---- | ---- | -------------- |
| id    | string | 是   |      | 节点 ID           |

示例：

```bash
irita tx node revoke b0fc236d4b3c39031640d203329e8b9cbb277ac1 --from=node0 --chain-id=test -b=block -y --home=testnet/node0/iritacli
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
              "value": "revoke_node"
            },
            {
              "key": "module",
              "value": "node"
            },
            {
              "key": "sender",
              "value": "iaa1mjwj7h8cln4m5aw7uuu4d4pkh9xwqjdvs7u94r"
            }
          ]
        },
        {
          "type": "revoke_node",
          "attributes": [
            {
              "key": "id",
              "value": "b0fc236d4b3c39031640d203329e8b9cbb277ac1"
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

## node

查询一个已授权节点。

```bash
irita tx node node [id]
```

**参数：**

| 名称    | 类型   | 必须 | 默认 | 描述           |
| ------- | ------ | ---- | ---- | -------------- |
| id    | string | 是   |      | 节点 ID           |

## nodes

查询所有已授权的节点。

```bash
irita query node nodes
```
