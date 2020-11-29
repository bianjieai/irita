<!--
order: 1
-->

# 概览

## 简介

控制台是 IRITA 平台提供的交互式客户端。通过此控制台，联盟成员可以便捷地与 IRITA 网络进行交互。控制台提供了所有必要的命令，覆盖了从节点、链、模块到管理的各式命令以供完成多种操作：获取节点状态、查询链上交易、查询各模块的状态、调用模块功能以及进行链的管理等。

## 安装

控制台功能集成在节点二进制 `irita` 中，[安装 IRITA](../installation/installation.md) 后即可使用。

## 命令通用标志

控制台命令支持一些通用的标志，包括全局标志以及分别特定于交易和查询的标志。

### 全局标志

控制台所有命令均可附加以下全局标志：

| 名称，速记        | 类型    | 描述                                                           |
| ----------------- | ------- | -------------------------------------------------------------- |
| --help，-h        |         | 打印命令帮助信息                                               |
| --node            | string  | host:port，节点的 RPC 接口地址，默认为 "tcp://localhost:26657" |
| --trust-node      | boolean | 信任连接的节点（不验证响应结果的证明），默认为真               |
| --chain-id        | string  | 节点的链 ID                                                    |
| --keyring-backend | string  | 指定 keyring 后端，支持 "os"、"file" 和 "test"，默认为 "os"    |
| --ledger          | boolean | 使用一个已连接的 ledger 设备                                   |
| --encoding，-e    | string  | 二进制编码方式，支持 "hex"、"b64" 和 "btc"，默认为 "hex"       |
| --home            | string  | 节点配置和数据的目录                                           |
| --output，-o      | string  | 输出格式，支持 "text" 和 "json"，默认为 "text"                 |
| --indent          | boolean | 为 JSON 输出增加缩进                                           |
| --trace           | boolean | 打印完整的错误栈追踪                                           |

### 交易命令通用标志

交易命令的通用标志如下：

| 名称，速记           | 类型    | 描述                                                                                      |
| -------------------- | ------- | ----------------------------------------------------------------------------------------- |
| --from               | string  | 签名私钥的名字或地址                                                                      |
| --fees               | string  | 支付的交易费用，例如 10point                                                              |
| --gas                | string  | 每个交易的 gas 限制；设置到 "auto" 将自动计算所需的 gas；默认为 200000                    |
| --gas-adjustment     | float   | gas 调整因子，即交易模拟器所估算gas 的倍数，默认为1；如果手动设置了 `--gas`，则忽略此标志 |
| --gas-prices         | string  | 决定交易费用的 gas 价格 (例如 10point)                                                    |
| --generate-only      | boolean | 构建未签名交易并写到标准输出（启用时，本地私钥不可访问，节点进行离线操作）                |
| --dry-run            | boolean | 如启用，将忽略 `--gas` 并模拟执行交易，但不广播                                           |
| --memo               | string  | 附加在交易中的备注                                                                        |
| --account-number，-a | uint    | 签名账户的 account number（仅用于离线模式）                                               |
| --sequence，-s       | uint    | 签名账户的 sequence number（仅用于离线模式）                                              |
| --broadcast-mode，-b | string  | 交易广播模式，支持 "sync"、"async" 和 "block"，默认为 "sync"                              |
| --offline            | boolean | 离线模式（将禁用任何在线功能）                                                            |
| --yes，-y            | boolean | 跳过交易广播确认提示                                                                      |

### 查询命令通用标志

查询命令的通用标志如下：

| 名称，速记 | 类型 | 描述                                   |
| ---------- | ---- | -------------------------------------- |
| --height   | int  | 指定状态查询的高度（不适用于快照模式） |

## 命令结果

交易命令结果包含以下主要字段：

| 字段       | 类型   | 描述                                                                                    |
| ---------- | ------ | --------------------------------------------------------------------------------------- |
| height     | string | 交易所在区块高度                                                                        |
| txhash     | string | 交易 hash                                                                               |
| raw_log    | string | 原生未解析的日志                                                                        |
| logs       | array  | 解析后的日志数组，包括各消息触发的事件以及其他元数据，仅当交易广播模式为 `block` 时可用 |
| gas_wanted | string | 交易的 gas 上限，仅当交易广播模式为 `block` 时可用                                      |
| gas_used   | string | 交易消耗的 gas，仅当交易广播模式为 `block` 时可用                                       |

> **_注意：_** 文档中所列交易命令示例的返回结果中均省略了 `raw_log` 的具体数据。

## 命令列表

`-h` 标志将打印出控制台支持的所有命令：

```bash
irita -h
```

控制台命令主要包括节点与链、核心模块以及管理等命令。

### 节点和链

| 名称                       | 描述                           |
| -------------------------- | ------------------------------ |
| [version](node.md#version) | 查询 app 版本                  |
| [status](node.md#status)   | 查询远程节点的状态             |
| [block](node.md#block) | 查询最新区块 |
| [tendermint-validator-set](node.md#tendermint-validator-set) | 查询验证人集合 |
| [tx](node.md#tx)           | 根据交易 Hash 查询交易        |

### 核心模块

#### 积分命令

| 名称                                  | 描述                                 |
| ------------------------------------- | ------------------------------------ |
| [issue](modules/token.md#issue)       | 发行积分                             |
| [edit](modules/token.md#edit)         | 编辑存在的积分                       |
| [mint](modules/token.md#mint)         | 增发积分到指定账户                   |
| [transfer](modules/token.md#transfer) | 转让积分所有权                       |
| [token](modules/token.md#token)       | 查询指定的积分                       |
| [tokens](modules/token.md#tokens)     | 查询所有积分或者指定所有者的积分列表 |
| [fee](modules/token.md#fee)           | 查询积分发行和增发费用               |

#### 数字资产建模命令

| 名称                                    | 描述                   |
| --------------------------------------- | ---------------------- |
| [issue](modules/nft.md#issue)           | 发行资产               |
| [mint](modules/nft.md#mint)             | 创建资产               |
| [edit](modules/nft.md#edit)             | 编辑已发行的资产       |
| [transfer](modules/nft.md#transfer)     | 转移指定资产           |
| [burn](modules/nft.md#burn)             | 销毁指定资产           |
| [supply](modules/nft.md#supply)         | 查询指定类别资产的总量 |
| [owner](modules/nft.md#owner)           | 查询指定账户的所有资产 |
| [collection](modules/nft.md#collection) | 查询指定类别的所有资产 |
| [denom](modules/nft.md#denom)           | 查询指定资产类别的定义 |
| [denoms](modules/nft.md#denoms)         | 查询所有资产类别       |
| [token](modules/nft.md#token)           | 查询指定资产           |

#### 存证命令

| 名称                                | 描述         |
| ----------------------------------- | ------------ |
| [create](modules/record.md#create)  | 创建存证     |
| [record](modules/record.md#record) | 查询指定存证 |

#### iService 命令

| 名称                                                       | 描述                                             |
| ---------------------------------------------------------- | ------------------------------------------------ |
| [define](modules/iservice.md#define)                       | 定义一个新的服务                                 |
| [bind](modules/iservice.md#bind)                           | 绑定一个服务                                     |
| [update-binding](modules/iservice.md#update-binding)       | 更新一个存在的服务绑定                           |
| [disable](modules/iservice.md#disable)                     | 禁用一个可用的服务绑定                           |
| [enable](modules/iservice.md#enable)                       | 启用一个不可用的服务绑定                         |
| [refund-deposit](modules/iservice.md#refund-deposit)       | 取回服务绑定的押金                               |
| [call](modules/iservice.md#call)                           | 发起服务调用                                     |
| [respond](modules/iservice.md#respond)                     | 响应指定的服务请求                               |
| [update](modules/iservice.md#update)                       | 更新请求上下文                                   |
| [pause](modules/iservice.md#pause)                         | 暂停一个正在进行的请求上下文                     |
| [start](modules/iservice.md#start)                         | 启动一个暂停的请求上下文                         |
| [kill](modules/iservice.md#kill)                           | 永久终止请求上下文                               |
| [set-withdraw-addr](modules/iservice.md#set-withdraw-addr) | 设置提取地址                                     |
| [withdraw-fees](modules/iservice.md#withdraw-fees)         | 提取服务费                                       |
| [definition](modules/iservice.md#definition)               | 查询服务定义                                     |
| [bindings](modules/iservice.md#bindings)                   | 查询指定服务的所有绑定或者指定 owner 的绑定列表  |
| [binding](modules/iservice.md#binding)                     | 查询服务绑定                                     |
| [request](modules/iservice.md#request)                     | 通过请求 ID 查询服务请求                         |
| [requests](modules/iservice.md#requests)                   | 通过服务绑定或请求上下文查询活跃的服务请求       |
| [response](modules/iservice.md#response)                   | 通过请求 ID 查询服务响应                         |
| [responses](modules/iservice.md#responses)                 | 通过请求上下文 ID 和批次计数器查询活跃的服务响应 |
| [request-context](modules/iservice.md#request-context)     | 查询请求上下文                                   |
| [withdraw-addr](modules/iservice.md#withdraw-addr)         | 查询服务费提取地址                               |
| [fees](modules/iservice.md#fees)                           | 查询指定服务提供者赚取的服务费                   |
| [schema](modules/iservice.md#schema)                       | 通过 schema 名称查询系统 schema                  |

#### 身份命令

| 名称                                | 描述         |
| ----------------------------------- | ------------ |
| [create](modules/identity.md#create)  | 创建身份     |
| [update](modules/identity.md#update)  | 更新身份     |
| [identity](modules/identity.md#identity) | 查询指定身份 |

### 管理

| 名称                                           | 描述         |
| ---------------------------------------------- | ------------ |
| [add-roles](admin.md#add-roles)     | 增加权限   |
| [remove-roles](admin.md#remove-roles) | 移除权限   |
| [block-account](admin.md#block-account)                 | 将指定账户加入黑名单   |
| [unblock-account](admin.md#unblock-account)           | 将指定账户移出黑名单   |
| [update](admin.md#update)       | 修改系统参数 |

### 其他

#### 轻客户端

启动 IRITA 轻客户端：

在节点配置文件 `config/app.toml` 中设置 `api.enable` 为 `true`，节点启动时将开启 `api` 服务，即轻客户端。
