<!--
order: 8
-->

# 管理智能合约-WASM（即将下线）

本教程将开发一个简单的应用来演示如何管理智能合约，包括：

- 上传合约代码
- 初始化合约状态
- 执行合约
- 升级合约
- 查询合约中的状态信息

本文档不包括智能合约代码的编写内容，有关`智能合约`的介绍请参考[这里](../core_modules/wasm.md)。

>**_需求：_** 开发前请完成[准备工作](prepare.md)。

## 开发步骤

### 初始化 SDK

参考[初始化 SDK](sdk_init.md)

### 上传合约代码

```go
baseTx := types.BaseTx{
    From:     iaccountName,
    Gas:      4000000,
    Fee:      types.NewDecCoins(types.NewInt64DecCoin("point",24)),
    Memo:     "test",
    Mode:     types.Commit,
    Password: accountPwd,
}
// 构建上传合约请求
request := wasm.StoreRequest{
    WASMFile: "./election.wasm",
}

codeID, err := client.WasmClient().Store(request, baseTx)
```

### 初始化合约状态

```go
args := wasm.NewArgs().
    Put("start", 1).
    Put("end", 100).
    Put("candidates", []string{"iaa1qvty8x0c78am8c44zv2n7tgm6gfqt78j0verqa", "iaa1zk2tse0pkk87p2v8tcsfs0ytfw3t88kejecye5"})

initReq := wasm.InstantiateRequest{
    CodeID:  codeID,
    Label:   "test wasm",
    InitMsg: args,
}

contractAddress, err := client.WasmClient().Instantiate(initReq, baseTx)
```

### 执行合约

下面以执行合约中的`vote`方法为例，参数`candidate`='iaa1qvty8x0c78am8c44zv2n7tgm6gfqt78j0verqa'

```go
execAbi := wasm.NewContractABI().
    WithMethod("vote").
    WithArgs("candidate", "iaa1qvty8x0c78am8c44zv2n7tgm6gfqt78j0verqa")
_, err = client.WasmClient().Execute(contractAddress, execAbi, nil, baseTx)
```

### 查询合约中的状态信息

在演示的合约中，我们定义了一个`get_vote_info`的查询方法，下面是调用示例：

```go
queryAbi := wasm.NewContractABI().
    WithMethod("get_vote_info")
bz, err := client.WasmClient().QueryContract(contractAddress, queryAbi)
```

`QueryContract`返回的结果是JSON编码的，需要用户自己去解码。
