<!--
order: 3
-->

# 发行工分

本教程将演示工分模块的下列功能：

- 发行工分
- 增发工分
- 转让工分
- 查询工分
- 查询账户工分

有关`工分`的介绍请参考[这里](../core_modules/token.md)。

>**_需求：_** 开发前请完成[准备工作](prepare.md)。

## 开发步骤

### 初始化 SDK

参考[初始化 SDK](sdk_init.md)

### 定义工分变量

```go
// 定义工分变量
symbol := "credit"
name := "test credit"
minUnit := "credit"
scale := 18
initialSupply := 10000
maxSupply := 100000
mintable := true
```

### 工分发行

调用 `Token` 模块的 `IssueToken` 方法发行工分。

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

// 构造发行工分请求
issueTokenReq := token.IssueTokenRequest(
    Symbol: symbol,
    Name: name,
    MinUnit: minUnit,
    Scale: scale,
    InitialSupply: initialSupply,
    MaxSupply: maxSupply,
    Mintable: mintable,
}

// 发行工分
_, err := client.Token.IssueToken(issueTokenReq, baseTx)
```

### 查询工分

根据 `Symbol` 查询发行的工分。

```go
// 查询指定的工分
res, err := client.Token.QueryToken(symbol)
```

### 增发工分

调用 `Token` 模块的 `MintToken` 方法进行增发。默认将增发到工分 Owner。

```go
// 构造增发工分请求
mintTokenReq := token.MintTokenRequest{
    Amount: 1000,
}

// 增发工分
_, err := client.Token.MintToken(mintTokenReq, baseTx)
```

### 转让工分所有权

调用 `Token` 模块的 `TransferToken` 方法转让工分所有权。

```go
// 构造转让工分请求
recipient := "iaa18up8anyjpal8rncm8rd4ukp5f7etga795gp33q"
transferTokenReq := token.TransferTokenRequest(
    Symbol: symbol,
    Recipient: recipient,
}

// 转让工分
_, err := client.Token.TransferToken(transferTokenReq, baseTx)
```

### 查询账户工分

查询 recipient 的全部工分。

```go
// 查询账户的全部工分
res, err := client.Token.QueryTokens(recipient)
```

## 完整示例代码

以下是此工分应用的完整代码：

```go
TODO
```
