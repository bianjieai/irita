<!--
order: 1
-->

# 积分

## issue

发行积分。

```bash
 irita tx token issue [flags]
 irita tx token issue --name="Kitty Token" --symbol="kitty" --min-unit="kitty" --scale=0 --initial-supply=100000000000 --max-supply=1000000000000 --mintable=true --from=<key-name> --chain-id=<chain-id> --fees=<fee>
```

**标志：**

| 名称，速记       | 类型    | 必须 | 默认          | 描述                                                                     |
| ---------------- | ------- | ---- | ------------- | ------------------------------------------------------------------------ |
| --name           | string  | 是   |               | 积分的名称；unicode字符，最大长度为32字节，例如 "IRITA Credit"             |
| --symbol         | string  | 是   |               | 积分的唯一标识符；长度在3到8之间，字母数字字符，以字符开始，不区分大小写 |
| --initial-supply | uint64  | 是   |               | 积分的初始供应；增发前的数量不应超过1000亿                               |
| --max-supply     | uint64  |      | 1000000000000 | 积分的最大供应，总供应不能超过最大供应；增发前的数量不应超过1万亿        |
| --scale          | uint8   | 是   |               | 积分的精度，最多可以有18位小数；为0将默认到18位小数                      |
| --min-unit       | string  | 是   |               | 最小单位；长度在3到10之间，字母数字字符，以字符开始，不区分大小写        |
| --mintable       | boolean |      | false         | 发行后是否可以增发                                                       |

### 发行积分示例

```bash
irita tx token issue --name="Kitty Token" --symbol="kitty" --min-unit="kitty" --scale=0 --initial-supply=100000000000 --max-supply=1000000000000 --mintable=true --from=node0 --chain-id=test  --home node0 -y
```

结果

```json
The token issuance transaction will consume extra fee: 13015000000uirita
{"height":"222","txhash":"804D3912D853A5286AF91B8270D5F2697C9AE429A8B567578785370D5D9BC413","codespace":"","code":0,"data":"0A0D0A0B69737375655F746F6B656E","raw_log":"[{\"events\":[{\"type\":\"issue_token\",\"attributes\":[{\"key\":\"symbol\",\"value\":\"kitty\"},{\"key\":\"creator\",\"value\":\"iaa1t07s27vgvgczpsvu5z75703azmmc9wcmje452z\"}]},{\"type\":\"message\",\"attributes\":[{\"key\":\"action\",\"value\":\"issue_token\"},{\"key\":\"sender\",\"value\":\"iaa1t07s27vgvgczpsvu5z75703azmmc9wcmje452z\"},{\"key\":\"sender\",\"value\":\"iaa183rfa8tvtp6ax7jr7dfaf7ywv870sykxxykejp\"},{\"key\":\"sender\",\"value\":\"iaa183rfa8tvtp6ax7jr7dfaf7ywv870sykxxykejp\"},{\"key\":\"module\",\"value\":\"token\"},{\"key\":\"sender\",\"value\":\"iaa1t07s27vgvgczpsvu5z75703azmmc9wcmje452z\"}]},{\"type\":\"transfer\",\"attributes\":[{\"key\":\"recipient\",\"value\":\"iaa183rfa8tvtp6ax7jr7dfaf7ywv870sykxxykejp\"},{\"key\":\"sender\",\"value\":\"iaa1t07s27vgvgczpsvu5z75703azmmc9wcmje452z\"},{\"key\":\"amount\",\"value\":\"13015000000uirita\"},{\"key\":\"recipient\",\"value\":\"iaa1k83ewmsh9t5ra60urmcj5jc8ev2agmfez0jawf\"},{\"key\":\"sender\",\"value\":\"iaa183rfa8tvtp6ax7jr7dfaf7ywv870sykxxykejp\"},{\"key\":\"amount\",\"value\":\"5206000000uirita\"},{\"key\":\"recipient\",\"value\":\"iaa1t07s27vgvgczpsvu5z75703azmmc9wcmje452z\"},{\"key\":\"sender\",\"value\":\"iaa183rfa8tvtp6ax7jr7dfaf7ywv870sykxxykejp\"},{\"key\":\"amount\",\"value\":\"100000000000kitty\"}]}]}]","logs":[{"msg_index":0,"log":"","events":[{"type":"issue_token","attributes":[{"key":"symbol","value":"kitty"},{"key":"creator","value":"iaa1t07s27vgvgczpsvu5z75703azmmc9wcmje452z"}]},{"type":"message","attributes":[{"key":"action","value":"issue_token"},{"key":"sender","value":"iaa1t07s27vgvgczpsvu5z75703azmmc9wcmje452z"},{"key":"sender","value":"iaa183rfa8tvtp6ax7jr7dfaf7ywv870sykxxykejp"},{"key":"sender","value":"iaa183rfa8tvtp6ax7jr7dfaf7ywv870sykxxykejp"},{"key":"module","value":"token"},{"key":"sender","value":"iaa1t07s27vgvgczpsvu5z75703azmmc9wcmje452z"}]},{"type":"transfer","attributes":[{"key":"recipient","value":"iaa183rfa8tvtp6ax7jr7dfaf7ywv870sykxxykejp"},{"key":"sender","value":"iaa1t07s27vgvgczpsvu5z75703azmmc9wcmje452z"},{"key":"amount","value":"13015000000uirita"},{"key":"recipient","value":"iaa1k83ewmsh9t5ra60urmcj5jc8ev2agmfez0jawf"},{"key":"sender","value":"iaa183rfa8tvtp6ax7jr7dfaf7ywv870sykxxykejp"},{"key":"amount","value":"5206000000uirita"},{"key":"recipient","value":"iaa1t07s27vgvgczpsvu5z75703azmmc9wcmje452z"},{"key":"sender","value":"iaa183rfa8tvtp6ax7jr7dfaf7ywv870sykxxykejp"},{"key":"amount","value":"100000000000kitty"}]}]}],"info":"","gas_wanted":"400000","gas_used":"153279","tx":null,"timestamp":""}
```

## edit

编辑存在的积分。可编辑的属性包括：名称、最大供应以及可增发性

```bash
irita tx token edit [symbol] [flags]
irita tx token edit <symbol> --name="Cat Token" --max-supply=100000000000 --mintable=true --from=<key-name> --chain-id=<chain-id> --fees=<fee>
```

**标志：**

| 名称，速记   | 类型   | 必须 | 默认  | 描述                                                  |
| ------------ | ------ | ---- | ----- | ----------------------------------------------------- |
| --name       | string |      |       | 积分名称，为空将不更新                                |
| --max-supply | uint   |      | 0     | 积分的最大供应量，应不小于当前的总供应量，为0将不更新 |
| --mintable   | bool   |      | false | 积分是否可以增发，默认为false                         |

### 编辑积分示例

```bash
irita tx token edit kitty --name="Cat" --max-supply=100000000000000 --mintable=true --from=node0 --chain-id=test --home node0 -y
```

结果

```json
{"height":"234","txhash":"A74F90F08F45D880B1BB8A1C8F25BA18DD0A5B8CBE781F8C1FDD756C2A76107D","codespace":"","code":0,"data":"0A0C0A0A656469745F746F6B656E","raw_log":"[{\"events\":[{\"type\":\"edit_token\",\"attributes\":[{\"key\":\"symbol\",\"value\":\"kitty\"},{\"key\":\"owner\",\"value\":\"iaa1t07s27vgvgczpsvu5z75703azmmc9wcmje452z\"}]},{\"type\":\"message\",\"attributes\":[{\"key\":\"action\",\"value\":\"edit_token\"},{\"key\":\"module\",\"value\":\"token\"},{\"key\":\"sender\",\"value\":\"iaa1t07s27vgvgczpsvu5z75703azmmc9wcmje452z\"}]}]}]","logs":[{"msg_index":0,"log":"","events":[{"type":"edit_token","attributes":[{"key":"symbol","value":"kitty"},{"key":"owner","value":"iaa1t07s27vgvgczpsvu5z75703azmmc9wcmje452z"}]},{"type":"message","attributes":[{"key":"action","value":"edit_token"},{"key":"module","value":"token"},{"key":"sender","value":"iaa1t07s27vgvgczpsvu5z75703azmmc9wcmje452z"}]}]}],"info":"","gas_wanted":"400000","gas_used":"60644","tx":null,"timestamp":""}
```

## mint

增发积分到指定地址。

```bash
irita tx token mint [symbol] [flags]
irita tx token mint <symbol> --amount=<amount> --to=<to> --from=<key-name> --chain-id=<chain-id> --fees=<fee>
```

**标志：**

| 名称，速记 | 类型   | 必须 | 默认 | 描述                                       |
| ---------- | ------ | ---- | ---- | ------------------------------------------ |
| --to       | string |      |      | 积分的接收地址，默认为交易发起者的账户地址 |
| --amount   | uint64 | 是   | 0    | 增发的数量                                 |

### 增发积分示例

```bash
irita tx token mint mycredit --to=iaa1lq8ye9aksqtyg2mn46esz9825zuxt5zatm5uxm --amount=1000 --from=node0 --chain-id=test -y --home=testnet/node0/iritacli
```

结果

```json
The token minting transaction will consume extra fee: 1301000000uirita
{"height":"239","txhash":"2CD0A71B2B694C374A47719B0D3B9DFE1D8ADEA62DDD2FB66939D6E29CFE637B","codespace":"","code":0,"data":"0A0C0A0A6D696E745F746F6B656E","raw_log":"[{\"events\":[{\"type\":\"message\",\"attributes\":[{\"key\":\"action\",\"value\":\"mint_token\"},{\"key\":\"sender\",\"value\":\"iaa1t07s27vgvgczpsvu5z75703azmmc9wcmje452z\"},{\"key\":\"sender\",\"value\":\"iaa183rfa8tvtp6ax7jr7dfaf7ywv870sykxxykejp\"},{\"key\":\"sender\",\"value\":\"iaa183rfa8tvtp6ax7jr7dfaf7ywv870sykxxykejp\"},{\"key\":\"module\",\"value\":\"token\"},{\"key\":\"sender\",\"value\":\"iaa1t07s27vgvgczpsvu5z75703azmmc9wcmje452z\"}]},{\"type\":\"mint_token\",\"attributes\":[{\"key\":\"symbol\",\"value\":\"kitty\"},{\"key\":\"amount\",\"value\":\"100\"},{\"key\":\"recipient\",\"value\":\"iaa1t07s27vgvgczpsvu5z75703azmmc9wcmje452z\"}]},{\"type\":\"transfer\",\"attributes\":[{\"key\":\"recipient\",\"value\":\"iaa183rfa8tvtp6ax7jr7dfaf7ywv870sykxxykejp\"},{\"key\":\"sender\",\"value\":\"iaa1t07s27vgvgczpsvu5z75703azmmc9wcmje452z\"},{\"key\":\"amount\",\"value\":\"1301000000uirita\"},{\"key\":\"recipient\",\"value\":\"iaa1k83ewmsh9t5ra60urmcj5jc8ev2agmfez0jawf\"},{\"key\":\"sender\",\"value\":\"iaa183rfa8tvtp6ax7jr7dfaf7ywv870sykxxykejp\"},{\"key\":\"amount\",\"value\":\"520400000uirita\"},{\"key\":\"recipient\",\"value\":\"iaa1t07s27vgvgczpsvu5z75703azmmc9wcmje452z\"},{\"key\":\"sender\",\"value\":\"iaa183rfa8tvtp6ax7jr7dfaf7ywv870sykxxykejp\"},{\"key\":\"amount\",\"value\":\"100kitty\"}]}]}]","logs":[{"msg_index":0,"log":"","events":[{"type":"message","attributes":[{"key":"action","value":"mint_token"},{"key":"sender","value":"iaa1t07s27vgvgczpsvu5z75703azmmc9wcmje452z"},{"key":"sender","value":"iaa183rfa8tvtp6ax7jr7dfaf7ywv870sykxxykejp"},{"key":"sender","value":"iaa183rfa8tvtp6ax7jr7dfaf7ywv870sykxxykejp"},{"key":"module","value":"token"},{"key":"sender","value":"iaa1t07s27vgvgczpsvu5z75703azmmc9wcmje452z"}]},{"type":"mint_token","attributes":[{"key":"symbol","value":"kitty"},{"key":"amount","value":"100"},{"key":"recipient","value":"iaa1t07s27vgvgczpsvu5z75703azmmc9wcmje452z"}]},{"type":"transfer","attributes":[{"key":"recipient","value":"iaa183rfa8tvtp6ax7jr7dfaf7ywv870sykxxykejp"},{"key":"sender","value":"iaa1t07s27vgvgczpsvu5z75703azmmc9wcmje452z"},{"key":"amount","value":"1301000000uirita"},{"key":"recipient","value":"iaa1k83ewmsh9t5ra60urmcj5jc8ev2agmfez0jawf"},{"key":"sender","value":"iaa183rfa8tvtp6ax7jr7dfaf7ywv870sykxxykejp"},{"key":"amount","value":"520400000uirita"},{"key":"recipient","value":"iaa1t07s27vgvgczpsvu5z75703azmmc9wcmje452z"},{"key":"sender","value":"iaa183rfa8tvtp6ax7jr7dfaf7ywv870sykxxykejp"},{"key":"amount","value":"100kitty"}]}]}],"info":"","gas_wanted":"400000","gas_used":"121308","tx":null,"timestamp":""}
```

## transfer

转让积分所有权。

```bash
irita tx token transfer [symbol] [flags]
irita tx token transfer <symbol> --to=<to> --from=<key-name> --chain-id=<chain-id> --fees=<fee>
```

**标志：**

| 名称，速记 | 类型   | 必须 | 默认 | 描述           |
| ---------- | ------ | ---- | ---- | -------------- |
| --to       | string | 是   |      | 新的所有者地址 |

### 转让积分所有权示例

```bash
irita tx token transfer kitty --to=iaa177w2evwnx3uje646k78zxlp82mc9eatuwkdwlh  --from=node0 --chain-id=test --home node0  -y
```

结果

```json
{"height":"253","txhash":"4499EBD8834791FB528B521AE7D4FDFB1EDCEE3A687337704E0CE5F218736674","codespace":"","code":0,"data":"0A160A147472616E736665725F746F6B656E5F6F776E6572","raw_log":"[{\"events\":[{\"type\":\"message\",\"attributes\":[{\"key\":\"action\",\"value\":\"transfer_token_owner\"},{\"key\":\"module\",\"value\":\"token\"},{\"key\":\"sender\",\"value\":\"iaa1t07s27vgvgczpsvu5z75703azmmc9wcmje452z\"}]},{\"type\":\"transfer_token_owner\",\"attributes\":[{\"key\":\"symbol\",\"value\":\"kitty\"},{\"key\":\"owner\",\"value\":\"iaa1t07s27vgvgczpsvu5z75703azmmc9wcmje452z\"},{\"key\":\"dst_owner\",\"value\":\"iaa177w2evwnx3uje646k78zxlp82mc9eatuwkdwlh\"}]}]}]","logs":[{"msg_index":0,"log":"","events":[{"type":"message","attributes":[{"key":"action","value":"transfer_token_owner"},{"key":"module","value":"token"},{"key":"sender","value":"iaa1t07s27vgvgczpsvu5z75703azmmc9wcmje452z"}]},{"type":"transfer_token_owner","attributes":[{"key":"symbol","value":"kitty"},{"key":"owner","value":"iaa1t07s27vgvgczpsvu5z75703azmmc9wcmje452z"},{"key":"dst_owner","value":"iaa177w2evwnx3uje646k78zxlp82mc9eatuwkdwlh"}]}]}],"info":"","gas_wanted":"400000","gas_used":"58630","tx":null,"timestamp":""}

```

## token

查询指定的积分。

```bash
irita query token [command]
```

### 查询积分示例

```bash
 irita query token  token kitty --chain-id test
```

结果

```json
'@type': /irismod.token.Token
initial_supply: "100000000000"
max_supply: "100000000000000"
min_unit: kitty
mintable: true
name: Cat
owner: iaa177w2evwnx3uje646k78zxlp82mc9eatuwkdwlh
scale: 0
symbol: kitty
```

## tokens

查询已发行的所有积分，包括系统原生积分。如指定 `owner` 参数，则查询该 `owner` 发行的积分列表

```bash
irita query token tokens [owner] [flags]
```

### 查询所有积分示例

```bash
 irita query token  tokens  --chain-id test
```

结果

```json
- type: irismod/token/Token
  value:
    initial_supply: "1000000000"
    max_supply: "18446744073709551615"
    min_unit: uirita
    mintable: true
    name: Irita base native token
    scale: 6
    symbol: irita
- type: irismod/token/Token
  value:
    initial_supply: "100000000000"
    max_supply: "100000000000000"
    min_unit: kitty
    mintable: true
    name: Cat
    owner: iaa177w2evwnx3uje646k78zxlp82mc9eatuwkdwlh
    symbol: kitty
- type: irismod/token/Token
  value:
    initial_supply: "1000000000"
    max_supply: "18446744073709551615"
    min_unit: upoint
    mintable: true
    name: Irita base native token
    scale: 6
    symbol: point
```

### 查询指定所有者的积分列表示例

```bash
irita query token tokens iaa177w2evwnx3uje646k78zxlp82mc9eatuwkdwlh --chain-id=test
```

结果

```json
- type: irismod/token/Token
  value:
    initial_supply: "100000000000"
    max_supply: "100000000000000"
    min_unit: kitty
    mintable: true
    name: Cat
    owner: iaa177w2evwnx3uje646k78zxlp82mc9eatuwkdwlh
    symbol: kitty
```

## fee

查询积分相关的费用，包括发行和增发。

```bash
irita query token fee [symbol] [flags]
irita query token fee <symbol>
```

### 查询发行和增发积分费用示例

```bash
irita query token fee credit  --chain-id=test
```

结果

```json
exist: false
issue_fee:
  amount: "8474000000"
  denom: uirita
mint_fee:
  amount: "847000000"
  denom: uirita
```

**_注：_**`exist` 指示此 `symbol` 是否已经存在。
