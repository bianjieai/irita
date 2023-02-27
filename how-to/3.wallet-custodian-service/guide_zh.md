# 硬件钱包集成文档

我们提供基于金融级可信硬件和软件的自托管解决方案，为数字资产服务提供商量身定制，并确保数字资产从发行、托管到使用的端到端安全，并符合监管机构对服务提供商的资产托管合规要求。联系我们了解更多信息。

## 一、 支持平台

| 链              | 代币 | 操作                                                  |
| --------------- | ---- | ----------------------------------------------------- |
| CosmosHub       | ATOM | 本链转账、跨链转账                                    |
| IRISHub         | IRIS | 本链转账、跨链转账、Token 模块、NFT模块、Coinswap模块 |
| IRISHub Testnet | NYAN | 本链转账、跨链转账、Token 模块、NFT模块、Coinswap模块 |
| IRITA           | UGAS | 本链转账、跨链转账、Token 模块、NFT模块               |
| IRITA Testnet   | UGAS | 本链转账、跨链转账、Token 模块、NFT模块               |

 

## 二、 支持操作

\- Token 模块：发行（Issue）、增发（Mint）、销毁（Burn）

\- NFT 模块：创建类别（Issue Denom）、转让类别（Transfer Denom）、发行 NFT（Mint）、转让 NFT（Transfer）、编辑 NFT（Edit）、销毁 NFT（Burn）

\- Coinswap 模块：注入流动性（Add Liquidity）、取回流动性（Remove Liquidity）、交易（Swap）

 

## 三、 API接口

1、创建地址（与原版本相关）：

| URL地址                                            | 类型 | 参数                                                         |
| -------------------------------------------------- | ---- | ------------------------------------------------------------ |
| https://www.xxx.com/wallet/irishub-testnet/address | POST | walletId: Stringplatform: String symbol: StringaddressName: String |

​	

2、提币（与原版本相关）：

| URL地址                                             | 类型 | 参数                                                         |
| --------------------------------------------------- | ---- | ------------------------------------------------------------ |
| https://www.xxx.com/wallet/irishub-testnet/withdraw | POST | walletId：StringrequestId：StringfromAddress：StringtoAddress：Stringamount：BigDecimalsymbol：StringgasFee：BigDecimalnote：String |

 

3、ibcTransfer: 

| URL地址                                                | 类型 | 参数                                                         |
| ------------------------------------------------------ | ---- | ------------------------------------------------------------ |
| https://www.xxx.com/wallet/irishub-testnet/ibcTransfer | POST | walletId：StringrequestId：StringownerAddress：StringtoAddress：Stringamount：BigDecimalsymbol：StringgasFee：BigDecimalnote：String |

 

4、链上message:

| URL地址                                            | 类型 | 参数                                                         |
| -------------------------------------------------- | ---- | ------------------------------------------------------------ |
| https://www.xxx.com/wallet/irishub-testnet/message | POST | walletId：StringrequestId：StringownerAddress：Stringmessage：JsongasFee：BigDecimalnote：StringmsgDesc：String |

​	参数中：message为JSON格式的消息体微调，去掉@type属性，增加type属性，type的value值做了简化，由“/irismod.token.MsgBurnToken”简化为：“MsgBurnToken”；msgDesc为type的value值MsgBurnToken；

​	注意：JSON体中需包含所有message的属性，可以为空，不能没有，没有则会解析报错。

| 类型               | msgDesc            | JSON体示例                                                   |
| ------------------ | ------------------ | ------------------------------------------------------------ |
| MsgIssueToken      | MsgIssueToken      | {"type": "MsgIssueToken","symbol": "rig","name": "rigToken","scale": 6,"min_unit": "rig","initial_supply": "100000","max_supply": "9999999","mintable": true,"owner": "iaa1rezgxhee2atzg6v2la0j8jsmflamzj9ypux8lg"} |
| MsgMintToken       | MsgMintToken       | {	"type": "MsgMintToken",	"symbol": "rig",	"amount": "99",	"to": "iaa1lekqsx4eh42grqey7hk6w74jpfkn36kfpcedgv",	"owner": "iaa1zq26lwpvc2q74kkhsy3s3cl77cpl5typrgqmfr"} |
| MsgBurnToken       | MsgBurnToken       | {	"type": "MsgBurnToken",	"symbol": "rig",	"amount": "1",	"sender": "iaa1lekqsx4eh42grqey7hk6w74jpfkn36kfpcedgv"} |
| MsgAddLiquidity    | MsgAddLiquidity    | {	"type": "MsgAddLiquidity",	"max_token": {		"denom": "rig",		"amount": "500000"	},	"exact_standard_amt": "1",	"min_liquidity": "1",	"deadline": "1655205453",	"sender": "iaa1zq26lwpvc2q74kkhsy3s3cl77cpl5typrgqmfr"} |
| MsgRemoveLiquidity | MsgRemoveLiquidity | {	"type": "MsgRemoveLiquidity",	"withdraw_liquidity": {		"denom": "lpt-11",		"amount": "1000"	},	"min_token": "1000",	"min_standard_amt": "1000",	"deadline": "1655202403",	"sender": "iaa1zq26lwpvc2q74kkhsy3s3cl77cpl5typrgqmfr"} |
| MsgSwapOrder       | MsgSwapOrder       | {	"type": "MsgSwapOrder",	"input": {		"address": "iaa1zq26lwpvc2q74kkhsy3s3cl77cpl5typrgqmfr",		"coin": {			"denom": "unyan",			"amount": "100000"		}	},	"output": {		"address": "iaa1zq26lwpvc2q74kkhsy3s3cl77cpl5typrgqmfr",		"coin": {			"denom": "rig",			"amount": "90000"		}	},	"deadline": "1655203526",	"is_buy_order": true} |
| MsgIssueDenom      | MsgIssueDenom      | {  "type": "MsgIssueDenom",  "id": "amber129c73c3017e2b0b884afb7d4cc9df069d",  "name": "植物园优惠券",  "sender": "iaa1rezgxhee2atzg6v2la0j8jsmflamzj9ypux8lg",  "schema": "{\"type\":\"/uptick.coupon\"}",  "symbol": "",  "mint_restricted": false,  "update_restricted": false,  "description": "创建优惠券分类",  "uri": "",  "uri_hash": "",  "data": ""} |
| MsgTransferDenom   | MsgTransferDenom   | {	"type": "MsgTransferDenom",	"id": "cardamber0550b544fd5d9ea7618ff4ce671e8881",	"sender": "iaa1zq26lwpvc2q74kkhsy3s3cl77cpl5typrgqmfr",	"recipient": "iaa1lekqsx4eh42grqey7hk6w74jpfkn36kfpcedgv"} |
| MsgMintNFT         | MsgMintNFT         | {  "type": "MsgMintNFT",  "sender": "iaa1rezgxhee2atzg6v2la0j8jsmflamzj9ypux8lg",  "recipient": "iaa1rezgxhee2atzg6v2la0j8jsmflamzj9ypux8lg",  "denom_id": "amber129c73c3017e2b0b884afb7d4cc9df069d",  "id": "amberpeter8cb924i4ldbj3ldwrpgc21dgk",  "uri": "https://zsvideo.86itn.cn/20220621172212836300572.png",  "data": "https://zsvideo.86itn.cn/20220621172212836300572.png",  "name": "概念风景插画","uri_hash": "1136e58ca7a04f6988a1f592f8b94a62f43385e6b91fd14a6df25e9257a21f0c"} |
| MsgEditNFT         | MsgEditNFT         | {  "type": "MsgEditNFT",  "sender": "iaa1rezgxhee2atzg6v2la0j8jsmflamzj9ypux8lg",  "denom_id": "amber129c73c3017e2b0b884afb7d4cc9df069d",  "id": "amberpeter8cb924i4ldbj3ldwrpgc21dgk",  "uri": "https://zsvideo.86itn.cn/20220621172212836300572.png",  "data": "https://zsvideo.86itn.cn/20220621172212836300572.png",  "name": "概念风景插画2",  "uri_hash": "1136e58ca7a04f6988a1f592f8b94a62f43385e6b91fd14a6df25e9257a21f0c" } |
| MsgTransferNFT     | MsgTransferNFT     | {  "type": "MsgTransferNFT",  "sender": "iaa1rezgxhee2atzg6v2la0j8jsmflamzj9ypux8lg",  "recipient": "iaa1n2rxxskf4ns4jqwce5maqrnygmkl5w70k83e47",  "denom_id": "amber129c73c3017e2b0b884afb7d4cc9df069d",  "id": "amberpeter8cb924i4ldbj3ldwrpgc21dgk",  "uri": "https://zsvideo.86itn.cn/20220621172212836300572.png",  "data": "https://zsvideo.86itn.cn/20220621172212836300572.png",  "name": "概念风景插画","uri_hash": "1136e58ca7a04f6988a1f592f8b94a62f43385e6b91fd14a6df25e9257a21f0c"} |
| MsgBurnNFT         | MsgBurnNFT         | {	"type": "MsgBurnNFT",	"id": "amberpeter8cb924i4ldbj3ldwrpgc21dgk ",	"denom_id": "amber129c73c3017e2b0b884afb7d4cc9df069d",	"sender": "iaa1rezgxhee2atzg6v2la0j8jsmflamzj9ypux8lg"} |

5、Cosmos策略设置:

Cosmos策略设置分为2种：转账策略和消息策略；转账策略同以往版本中，用于设置币种不同金额的签数；消息策略，设置某平台下某类型消息的固定签章盾数，如图：

![image](https://user-images.githubusercontent.com/31681438/218011598-ee9a5b3b-95ea-4a83-aede-94fda6482f47.png)

其中，名称为消息类型的最后一部分，展示在签章盾上，消息类型与message中@type完全一致，例如：“/irismod.token.MsgBurnToken”，上图设置中，消息策略0代表该类型0签。

 

6、Cosmos代币管理：

​	增加代币地址：如果是IBC转账的代币，该栏目为链上代币的demon，例如：“ibc/96F569ACD14BC391AF3A9A81C1ACB03E8BF3F8FDF7A97CAC05EBAAE2095397B6”，建议代币名称为“ibc-xxx”，用于区分资产，非IBC转账代币，地址与代币名称相同。在进行代币转账或ibc转账时，symbol参数为代币名称。如图：

![image](https://user-images.githubusercontent.com/31681438/218011659-2aeec91f-f72c-41a3-a884-ef9551054f5d.png)
