<!--
order: 9
-->

# 管理智能合约-EVM

本教程介绍三种不同的开发环境部署和管理 EVM 智能合约，包括：
- Remix
- Hardhat
- Truffle

## Remix

Remix 是 Solidity 智能合约的浏览器 IDE 。本指南中，我们将学习用 Remix 将智能合约部署到正在运行的 IRITA 网络，并与其进行交互。

### 连接 Remix

在连接 Remix 之前，需要将 IRITA 私钥导入到 Metamask，启动 IRITA 守护进程和 REST 服务。

访问 [Remix](http://remix.ethereum.org/)，文件资源管理器中已有一些合约。将它们替换为下方的源代码`Couter.sol`。在左边栏中，选择 Solidity Compiler 编译合约。

```solidity
pragma solidity >=0.7.0 <0.9.0;

contract Counter {
  uint256 counter = 0;

  function add() public {
    counter++;
  }

  function subtract() public {
    counter--;
  }

  function getCounter() public view returns (uint256) {
    return counter;
  }
}
```

接下来，选择 `Deploy and Run` 选项。选择 `Injected Web3` 作为 `Environment`。这将打开一个 Metamask 弹窗，让你连接 Metamask 到 Remix。选择 `Connect` 进行确认。

完成上述步骤后，可以在左侧面板看到自己的账户。

### 部署与交互

账户连接后，即可进行合约的部署。点击 `Deploy` 按钮，Metamask 窗口弹出确认请求。确认交易后，即可在 IRITA 守护进程日志中看到部署交易的日志：

```
I[2020-07-15|17:26:43.155] Added good transaction                       module=mempool tx=877A8E6600FA27EC2B2362719274314977B243671DC4E5F8796ED97FFC0CBE42 res="&{CheckTx:log:\"[]\" gas_wanted:121193 }" height=31 total=1
```

一旦合约成功部署，即可在左边栏的`Deployed Section`部分看到合约。Remix 控制台的绿色选框展示了交易的详细信息。

随后，即可与 Remix 进行交互。对 `Couter.sol`，点击 `add`。Metamask 窗口弹出确认请求。确认交易后，点击 `getCounter` 获取计数，该计数为 1。

## Hardhat

Hardhat 是一个灵活的开发环境，用户部署 EVM 智能合约。其设计考虑了集成和拓展性。

### 安装依赖

在开始前，需要安装 Node.js 环境和 npm 包管理器。

```
curl -sL https://deb.nodesource.com/setup_16.x | sudo -E bash -

sudo apt install -y nodejs
```

通过下列方式查询是否正确安装

```
$ node -v
...

$ npm -v
...
```

### 创建 Hardhat 项目

创建新项目：
```
$ npx hardhat

888    888                      888 888               888
888    888                      888 888               888
888    888                      888 888               888
8888888888  8888b.  888d888 .d88888 88888b.   8888b.  888888
888    888     "88b 888P"  d88" 888 888 "88b     "88b 888
888    888 .d888888 888    888  888 888  888 .d888888 888
888    888 888  888 888    Y88b 888 888  888 888  888 Y88b.
888    888 "Y888888 888     "Y88888 888  888 "Y888888  "Y888

👷 Welcome to Hardhat v2.9.3 👷‍

? What do you want to do? …
  Create a basic sample project
❯ Create an advanced sample project
  Create an advanced sample project that uses TypeScript
  Create an empty hardhat.config.js
  Quit
```

根据提示符创建新项目。可以参考 [Hardhat 配置页](https://hardhat.org/hardhat-runner/docs/config) 获取配置选项列表，并在`hardhat.config.js` 中指定选项。最重要的一步，设置 `defaultNetwork` 条目指向你希望的 JSON-RPC 网络：

本地节点
```
module.exports = {
  defaultNetwork: "local",
  networks: {
    hardhat: {
    },
    local: {
      url: "http://localhost:8545/",
      accounts: [privateKey1, privateKey2, ...]
    }
  },
  ...
}

```

测试网
```
module.exports = {
  defaultNetwork: "testnet",
  networks: {
    hardhat: {
    },
    testnet: {
      url: "https://eth.bd.evmos.dev:8545",
      accounts: [privateKey1, privateKey2, ...]
    }
  },
  ...
}

```

为确保指向了正确的网络，可从默认的网络供应商查询可供使用的账户列表

```
$ npx hardhat accounts
0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266
0x70997970C51812dc3A010C7d01b50e0d17dc79C8
0x3C44CdDdB6a900fa2b585dd299e03d12FA4293BC
0x90F79bf6EB2c4f870365E785982E1f101E93b906
...
```

### 部署智能合约

`contracts/Greeter.sol` 下提供了由 Solidity 编写的一份默认的智能合约。

```
 pragma solidity ^0.8.0;

import "hardhat/console.sol";

contract Greeter {
    string private greeting;

    constructor(string memory _greeting) {
        console.log("Deploying a Greeter with greeting:", _greeting);
        greeting = _greeting;
    }

    function greet() public view returns (string memory) {
        return greeting;
    }

    function setGreeting(string memory _greeting) public {
        console.log("Changing greeting from '%s' to '%s'", greeting, _greeting);
        greeting = _greeting;
    }
}

```

该合约允许你设置并查询字符串 `greeting`。Hardhat 同样提供了将合约部署到目标网络的脚本，可按如下命令操作：

```
npx hardhat run scripts/deploy.js
```

也可以手动指定目标网络：

本地节点
```
npx hardhat run --network {{ $themeConfig.project.rpc_url_local }} scripts/deploy.js
```

测试网
```
npx hardhat run --network {{ $themeConfig.project.rpc_url_testnet }} scripts/deploy.js

```

最后，尝试运行 Hardhat 测试

```
$ npx hardhat test
Compiling 1 file with 0.8.4
Compilation finished successfully


  Greeter
Deploying a Greeter with greeting: Hello, world!
Changing greeting from 'Hello, world!' to 'Hola, mundo!'
    ✓ Should return the new greeting once it's changed (803ms)


  1 passing (805ms)

```
## Truffle

Truffle 是用于部署和管理 Solidity 智能合约的开发框架。

### 安装依赖

安装最新版本的 Truffle。

```
yarn install truffle -g
```

### 创建 Truffle 项目

创建新目录并初始化：

```
mkdir evmos-truffle
cd evmos-truffle
```

初始化 Truffle 套件：

```
truffle init
```

创建智能合约 `contracts/Counter.sol`：

```
pragma solidity >=0.7.0 <0.9.0;

contract Counter {
  uint256 counter = 0;

  function add() public {
    counter++;
  }

  function subtract() public {
    counter--;
  }

  function getCounter() public view returns (uint256) {
    return counter;
  }
}
```

编译合约：

```
truffle compile
```

创建测试脚本：

```
const Counter = artifacts.require("Counter")

contract('Counter', accounts => {
  const from = accounts[0]
  let counter

  before(async() => {
    counter = await Counter.new()
  })

  it('should add', async() => {
    await counter.add()
    let count = await counter.getCounter()
    assert(count == 1, `count was ${count}`)
  })
})
```

### 配置 Truffle

打开`truffle-config.js`，取消`network`中`development`部分的注释：

```
    development: {
      host: "127.0.0.1",     // Localhost (default: none)
      port: 8545,            // Standard Ethereum port (default: none)
      network_id: "*",       // Any network (default: none)
    },
```

这使得合约可以连接到本地节点。

### 部署合约

在 Truffle 终端，迁移合约：

```
truffle migrate --network development
```

随后会在 IRITA 守护进程中看到每个交易传入的部署日志（一个是`Migrations.sol`，一个是`Counter.sol`）。

```
$ I[2020-07-15|17:35:59.934] Added good transaction                       module=mempool tx=22245B935689918D332F58E82690F02073F0453D54D5944B6D64AAF1F21974E2 res="&{CheckTx:log:\"[]\" gas_wanted:6721975 }" height=3 total=1
I[2020-07-15|17:36:02.065] Executed block                               module=state height=4 validTxs=1 invalidTxs=0
I[2020-07-15|17:36:02.068] Committed state                              module=state height=4 txs=1 appHash=76BA85365F10A59FE24ADCA87544191C2D72B9FB5630466C5B71E878F9C0A111
I[2020-07-15|17:36:02.981] Added good transaction                       module=mempool tx=84516B4588CBB21E6D562A6A295F1F8876076A0CFF2EF1B0EC670AD8D8BB5425 res="&{CheckTx:log:\"[]\" gas_wanted:6721975 }" height=4 total=1
```

### 运行 Truffle 测试

使用测试命令运行 Truffle 测试：

```
$ truffle test --network development

Using network 'development'.


Compiling your contracts...
===========================
> Everything is up to date, there is nothing to compile.



  Contract: Counter
    ✓ should add (5036ms)


  1 passing (10s)
```
