<!--
order: 2
-->

# 参数管理

IRITA 管理员可以根据达成的共识更新系统的相关参数。

```bash
irita tx params update [change-file] [flags]
```

**参数：**

| 名称        | 类型   | 必须 | 默认 | 描述             |
| ----------- | ------ | ---- | ---- | ---------------- |
| change-file | string | 是   |      | 参数改变文件路径 |

### 修改系统参数示例

```bash
echo '[{"subspace":"service","key":"MaxRequestTimeout","value":"150"}]' > paramschange.json

irita tx params update paramschange.json --from=node0 --chain-id=irita-test -b=block -o=json --indent -y --home=testnet/node0/iritacli
```

结果：

```json
{
  "height": "650",
  "txhash": "1740D8F2EB27DF05A618186B37CD081B8CF9C55485EA2F4C1380EF2D5293D777",
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
              "value": "update_params"
            },
            {
              "key": "module",
              "value": "params"
            },
            {
              "key": "sender",
              "value": "iaa1mjwj7h8cln4m5aw7uuu4d4pkh9xwqjdvs7u94r"
            }
          ]
        },
        {
          "type": "update_params",
          "attributes": [
            {
              "key": "param_key",
              "value": "MaxRequestTimeout"
            }
          ]
        }
      ]
    }
  ],
  "gas_wanted": "200000",
  "gas_used": "43313"
}
```
