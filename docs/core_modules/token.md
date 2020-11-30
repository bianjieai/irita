<!--
order: 1
-->

# 积分

## 简介

积分模块旨在将各种权益通证化，从而让联盟成员能够基于 IRITA 平台进行公平、透明以及可追溯的交互与协作。数字化的积分代表着特定的权利，被不可篡改地记录在 IRITA 链上。

积分是一种可替代性（同质化）通证。

积分主要包括以下一些属性：

- _符号和名称_
  - 符号是积分的唯一标识符
  - 名称是积分的描述性名字

- _供应量相关_
  - 初始供应量：初次发行的积分数量
  - 最大供应量：积分的总量
  - 可增发性：初次发行后是否可以增发

- _分割性_：即允许积分拥有的最大小数位

- _最小单位_：积分在 IRITA 平台的存储和交易单位

## 功能

### 发行

通过指定积分相关参数，即可发行积分。

`CLI`

```bash
irita tx token issue --symbol=<symbol> --name=<name> --initial-supply=<initial-supply> --max-supply=<max-supply> --scale=<decimals> --min-unit=<min-unit> --mintable=<mintable>
```

### 查询

在发行积分之后，可通过如下命令查询：

`CLI`

```bash
irita query token <symbol>
```

### 增发

如果积分在发行时指定为可增发（`mintable` 为 `true`），则可进行增发操作：

`CLI`

```bash
irita tx token mint <symbol> --amount=<amount>
```

### 编辑

可对已发行积分的相关属性进行更新。可更新属性包括：积分名、最大供应量以及可增发性。

`CLI`

```bash
irita tx token edit <symbol> --name=<name> --max-supply=<max-supply> --mintable=<mintable>
```

### 转让所有权

积分所有权可以进行转让。

`CLI`

```bash
irita tx token transfer <symbol> --to=<new-owner>
```

## 费用

### 相关参数

| 名称              | 类型 | 默认值   | 描述                     |
| ----------------- | ---- | --------- | ------------------------------- |
| TokenTaxRate      | Dec  | 0.4       | 积分税率，即进行积分发行与增发时的 `Community Tax` 比例 |
| IssueTokenBaseFee | Coin | 60000point | 发行积分的基准费用，即 `symbol`（积分唯一标识）长度为最小（3）时的费用 |
| MintTokenFeeRatio | Dec  | 0.1       | 增发积分的费率（相对于发行费用）|

> **_提示：_** 以上参数是可以更改的系统参数。

### 发行费用

- 费用因子计算公式：`(ln(len({symbol}))/ln3)^4`
- 发行积分费用计算公式：`IssueTokenBaseFee` / 费用因子；结果取整到 `point`（大于1时四舍五入，小于等于1时取值为1）

### 增发费用

- 增发积分费用计算公式：发行积分费用 * `MintTokenFeeRatio`；结果取整到 `point`（大于1时四舍五入，小于等于1时取值为1）

### 费用扣除方式

- Community Tax：积分相关的操作费用一部分将作为 Community Tax，比例由 `TokenTaxRate` 决定
- 销毁：剩余部分将被销毁
