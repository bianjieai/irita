<!--
order: 5
-->

# 身份

## 简介

身份模块构建了一个去中心化身份体系（`DID`），实现并扩展了 [W3C DID 规范](https://www.w3.org/TR/did-core/)。

主要特征包括：

- DID 方法为 `irita`，完整的身份 DID 形式表示为：did:irita:id

- 身份的密码学材料包括公钥以及公钥证书

- 身份可以包含额外的凭证信息

一个身份由以下几个部分组成：

- _ID_：全局唯一的身份标识符

- _公钥列表_：身份主体的公钥列表

- _公钥证书列表_：身份主体的公钥证书列表

- _身份凭证 URI_：身份主体链外凭证信息的 URI

## 功能

### 创建

提供身份 ID、公钥、证书以及凭证 URI，即可创建一个身份。上述参数均为可选。如身份 ID 未指定，将自动生成。

当前支持的公钥算法包括：

- `RSA`：`DER` 编码的公钥
- `DSA`：`DER` 编码的公钥
- `ECDSA`：33字节的压缩公钥
- `ED25519`：32字节的压缩公钥
- `SM2`：33字节的压缩公钥

所有公钥均采用 `Hex` 字符串表示。

`CLI`

```bash
irita tx identity create --id=<identity-id> --pubkey=<public-key> --pubkey-algo=<public-key-algorithm> --cert-file=<certificate-file-path> --credentials=<credentials-uri>
```

### 更新

更新指定的身份。更新操作包括：增加公钥、增加公钥证书以及更改凭证 URI。

`CLI`

```bash
irita tx identity update <identity-id> --pubkey=<public-key> --pubkey-algo=<public-key-algorithm> --cert-file=<certificate-file-path> --credentials=<new-credentials-uri>
```

### 查询

查询指定的身份。

`CLI`

```bash
irita query identity identity <id>
```
