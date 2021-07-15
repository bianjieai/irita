## IRITA 节点运维 Q&A

### 1. IRITA 有哪些不同类型的密钥

IRITA有两种类型的密钥：

- Tendermint 密钥，存储在 `config/priv_validator_key.json` 中
  - 由 `irita init` 创建节点时生成

  - 查询bech32前缀为 `icp` 的共识公钥

    ```shell
    irita tendermint show-validator
    ```

  - 查询bech32前缀为 `ica` 的节点地址

    ```shell
    irita tendermint show-address
    ```

  - 查询对应的节点 Hex 地址

    ```shell
    irita status | jq .validator_info.address
    ```


- 应用程序密钥：

  该密钥可以通过`irita`创建，用于签名交易。应用程序密钥与以`iap`为前缀的公钥和以`iaa`为前缀的地址相关联。两者都是由`irita keys add`生成的帐户密钥派生出来的。

注意：验证人操作员的密钥直接与应用程序密钥绑定，地址和公钥分别使用保留的前缀：iva 和 ivp。

### 2. 如何备份验证人节点

安全备份验证人节点私钥非常**重要**，这是恢复验证人节点的唯一方法。请注意，这里指的是 Tendermint 密钥。

如果您使用的是软件签名（ tendermint 的默认签名方法），则您的 Tendermint 密钥)位于`<irita-home>/config/priv_validator_key.json`中。最简单的方法是备份整个 config 文件夹。

或者，您可以使用硬件更安全地管理 Tendermint 密钥，例如[YubiSM2](https://developers.yubico.com/YubiHSM2/)

### 3. 如何迁移/恢复验证人节点

迁移验证人的方法有很多，最推荐的方法是：

1. 在新服务器上运行运行全节点
2. 追赶上最新区块之后，停止验证人节点和全节点
3. 将全节点的 `config` 文件夹替换为验证人的
4. 启动新的验证人节点

### 4. 如何查询我的验证人地址

验证人地址有2种：

- 验证人操作员地址，即用来创建验证人节点的**应用程序密钥**

  查询验证人的操作员地址（iva ...）和 pubkey（ivp ...）：

  ```bash
  irita keys show MyKey --bech=val
  ```

- 验证人节点地址，即 **Tendermint 密钥**

  查询关联的地址（ica ...）和共识pubkey（icp ...），请参考Tendermint密钥

### 5. 验证人的投票权为 0 怎么办

可能是您的验证人被监禁，或由于验证人数量已达上限，而您的验证人权重排名在数量上限以外.

如果您的验证人被监禁，可以按以下步骤操作来恢复：

- 如果`irita`没有运行，请重新启动：

  ```bash
  irita start
  ```

- 等待节点赶上最新的区块，检查验证人会被监禁到什么时间：

  ```bash
  # 查询验证人节点共识公钥
  irita tendermint show-validator --home=<irita-home>
  
  # 使用共识公钥查询节点状态
  irita query slashing signing-info [validator-conspub]
  ```

  您将可以看到 `Jailed Until` 的时间，只有在该时间之后，您才可以执行接下来的步骤。

- 如果当前时间已经超过了 `Jailed Until`，即可执行解禁操作：

  ```bash
  irita tx slashing  unjail --from=<key-name> --fees=0.3uirita --chain-id=irita
  ```

- 再次检查您的验证人，看看您的投票权是否恢复。

  ```bash
  irita status
  ```

  您可能会注意到您的投票权比以前低，那是因为您被惩罚了。

### 6. irita 异常退出：too many open files

Linux 可以打开（每个进程）的默认文件数是 `1024`，而 `irita` 进程会打开超过1024个文件，进而导致进程崩溃。一个快速的解决方法是执行 `ulimit -n 4096`（增加允许的打开文件数量，仅对当前会话有效），然后使用 `irita start` 重新启动。如果您使用的 systemd 或其他进程管理器来启动 `irita `，则最好在该级别进行一些配置。好在该级别进行一些配置。

- 示例`systemd`配置：

  ```toml
  # /etc/systemd/system/irita.service
  [Unit]
  Description=irita Node
  After=network.target
  
  [Service]
  Type=simple
  User=ubuntu
  WorkingDirectory=/home/ubuntu
  ExecStart=/home/ubuntu/go/bin/irita start
  Restart=on-failure
  RestartSec=3
  LimitNOFILE=65535
  
  [Install]
  WantedBy=multi-user.target
  ```

- 在 Ubuntu 系统中修改全局 ulimit 示例：

  ```bash
  # Edit limits.conf
  vim /etc/security/limits.conf
  
  # Append the following lines at the bottom
  * hard nofile 65535
  * soft nofile 65535
  root hard nofile 65535
  root soft nofile 65535
  
  # Reboot the system
  reboot
  
  # Re-login & Check whether ulimit is updated to 65535
  ulimit -n
  ```

### 7. Uptime始终为0％，即使节点已经完成同步

比较两个`Consensus Pubkey`：

- 从[区块浏览器](https://irishub.iobscan.io/#/staking)中，您可以在“验证人详情”页中找到该验证人声明的`Consensus Pubkey`。
- 通过 `irita tendermint show-validator --home=<irita-home>` 检查正在使用的`Consensus Pubkey`。

如果它们不相同，则意味着您正在运行的只是一个普通全节点，而不是验证人节点。

#### 最好的情况是您已经备份了Tendermint 密钥

那么您可以执行以下操作：

- 停止节点
- 用您的备份替换当前的 `<irita-home>/config/priv_validator.json`
- 通过`irita tendermint show-validator --home=<irita-home>` 确认 `Consensus Pubkey` 是正确的
- 启动节点
- 完成同步后，检查`voting_power`现在应该大于0：`irita status`

#### 如果我丢失了Tendermint密钥怎么办

这意味着您 **永远失去了您的验证人！**您只能创建一个新的验证人，并记得做好备份。

