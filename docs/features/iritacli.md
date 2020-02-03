<!--
order: 1
-->

# IRITA Client

## IRITA CLI

`iritacli` is the tool that enables you to interact with the node that runs on the IRITA network, whether you run it yourself or not.

### Setting up iritacli

The main command used to set up `iritacli` is the following:

```bash
iritacli config <flag> <value>
```

It allows you to set a default value for each given flag.

First, set up the address of the full-node you want to connect to:

```bash
iritacli config node <host>:<port>

# example: iritacli config node https://77.87.106.33:26657
```

If you run your own full-node, just use `tcp://localhost:26657` as the address.

Then, let us set the default value of the `--trust-node` flag:

```bash
iritacli config trust-node true

# Set to true if you trust the full-node you are connecting to, false otherwise
```

Finally, let us set the `chain-id` of the blockchain we want to interact with:

```bash
iritacli config chain-id irita-2
```

### Keys

#### Migrate Keys From Legacy On-Disk Keybase To OS Built-in Secret Store

Older versions of `iritacli` used store keys in the user's home directory. If you are migrating
from an old version of `iritacli` you will need to migrate your old keys into your operating system's
credentials storage by running the following command:

```bash
iritacli keys migrate
```

The command will prompt for every passphrase. If a passphrase is incorrect, it will skip the
respective key.

#### Generate Keys

You'll need an account private and public key pair \(a.k.a. `sk, pk` respectively\) to be able to receive funds, send txs, bond tx, etc.

To generate a new _secp256k1_ key:

```bash
iritacli keys add <account_name>
```

The output of the above command will contain a _seed phrase_. It is recommended to save the _seed
phrase_ in a safe place so that in case you forget the password of the operating system's
credentials store, you could eventually regenerate the key from the seed phrase with the
following command:

```bash
iritacli keys add --recover
```

If you check your private keys, you'll now see `<account_name>`:

```bash
iritacli keys show <account_name>
```

View the validator operator's address via:

```shell
iritacli keys show <account_name> --bech=val
```

You can see all your available keys by typing:

```bash
iritacli keys list
```

View the validator pubkey for your node by typing:

```bash
irita tendermint show-validator
```

Note that this is the Tendermint signing key, _not_ the operator key you will use in delegation transactions.

::: danger Warning
We strongly recommend _NOT_ using the same passphrase for multiple keys. The Tendermint team and the Interchain Foundation will not be responsible for the loss of funds.
:::

#### Generate Multisig Public Keys

You can generate and print a multisig public key by typing:

```bash
iritacli keys add --multisig=name1,name2,name3[...] --multisig-threshold=K new_key_name
```

`K` is the minimum number of private keys that must have signed the
transactions that carry the public key's address as signer.

The `--multisig` flag must contain the name of public keys that will be combined into a
public key that will be generated and stored as `new_key_name` in the local database.
All names supplied through `--multisig` must already exist in the local database. Unless
the flag `--nosort` is set, the order in which the keys are supplied on the command line
does not matter, i.e. the following commands generate two identical keys:

```bash
iritacli keys add --multisig=foo,bar,baz --multisig-threshold=2 multisig_address
iritacli keys add --multisig=baz,foo,bar --multisig-threshold=2 multisig_address
```

Multisig addresses can also be generated on-the-fly and printed through the which command:

```bash
iritacli keys show --multisig-threshold K name1 name2 name3 [...]
```

For more information regarding how to generate, sign and broadcast transactions with a
multi signature account see [Multisig Transactions](#multisig-transactions).

### Tx Broadcasting

When broadcasting transactions, `iritacli` accepts a `--broadcast-mode` flag. This
flag can have a value of `sync` (default), `async`, or `block`, where `sync` makes
the client return a CheckTx response, `async` makes the client return immediately,
and `block` makes the client wait for the tx to be committed (or timing out).

It is important to note that the `block` mode should **not** be used in most
circumstances. This is because broadcasting can timeout but the tx may still be
included in a block. This can result in many undesirable situations. Therefore, it
is best to use `sync` or `async` and query by tx hash to determine when the tx
is included in a block.

### Fees & Gas

Each transaction may either supply fees or gas prices, but not both.

Validator's have a minimum gas price (multi-denom) configuration and they use
this value when when determining if they should include the transaction in a block during `CheckTx`, where `gasPrices >= minGasPrices`. Note, your transaction must supply fees that are greater than or equal to **any** of the denominations the validator requires.

**Note**: With such a mechanism in place, validators may start to prioritize
txs by `gasPrice` in the mempool, so providing higher fees or gas prices may yield higher tx priority.

e.g.

```bash
iritacli tx send ... --fees=50000uatom
```

or

```bash
iritacli tx send ... --gas-prices=0.025uatom
```

### Account

#### Get Tokens

On a testnet, getting tokens is usually done via a faucet.

#### Query Account Balance

After receiving tokens to your address, you can view your account's balance by typing:

```bash
iritacli query account <account_irita>
```

::: warning Note
When you query an account balance with zero tokens, you will get this error: `No account with address <account_irita> was found in the state.` This can also happen if you fund the account before your node has fully synced with the chain. These are both normal.

:::

### Send Tokens

The following command could be used to send coins from one account to another:

```bash
iritacli tx send <sender_key_name_or_address> <recipient_address> 10irita \
  --chain-id=<chain_id>
```

::: warning Note
The `amount` argument accepts the format `<value|coin_name>`.
:::

::: tip Note
You may want to cap the maximum gas that can be consumed by the transaction via the `--gas` flag.
If you pass `--gas=auto`, the gas supply will be automatically estimated before executing the transaction.
Gas estimate might be inaccurate as state changes could occur in between the end of the simulation and the actual execution of a transaction, thus an adjustment is applied on top of the original estimate in order to ensure the transaction is broadcasted successfully. The adjustment can be controlled via the `--gas-adjustment` flag, whose default value is 1.0.
:::

Now, view the updated balances of the origin and destination accounts:

```bash
iritacli query account <account_irita>
iritacli query account <destination_irita>
```

You can also check your balance at a given block by using the `--block` flag:

```bash
iritacli query account <account_irita> --block=<block_height>
```

You can simulate a transaction without actually broadcasting it by appending the
`--dry-run` flag to the command line:

```bash
iritacli tx send <sender_key_name_or_address> <destination_iritaaccaddr> 10irita \
  --chain-id=<chain_id> \
  --dry-run
```

Furthermore, you can build a transaction and print its JSON format to STDOUT by
appending `--generate-only` to the list of the command line arguments:

```bash
iritacli tx send <sender_address> <recipient_address> 10irita \
  --chain-id=<chain_id> \
  --generate-only > unsignedSendTx.json
```

```bash
iritacli tx sign \
  --chain-id=<chain_id> \
  --from=<key_name> \
  unsignedSendTx.json > signedSendTx.json
```

::: tip Note
The `--generate-only` flag prevents `iritacli` from accessing the local keybase.
Thus when such flag is supplied `<sender_key_name_or_address>` must be an address.
:::

You can validate the transaction's signatures by typing the following:

```bash
iritacli tx sign --validate-signatures signedSendTx.json
```

You can broadcast the signed transaction to a node by providing the JSON file to the following command:

```bash
iritacli tx broadcast --node=<node> signedSendTx.json
```

### Query Transactions

#### Matching a Transaction's Hash

You can also query a single transaction by its hash using the following command:

```bash
iritacli query tx [hash]
```