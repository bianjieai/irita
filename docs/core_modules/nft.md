<!--
order: 2
-->

# 数字资产建模

## 简介

数字资产建模为联盟链成员提供了将资产进行数字化的能力。通过该模块，每个链外资产将被建模为唯一的 IRITA 链上资产。

链上资产用 `ID` 进行标识，借助 IRITA 安全、不可篡改的特性，资产的所有权将得到明确。资产在成员间的交易过程也将被公开地记录，以便于追溯以及争议处理。

资产的元数据（`metadata`）可以直接存储在链上，也可以将其在链外存储源的 `URI` 存储在链上。资产元数据按照特定的 [JSON Schema](https://JSON-Schema.org/) 进行组织。[这里](./schemas/nft-metadata.md)是一个元数据 JSON Schema 示例。

资产在创建前需要发行，用以声明其抽象属性：

- _Denom_：即全局唯一的资产类别标识符

- _元数据规范_：资产元数据应遵循的 JSON Schema

每一个具体的资产由以下元素描述：

- _Denom_: 该资产的类别

- _ID_：资产的标识符，在此资产类别中唯一；此 ID 在链外生成

- _元数据_：包含资产具体数据的结构

- _元数据 URI_：当元数据存储在链外时，此 URI 表示其存储位置

## 功能

### 发行

指定资产 Denom（资产类别）、元数据 JSON Schema，即可发行资产。

`CLI`

```bash
irita tx nft issue [denom] [flags]
irita tx nft issue <denom> --from=<key-name> --name=<denom-name> --schema=<schema-content or path to schema.json> --chain-id=<chain-id> --fees=<fee>
```

​			**参数：**

| 名称  | 类型   | 必须 | 默认 | 描述                                                        |
| :---- | :----- | :--- | :--- | :---------------------------------------------------------- |
| denom | string | 是   |      | 资产的类别，全局唯一；长度为3到64，字母数字字符，以字母开始 |

​			**标志：**

| 名称，速记 | 类型   | 必须 | 默认 | 描述                                                         |
| :--------- | :----- | :--- | :--- | :----------------------------------------------------------- |
| --schema   | string | 否   |      | 资产元数据 [JSON Schema (opens new window)](https://json-schema.org/)规范 |

​		**使用示例：**

```bash
irita tx nft issue security --name security --schema='{"type":"object","properties":{"name":{"type":"string"}}}' --from=node0 --chain-id=test -b=block   -y --home=node0
```



### 增发

在发行资产之后即可增发（创建）该类型的具体资产。需指定资产 ID、接收者地址、元数据或其 URI。

`CLI`

```bash
irita tx nft mint [denom] [token-id] [flags]
irita tx nft mint <denom> <token-id> --uri=<uri> --recipient=<recipient> --from=<key-name> --chain-id=<chain-id> --fees=<fee>
```

**参数：**

| 名称    | 类型   | 必须 | 默认 | 描述                                              |
| :------ | :----- | :--- | :--- | :------------------------------------------------ |
| denom   | string | 是   |      | 资产的类别；长度为3到64，字母数字字符，以字母开始 |
| tokenID | string | 是   |      | 资产的唯一 ID，如 `UUID`                          |

​		**标志：**

| 名称，速记  | 类型   | 必须 | 默认 | 描述                                 |
| :---------- | :----- | :--- | :--- | :----------------------------------- |
| --uri       | string | 否   |      | 资产元数据的 `URI`                   |
| --data      | string | 否   |      | 资产元数据                           |
| --recipient | string | 否   |      | 资产接收者地址，默认为交易发起者地址 |

​		**使用示例：**

```bash
irita tx nft mint security test --uri=https://test.com --data='{"name":"test security"}' --from=node0 --chain-id=test -b=block -y --home=node0
```



### 编辑

可对指定资产的元数据进行更新。

`CLI`

```bash
irita tx nft edit [denom] [token-id] [flags]
irita tx nft edit <denom> <token-id> --uri=<uri> --from=<key-name> --chain-id=<chain-id> --fees=<fee>
```

​	**参数：**

| 名称    | 类型   | 必须 | 默认 | 描述                                              |
| :------ | :----- | :--- | :--- | :------------------------------------------------ |
| denom   | string | 是   |      | 资产的类别；长度为3到64，字母数字字符，以字母开始 |
| tokenID | string | 是   |      | 资产的唯一 ID                                     |

​	**标志：**

| 名称，速记 | 类型   | 必须 | 默认 | 描述               |
| :--------- | :----- | :--- | :--- | :----------------- |
| --uri      | string | 否   |      | 资产元数据的 `URI` |
| --data     | string | 否   |      | 资产元数据         |

​	**使用示例：**

```bash
irita tx nft edit security test --data='{"name":"new test security"}' --from=node0 --chain-id=test -b=block  -y --home=node0
```





### 转移

转移指定资产。

`CLI`

```bash
irita tx nft transfer [recipient] [denom] [token-id] [flags]
irita tx nft transfer <recipient> <denom> <token-id> --uri=<uri> --from=<key-name> --chain-id=<chain-id> --fees=<fee>
```

**参数：**

| 名称      | 类型   | 必须 | 默认 | 描述                                              |
| :-------- | :----- | :--- | :--- | :------------------------------------------------ |
| recipient | string | 是   |      | 积分的唯一标识符                                  |
| denom     | string | 是   |      | 资产的类别；长度为3到64，字母数字字符，以字母开始 |
| tokenID   | string | 是   |      | 资产的唯一 ID                                     |

​		**使用示例：**

```bash
irita tx nft transfer  iaa177w2evwnx3uje646k78zxlp82mc9eatuwkdwlh<node1>   security a4c74c4203af41619d00bb3e2f462c10 --from=node0 --chain-id=test -b=block  -y --home=node0
```



### 销毁

可以销毁已创建的资产。

`CLI`

```bash
irita tx nft burn [denom] [token-id] [flags]
irita tx nft burn <denom> <token-id> --from=<key-name> --chain-id=<chain-id> --fees=<fee>
```

​	**参数：**

| 名称    | 类型   | 必须 | 默认 | 描述                                              |
| :------ | :----- | :--- | :--- | :------------------------------------------------ |
| denom   | string | 是   |      | 资产的类别；长度为3到64，字母数字字符，以字母开始 |
| tokenID | string | 是   |      | 资产的唯一 ID                                     |

​	**使用示例：**

```bash
irita tx nft burn security a4c74c4203af41619d00bb3e2f462c10 --from=iaa177w2evwnx3uje646k78zxlp82mc9eatuwkdwlh<node1>  --chain-id=test -b=block  -y --home=node0
```



### 查询指定的资产类别

根据 Denom 查询资产类别信息。

`CLI`

```bash
irita query nft denom [denom-id] [flags]
irita query nft denom <denom-id>
```

​	**参数：**

| 名称  | 类型   | 必须 | 默认 | 描述                                              |
| :---- | :----- | :--- | :--- | :------------------------------------------------ |
| denom | string | 是   |      | 资产的类别；长度为3到64，字母数字字符，以字母开始 |

​		**使用示例：**

```bash
irita query nft denom security  --chain-id=test 
```



### 查询所有资产类别信息

查询已发行的所有资产类别信息。

`CLI`

```bash
irita query nft denoms [flags]
irita query nft denoms
```

​	**使用示例：**

```bash
irita query nft denoms  --chain-id=test
```



### 查询指定类别资产的总量

根据 Denom 查询资产总量；接受可选的 owner 参数。

`CLI`

```bash
irita query nft supply [denom] [flags]
irita query nft supply <denom>
```

**参数：**

| 名称  | 类型   | 必须 | 默认 | 描述                                              |
| :---- | :----- | :--- | :--- | :------------------------------------------------ |
| denom | string | 是   |      | 资产的类别；长度为3到64，字母数字字符，以字母开始 |

**标志：**

| 名称，速记 | 类型   | 必须 | 默认 | 描述           |
| :--------- | :----- | :--- | :--- | :------------- |
| --owner    | string | 否   |      | 资产所有者地址 |

**使用示例：**

```bash
irita query nft supply security --owner iaa1t07s27vgvgczpsvu5z75703azmmc9wcmje452z  --chain-id=test
```





### 查询指定账户的所有资产

查询某一账户所拥有的全部资产；可以指定 Denom 参数。

`CLI`

```bash
irita query nft owner [address] [flags]
irita query nft owner <address> --denom-id=<denom>
```

​	**参数：**

| 名称    | 类型   | 必须 | 默认 | 描述         |
| :------ | :----- | :--- | :--- | :----------- |
| address | string | 是   |      | 目标账户地址 |

**标志：**

| 名称，速记 | 类型   | 必须 | 默认 | 描述                                              |
| :--------- | :----- | :--- | :--- | :------------------------------------------------ |
| --denom    | string | 是   |      | 资产的类别；长度为3到64，字母数字字符，以字母开始 |

**使用示例：**

```bash
irita query nft owner iaa1t07s27vgvgczpsvu5z75703azmmc9wcmje452z  --chain-id=test
```



### 查询指定类别的所有资产

根据 Denom 查询所有资产。

`CLI`

```bash
irita query nft collection [denom] [flags]
irita query nft collection <denom>
```

**参数：**

| 名称  | 类型   | 必须 | 默认 | 描述                                              |
| :---- | :----- | :--- | :--- | :------------------------------------------------ |
| denom | string | 是   |      | 资产的类别；长度为3到64，字母数字字符，以字母开始 |

**使用示例：**

```bash
irita query nft collection security  --chain-id=test
```



### 查询指定资产

根据 Denom 以及 ID 查询具体资产。

`CLI`

```bash
 irita query nft token [denom] [token-id] [flags]
 irita query nft token <denom> <token-id>
```

**参数：**

| 名称    | 类型   | 必须 | 默认 | 描述                                              |
| :------ | :----- | :--- | :--- | :------------------------------------------------ |
| denom   | string | 是   |      | 资产的类别；长度为3到64，字母数字字符，以字母开始 |
| tokenID | string | 是   |      | 资产的唯一 ID                                     |

**使用示例：**

```bash
 irita query nft token security a4c74c4203af41619d00bb3e2f462c10<tokenid>  --chain-id=test
```

