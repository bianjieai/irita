<!--
order: 2
-->

# 部署开发环境

本文档介绍三种部署 IRITA 开发环境的方式，分别对应不同的用例。

- 单节点本地手动测试网
- 多节点本地自动测试网
- 多节点远程手动测试网

::: tip
**注意**：由于 IRITA 使用了国密 `sm2` 加密算法，本文档中涉及的 `openssl` 工具需要使用支持 `sm2` 算法的版本。从源码安装步骤如下：

```bash
git clone -b openssl-3.0.0-alpha4 https://github.com/openssl/openssl.git
cd openssl && ./config
sudo make install
```

:::

## 单节点本地手动测试网

将指导你创建一个由单个验证人节点构成的本地网络，用于测试以及其他开发相关的目的。

### 需求

- [安装 IRITA](./installation.md)
- [安装 `jq`](https://stedolan.github.io/jq/download/) (可选)
- 安装openssl

### 创建 Genesis 文件并启动网络

1. **初始化 genesis.json 文件**

   ```bash
   irita init node0 --chain-id=irita-test --home=testnet
   ```

2. **创建一个初始化账户 v1**

   ```bash
   irita keys add v1 --home=testnet
   ```

3. **将 v1 添加到 genesis.json 文件，并为该账户添加'RootAdmin'权限**

   ```bash
   irita add-genesis-account $(irita keys show v1 -a --home=testnet) 1000000000000upoint,1000000000000upointuirita --home=testnet --root-admin
   ```

4. **导出验证节点 node0（步骤1生成的）私钥为 pem 格式，方便用于申请节点证书**

   ```bash
   irita genkey --home=testnet --out-file priv_validator.pem
   ```

5. **使用`步骤4`中的私钥文件生成证书请求并申请[签发证书](../node_identity_management/cert.md)**

6. **导入 IRITA 网络的企业根证书**(需要先获取根证书)

   ```bash
   irita set-root-cert root.crt --home=testnet
   ```

7. **添加 node0 到 genesis.json 文件**

   ```bash
   irita add-genesis-validator --name node0 --cert node0.crt --power 100 --home=testnet --from v1
   ```

8. **在genesis.json文件中加入upoint并重置irita**

   添加位置位于Token	
   
   owner即为密钥 address 
   
   ```bash
   ,	
    		{
   		  "symbol": "point",
             "name": "Irita point token",
             "scale": 6,
             "min_unit": "upoint",
             "initial_supply": "1000000000",
             "max_supply": "18446744073709551615",
             "mintable": true,
             "owner": ""
           }	
   ```
   
   重置irita
   
   ```bash
   irita unsafe-reset-all --home testnet
   ```
   
9. **启动节点**

   ```bash
   irita start --home=testnet --pruning=nothing
   ```

**以上步骤也可以简化，同样适合多节点的手动部署**

```bash
irita testnet --v 1 --output-dir ./testnet --chain-id=test
```

该命令会生成一个节点，并配置节点证书和根证书(自签证书)。

## 多节点本地自动测试网

将指导你创建一个由多个验证人节点构成的本地网络，可以用于性能相关的测试。

### 需求

- [安装 irita](./installation.md)

### 构建

执行命令获取多节点，以下命令生成三个节点并创建对应的验证人身份

```bash
irita testnet -o ./testnet --chain-id test --v 3
```

上述命令将在 `./testnet` 目录为一个三节点测试网创建文件。这将在 `./testnet` ` 目录产生如下的文件结构：

```bash
testnet/
├── gentxs
│   ├── node0.json
│   ├── node1.json
│   └── node2.json
├── node0
│   ├── iritacli
│   │   ├── key_seed.json
│   │   └── keys
│   └── irita
│       ├── config
│       └── data
├── node1
│   ├── iritacli
│   │   └── key_seed.json
│   └── irita
│       ├── config
│       └── data
├── node2
│   ├── iritacli
│   │   └── key_seed.json
│   └── irita
│       ├── config
│       └── data
│
└── root_cert.pem 
│
└──root_cert.srl
│       
└──root_key.pem
```

### 修改配置

testnet生成的多节点是以不同服务器为目标生成其ip地址为默认状态的192.168.0.2为起始地址，仅本地使用且使用同一个ip则需要进行以下配置。

##### 	对node0修改config

```shell
vim testnet/node1/irita/config/config.toml
```

修改与其他节点连接的端口，多个节点由逗号分隔开

```bash
#peer规则：`{node_id}@ip:port`
```

```shell
persistent_peers = "<node2的id>@localhost:<node2的端口号>,<node1的id>@localhost:<node1的端口号>"
```

​	关闭多节点IP保护设为true

```shell
allow_duplicate_ip = true
```

##### 	对node1修改config

修改ABCI应用程序的TCP或UNIX套接字地址端口号

```shell
原:proxy_app = "tcp://127.0.0.1:26658"
改:proxy_app = "tcp://127.0.0.1:36658"
//原端口号与node0节点冲突
```

修改RPC服务器侦听的TCP或UNIX套接字地址端口号

```shell
原:laddr = "tcp://0.0.0.0:26657"
改:laddr = "tcp://0.0.0.0:36657"
//原端口号与node0节点冲突
```

修改用于侦听传入连接的地址端口号

```shell
原: laddr = "tcp://0.0.0.0:26656"
改: laddr = "tcp://0.0.0.0:36656"
//原端口号与node0节点冲突
```

与其他节点连接端口修改，多个节点由逗号分隔开

```shell
persistent_peers = "<node0的id>@localhost:<node0的端口号>,<node2的id>@localhost:<node2的端口号>"
```

关闭多节点IP保护设为true

```shell
allow_duplicate_ip = true
```

#####   **对node1修改app.toml文件**

​	修改Address定义要绑定的gRPC服务器地址端口号

```shell
原:	address = "0.0.0.0:9090"
改: 	address = "0.0.0.0:9091"
//原端口号与node0节点冲突
```

​	***请对node2进行和node1相同地方进行修改，并保证端口号不冲突***

​	修改连接端口时需要保证连接到除自己以外的其他端口信息正确

```shell
persistent_peers = "<node0的id>@localhost:<node0的端口号>,<node2的id>@localhost:<node2的端口号>"
```

请分别启动三个节点

```bash
irita start --home testnet/node0/irita
irita start --home testnet/node1/irita
irita start --home testnet/node2/irita
每个节点的端口推荐如下表所示：
```

### 密钥和账户

为使用 `irita` 查询状态或者创建交易，需使用给定节点的 `iritacli` 目录作为 `home`，例如：

```bash
irita keys list --home testnet/node0/iritacli
```

现在账户已经存在，可以创建新的账户并且发送资金到新账户。



### 可能出现的错误

1. 9090端口占用

​	 请修改相应节点下app.toml文件中Address定义要绑定的gRPC服务器地址端口号。

2.  未找到 GenesisDoc

 ​	启动节点--home未设置或目录错误

3.  pprof服务端口被占用

 ​	修改相应节点下config.toml文件中proxy_app设置的地址端口号。

4.  26657端口被占用

 ​	修改相应节点下config.toml文件中RPC服务器侦听的TCP或UNIX套接字地址laddr指向地址的端口号

5.  26656端口被占用

 ​	修改相应节点下侦听传入连接的地址laddr指向地址的端口号

6.  连接节点出错

 ​	检查所连节点的address与端口是否对应正确

 ​	检查config.toml文件下allow_duplicate_ip 是否设置为true

 7. 节点访问被拒绝

    ​	删除所有节点下irita/config/addrbook.json文件并重新启动所有节点



## 多节点远程手动测试网

测试使用环境为虚拟机搭建 Ubuntu 

**1. 虚拟机环境搭建**

完成对虚拟机环境的搭建，以保证irita单节点手动部署成功

**2. 节点部署**

在虚拟机0中部署单节点，设置节点名为node0并运行

```bash
irita start --home node0
```

**3. 节点连接**

在虚拟机1中初始化节点

```shell
irita init <moniker> --chain-id=<chain-id>
```

拷贝虚拟机0中node0节点genesis.json文件到虚拟机1的irita运行目录中

添加node0网络节点到该节点配置文件`<home>/config/config.toml`的`persistent_peers`中。

```shell
#peer规则：`{node_id}@ip:port`

# 获取节点id
irita tendermint show-node-id --home <home-dir>
```

 运行node1节点，节点加入到链上

```bash
irita start --home node1
```

**4. 升级为验证人节点**

 在虚拟机1中导出验证节点私钥为pem格式，用于申请节点证书

```bash
irita genkey --out-file priv.pem
```

使用上一步骤中的私钥文件生成证书请求

```bash
openssl req -new -key priv.pem -out req1.csr -sm3 -sigopt "distid:1234567812345678"
```

使用上一步骤中的证书请求在根节点进行签发证书

```bash
openssl x509 -req -in req1.csr -out node1.crt -sm3 -sigopt "distid:1234567812345678" -vfyopt "distid:1234567812345678" -CA root.crt -CAkey root.key -CAcreateserial
```

在根节点使用签发的证书 req1.csr 授权升级为验证人节点

```bash
irita tx node create-validator --name node1 --cert node1.crt --power 100 --description "test" --from node1  --chain-id test --home node1
```

查看验证人节点

```bash
irita query node  validators -o json
```

再加入更多节点同以上步骤进行

