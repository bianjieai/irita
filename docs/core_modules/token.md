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

```bash
 irita tx token issue [flags]
 irita tx token issue --name="Kitty Token" --symbol="kitty" --min-unit="kitty" --scale=0 --initial-supply=100000000000 --max-supply=1000000000000 --mintable=true --from=<key-name> --chain-id=<chain-id> --fees=<fee>
```

**标志：**

| 名称，速记       | 类型    | 必须 | 默认          | 描述                                                         |
| :--------------- | :------ | :--- | :------------ | :----------------------------------------------------------- |
| --name           | string  | 是   |               | 积分的名称；unicode字符，最大长度为32字节，例如 "IRITA Credit" |
| --symbol         | string  | 是   |               | 积分的唯一标识符；长度在3到8之间，字母数字字符，以字符开始，不区分大小写 |
| --initial-supply | uint64  | 是   |               | 积分的初始供应；增发前的数量不应超过1000亿                   |
| --max-supply     | uint64  |      | 1000000000000 | 积分的最大供应，总供应不能超过最大供应；增发前的数量不应超过1万亿 |
| --scale          | uint8   | 是   |               | 积分的精度，最多可以有18位小数；为0将默认到18位小数          |
| --min-unit       | string  | 是   |               | 最小单位；长度在3到10之间，字母数字字符，以字符开始，不区分大小写 |
| --mintable       | boolean |      | false         | 发行后是否可以增发                                           |

**使用示例 ：**

```bash
irita tx token issue --name="Kitty Token" --symbol="kitty" --min-unit="kitty" --scale=0 --initial-supply=100000000000 --max-supply=1000000000000 --mintable=true --from=node0 --chain-id=test  --home node0 -y
```

#### 		**查询**

在发行积分之后，可通过如下命令查询

`CLI`

```bash
irita query token [command]
```

**使用示例 ：**

```bash
 irita query token  token kitty 
 irita query token  tokens 
 //使用tokens可以查看所有的token
```

执行结果为 : 

```bash
'@type': /irismod.token.Token
initial_supply: "100000000000"
max_supply: "1000000000000"
min_unit: kitty
mintable: true
name: Kitty Token
owner: iaa1t07s27vgvgczpsvu5z75703azmmc9wcmje452z
scale: 0
symbol: kitty
```



### 增发

如果积分在发行时指定为可增发（`mintable` 为 `true`），则可进行增发操作：

`CLI`

```bash
irita tx token mint [symbol] [flags]
irita tx token mint <symbol> --amount=<amount> --to=<to> --from=<key-name> --chain-id=<chain-id> --fees=<fee>
```

**使用示例 ：**

```bash
irita tx token mint kitty --amount=100 --from node0 --home node0 --chain-id test -y
```

增发后可以在发行者名下查询到对应的数量

```bash
irita q bank balances iaa1n0t9jn2dmyzxedkxztpqsrdgq9wn3nhr3ukfw6
```

查询结果为 ：

```bash
balances:
- amount: "100000000100"
  denom: kitty
- amount: "9985684000000"
  denom: uirita
- amount: "10000000000000"
  denom: upoint
```

可以看到kitty的数量相较于初始化的数量增加了100



### 编辑

可对已发行积分的相关属性进行更新。可更新属性包括：积分名、最大供应量以及可增发性。

`CLI`

```bash
irita tx token edit [symbol] [flags]
irita tx token edit <symbol> --name="Cat Token" --max-supply=100000000000 --mintable=true --from=<key-name> --chain-id=<chain-id> --fees=<fee>
```

**标志：**

| 名称，速记   | 类型   | 必须 | 默认  | 描述                                                  |
| :----------- | :----- | :--- | :---- | :---------------------------------------------------- |
| --name       | string |      |       | 积分名称，为空将不更新                                |
| --max-supply | uint   |      | 0     | 积分的最大供应量，应不小于当前的总供应量，为0将不更新 |
| --mintable   | bool   |      | false | 积分是否可以增发，默认为false                         |

**使用示例 ：**

```bash
irita tx token edit kitty --name="Cat" --max-supply=100000000000000 --mintable=true --from=node0 --chain-id=test --home node0 -y
```

编辑完成之后再次查看：

```bash
 irita query token  token kitty
```

查询结果为：

```bash
'@type': /irismod.token.Token
initial_supply: "100000000000"
max_supply: "100000000000000"
min_unit: kitty
mintable: true
name: Cat
owner: iaa1t07s27vgvgczpsvu5z75703azmmc9wcmje452z
scale: 0
symbol: kitty
```

可以看到 name 已经由 Kitty Token 变为 Cat 



### 转让所有权

积分所有权可以进行转让。

`CLI`

```bash
irita tx token transfer [symbol] [flags]
irita tx token transfer <symbol> --to=<to> --from=<key-name> --chain-id=<chain-id> --fees=<fee>
```





## 费用使用示例 ：

```bash
irita tx token transfer kitty --to=iaa1j8ykt4yh2dzv42467vxjzq924z5paz4pk4m3d6  --from=node0 --chain-id=test --home node0  -y
```

转让前查询 kitty积分信息：

```bash
irita query token token kitty
```

查询结果 ：

```bash
'@type': /irismod.token.Token
initial_supply: "100000000000"
max_supply: "100000000000000"
min_unit: kitty
mintable: true
name: Cat
owner: iaa15t9d5gqfqdh86xkyygq5kheetaut5pz87gew6d
scale: 0
symbol: kitty
```

查询bank发现还是在发行者名下，其转让所有权仅仅将所有者owner更换

### 可能遇见的问题

问题 ：账户中uirita为0，积分发行需要13015000000uirita

解决方案:  在账户创建时为账户增加初始化uirita数量。



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
