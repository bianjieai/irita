<!--
order: 2
-->

# 节点与链

## 查询节点信息

**API：**

```bash
GET /node_info
```

**参数：**

无

**返回值：**

- 节点信息：object

**请求示例：**

```bash
curl -X GET "http://localhost:1317/node_info" -H "accept: application/json" | jq
```

```json
{
  "node_info": {
    "protocol_version": {
      "p2p": "7",
      "block": "10",
      "app": "0"
    },
    "id": "b2aa2d7a69cce7efeef554c23cab293ed64bd20a",
    "listen_addr": "tcp://0.0.0.0:26656",
    "network": "irita-test",
    "version": "0.33.4",
    "channels": "4020212223303800",
    "moniker": "node0",
    "other": {
      "tx_index": "on",
      "rpc_address": "tcp://0.0.0.0:26657"
    }
  },
  "application_version": {
    "name": "irita",
    "server_name": "irita",
    "version": "",
    "commit": "3daa57798eaf3cd13007fbb356c5760be771f393",
    "build_tags": "netgo,ledger",
    "go": "go version go1.13.1 darwin/amd64"
  }
}
```

## 查询最新的区块

**API：**

```bash
GET /blocks/latest
```

**参数：**

无

**返回值：**

- 区块信息：object，最新的区块数据

**请求示例：**

```bash
curl -X GET "http://localhost:1317/blocks/latest" -H "accept: application/json" | jq
```

```json
{
  "block_id": {
    "hash": "DBE96AA9FED168751A6E6C14EE8EC4486746B93225DEDB5A8B8FCC8044D4C74C",
    "parts": {
      "total": "1",
      "hash": "0DC2CDF5500D73B673903C7D3592FBE69B25AFD9CA1DA0F83A756DE3C1BE6025"
    }
  },
  "block": {
    "header": {
      "version": {
        "block": "10",
        "app": "0"
      },
      "chain_id": "irita-test",
      "height": "1689",
      "time": "2020-05-29T06:45:32.338892Z",
      "last_block_id": {
        "hash": "87F0AA90A763AB65399FD7983C586E12A2CB0A85A65F90AE706A7C2605B1D3B3",
        "parts": {
          "total": "1",
          "hash": "5CA7CECEE1CE8EA5F87347E7E1E62CE26E9C7DB49EE5A04A11A03E17867AD17C"
        }
      },
      "last_commit_hash": "721C706AA34ED59F948F7B9CA4CB62ECA9EE3CE5DD70F89B2C80F7AA9C71FE08",
      "data_hash": "",
      "validators_hash": "AAFAD29903844C650B1F8301E71AC13180D20EB4EE05554A35E958914C2A36D6",
      "next_validators_hash": "AAFAD29903844C650B1F8301E71AC13180D20EB4EE05554A35E958914C2A36D6",
      "consensus_hash": "048091BC7DDC283F77BFBF91D73C44DA58C3DF8A9CBC867405D8B7F3DAADA22F",
      "app_hash": "EAB2207D65D83E275489A2EB6589D68816221C9B6BB88DF15620966DFC8DF6A9",
      "last_results_hash": "",
      "evidence_hash": "",
      "proposer_address": "2021B3614D90C26CFC1E6D4786920373073D5FB6"
    },
    "data": {
      "txs": null
    },
    "evidence": {
      "evidence": null
    },
    "last_commit": {
      "height": "1688",
      "round": "0",
      "block_id": {
        "hash": "87F0AA90A763AB65399FD7983C586E12A2CB0A85A65F90AE706A7C2605B1D3B3",
        "parts": {
          "total": "1",
          "hash": "5CA7CECEE1CE8EA5F87347E7E1E62CE26E9C7DB49EE5A04A11A03E17867AD17C"
        }
      },
      "signatures": [
        {
          "block_id_flag": 2,
          "validator_address": "2021B3614D90C26CFC1E6D4786920373073D5FB6",
          "timestamp": "2020-05-29T06:45:32.338892Z",
          "signature": "JTdExuVtlmpI/la8t8PLAcw1RJwCyHRhDA0o8NMyij/WZIhF/7jQrhwyGkV1l83sKfJKmCvoa3zRzRUhKlZ9yw=="
        }
      ]
    }
  }
}
```

## 查询交易

根据交易 Hash 查询交易。

**API：**

```bash
GET /txs/{hash}
```

**参数：**

- hash：string，交易的 Hash

**返回值：**

- TxQuery：object，交易查询结果对象

**请求示例：**

```bash
curl -X GET "http://localhost:1317/txs/F266F1734DDE56A8C4A5676BD929C414F0E9C874131856AF362762E4474B489D" -H "accept: application/json" | jq
```

```json
{
  "height": "388",
  "txhash": "F266F1734DDE56A8C4A5676BD929C414F0E9C874131856AF362762E4474B489D",
  "raw_log": "<raw-log>",
  "logs": [
    {
      "msg_index": 0,
      "log": "",
      "events": [
        {
          "type": "block_account",
          "attributes": [
            {
              "key": "account",
              "value": "iaa18up8anyjpal8rncm8rd4ukp5f7etga795gp33q"
            }
          ]
        },
        {
          "type": "message",
          "attributes": [
            {
              "key": "action",
              "value": "block_account"
            },
            {
              "key": "module",
              "value": "admin"
            },
            {
              "key": "sender",
              "value": "iaa12v3r0unp6nprp9zur0fn446n832dfag5l3w38e"
            }
          ]
        }
      ]
    }
  ],
  "gas_wanted": "200000",
  "gas_used": "41058",
  "tx": {
    "type": "cosmos-sdk/StdTx",
    "value": {
      "msg": [
        {
          "type": "irita/modules/MsgBlockAccount",
          "value": {
            "address": "iaa18up8anyjpal8rncm8rd4ukp5f7etga795gp33q",
            "operator": "iaa12v3r0unp6nprp9zur0fn446n832dfag5l3w38e"
          }
        }
      ],
      "fee": {
        "amount": [],
        "gas": "200000"
      },
      "signatures": [
        {
          "pub_key": "61rphyECERx9oolgDwPYd7r8rqra8RRY9coPnNLrrHvmf1xhnoY=",
          "signature": "uxeCOa/+qGevQmoiP0HJ9KptxySPN0cku3qteGVCpCpvLVl/7Is8Ir++yTlpRZK4n05fFFu/vPiov9wTmRbDBw=="
        }
      ],
      "memo": ""
    }
  },
  "timestamp": "2020-05-29T03:38:39Z"
}
```

## 广播交易

将一个已签名交易发送到节点。

**API：**

```bash
POST /txs
```

**参数：**

- txBroadcast：object，交易广播的消息体

**返回值：**

- BroadcastTxCommitResult：object，交易提交结果对象

**请求示例：**

```bash
curl -X POST "http://localhost:1317/txs" -d '{"tx":{"msg":[{"type":"irismod/nft/MsgIssueDenom","value":{"sender":"iaa12v3r0unp6nprp9zur0fn446n832dfag5l3w38e","denom":"testasset","schema":"{\"type\":\"object\",\"properties\":{\"name\":{\"type\":\"string\"}}}"}}],"fee":{"amount":[],"gas":"200000"},"signatures":[{"pub_key":"61rphyECERx9oolgDwPYd7r8rqra8RRY9coPnNLrrHvmf1xhnoY=","signature":"+j1wRbFunzm7nOGTe5/eOZEJAugkTYpqMa6YK6pf2ctNty6/XNCba7oYSgqjLO4CiohVKnLZPPnD9UPgJpednA=="}],"memo":""},"mode":"block"}' -H "accept: application/json" | jq
```

```json
{
  "height": "2733",
  "txhash": "FD0592E9B8130ECA28C172CAC784AE4B601060F5736C4BD33BF2BBF943866D2A",
  "raw_log": "<raw-log>",
  "logs": [
    {
      "msg_index": 0,
      "log": "",
      "events": [
        {
          "type": "issue_denom",
          "attributes": [
            {
              "key": "denom",
              "value": "testasset"
            }
          ]
        },
        {
          "type": "message",
          "attributes": [
            {
              "key": "action",
              "value": "issue_denom"
            },
            {
              "key": "module",
              "value": "nft"
            },
            {
              "key": "sender",
              "value": "iaa12v3r0unp6nprp9zur0fn446n832dfag5l3w38e"
            }
          ]
        }
      ]
    }
  ],
  "gas_wanted": "200000",
  "gas_used": "43238"
}
```
