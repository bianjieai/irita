# iritacli bank

`Bank`模块用于管理在`IRITA`上原生代币，可以通过本模块实现转账，余额查询等功能。

## 可用命令

| 名称                                      | 描述               |
| ----------------------------------------- | ------------------ |
| [send](#iriatacli-tx-send)                | 给指定账户转账     |
| [sign](#iriatacli-tx-sign)                | 为交易进行离线签名 |
| [balances](#iriatacli-q-account-balances) | 查询账户余额       |

## iriatacli tx send

给指定的地址转账

```bash
iritacli tx send [from_key_or_address] [to_address] [amount] --from=node0 --chain-id=test --fees=100stake -b=block
```

## iriatacli tx sign

为了提高账户安全性，`irita`支持交易离线签名保护账户的私钥。在任意交易中，使用参数`--generate-only`可以构建一个未签名的交易。这里以转账交易作为示例：

```bash
iritacli tx send iaa1f46tn5v6hm3av9rcwjk8add0xj0kstpdmxrtse iaa1k3j7texqrf3g5zu7ac227mv5sc0vwpju7lm9j6 10000000stake --chain-id=chain-vCna0J --fees=4stake -b block -y --generate-only
```

以上命令将构建一条未签名交易：

```json
{
  "type": "cosmos-sdk/StdTx",
  "value": {
    "msg": [
      {
        "type": "cosmos-sdk/MsgSend",
        "value": {
          "from_address": "iaa1f46tn5v6hm3av9rcwjk8add0xj0kstpdmxrtse",
          "to_address": "iaa1k3j7texqrf3g5zu7ac227mv5sc0vwpju7lm9j6",
          "amount": [
            {
              "denom": "stake",
              "amount": "10000000"
            }
          ]
        }
      }
    ],
    "fee": {
      "amount": [
        {
          "denom": "stake",
          "amount": "4"
        }
      ],
      "gas": "200000"
    },
    "signatures": null,
    "memo": ""
  }
}
```

将结果保存到文件`<file>`。对上述的离线交易进行签名：

```bash
iritacli tx sign tx.json --from node0 --home ./irita/node0/iritacli/ --chain-id=chain-vCna0J | jq .
```

将返回已签名的交易：

```json
{
  "type": "cosmos-sdk/StdTx",
  "value": {
    "msg": [
      {
        "type": "cosmos-sdk/MsgSend",
        "value": {
          "from_address": "iaa1f46tn5v6hm3av9rcwjk8add0xj0kstpdmxrtse",
          "to_address": "iaa1k3j7texqrf3g5zu7ac227mv5sc0vwpju7lm9j6",
          "amount": [
            {
              "denom": "stake",
              "amount": "10000000"
            }
          ]
        }
      }
    ],
    "fee": {
      "amount": [
        {
          "denom": "stake",
          "amount": "4"
        }
      ],
      "gas": "200000"
    },
    "signatures": [
      {
        "pub_key": {
          "type": "tendermint/PubKeySecp256k1",
          "value": "AptSxKaG4bUVCqitGmqlS90sOYXnXbqzSivu6B4A01EB"
        },
        "signature": "OLd8RmVJaCPii8esjSh4mBLn1njbulv/FSV++Oj/MlUW74JovtfIK9XBzm1P+ypdnIWe13cVRoc+/sQ2bC5Qvw=="
      }
    ],
    "memo": ""
  }
}
```

将结果保存到文件`<file>`。使用`broadcast`广播交易：

```bash
iritacli tx broadcast signed_tx.json -b block
```

## iriatacli q account balances

查询指定账户余额

```bash
iritacli query account balances
```
