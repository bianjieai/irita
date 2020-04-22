# Token管理

## 简介

该规范描述了`IRITA`上的资产管理，任何人都可以在`IRITA`上发布新资产。本文章所涉及的`Token`指代都是`Fungible Token`。

## 概念

### 资产(Token)

`IRITA` 允许个人和公司创建和发行他们自己的资产，用于他们可以想象的任何场景。内部资产的潜在用例是数不胜数的。一方面，可用作存放在客户手机上的门票，以通过音乐会的入口。另一方面，它们可用于众筹，权益跟踪甚至以股票形式出售公司股权。

想要创建新的资产，您需要做的仅仅是执行一行命令，为您的资产定义初始化参数，例如总量，符号，描述等。然后，您可以发送一些您自己发行的 Token 到任何人的账户，就像 转账一样简单。作为该资产的所有者，您无需处理区块链的任何技术细节，例如分布式共识算法，区块链开发或集成，而且也不需要运行任何挖矿设备或服务器。

### 费用

#### 相关参数

| name              | Type | Default   | Description                     |
| ----------------- | ---- | --------- | ------------------------------- |
| TokenTaxRate      | Dec  | 0.4       | 资产税率，即Community Tax的比例 |
| IssueTokenBaseFee | Coin | 60000iris | 发行Token的基准费用             |
| MintTokenFeeRatio | Dec  | 0.1       | 增发Token的费率(相对于发行费用) |

注：以上参数均为可治理参数

#### 发行 Token 费用

- 基准费用：发行Token所需的基本费用，即Token的`Symbol`长度为最小(3)时的费用
- 费用因子计算公式：(ln(len({symbol}))/ln3)^4
- 发行FT费用计算公式：IssueTokenBaseFee/费用因子；结果取整（大于1时四舍五入，小于等于1时取值为1）

#### 增发 Token 费用

- 增发Token费率：相对于发行FT时的费率
- 增发Token费用计算公式：发行FT费用 * MintTokenFeeRatio；结果取整（大于1时四舍五入，小于等于1时取值为1）

#### 费用扣除方式

- Community Tax：资产相关的操作费用一部分将作为Community Tax，比例由`TokenTaxRate`决定。
- Burned：剩余部分将被销毁

## 操作

`Token`模块的相关命令请参考[文档](../cli-client/token.md)

