<!--
order: 8
-->

# 智能合约-EVM

## 简介

Smart Contract，是一种旨在以信息化方式传播、验证或执行合同的计算机协议，是对区块链功能的一种可定制化的扩展。目前流行的智能合约实施方法主要是以太坊的`EVM`。 IRITA 对 `EVM` 进行了支持。

主要特征包括：

- 默认的 `EVM` 的 `chain-id` 为 `1223`；如果想要修改此端口，请在编译前修改 makefile 中的 `github.com/tharsis/ethermint/types.EvmChainID=<your_chain_id>`
- `EVM` 的相关的 `API` 端口是: `8545` 和 `8546`
- 默认是开启 `EVM`的相关功能。默认开启的 `namespace`  有：`"eth,net,web3"`
- 兼容 web3 相关的组件。例如： `metamask` 和 `Remix` 等其他相关的开发组件
- 其他配置项请参考 `app.toml` 的 `EVM Configuration`

### 注意事项

`EVM` 模块仅支持以 `eth_secp256k1` 算法生成的账户。 生成以 `eth_secp256k1` 算法账户的方式为：

```shell
irita keys add [account_name] --algo eth_secp256k1
```

## 功能

### API 相关功能

IRITA 支持了 `EVM` 的所有功能。相关的 `API` 使用文档，可以参考：[EVM API](https://eth.wiki/json-rpc/API)

### 导出账户的 `ETH` 私钥

```shell
irita keys unsafe-export-eth-key [name]
```

### 导入账户的 `ETH` 私钥

```shell
irita keys unsafe-import-eth-key [name] [pk]
```

### 获取智能合约的 `code`

允许用户在给定地址查询智能合约代码。

```shell
irita query evm code [address] [flags]
```

### 获取智能合约的 `code`

允许用户使用给定的 `key` 和 `height` 查询 `address` 的存储。

```shell
irita query evm storage [address] [key] [flags]
```

### 发送交易

允许用户从原始的 `ETH` 交易构建 `Cosmos` 交易。

```shell
irita tx evm raw [tx-hex] [flags]
```