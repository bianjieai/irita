## Cosmos Enterprise - IRITA Performance Tuning Strategy

More and more traditional web applications have started to accept Web3 solutions for innovating and expanding their businesses. Because of the close ties between these businesses and real-life customers and scenarios, high performance (often referred to as TPS or Transactions Per Second) has become a key preference. One example would be high concurrent volume in ticketing systems when dealing with large public events or mass customer loyalty program activities.

The Cosmos Enterprise product - IRITA offers secure, stable, and high performance in supporting enterprise-level applications, and we are continually putting in resources to optimize it.

To address the high TPS required in a public park ticketing insurance application, we tuned the blockchain with various strategies and are able to achieve TPS  up to 17.5K which is satisfactory to the business requirements.

We would like to share the strategy we used and the settings and information that we believe could be helpful to other enterprise developers. Details and settings are described below:


---
## Node Network Connections Adjustment

Consensus node p2p interconnection, enabling node address exchange but not broadcasting transactions. Full nodes only connect to consensus nodes, and broadcast transactions but do not exchange node addresses. Parameters involved in modifying the config.toml file are as follows:

1. Consensus nodes:
pex = true # Enable node address exchange

broadcast = false # Do not broadcast transactions

persistent_peers = "consensus node ID@consensus node IP:26656" # Other consensus node peers

size # Transaction memory pool is not less than the total memory pool of full nodes

2. Full nodes:
pex = false # Do not exchange node addresses

broadcast = true # Broadcast transactions

persistent_peers = "consensus node ID@consensus node IP:26656" # Consensus node peers

size # Transaction memory pool is smaller than that of consensus nodes

## Node Size Limits Adjustment

Connection numbers, memory pool, single transaction, total transactions in memory pool, request body, requesting/receiving rate, block time, and other size limits:

1. Main parameters in app.toml:
iavl-cache-size # The default size of the node iavl cache is 50M, which can be increased moderately to improve transaction processing speed.

2. Main parameters in config.toml:
size # Default 5000. Increase moderately with configuration to pack more transactions in each block;

max_tx_bytes # Default 1048576 (1M). Recommend to limit the size of a single transaction or the number of msgs in a transaction;

max_txs_bytes # Default 1073741824 (1G), which can be adjusted according to needs, generally less than size * max_tx_bytes;

send_rate # Default 5120000 (5M/s). Applicable for public chains on public networks. For private chains, due to the higher performance of intranet, the default value can be increased to over 50M/s;

recv_rate # Same as above;

max_body_bytes # Default 1000000 (1M). Need to increase for large transactions or multiple msgs;

timeout_commit # Default 5s. Try to reduce while maintaining the single-round block time;

timeout_propose # Default 3s. Try to reduce in the case of non-empty proposal blocks;

flush_throttle_timeout # Default 100ms. Can shorten the waiting time for caching messages in an intranet environment;

max_packet_msg_payload_size # Default 1024 (1K). Can increase message packet size in an intranet environment.

## Node Resource Configuration Adjustment

CPU, memory, disk, and bandwidth:

1. Multi-core CPU processing does not contribute much to improving the current transaction block consensus. Increasing the frequency is more effective;
2. Memory mainly depends on the amount of data. For new chains, the initial 4G is sufficient, and it can be increased as needed later;
3. Disk has a significant impact on node performance. It is recommended to use SSD or higher IOPS;
4. Network bandwidth can accelerate transaction broadcasting and improve consensus voting efficiency.
## On-Chain Block Size Adjustment

While maintaining or approaching the speed of producing empty blocks, ensure that the total size of pending transactions in the mempool to be packaged does not exceed the current block size, then gradually increasing the block size to package more transactions and achieve higher TPS. 

The following are the block debugging sizes achieved on a server with the following configuration: 8-core CPU, 32GB RAM, high-performance SSD disks with high I/O from Alibaba Cloud, and 10Gbps network bandwidth. The TPS achieved for transactions under different message (msgs) sizes are also listed below:

1. iris query params subspace baseapp BlockParams
  value: '{"max_bytes":"819200","max_gas":"-1"}'

2. iris query params subspace baseapp EvidenceParams
  value: '{"max_age_num_blocks":"100000","max_age_duration":"172800000000000","max_bytes":"818176"}'

3. Each transaction contains 7500 msgs with an average TPS of 16.4K. 

![image](https://user-images.githubusercontent.com/31681438/219251944-92d0ff31-6417-447b-9a63-0bbe9af85bc3.png)

4. Each transaction contains 5000 msgs with an average TPS of 17.5K. 

![image](https://user-images.githubusercontent.com/31681438/219252001-5b09f486-625f-46dc-b129-31aad3219b96.png)

5. Each transaction contains 1500 msgs with an average TPS of 16.7K.

![image](https://user-images.githubusercontent.com/31681438/219252041-d3036a58-6130-4208-b4d5-01f03a635fc8.png)

6. Each transaction contains 1000 msgs with an average TPS of 17.4K. 

![image](https://user-images.githubusercontent.com/31681438/219252119-02b1e661-fc9a-4480-aea2-7505eeab5208.png)

7. Each transaction contains 500 msgs with an average TPS of 7.86K. 

![image](https://user-images.githubusercontent.com/31681438/219252152-da1c0c92-3aa1-4448-a630-fe5564bb8ce6.png)

8. Each transaction contains 100 msgs with an average TPS of 8.49K.

![image](https://user-images.githubusercontent.com/31681438/219252197-4cd742c0-5a1c-4e1b-b32a-9b7384737be9.png)

## Transaction Sending Methods:

1. Multi-msg transactions can help increasing TPS, but the larger the number of messages in each transaction, the larger the pending transactions in the block, and the more gas consumption required for each transaction.
2. If the message is too large, the TPS performance may decrease due to longer block production times as the block size increases.
3. Control the transaction sending speed based on the size and occupancy of the mempool, otherwise sending transactions too quickly will result in a slowdown of the block production speed.

---


We are glad to share our performance tuning details with the community and hope it could be of certain help to Cosmos enterprise application developers in addressing the performance requirements of the internet-scale business applications.

We are also eager to learn more practices, if you feel interested in sharing or have any questions/suggestions, please reach out to us at contact@bianjie.ai



