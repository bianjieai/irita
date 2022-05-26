<!--
order: 7
-->

# 智能合约-WASM（即将下线）

## 简介

Smart Contract，是一种旨在以信息化方式传播、验证或执行合同的计算机协议，是对区块链功能的一种可定制化的扩展。目前流行的智能合约实施方法主要是以太坊的[EVM](./evm.md)（IRITA已支持）。目前IRITA同时支持的另一种智能合约方案WASM(即将下线)，底层采用的是[Webassembly](https://www.wasm.com.cn/) 技术，理论上支持
Webassembly所支持的各种语言来编写智能合约代码，减少了使用其他语言编写智能合约的学习成本。由于**WASM**还处于快速发展阶段，vm底层代码采用的是Rust语言编写，所以目前支持的智能合约语言是**Rust**。<font color=#FF0000>当前已推出以太坊的[EVM](./evm.md)来替换即将下线的**WASM**</font>。

## 功能

### 上传合约代码

当合约代码编写完成后，将源码编译为.wasm格式，使用store命令来上传合约二进制文件到 irita 。

`CLI`

```bash
irita tx wasm store [wasm-file] --source [source] --builder [builder]
```

### 初始化合约状态

合约代码在初始化之前没有任何状态信息，可以使用`Instantiate`命令来完成合约的初始化工作。

`CLI`

```bash
irita tx wasm instantiate [code_id] [json_encoded_init_args] --label [text] --admin [address] --amount [coins] [flags]
```

### 执行合约

合约方法签名使用的是JSON编码（以太坊solidity采用的是ABI），可以根据的生成schema构造合约的方法调用参数。

`CLI`

```bash
irita tx wasm execute [contract_addr_bech32] [json_encoded_send_args] [flags]
```

### 升级合约

当合约代码中存在Bug时，可以通过合约管理者上传新的合约代码，然后通过`Migrate`来更新旧的合约，而且合约地址不会变更（以太坊大多采用的是代理合约的方式）。

`CLI`

```bash
irita tx wasm migrate [contract_addr_bech32] [new_code_id_int64] [json_encoded_migration_args] [flags]
```

### 查询合约中的状态信息

合约中一般会存储我们需要持久化的数据，可以使用`contract-state smart`命令查询当前合约中的某些数据。

`CLI`

```bash
irita query wasm contract-state smart [contract_addr_bech32] [json_encoded_query_args]
```
