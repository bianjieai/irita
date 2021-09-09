<!--
order: 3
-->

# 随机数

随机数模块允许你向 IRITA 发送随机数请求，查询随机数或待处理的随机数请求队列。

## 可用命令

|                             名称                             |                描述                |
| :----------------------------------------------------------: | :--------------------------------: |
| [request](https://www.irisnet.org/docs/zh/cli-client/random.html#iris-tx-random-request) |     请求具有可选块间隔的随机数     |
| [random](https://www.irisnet.org/docs/zh/cli-client/random.html#iris-query-random-random) |     使用ID查询链上生成的随机数     |
| [queue](https://www.irisnet.org/docs/zh/cli-client/random.html#iris-query-random-queue) | 查询随机数请求队列，支持可选的高度 |

## irita tx random request

请求一个随机数。

```bash
irita tx random request [flags]
```

**标志：**

| 名称，速记        | 类型   | 必须 | 默认  | 描述                                       |
| :---------------- | ------ | ---- | ----- | ------------------------------------------ |
| --block-interval  | uint64 |      | 10    | 请求的随机数将在指定的区块间隔后生成       |
| --oracle          | bool   |      | false | 是否使用 Oracle 方式                       |
| --service-fee-cap | string |      | ""    | 最大服务费用（如果使用 Oracle 方式则必填） |

### 请求一个随机数

向 IRIS Hub 发送随机数请求，该随机数将在`--block-interval`指定块数后生成。

```bash
# without oracle
irita tx random request --block-interval=100 --from=<key-name> --chain-id=irita --fees=0.3point

# with oracle
irita tx random request --block-interval=100 --oracle=true --service-fee-cap=1point --from=<key-name> --chain-id=irita --fees=0.3point
```

> **提示** 
>
> 如果交易已被执行，你将获得一个唯一的请求 ID，该 ID 可用于查询请求状态。你也可以通过[查询交易详情](/offchain_facilities/explorer.html)获取请求 ID。

##  irita query random random

使用ID查询链上生成的随机数。

```bash
irita query random random <request-id> [flags]
```

## irita query random queue

查询随机数请求队列，支持可选的高度。

```bash
irita query random queue <gen-height> [flags]
```

### 查询随机数请求队列

查询尚未处理的随机数请求，可指定将要生成随机数（或请求 Service）的区块高度。

```bash
irita query random queue 100000
```