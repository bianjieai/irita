<!--
order: 2
-->

# 节点与链

控制台提供了节点与链的操作命令，可以通过这些命令查询节点状态、最新区块以及交易等。

## 查询 app 版本

查询 app 的版本信息。

```bash
irita version [flags]
```

### 查询 app 版本示例

```bash
irita version
```

结果

```text
0.5.0-12-gd06c1ae
```

## 查询节点状态

查询远程节点的状态。

```bash
irita status [flags]
```

### 查询节点状态示例

```bash
irita status -o json --indent
```

结果

```json
{
  "node_info": {
    "protocol_version": {
      "p2p": "7",
      "block": "10",
      "app": "0"
    },
    "id": "cc7adb370b3d20d83c0c39be7322909ae8ef1258",
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
  "sync_info": {
    "latest_block_hash": "5CA1B3634D1984C57B6073B1CDDD2C9702906C9393D5AD1FDA38BAD3BA6FFEC0",
    "latest_app_hash": "9F431407917DB5B407E2DE38726103ACB27DFAD0836D18F2EED74387D77D7E12",
    "latest_block_height": "1932",
    "latest_block_time": "2020-06-18T06:22:05.281994Z",
    "earliest_block_hash": "18BC6E63B681920F1EFB209C1681AD21877D4F8F557800EBAE5B0F5F472FF5B2",
    "earliest_app_hash": "",
    "earliest_block_height": "1",
    "earliest_block_time": "2020-05-28T04:00:21.211932Z",
    "catching_up": false
  },
  "validator_info": {
    "address": "15325D5713194219CD9297E1E37CAC87B76852B2",
    "pub_key": {
      "type": "tendermint/PubKeySm2",
      "value": "AFjXg1oEURbDV69h5gn7Ou+aklM0zgir0K0qybTd/kjZ"
    },
    "voting_power": "100"
  }
}
```

## 查询区块

查询指定高度的区块数据。如未指定高度，则查询最新区块。

```bash
irita query block [height] [flags]
```

**参数：**

| 名称      | 类型    | 必须 | 默认          | 描述                                                                     |
| ---------------- | ------- | ---- | ------------- | ------------------------------------------------------------------------ |
| height  | uint64  | 否   |             | 区块高度，默认为最新区块高度 |

### 查询区块示例

```bash
irita query block -o json --indent --chain-id=irita-test
```

结果

```json
```

## 查询交易

通过指定的交易 Hash 查询交易。

```bash
irita query tx [hash] [flags]
```

**参数：**

| 名称      | 类型    | 必须 | 默认          | 描述                                                                     |
| ---------------- | ------- | ---- | ------------- | ------------------------------------------------------------------------ |
| hash  | string  | 是   |             | 交易 Hash |

### 查询交易示例

```bash
irita query tx 497921D9C76C7FE17960CC3C3072ED05367B136036620ABB5E50A562A9DAC532 --chain-id=irita-test
```

结果

```json
```
