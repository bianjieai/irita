# Cosmos VRF Provider

VRF Provider Application for Cosmos Enterprise.

Requires: [cosmos-vrf](https://github.com/bianjieai/cosmos-vrf)

## 配置

配置请参考 `configs/example.toml`

## 生产部署

### 配置更改

1. 链配置：`chain_id` 和 `uri` 替换为主网配置
2. `vrf_admin_key`：使用 `vrf-provider genkey`命令生成 `vrf_admin_key`，示例：`vrf-provider genkey -p 123456789`（默认密码是`12345678`）
3. `contract_services.vrf.contract.addr`：替换为自己部署的 `VRF`合约地质
4. `contract_services.vrf.contract.opt_priv_key`：替换为要设置的 `provider` 地质

### 启动

启动命令： `vrf-provider start -c configs/example.toml`
