<!--
order: 5
-->

# 区块链存证应用

本教程将开发一个简单的存证应用来演示存证相关的操作，包括存证的创建和查询。

有关`存证`的介绍请参考[这里](../core_modules/record.md)。

>**_需求：_** 开发前请完成[准备工作](prepare.md)

## 开发步骤

### 初始化 SDK

参考[初始化 SDK](sdk_init.md)

### 定义存证

定义存证相关的变量。

```go
// 存证元数据
recordMetadata := `{"data":"hello,world"}`

// 存证元数据摘要
digestAlgo := "SHA256"
recordDigestBz := tmhash.Sum([]byte(recordMetadata))
recordDigest := hex.EncodeToString(recordDigestBz)

// 生成存证内容
recordContent := record.Content{
    Digest: recordDigest,
    DigestAlgo: digestAlgo,
    URI: "",
    Meta: recordMetadata,
}
```

### 创建存证

调用 `Record` 模块的 `CreateRecord` 方法发起创建存证交易。

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

// 构造创建存证请求
createRecordReq := record.CreateRecordRequest{
    Contents: []record.Content{recordContent},
}

// 创建存证
recordID, err := client.Record.CreateRecord(createRecordReq, baseTx)
```

### 查询存证

根据 `recordID` 查询存证信息。

```go
// 查询存证
res, err := client.Record.QueryRecord(recordID)
```

### 完整示例代码

完整的存证示例应用代码如下：

```go
TODO
```
