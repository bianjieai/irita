# VRF Provider

VRF Provider Application

Requires: [vrf](../vrf)

## Configuration

Please refer to `configs/example.toml` for configuration.

## Deployment

#### Configuration Changes

1. Chain configuration: Replace `chain_id` and `uri` with mainnet configuration 
2. `vrf_admin_key`: Use the command `vrf-provider genkey` to generate `vrf_admin_key`. For example, `vrf-provider genkey -p 123456789` where `123456789` is the default password.
3. `contract_services.vrf.contract.addr`: Replace it with the deployed `VRF` contract address.
4. `contract_services.vrf.contract.opt_priv_key`: Replace it with the address of `provider` that you want to set.
#### Startup

Startup command: `vrf-provider start -c configs/example.toml.`
