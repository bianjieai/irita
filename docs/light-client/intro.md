---
order: 1
---

# 简介

IRITA API服务器也称为LCD（Light Client Daemon）。LCD实例是IRITA的轻节点。与IRITA全节点不同，它不会存储所有块并执行所有事务，这意味着它仅需要最少的带宽，计算和存储资源。在不信任模式下，它将追踪验证人集更改的过程，并要求全节点返回共识证明和Merkle证明。除非具有超过2/3投票权的验证者采取拜占庭式行为，否则LCD证明验证算法可以检测到所有潜在的恶意数据，这意味着ILCD实例可以提供与全节点相同的安全性。

lcd的默认主文件夹为`$HOME/.iritacli`。一旦启动LCD，它将创建两个目录`keys`和`trust-base.db`，密钥存储db位于`keys`中。`trust-base.db`存储所有受信任的验证人集合以及其他与验证相关的文件。

当LCD以非信任模式启动时，它将检查`trust-base.db`是否为空。如果为true，它将获取最新的块作为其信任基础，并将其保存在`trust-base.db`下。IRISLCD实例始终信任该基础。所有查询证明将在此信任的基础上进行验证，有关详细的证明验证算法，请参阅[tendermint lite](https://github.com/tendermint/tendermint/blob/master/docs/tendermint-core/light-client-protocol.md)。

## 基本功能

- 提供restful APIs并使用swagger-ui列出这些APIs。
- 验证查询证明

## 启动

启动LCD的子命令:

```bash
iritacli rest-server [flags]
```

标识:

| 标识       | 类型   | 默认值                  | 必须 | 描述                                       |
| ---------- | ------ | ----------------------- | ---- | ------------------------------------------ |
| chain-id   | string |                         | 是   | Tendermint节点的chain ID                   |
| home       | string | "$HOME/.iritacli"       |      | 配置home目录，key和proof相关的信息都存于此 |
| node       | string | "tcp://localhost:26657" |      | 全节点的rpc地址                            |
| laddr      | string | "tcp://localhost:1317"  |      | 侦听的地址和端口                           |
| trust-node | bool   | false                   |      | 是否信任全节点                             |
| max-open   | int    | 1000                    |      | 最大连接数                                 |

默认情况下，LCD不信任连接的完整节点。但是，如果确定所连接的完整节点是可信任的，则应使用`--trust-node`标识运行LCD：

```bash
iritacli rest-server --node=tcp://localhost:26657 --chain-id=irita --trust-node
```

要公开访问你的LCD实例，您需要指定`--ladder`：

```bash
iritacli rest-server --node=tcp://localhost:26657 --chain-id=irishub --laddr=tcp://0.0.0.0:1317 --trust-node
```

## REST APIs

一旦启动LCD，就可以在浏览器中打开<http://localhost:1317/swagger-ui/>，然后可以浏览可用的restful APIs。swagger-ui页面包含有关APIs功能和所需参数的详细说明。在这里，我们仅列出所有API并简要介绍其功能。

### Token模块的APIs

1. `POST /token/tokens`: 发行一个`Token`
2. `PUT /token/tokens/{symbol}`: 编辑一个已存在的`Token`
3. `POST /token/tokens/{symbol}/mint`: 增发`Token`到指定地址
4. `POST /token/tokens/{symbol}/transfer`: 转让`Token`的所有权
5. `GET /token/tokens/{symbol}`: 查询`Token`
6. `GET /token/tokens`: 查询指定所有者的`Token`集合
7. `GET /token/tokens/{symbol}/fee`: 查询发行和铸造指定`Token`的费用

### NFT模块的APIs

1. `POST /nft/nfts/mint`: 发行一个新的`NFT`资产
2. `PUT /nft/nfts/{denom}/{id}`: 编辑一个已经存在的`NFT`
3. `PUT /nft/nfts/{denom}/{id}/transfer`: 转让`NFT`
4. `POST /nft/nfts/{denom}/{id}/burn`: 销毁`NFT`
5. `GET /nft/nfts/{denom}/{id}`: 查询单个`NFT`
6. `GET /nft/nfts/supplies/{denom}`: 查询某个分类下`NFT`的数量
7. `GET /nft/nfts/owners/{address}`: 查询某人拥有的`NFT`
8. `GET /nft/nfts/collections/{denom}`: 查询某个分类下所有的`NFT`
9. `GET /nft/nfts/denoms`: 查询所有`NFT`的分类

### Service模块的APIs

1. `POST /service/definitions`: 定义一个新的服务
2. `GET /service/definitions/{service-name}`: 查询服务定义
3. `POST /service/bindings`: 绑定一个服务
4. `GET /service/bindings/{service-name}/{provider}`: 查询服务绑定
5. `GET /service/bindings{service-name}`: 查询服务绑定列表
6. `POST /service/providers/{provider}/withdraw-address`: 设置提取地址
7. `GET /service/providers/{provider}/withdraw-address`: 查询提取地址
8. `PUT /service/bindings/{service-name}/{provider}`: 更新一个存在的服务绑定
9. `POST /service/bindings/{service-name}/{provider}/disable`: 禁用一个可用的服务绑定
10. `POST /service/bindings/{service-name}/{provider}/enable`: 启用一个不可用的服务绑定
11. `POST /service/bindings/{service-name}/{provider}/refund-deposit`: 取回一个服务绑定的所有押金
12. `POST /service/contexts`: 发起服务调用
13. `GET /service/contexts/{request-context-id}`: 查询请求上下文
14. `PUT /service/contexts/{request-context-id}`: 更新请求上下文
15. `POST /service/contexts/{request-context-id}/pause`: 暂停一个正在进行的请求上下文
16. `POST /service/contexts/{request-context-id}/start`: 启动一个暂停的请求上下文
17. `POST /service/contexts/{request-context-id}/kill`: 终止请求上下文
18. `GET /service/requests/{request-id}`: 查询服务请求
19. `GET /service/requests/{service-name}/{provider}`: 查询一个服务绑定的活跃请求
20. `GET /service/requests/{request-context-id}/{batch-counter}`: 根据请求上下文ID和批次计数器查询请求列表
21. `POST /service/responses`: 响应服务请求
22. `GET /service/responses/{request-id}`: 查询服务响应
23. `GET /service/responses/{request-context-id}/{batch-counter}`: 根据请求上下文ID和批次计数器查询服务响应列表
24. `GET /service/fees/{provider}`: 查询服务提供者的收益
25. `POST /service/fees/{provider}/withdraw`: 提取服务提供者的收益

### Record模块的APIs

1. `POST /record/records`: 创建一条记录
2. `GET /record/records/{record-id}`: 通过`id`查询指定的记录
