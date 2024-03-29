<!--
order: 4
-->

# 存证

## 简介

存证用于将任何链外数据可信地映射在 IRITA 链上，作为对原始数据的证明。其可信性是通过密码学算法与区块链的安全性来保证的。

数据存证分为链外和链上两个过程：

- 将数据通过密码学安全的摘要算法进行处理，产生唯一的数据摘要。安全摘要算法包括 `SHA256`、`SHA512`、`SHA3` 等。

- 将上述摘要及其算法名称、原生数据或其 `URI` 存储在链上，以用于数据的真实性验证。

存证包含以下属性：

- _数据摘要_：数据的密码学证明

- _摘要算法_：用于摘要生成的密码学算法名称

- _元数据_：欲存证的原生数据，可直接存储在链上

- _元数据 URI_：元数据在链外存储的 URI

## 功能

### 创建

创建存证需要提供数据摘要、摘要算法、元数据或元数据的 URI。成功创建后，将产生此存证唯一的 ID。

`CLI`

```bash
irita tx record create [digest] [digest-algo] --uri=<metadata-uri> --meta=<metadata>
```

### 查询

根据存证 ID 可查询相应的存证记录。

`CLI`

```bash
irita query record record [record-id]
```
