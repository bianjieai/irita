<!--
order: 2
-->

# 部署开发环境

本文档介绍三种部署 IRITA 开发环境的方式，分别对应不同的用例。

- 单节点本地手动测试网
- 多节点本地自动测试网
- 多节点远程自动测试网

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
- openssl

### 创建 Genesis 文件并启动网络

1. **初始化 genesis.json 文件**

   ```bash
   irita init node0 --chain-id=irita-test --home=testnet
   ```

2. **创建一个初始化账户 v1**

   ```bash
   irita keys add v1
   ```

3. **将 v1 添加到 genesis.json 文件，并为该账户添加'RootAdmin'权限**

   ```bash
   irita add-genesis-account $(irita keys show v1 -a) 1000000000point --home=testnet --root-admin
   ```

4. **导出验证节点 node0（步骤1生成的）私钥为 pem 格式，方便用于申请节点证书**

   ```bash
   irita genkey --home=testnet --out-file priv_validator.pem
   ```

5. **使用`步骤4`中的私钥文件生成证书请求并申请[签发证书](../node_identity_management/cert.md)**

6. **导入 IRITA 网络的企业根证书**(需要先获取根证书)

   ```bash
   irita set-root-cert ca.crt --home=testnet
   ```

7. **添加 node0 到 genesis.json 文件**

   ```bash
   irita add-genesis-validator --name node0 --cert node0.crt --power 100 --home=testnet --from node0
   ```

8. **启动节点**

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
- [安装 docker](https://docs.docker.com/engine/installation/)
- [安装 docker-compose](https://docs.docker.com/compose/install/)

### 构建

为运行 `localnet` 命令，需构建 `irita` 二进制 (linux) 以及 `irita/iritanode` docker 镜像。二进制将被挂载到容器，能通过重新构建镜像来更新，因此你只需构建镜像一次。

```bash
# 克隆 irita
git clone https://github.com/bianjieai/irita.git

# 进入 irita 目录
cd irita

# 构建 linux 二进制
make build-linux

# 构建 irita/iritanode 镜像
make build-docker-iritanode
```

### 运行测试网

启动一个四节点测试网：

```bash
make localnet-start
```

这个命令使用 iritanode 镜像创建一个四节点网络。
每个节点的端口如下表所示：

| 节点 ID     | P2P 端口 | RPC 端口 |
| ----------- | -------- | -------- |
| `iritanode0` | `26656`  | `26657`  |
| `iritanode1` | `26659`  | `26660`  |
| `iritanode2` | `26661`  | `26662`  |
| `iritanode3` | `26663`  | `26664`  |

为更新二进制文件，只需重新构建镜像并且启动这些节点：

```bash
make build-linux localnet-start
```

### 配置

`make localnet-start` 将调用 `irita testnet` 命令在 `./build` 目录为一个四节点测试网创建文件。这将在 `./build` 目录产生如下的文件结构：

```bash
$ tree -L 2 build/
build/
├── irita
├── gentxs
│   ├── node0.json
│   ├── node1.json
│   ├── node2.json
│   └── node3.json
├── node0
│   ├── iritacli
│   │   ├── key_seed.json
│   │   └── keys
│   └── irita
│       ├── ${LOG:-irita.log}
│       ├── config
│       └── data
├── node1
│   ├── iritacli
│   │   └── key_seed.json
│   └── irita
│       ├── ${LOG:-irita.log}
│       ├── config
│       └── data
├── node2
│   ├── iritacli
│   │   └── key_seed.json
│   └── irita
│       ├── ${LOG:-irita.log}
│       ├── config
│       └── data
└── node3
    ├── iritacli
    │   └── key_seed.json
    └── irita
        ├── ${LOG:-irita.log}
        ├── config
        └── data
```

每个 `./build/nodeN` 目录被挂载到对应容器的 `/irita` 目录。

### 日志

日志被保存在每个节点的 `./build/nodeN/irita/irita.log` 文件。可以直接通过 Docker 监视日志，例如：

```bash
docker logs -f node0
```

### 密钥和账户

为使用 `irita` 查询状态或者创建交易，需使用给定节点的 `iritacli` 目录作为 `home`，例如：

```bash
irita keys list --home ./build/node0/iritacli
```

现在账户已经存在，可以创建新的账户并且发送资金到新账户。

::: tip
**注意**：每个节点的 seed 存放在 `./build/nodeN/iritacli/key_seed.json`，可以使用 `irita keys add --restore` 命令将其恢复到 CLI 。 
:::

## 多节点远程自动测试网

多节点远程自动测试网将通过 `ssh` 命令部署一个四节点测试网。请确保各服务器 `ssh` 的权限已配置并且 `docker` 已安装。

### 脚本代码

> 执行脚本前需修改相应参数

```bash
ChainID=testnet # chain-id
ChainCMD=irita
NodeName=irita-node # node name
DockerIP=(tcp://192.168.0.1 tcp://192.168.0.2 tcp://192.168.0.3 tcp://192.168.0.4)
Names=("node0" "node1" "node2" "node3")
Mnemonics=("1 2 ... 24" "1 2 ... 24" "1 2 ... 24" "1 2 ... 24")
Stake=point
TotalStake=1000000000000000${Stake} # total stake in genesis
SendStake=100000000${Stake}
DataPath=/tmp

for i in `seq 0 $[ ${#DockerIP[*]} -1 ]`; do docker -H ${DockerIP[$i]} run -itd -e ChainCMD=$ChainCMD -e NodeName=${Names[$i]} -v $DataPath/$NodeName-$i:/root --name $NodeName-$i bianjie/irita:v2.0.0-alpha; done
for i in `seq 0 $[ ${#DockerIP[*]} -1 ]`; do docker -H ${DockerIP[$i]} exec -it $NodeName-$i $ChainCMD version; done
for i in `seq 0 $[ ${#DockerIP[*]} -1 ]`; do docker -H ${DockerIP[$i]} exec -it $NodeName-$i rm -rf /root/.${ChainCMD} /root/.${ChainCMD}; done
for i in `seq 0 $[ ${#DockerIP[*]} -1 ]`; do docker -H ${DockerIP[$i]} exec -it $NodeName-$i bash -c "echo \"${Mnemonics[$i]}\n12345678\n12345678\" | $ChainCMD keys add validator --recover 2>&1 | tee /root/seed.key"; done
for i in `seq 0 $[ ${#DockerIP[*]} -1 ]`; do docker -H ${DockerIP[$i]} exec -it $NodeName-$i bash -c "echo 12345678 | ${ChainCMD} keys list"; done
for i in `seq 0 $[ ${#DockerIP[*]} -1 ]`; do docker -H ${DockerIP[$i]} exec -it $NodeName-$i $ChainCMD init moniker --chain-id $ChainID; done
for i in `seq 0 $[ ${#DockerIP[*]} -1 ]`; do docker -H ${DockerIP[$i]} exec -it $NodeName-$i $ChainCMD genkey --out-file /root/priv_validator.pem; done
for i in `seq 0 $[ ${#DockerIP[*]} -1 ]`; do docker -H ${DockerIP[$i]} exec -it $NodeName-$i $ChainCMD genkey --type node --out-file /root/priv_node.pem; done

docker -H ${DockerIP[0]} exec -it $NodeName-0 sed -i 's/127.0.0.1:26657/0.0.0.0:26657/g' /root/.$ChainCMD/config/config.toml
docker -H ${DockerIP[0]} exec -it $NodeName-0 sed -i 's/timeout_commit = "5s"/timeout_commit = "2s"/' /root/.$ChainCMD/config/config.toml
docker -H ${DockerIP[0]} exec -it $NodeName-0 sed -i "s/stake/$Stake/g" /root/.$ChainCMD/config/genesis.json
docker -H ${DockerIP[0]} exec -it $NodeName-0 bash -c 'sed -i "s/owner\": \"iaa183rfa8tvtp6ax7jr7dfaf7ywv870sykxxykejp/owner\": \"$(echo 12345678 | $ChainCMD keys show validator | grep address | cut -b 12-)/" /root/.$ChainCMD/config/genesis.json'
docker -H ${DockerIP[0]} exec -it $NodeName-0 bash -c 'sed -i "s/nodes\": \[/nodes\": \[{\"id\": \"$($ChainCMD tendermint show-node-id)\", \"name\": \"$NodeName\"}/" /root/.$ChainCMD/config/genesis.json'
docker -H ${DockerIP[0]} exec -it $NodeName-0 bash -c "$ChainCMD add-genesis-account \$(echo 12345678 | ${ChainCMD} keys show validator -a) ${TotalStake} --root-admin"
docker -H ${DockerIP[0]} exec -it $NodeName-0 openssl ecparam -genkey -name SM2 -out /root/root.key
docker -H ${DockerIP[0]} exec -it $NodeName-0 bash -c 'echo -e "CN\nSH\nSH\nIT\nDEV\n'${Names[0]}'\n\n" | openssl req -new -x509 -sm3 -sigopt "distid:1234567812345678" -key /root/root.key -out /root/root.crt -days 3650'
docker -H ${DockerIP[0]} exec -it $NodeName-0 $ChainCMD set-root-cert /root/root.crt

docker -H ${DockerIP[0]} cp $NodeName-0:/root/root.key .
docker -H ${DockerIP[0]} cp $NodeName-0:/root/root.crt .
for i in `seq 1 $[ ${#DockerIP[*]} -1 ]`; do docker -H ${DockerIP[$i]} cp root.crt $NodeName-$i:/root/; done
for i in `seq 1 $[ ${#DockerIP[*]} -1 ]`; do docker -H ${DockerIP[$i]} cp root.key $NodeName-$i:/root/; done
for i in `seq 0 $[ ${#DockerIP[*]} -1 ]`; do docker -H ${DockerIP[$i]} exec -it $NodeName-$i bash -c 'echo -e "CN\nSH\nSH\nIT\nDEV\n'"${Names[$i]}"'\n\n\n\n" | openssl req -new -key /root/priv_validator.pem -out /root/validator_req.csr -sm3 -sigopt "distid:1234567812345678"'; done
for i in `seq 0 $[ ${#DockerIP[*]} -1 ]`; do docker -H ${DockerIP[$i]} exec -it $NodeName-$i openssl x509 -req -in /root/validator_req.csr -out /root/validator.crt -sm3 -sigopt "distid:1234567812345678" -vfyopt "distid:1234567812345678" -CA /root/root.crt -CAkey /root/root.key -CAcreateserial; done
for i in `seq 0 $[ ${#DockerIP[*]} -1 ]`; do docker -H ${DockerIP[$i]} exec -it $NodeName-$i bash -c 'echo -e "CN\nSH\nSH\nIT\nDEV\n'"${Names[$i]}"'\n\n\n\n" | openssl req -new -key /root/priv_node.pem -out /root/node_req.csr -sm3 -sigopt "distid:1234567812345678"'; done
for i in `seq 0 $[ ${#DockerIP[*]} -1 ]`; do docker -H ${DockerIP[$i]} exec -it $NodeName-$i openssl x509 -req -in /root/node_req.csr -out /root/node.crt -sm3 -sigopt "distid:1234567812345678" -vfyopt "distid:1234567812345678" -CA /root/root.crt -CAkey /root/root.key -CAcreateserial; done

docker -H ${DockerIP[0]} exec -it $NodeName-0 bash -c "echo 12345678 | $ChainCMD add-genesis-validator --name ${Names[0]} --cert /root/validator.crt --power 10000 --from validator"
docker -H ${DockerIP[0]} cp $NodeName-0:/root/.$ChainCMD/config/config.toml .
docker -H ${DockerIP[0]} cp $NodeName-0:/root/.$ChainCMD/config/genesis.json .
sed -i "s/persistent_peers = \"\"/persistent_peers = \"$(docker -H ${DockerIP[0]} exec -it $NodeName-0 $ChainCMD tendermint show-node-id | cat -vet | sed 's/\^M\$//')@`echo ${DockerIP[0]} | awk -F // '{print $2}'`:26656\"/" config.toml
for i in `seq 1 $[ ${#DockerIP[*]} -1 ]`; do docker -H ${DockerIP[$i]} cp config.toml $NodeName-$i:/root/.$ChainCMD/config/; done
for i in `seq 1 $[ ${#DockerIP[*]} -1 ]`; do docker -H ${DockerIP[$i]} cp genesis.json $NodeName-$i:/root/.$ChainCMD/config/; done
for i in `seq 1 $[ ${#DockerIP[*]} -1 ]`; do docker -H ${DockerIP[$i]} cp $NodeName-$i:/root/node.crt node$i.crt; done
for i in `seq 0 $[ ${#DockerIP[*]} -1 ]`; do docker -H ${DockerIP[$i]} stop $NodeName-$i; done
for i in `seq 0 $[ ${#DockerIP[*]} -1 ]`; do docker -H ${DockerIP[$i]} rm $NodeName-$i; done
for i in `seq 0 $[ ${#DockerIP[*]} -1 ]`; do docker -H ${DockerIP[$i]} run -itd -p26656:26656 -p26657:26657 -p9090:9090 -v $DataPath/$NodeName-$i:/root --name $NodeName-$i bianjie/irita:v2.0.0-alpha $ChainCMD start --pruning=nothing; done

sleep 5
for i in `seq 1 $[ ${#DockerIP[*]} -1 ]`; do
address=$(docker -H ${DockerIP[$i]} exec -it $NodeName-$i bash -c "echo 12345678 | ${ChainCMD} keys show validator | grep address" | awk '{print $2}');
echo $address
docker -H ${DockerIP[0]} exec -it $NodeName-0 bash -c "echo -e \"12345678\n12345678\" | ${ChainCMD} tx bank send validator \$(echo $address | cat -A | sed 's/\\^M\\$//') ${SendStake} --chain-id $ChainID -y";
sleep 5
docker -H ${DockerIP[0]} exec -it $NodeName-0 bash -c "${ChainCMD} q bank balances \$(echo $address | cat -A | sed 's/\\^M\\$//') --chain-id $ChainID";
docker -H ${DockerIP[0]} exec -it $NodeName-0 bash -c "echo -e \"12345678\n12345678\" | ${ChainCMD} tx admin add-roles --from validator \$(echo $address | cat -A | sed 's/\\^M\\$//') NODE_ADMIN --chain-id $ChainID -y";
sleep 5
docker -H ${DockerIP[0]} exec -it $NodeName-0 bash -c "${ChainCMD} q admin roles \$(echo $address | cat -A | sed 's/\\^M\\$//') --chain-id $ChainID";
docker -H ${DockerIP[0]} cp node$i.crt $NodeName-0:/root/;
docker -H ${DockerIP[0]} exec -it $NodeName-0 bash -c "echo -e \"12345678\n12345678\" | ${ChainCMD} tx node grant --name \"${Names[$i]}\" --cert /root/node$i.crt --from validator --chain-id $ChainID -b block -y";
docker -H ${DockerIP[$i]} exec -it $NodeName-$i bash -c "echo -e \"12345678\n12345678\" | ${ChainCMD} tx node create-validator --name \"${Names[$i]}\" --from validator --cert /root/node.crt --power 100 --chain-id $ChainID --node=${DockerIP[0]}:26657 -y";
done
sleep 5
for i in `seq 0 $[ ${#DockerIP[*]} -1 ]`; do docker -H ${DockerIP[$i]} exec -it $NodeName-$i sed -i 's/filter_peers = false/filter_peers = true/' /root/.$ChainCMD/config/config.toml; docker -H ${DockerIP[$i]} restart $NodeName-$i; done
```
