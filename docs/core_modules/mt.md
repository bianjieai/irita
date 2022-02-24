<!--
order: 7
-->

# 资产批量数字化建模

## 简介

资产数字化建模为联盟链成员提供了将资产进行批量数字化的能力。通过该模块，可以批量，每个链外资产将被建模为唯一的 IRITA 链上资产。

链上资产用 `ID` 进行标识，借助 IRITA 安全、不可篡改的特性，资产的所有权将得到明确。资产在成员间的交易过程也将被公开地记录，以便于追溯以及争议处理。

资产的元数据（`data`）可以直接存储在链上，也可以将其在链外存储源的地址记录在元数据中。资产元数据按照特定的 [JSON Schema](https://JSON-Schema.org/) 进行组织。[这里](./schemas/nft-metadata.md)是一个元数据 JSON Schema 示例。

资产在创建前需要发行，用以声明其抽象属性：

- _DenomID_：即全局唯一的资产类别标识符；此 ID 在链上生成

- _DenomName_:资产类别名称

- _元数据规范_：资产元数据应遵循的 JSON Schema

每一个具体的资产由以下元素描述：

- _DenomID_: 该资产的类别

- _ID_：资产的标识符，在此资产类别中唯一；此 ID 在链上生成

- _元数据_：包含资产具体数据的结构

## 功能

### 发行

指定资产 DenomName（资产名称）、创建者，即可创造资产空间。

`CLI`

```bash
irita tx mt issue --name=<denom-name> --from=<sender-address>
```

### 生产

在发行资产之后即可创建该类型的具体资产，需指定资产 DenomID、发型数量、元数据、发行者地址（Denom拥有者）和接收者地址。

`CLI`

```bash
irita tx mt mint <denom-id> --amount=<amount> --data=<data> --from=<sender-address> --recipient=<recipient-address>
```

### 增发

在发行资产之后即可增发该类型的具体资产，需指定资产 DenomID、增发数量、发行者地址（Denom拥有者）和接收者地址。

`CLI`

```bash
irita tx mt mint <denom-id> --mt-id=<mt-id> --amount=<amount> --from=<sender-address> --recipient=<recipient-address>
```

### 编辑

可对指定资产的元数据进行更新。

`CLI`

```bash
irita tx mt edit <denom-id> <mt-id> --data=<data> --from=<sender-address>
```

### 转移

转移指定资产；可以指定转移数量。

`CLI`

```bash
irita tx mt transfer <sender> <recipient> <denom-id> <mt-id> <amount>
```

### 销毁

销毁已创建的资产；可以指定销毁数量。

`CLI`

```bash
irita tx mt burn <denom-id> <mt-id> <amount> --from=<sender-address>
```

### 查询指定的资产类别

根据 DenomID 查询资产类别信息。

`CLI`

```bash
irita query mt denom <denom-id>
```

### 查询所有资产类别信息

查询已发行的所有资产类别信息。

`CLI`

```bash
irita query mt denoms
```

### 查询指定类别资产的总量

根据 DenomID 查询资产总量。

`CLI`

```bash
irita query mt supply <denom-id> <mt-id>
```

### 查询指定账户的所有资产

查询账户在指定资产类别中所拥有的全部资产。

`CLI`

```bash
irita query mt balances <owner> <denom-id>
```

### 查询指定资产

根据 DenomID 以及 MtID 查询具体资产信息。

`CLI`

```bash
irita query mt token <denom-id> <mt-id>
```

### 查询指定类别的所有资产

根据 DenomID 查询所有资产信息。

`CLI`

```bash
irita query mt tokens <denom-id>
```
