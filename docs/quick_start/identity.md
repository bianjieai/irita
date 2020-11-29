<!--
order: 7
-->

# 创建去中心化身份

本教程将开发一个简单的应用来演示去中心化身份相关的操作，包括：

- 创建身份
- 更新身份
- 查询身份

有关`身份`的介绍请参考[这里](../core_modules/identity.md)。

>**_需求：_** 开发前请完成[准备工作](prepare.md)

## 开发步骤

### 初始化 SDK

参考[初始化 SDK](sdk_init.md)

### 定义身份变量

定义身份相关的变量。

```go
// 身份 ID
id := uuid.NewV4().Bytes()

// 身份公钥
pubKeyECDSA := "03576aac14e47fc165789df8f86268faaa8f012a9bfbdef3ae18c22a63cfe5eac0"
pubKeySM2 := "03281ce4ba0b8c97e5b1434f8f298b064f03d4c1d21aae9276065e170fc90a5d51"

// 身份凭证
credentials1 := "https://security.com/kyc/10001/"
credentials2 := "https://security.com/aml/10001/"
```

### 创建身份

调用 `Identity` 模块的 `CreateIdentity` 方法创建身份。

```go
// 构造 BaseTx
baseTx := types.BaseTx{
    From:     accountName,
    Gas:      uint64(gas),
    Fee:      fee,
    Memo:     "",
    Mode:     mode,
    Password: password,
}

// 构造创建身份请求
pubKeyInfo := identity.PubkeyInfo{
    PubKey: pubKeyECDSA,
    PubKeyAlgo: identity.ECDSA,
}
createIdentityReq := identity.CreateIdentityRequest{
    ID: id,
    PubkeyInfo: &pubKeyInfo,
    Certificate: "",
    Credentials: &credentials1,
}

// 创建身份
_, err := client.Identity.CreateIdentity(createIdentityReq, baseTx)
```

### 更新身份

调用 `Identity` 模块的 `UpdateIdentity` 方法更新身份。

```go
// 构造更新身份请求
pubKeyInfo = identity.PubkeyInfo{
    PubKey: pubKeySM2,
    PubKeyAlgo: identity.SM2,
}
updateIdentityReq := identity.UpdateIdentityRequest{
    ID: id,
    PubkeyInfo: &pubKeyInfo,
    Certificate: "",
    Credentials: &credentials2,
}

// 更新身份
_, err := client.Identity.UpdateIdentity(updateIdentityReq, baseTx)
```

### 查询身份

根据身份 `ID` 查询指定身份的信息。

```go
// 查询身份
res, err := client.Identity.QueryIdentity(id)
```

### 完整示例代码

此去中心化身份应用的完整代码如下：

```go
TODO
```
