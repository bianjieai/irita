# Start wasm on irita

### Compile smart contract

```shell
# install wasm-pack
curl https://rustwasm.github.io/wasm-pack/installer/init.sh -sSf | sh 

# clone example
git clone https://github.com/confio/cosmwasm-examples
cd cosmwasm-examples/escrow
git checkout escrow-0.2.1

# complie
wasm-pack build
du -h ./pkg/cw_escrow_bg.wasm

# cp wasm to work dir
mkdir ~/wasm-test
cp ./pkg/cw_escrow_bg.wasm ~/wasm-test/contract.wasm
```

### Init testnet

```shell
cd ~/wasm-test
irita testnet --v 1 --chain-id test
irita start --home mytestnet/node0/irita
```

## deploy the contract

```shell
# add key node1
iritacli keys add node1

# both should be empty
iritacli query wasm list-code
iritacli query wasm list-contracts

# upload and see we create code 1
# gas is huge due to wasm size... but auto-zipping reduced this from 800k to around 260k
iritacli tx wasm store contract.wasm --gas 3000000 --from node0  -y -b block --chain-id test
iritacli query wasm list-code
```

### Instantiating the contract

```shell
# instantiate contract and verify
INIT="{\"arbiter\":\"$(iritacli keys show node0 -a)\", \"recipient\":\"$(iritacli keys show node1 -a)\", \"end_time\":0, \"end_height\":0}"

iritacli tx wasm instantiate 1 "$INIT" --from node0 --amount=50000stake  -y --from node0 -b block --chain-id test

# check the contract state (and account balance)
iritacli query wasm list-contracts
iritacli query wasm contract <contract-id>
iritacli query account <contract-id>
```