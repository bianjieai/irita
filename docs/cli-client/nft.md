# iritacli nft

`nft`模块用于管理你在`IRITA`上发行的`NFT`资产。

## 可用命令

| 名称                                         | 描述                          |
| -------------------------------------------- | ----------------------------- |
| [nft mint](#iritacli-tx-nft-mint)            | 发行`nft`资产                 |
| [nft edit](#iritacli-tx-nft-edit)            | 编辑`nft`资产                 |
| [nft transfer](#iritacli-tx-nft-transfer)    | 转让`nft`资产                 |
| [nft burn](#iritacli-tx-nft-burn)            | 销毁`nft`资产                 |
| [nft supply](#iritacli-q-nft-supply)         | 查询某类资产的总量            |
| [nft owner](#iritacli-q-nft-owner)           | 查询某人所有的`nft`资产       |
| [nft collection](#iritacli-q-nft-collection) | 查询某中类型下的所有`nft`资产 |
| [nft denoms](#iritacli-q-nft-denoms)         | 查询`nft`资产的所有分类       |
| [nft token](#iritacli-q-nft-token)           | 查询某种`nft`资产             |

## iritacli tx nft mint

发行一个新的`nft`资产。

```bash
iritacli tx nft edit kitty xiaobai --token-uri=<metadata> --from=node0 --chain-id=test --fees=100stake -b=block
```

**标识：**

| 名称，速记  | 类型   | 必须 | 默认            | 描述                 |
| ----------- | ------ | ---- | --------------- | -------------------- |
| --token-uri | string | 否   | [do-not-modify] | 资产元数据的网络地址 |
| --recipient | string | 否   | ""              | 资产接收人           |

## iritacli tx nft edit

编辑`nft`资产的元数据信息(通过修改uri指向)。

```bash
iritacli tx nft edit kitty xiaobai --token-uri=<edit uri> --from=node0 --chain-id=test --fees=100stake -b=block
```

**标识：**

| 名称，速记  | 类型   | 必须 | 默认            | 描述                 |
| ----------- | ------ | ---- | --------------- | -------------------- |
| --token-uri | string | 否   | [do-not-modify] | 资产元数据的网络地址 |

## iritacli tx nft transfer

转让某个`nft`资产。

```bash
iritacli tx nft transfer iaa13ttyazvndnyulurwyajjsd77amprzpkklkx650 cat kitty --from=node0 --chain-id=test --fees=100stake -b=block
```

**标识：**

| 名称，速记  | 类型   | 必须 | 默认            | 描述                 |
| ----------- | ------ | ---- | --------------- | -------------------- |
| --token-uri | string | 否   | [do-not-modify] | 资产元数据的网络地址 |

## iritacli tx nft burn

销毁指定的`nft`资产。

```bash
iritacli tx nft burn kitty xiaobai --from=node0 --chain-id=test --fees=100stake -b=block
```

## iritacli q nft supply

查询某种类型资产的总量

```bash
iritacli q supply kitty
```

同样，也可以查询某个人某一类资产的总量，只需要添加`--owner`标记

```bash
iritacli q supply kitty --owner=<owner-address>
```

## iritacli q nft owner

查询某个人的所有资产

```bash
iritacli q nft owner iaa13ttyazvndnyulurwyajjsd77amprzpkklkx650
```

也可以查询某个人某类所有资产

```bash
iritacli q nft owner iaa13ttyazvndnyulurwyajjsd77amprzpkklkx650 --denom=kitty
```

## iritacli q nft collection

查询某个某个分类下的所有资产

```bash
iritacli q nft collection kitty
```

## iritacli q nft denoms

查询`nft`资产的所有分类，返回分类名称

```bash
iritacli q nft denoms
```

## iritacli q nft token

查询指定的`nft`资产。

```bash
iriscli q nft token kitty xiaobai
```