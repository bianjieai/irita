# Using Cosmos technology to deploy an Enterprise Consortium Chain on AWS

> The original article was published on the AWS blog, co-authored by Bianjie and AWS Teams. You can find the source link here: [https://aws.amazon.com/cn/blogs/china/deploying-cosmos-technology-based-enterprise-federation-chains-on-aws/](https://aws.amazon.com/cn/blogs/china/deploying-cosmos-technology-based-enterprise-federation-chains-on-aws/)

## 1. Overview

Cosmos is one of the largest and globally renowned blockchain networks recognized for its open and highly scalable Interchain stack. With effective support for the cross-chain interaction between homogeneous and heterogeneous blockchains, Cosmos aims to extend interoperability to a broader landscape.

Interoperability is essential to blockchains in the multichain era. Cosmos's Interchain stack is well-suited for public chains that focus on vertical domains, providing convenience to dApp builders by offering modular Cosmos SDK tailored to their needs. Applications and protocols within the Cosmos ecosystem are interconnected using the Inter-Blockchain Communication (IBC) protocol, enabling sovereign assets and data exchange between sovereign blockchains. The ultimate goal of Cosmos is to create an Internet of Blockchains that allows for the extensive expansion and interaction of autonomous blockchains.
 
This article primarily focuses on the value and technical architecture of Cosmos, as well as provides a detailed tutorial on the quick deployment of the Cosmos Enterprise Framework — IRITA within the AWS environment.

## 2. Value Proposition
With the continuous development and prosperity of the Cosmos ecosystem, Cosmos-related technologies and communities have garnered increasing attention. 
 
In theory, Cosmos addresses the three most challenging problems in the modern blockchain realm:

1. Scalability: The CometBFT (former Tendermint Core) consensus can be seen as a voting consensus system. Proposers are selected based on Proof-of-Stake (PoS) and Byzantine Fault Tolerance (BFT) algorithms. Under the BFT mechanism, as long as 2/3 of the nodes are honest, the consistency of voting results can be guaranteed.
 
2. Usability: The modular framework — Cosmos SDK — allows for the convenient construction of highly interoperable application-specific blockchains.

3. Interoperability: Cosmos achieves cross-chain communication between L1 networks through the IBC protocol, similar to the role of TCP/IP.

With these unique designs, projects in the Cosmos ecosystem are granted more autonomy, flexibility, and superior performance.

The Cosmos ecosystem provides the necessary framework and infrastructure tools to realize an interoperable multi-chain world. With a focus on autonomy, sovereignty, and scalability, Cosmos offers developers and entrepreneurs a convenient way to experiment and innovate without significant upfront investment. The evolving community, on-chain governance, and decentralized development teams make Cosmos a truly decentralized ecosystem. While Cosmos presents significant opportunities, competition between L1 networks is intensifying. In a minimalist world, Cosmos is constructing a more inclusive multi-chain world and steadily gaining traction.

## 3. Solution

The core products of Cosmos include CometBFT consensus, Cosmos SDK, and the Inter-Blockchain Communication (IBC) protocol as well as Interchain Security.

### 3.1 CometBFT Consensus

CometBFT consists of two main technical components: the blockchain consensus engine (CometBFT) and the Application Blockchain Interface (ABCI). CometBFT ensures all nodes record transactions in the same order. It adopts a typical Byzantine fault-tolerant approach and is a hybrid consensus combining PBFT (Practical Byzantine Fault Tolerance) and Bonded Proof of Stake (Bonded PoS).

![deploying-cosmos-technology-based-enterprise-federation-chains-on-aws1](https://github.com/eviefushigi/irita/assets/68934624/227c5490-16fc-4d46-9968-8d52fd14adc6)

### 3.2 Cosmos SDK

Cosmos SDK is a toolkit that helps developers accelerate the development process, characterized by modularity and pluggability. By using Cosmos SDK, developers can build their own blockchains or functions based on the CometBFT consensus algorithm. Cosmos SDK provides great convenience to developers by significantly shortening their development cycles.

![deploying-cosmos-technology-based-enterprise-federation-chains-on-aws2](https://github.com/eviefushigi/irita/assets/68934624/b40913e1-0281-4343-a9b4-6a78cf2573de)

### 3.3 Inter-Blockchain Communication (IBC) Protocol

Cosmos is a decentralized network composed of multiple sovereign blockchains, which achieve cross-chain communication between different blockchains through the IBC protocol and relayers. Designed by Cosmos, IBC is the most crucial part of the Interchain landscape.

![deploying-cosmos-technology-based-enterprise-federation-chains-on-aws3](https://github.com/eviefushigi/irita/assets/68934624/ac1bdba1-fb7d-40f0-b9d8-0d2b2670edf2)

### 3.4 Interchain Security

[Interchain Security](https://interchainsecurity.dev/) is a complete security solution for Cosmos chains that look to get instant security at launch. With Interchain Security,  "consumer chains" can be secured by the full validator set and multi-billion dollar market cap of the "provider chain" (Cosmos Hub). 

Deploying a consumer chain can be as seamless as deploying a smart contract on a platform such as Ethereum, or the chain can be customized at a very low level using Cosmos SDK.

### 3.5 IRITA

The Inter-Realm Industry Trust Alliance (IRITA) is the first enterprise-level consortium chain product in Cosmos. Built with the modern blockchain framework of CometBFT and IRIS SDK, and backed by the years of experience of the Bianjie in the Cross-chain, NFT, and big data privacy protection fields, IRITA supports next-generation distributed business systems in the form of an enterprise-level consortium chain product line.

IRITA has six core technological advantages: privacy-preserving data encryption and sharing, efficient consensus protocol, advanced cross-chain technology, highly practical on-chain/off-chain system interaction and multi-party collaborative business flow integration capabilities, flexible asset digital modeling and trusted exchange support, as well as big data storage. It can be widely applied in various business scenarios such as finance, healthcare, supply chain, and the Internet of Vehicles, providing value empowerment to the real economy based on blockchain trust machines.

IRITA supports the Chinese national cryptographic standards and provides comprehensive SDK as well as operation and maintenance tool support. It meets enterprise-level application requirements in terms of performance, security, authentication and permissions, maintainability, scalability, and operational monitoring.

As a core contributor to Cosmos technology, the Bianjie team has long been devoted to the field of cross-chain technology. Their open-source code has been adopted by dozens of global blockchain networks. The Bianjie team has contributed functional modules such as NFT module to the Cosmos SDK, completed the implementation of ICS-20 code, and led the development of the ICS-721 Interchain NFTs standard and its code implementation. This functionality introduces cross-chain NFT interoperability based on IBC and extends the capabilities of IBC.

![deploying-cosmos-technology-based-enterprise-federation-chains-on-aws4](https://github.com/eviefushigi/irita/assets/68934624/071049ae-c141-419a-b8e4-c4de8d500d4e)

## 4. Deployment

### 4.1 Node Preparation

Node configuration requirements:

Testing environment: 2 vCPUs, 8GB RAM, 100GB disk. Recommended instance types: m6a.large, m5.large, t3.large.

Production environment: 4 vCPUs, 16GB RAM, 100GB disk. Recommended instance types: m6a.xlarge, m5a.xlarge, m6i.xlarge.

Create 4 blockchain node servers, with the following key steps:

(1) Choose the operating system: Amazon Linux 2 AMI - Kernel 5.10.

![deploying-cosmos-technology-based-enterprise-federation-chains-on-aws6](https://github.com/eviefushigi/irita/assets/68934624/6af5fe3b-4183-4b83-aa9e-f74020266fa6)

(2) Select the instance type, for example, m6a.large.

![deploying-cosmos-technology-based-enterprise-federation-chains-on-aws7](https://github.com/eviefushigi/irita/assets/68934624/7ba3a991-78a0-4a87-929e-b0eb6aa45f4b)

(3) Configure local storage by creating 1 system disk (20GB gp3 EBS volume) and 1 data disk (100GB gp3 EBS volume). Mount the data disk to the /data directory.

![deploying-cosmos-technology-based-enterprise-federation-chains-on-aws8](https://github.com/eviefushigi/irita/assets/68934624/e2cb69c2-3a29-4632-85a8-08918375e7e4)

(4) Enter the following content in the “user data” field under the “Advanced” option to execute node initialization operations.
```ruby
#!/bin/bash -ex
# Install Docker
amazon-linux-extras install docker -y
# Start the docker service
systemctl start docker.service
systemctl enable docker.service
# Add user to docker group
usermod -a -G docker ec2-user
# Format and mount data volume
mkfs -t xfs /dev/nvme1n1
mkdir /data
echo "/dev/nvme1n1    /data           xfs    defaults,nofail  0   2" | tee -a /etc/fstab
mount -a
chown ec2-user.ec2-user /data
# Reboot instance
reboot
```

(5) 4 blockchain node servers will appear after completing the above steps. Name them from “node0” to node3 respectively.

![deploying-cosmos-technology-based-enterprise-federation-chains-on-aws9](https://github.com/eviefushigi/irita/assets/68934624/a5c9d85b-863f-42c8-adbf-bc76a70e9844)

(6) Configure security groups and enable the following port access for each node:
- 1317: Provides external access to the RESTful API (consensus nodes may not open, full nodes optional).
- 8545-8546: Provides RPC and WebSockets interfaces for EVM (consensus nodes may not open, full nodes optional).
- 9090: External gRPC interface for nodes (consensus nodes may not open, full nodes optional).
- 26656: P2P network between nodes.
- 26657: External RPC interface for nodes (consensus nodes may not open, full nodes optional).
- 26660: Provides monitoring metrics.
Created security groups:

![deploying-cosmos-technology-based-enterprise-federation-chains-on-aws10](https://github.com/eviefushigi/irita/assets/68934624/270c066e-9644-4c16-9fd3-64c4292cde75)

Associate the security group with each node:

![deploying-cosmos-technology-based-enterprise-federation-chains-on-aws11](https://github.com/eviefushigi/irita/assets/68934624/29f806a1-787b-4069-9008-4ed299d4873d)

### 4.2 Installation and Deployment

(1) Generate node data on the first machine (node0) with chain-id: testnet.
```ruby
# Generate 4 node data，prompt for password input 8 times.
$ docker run -it --rm -v /data:/root bianjie/irita:v3.2.2-wenchangchain irita testnet --v 4 --chain-id testnet --output-dir /root
Enter keyring passphrase:
Re-enter keyring passphrase:
…
…
…
Successfully initialized 4 node directories
```
(2) Copy data

At this point, the /data directory of node0 will contain data for node0, node1, node2, and node3. Copy the data to the corresponding /data directory of each node and name it “node”. Pay attention to the directory's permission settings during the copying process.
```ruby
# You can use scp to copy the data to the /data directory of other nodes (node1~3).
$ scp -i ~/.ssh/your-private-key.pem -r /data/nodeX ec2-user@172.31.17.118:/data/node
```

(3) Modify the configuration (perform the following steps on all four nodes from this point on).
   
a. Configure the peers of other nodes.

Determine the internal IP and peer ID of each machine.
```ruby
$ ifconfig
$ docker run -it --rm -v /data/node/irita:/root/.irita bianjie/irita:v3.2.2-wenchangchain irita tendermint show-node-id
```
At the end we can obtain the peer configuration of each node in the format of node_peerid@node_ip:26656.

It is recommended to organize the obtained information in a table for convenient modification of the configuration file later, for example:
| Name  | node_peerid@node_ip:26656 |
| --- | --- |
| Node0 | cf14286aef99e49c702cba4bd31d3529b8b3c01a@172.31.17.118:26656 |
| Node1 | 53aabd9bf37c54a2c21ece0671d08131e1b121cf@172.31.24.153:26656 |
| Node2 | a35ec424702ed479247c2a8adc546f238f9bc2d5@172.31.17.206:26656 |
| Node3 | 3368c48c13fb26ca0bcef27680c252877d94eddf@172.31.29.175:26656 | 

Modify the persistent_peers configuration in /data/node/irita/config/config.toml, where each node needs to configure the peers of other nodes (excluding its own configuration).

For example, the configuration for node0 should be as follows:
```ruby
# node0
persistent_peers = "53aabd9bf37c54a2c21ece0671d08131e1b121cf@172.31.24.153:26656,a35ec424702ed479247c2a8adc546f238f9bc2d5@172.31.17.206:26656,3368c48c13fb26ca0bcef27680c252877d94eddf@172.31.29.175:26656"
```

b. Modify /data/node/irita/config/config.toml.
```ruby
# Change it to false. The node can be started within the intranet.
addr_book_strict = false
```

(4) Start the nodes on each host.
```ruby
$ docker run -itd -p1317:1317 -p26656-26660:26656-26660 -p9090:9090 -p8545:8545 -p8546:8546 -v /data/node/irita:/root/.irita --name node bianjie/irita:v3.2.2-wenchangchain irita start
```

(5) Import validators

There is an iritacli directory in the /data/node directory of each node. Perform the following steps on each node:

View the mnemonic:
```ruby
# Copy the value of "secret" inside the file, which is the mnemonic phrase 
cat /data/node0/iritacli/key_seed.json
```

Recover the address using the mnemonic:
```ruby
# Open an interactive mode terminal within the running node container
$ docker exec -it 951f0200849e bash
# To recover the address using the mnemonic phrase, paste the copied mnemonic phrase and press Enter. Then enter the password set in Step 1) when prompted
$ irita keys add validator --recover
```
![deploying-cosmos-technology-based-enterprise-federation-chains-on-aws12](https://github.com/eviefushigi/irita/assets/68934624/bcb06952-8828-4fd4-a076-9ef2441f5594)

List local addresses:
```ruby
# To view the local address, enter the password set in Step 1) when prompted.
$ irita keys list
```
![deploying-cosmos-technology-based-enterprise-federation-chains-on-aws13](https://github.com/eviefushigi/irita/assets/68934624/9a63d847-ae7f-4561-9e5c-75d8048723e3)

(6) Check the status

Access [http://node-ip-address:26657/status](http://node-ip-address:26657/status) to view blockchain status information.
![deploying-cosmos-technology-based-enterprise-federation-chains-on-aws14](https://github.com/eviefushigi/irita/assets/68934624/04c38cff-3ffc-4298-b660-0e9e4cfed774)

### 4.3 Node Configuration Guide

The main configuration files for nodes are: config.toml, app.toml, genesis.json. Among them, genesis.json is the genesis block file, which can be modified before chain startup and contains consensus parameters such as chain-id, block size, and the number of consensus nodes.

(1) config.toml common configurations, effective after node restart:
```ruby
# The address of the peer node is in the format: ip@id:port. Use "seeds" to obtain addresses of other broadcast nodes.
seeds = ""
persistent_peers = ""

# Whether to broadcast the address of this node.
pex = true

# Private protected peer addresses, mainly used for sentinel nodes protecting consensus nodes that are not exposed to the public network.
private_peer_ids = ""

[mempool]
# Maximum number of transactions in the mempool.
size = 5000

# Whether the node broadcasts transactions.
broadcast = true

# Overall size limit of transactions in the mempool.
max_txs_bytes = 1073741824

# Size limit of each transaction in the mempool.
max_tx_bytes = 1048576

[statesync]
# State sync allows starting a node without copying data. It requires a node with snapshot functionality to start from the latest state without historical block data.
enable = false

[consensus]
# Block time. The actual block time depends on all nodes.
timeout_commit = "5s"

[tx_index]
# Whether the node indexes block height, transactions, etc. "null" disables indexing to improve node performance but prevents querying transactions by TX hash.
indexer = "kv"

[instrumentation]
# Enable node monitoring and listening port.
prometheus = false
prometheus_listen_addr = ":26660"
```

(2) app.toml common configuration items, effective after node restart:
```ruby
# Specify the minimum gas price threshold for a node to accept a transaction. For example, if the transaction has a gas value of 200000, the transaction will only be broadcasted by the node if fees >= gas * minimum-gas-prices = 200000ugas.
minimum-gas-prices = "1ugas" 
# Nodes clean historical block state to reduce disk storage. There are multiple pruning strategies to choose from,
pruning = "default"
# Block height or time at which a node stops, typically used for chain halting upgrades. 
halt-height = 0 
halt-time = 0 
# Whether the node has enabled the REST API server and its corresponding Swagger interface documentation. 
[api] 
enable = false 
swagger = false 
[grpc] 
# Whether the node has enabled the grpc server
enable = true 
# Whether the node has enabled the snapshot feature for quick startup by other nodes. When snapshot-interval is non-zero, the node starts generating snapshots at block heights that are multiples of the configured value. The number of snapshots to retain is specified by snapshot-keep-recent
[state-sync] 
snapshot-interval = 0 
snapshot-keep-recent = 2
```

### 4.4 Block Data Synchronization

For newly started nodes, besides synchronizing from the genesis block, there are three methods for fast synchronization:

(1) Stop a running node, and package the /root/.irita/data directory, and then copy and decompress it to the target host.

(2) For nodes running on the Amazon cloud platform, take a snapshot of the data disk (EBS volume) and use the snapshot to quickly restore data for other nodes. The snapshot can also be shared with other accounts to accelerate node data synchronization. Ensure that no other data/files outside the data are left in the snapshot.

![deploying-cosmos-technology-based-enterprise-federation-chains-on-aws15 jpg](https://github.com/eviefushigi/irita/assets/68934624/fe16c638-84f6-42cb-87dd-bbc776d34a7b)

(3) Use state sync to quickly start a new node. This method requires a node to enable block state snapshot as the data source for the new joining node. The new node can start with state sync mode and directly synchronize the latest block height state (skipping historical blocks) for fast startup.

### 4.5 Blockchain Monitoring

(1) Enable chain monitoring.

Modify the prometheus option in /data/node/irita/config/config.toml, with the default port being 26660.
```ruby
# Change to true to enable prometheus monitoring
prometheus = true
```
Save and restart the node service.
```ruby
# restart $ docker restart node
```
Access [http://<node-ip-address>:26660/metrics](http://<node-ip-address>:26660/metrics) to view the monitored metrics.
![deploying-cosmos-technology-based-enterprise-federation-chains-on-aws16](https://github.com/eviefushigi/irita/assets/68934624/c50853cf-f81a-41a1-bbc4-21976d5c0549)

Common monitoring rules are as follows:
```ruby
   - "name": "chain"
"rules":
     - "alert": "chain Consensus Halt"
"annotations":
"cluster": "Irita nodes"
"message": "chain consensus has halted for {{ $value }} rounds."
"expr": |
         max(tendermint_consensus_rounds) > 10
"for": "1m"
"labels":
"severity": "critical"
     - "alert": "chain Node Out Of Sync"
"annotations":
"cluster": "Irita nodes"
"message": "chain node '{{ $labels.instance }}' out of block sync for over 5 minutes."
"expr": |
         changes(tendermint_consensus_latest_block_height{job="chain"}[5m]) == 0 and tendermint_consensus_rounds{job="chain"} == 0
"for": "1m"
"labels":
"severity": "critical"
     - "alert": "chain Validator Jailed" 
"annotations":
"cluster": "Irita nodes"
"message": "{{ $value }} validators are jailed"
"expr": |
         sum(iris_module_stake_jailed{instance="validator"} == 1) by (namespace, instance)
"for": "1m"
"labels":
"severity": "critical"
     - "alert": "chain Online Voting Power Waring"
"annotations":
"cluster": "Irita nodes"
"message": "chain online voting power is less than 70%."
"expr": |
         (1 - (tendermint_consensus_byzantine_validators_power + tendermint_consensus_missing_validators_power) / tendermint_consensus_validators_power) <= 0.7
"for": "1m"
"labels":
"severity": "critical"
     - "alert": "chain Byzantine Validators Waring"
"annotations":
"cluster": "Irita nodes"
"message": "chain has found {{ $value }} byzantine validator(s)."
"expr": |
         tendermint_consensus_byzantine_validators{instance="chain-validator"} > 0
"for": "1m"
"labels":
"severity": "critical"
     - "alert": "chain Node Down"
"annotations":
"cluster": "Irita nodes"
"message": "chain node '{{ $labels.instance }}' has disappeared from Prometheus target discovery."
"expr": |
         up{job="chain"} == 0
"for": "1m"
"labels":
"severity": "critical"
```

(2) Prometheus configuration.
```ruby
- "job_name": "irita"
"static_configs":
- "targets": ["192.168.0.160:26660"]
```

(3) Grafana configuration template.

You can use a fully managed Grafana service (Amazon Managed Grafana) to visualize and monitor metric data from the Prometheus data source.

In the Amazon Grafana console, create a new workspace for Irita and complete the basic configuration:

![deploying-cosmos-technology-based-enterprise-federation-chains-on-aws17](https://github.com/eviefushigi/irita/assets/68934624/8a355bc7-4f27-4fa6-8791-6164c03cc722)

Once ready, log in with a user having Admin privileges. In the Data source section, add the Prometheus data source.

![deploying-cosmos-technology-based-enterprise-federation-chains-on-aws18](https://github.com/eviefushigi/irita/assets/68934624/bdc6364d-39a1-4684-8e36-86108e43e7bd)

Create a dashboard template and add panels, select Irita as the data source, and choose the desired metrics in the metric browser, such as block height, transaction count, block size, average block time, and other metrics.

![deploying-cosmos-technology-based-enterprise-federation-chains-on-aws19](https://github.com/eviefushigi/irita/assets/68934624/520a233d-cd1a-4a43-b5ba-1ae8a9967642)

Example of an Irita monitoring dashboard:

![deploying-cosmos-technology-based-enterprise-federation-chains-on-aws20](https://github.com/eviefushigi/irita/assets/68934624/4d9e21be-2277-4e2d-a61f-f0cbf8ac4950)

### 4.6 Resource Cleanup

(1) To keep the test data, unmount the data disk.
```ruby
- "job_name": "irita"
  "static_configs":
  - "targets": ["192.168.0.160:26660"]
```

In the EC2 console, click on the instance, go to the Storage section, click on the corresponding Volume ID, select the volume, and click on "Detach Volume" in the upper-right corner. If you want to delete the volume, select the volume and click on "Delete" in the Actions menu.

(2) Delete the instance.

In the EC2 console, click on the instance, select the corresponding instance, and click on "Terminate Instance" in the upper-right corner of the instance status.

![deploying-cosmos-technology-based-enterprise-federation-chains-on-aws21](https://github.com/eviefushigi/irita/assets/68934624/83d96312-55fd-4b29-99d0-c46f2d6583cc)

## 5. Summary

Deploying Cosmos-based blockchain nodes with AWS EC2 allows for convenient dynamic adjustment of resource configurations such as CPU, memory, and disk IO based on the workload. Dynamic resource adjustment according to business needs can effectively achieve cost reduction and increased efficiency.

At the network level, interconnecting consortium chain nodes through the public network usually compromise security and performance while incurring high costs. By utilizing AWS VPC peering connections, different node operators can deploy their nodes on Amazon Cloud and achieve interconnection across accounts through the internal backbone network, ensuring both security and performance, while virtually neglecting network traffic costs. The resource isolation and permission isolation between AWS accounts also effectively support the decentralized operation of the consortium chain.

In terms of data security, utilizing AWS EC2 Data Lifecycle Manager allows for defining snapshot policies for node data disks and performing incremental snapshots on a scheduled basis without disrupting operations. Additionally, disk recovery based on snapshots can be completed in seconds, significantly improving fault recovery and the efficiency of new node startup. This approach avoids the time-consuming and resource-intensive process of synchronizing and executing historical block data from scratch. Moreover, snapshots can be shared across AWS accounts, playing a crucial role in the deployment architecture of consortium chains involving multiple participants.

Using Amazon Managed Grafana as a managed service enables easier visualization and monitoring of node operation status. Facing increasing usage demands, it can automatically scales computing and database infrastructure, performs automatic version updates and security patches, thereby reducing the operational management burden of Grafana.

Bianjie combines multiple products from AWS to realize one-click deployment and automated operation and maintenance of enterprise-level consortium chains based on Cosmos, and plans to add the deployment of consortium chain nodes as a managed service to the AWS Marketplace, allowing institutional customers to easily join existing networks or launch customized consortium chain networks with just a few clicks.

