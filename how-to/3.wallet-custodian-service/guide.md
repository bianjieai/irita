# Hardware Wallet Integration Guide

Hardware Wallet Integration Documentation for Cosmos Networks

We provide a self-custody solution based on finance-grade trusted hardware and software, tailored for digital asset service providers, and ensure end-to-end security of digital assets from issuance, and custody, to utilization. At the same time, it meets the regulator's asset custody compliance requirements for service providers. Contact us to learn more.

1. **Supported Platforms**

|**Chain name**|**Token**|**Operation**|
|:----|:----|:----|
|CosmosHub|ATOM|Native/cross-chain transfers|
|IRISHub|IRIS|Native/cross-chain transfers, token module, nft module and coinswap module|
|IRISHub Testnet|NYAN|Native/cross-chain transfers, token module, nft module and coinswap module|
|IRITA|UGAS|Native/cross-chain transfers, token module, nft module|
|IRITA Testnet|UGAS|Native/cross-chain transfers, token module, nft module|

2. **Supported Operations**

* Token module: Issue, mint, and burn
* nft module: Issue denom, transfer denom, mint, transfer, edit, and burn
* coinswap module: Add liquidity, remove liquidity, and swap
 
3. **API Interfaces**

3.1 Create address (related to the original version)

|URL address|Type|Parameters|
|:----|:----|:----|
|[https://www.xxx.com/wallet/irishub-testnet/address](https://www.xxx.com/wallet/irishub-testnet/address)|POST|walletId: Stringplatform: String symbol: StringaddressName: String|



3.2 Withdraw token (related to the original version)

|URL address|Type|Parameters|
|:----|:----|:----|
|[https://www.xxx.com/wallet/irishub-testnet/withdraw](https://www.xxx.com/wallet/irishub-testnet/withdraw)|POST|walletId：StringrequestId：StringfromAddress：StringtoAddress：Stringamount：BigDecimalsymbol：StringgasFee：BigDecimalnote：String|

3.3 ibcTransfer:

|URL address|Type|Parameters|
|:----|:----|:----|
|[https://www.xxx.com/wallet/irishub-testnet/ibcTransfer](https://www.xxx.com/wallet/irishub-testnet/ibcTransfer)|POST|walletId：StringrequestId：StringownerAddress：StringtoAddress：Stringamount：BigDecimalsymbol：StringgasFee：BigDecimalnote：String|

3.4 on-chain message:

|URL address|Type|Parameters|
|:----|:----|:----|
|[https://www.xxx.com/wallet/irishub-testnet/message](https://www.xxx.com/wallet/irishub-testnet/message)|POST|walletId：StringrequestId：StringownerAddress：Stringmessage：JsongasFee：BigDecimalnote：StringmsgDesc：String|

In the parameters:

* The message is a JSON format message body that needs to be adjusted. The "@type" attribute should be removed, and a "type" attribute should be added. The value of "type" has been simplified from "/irismod.token.MsgBurnToken" to "MsgBurnToken".
* The msgDesc is the value of the "type" attribute, which is "MsgBurnToken" in this case.
**Note that all the attributes of the message must be included in the JSON body, which can be empty but cannot be missing. Otherwise, parsing errors may occur.*

|**Type**|**msgDesc**|**Example of JSON**|
|:----|:----|:----|
|MsgIssueToken|MsgIssueToken|{"type": "MsgIssueToken","symbol": "rig","name": "rigToken","scale": 6,"min_unit": "rig","initial_supply": "100000","max_supply": "9999999","mintable": true,"owner": "iaa1rezgxhee2atzg6v2la0j8jsmflamzj9ypux8lg"}|
|MsgMintToken|MsgMintToken|{ "type": "MsgMintToken", "symbol": "rig", "amount": "99", "to": "iaa1lekqsx4eh42grqey7hk6w74jpfkn36kfpcedgv", "owner": "iaa1zq26lwpvc2q74kkhsy3s3cl77cpl5typrgqmfr"}|
|MsgBurnToken|MsgBurnToken|{ "type": "MsgBurnToken", "symbol": "rig", "amount": "1", "sender": "iaa1lekqsx4eh42grqey7hk6w74jpfkn36kfpcedgv"}|
|MsgAddLiquidity|MsgAddLiquidity|{ "type": "MsgAddLiquidity", "max_token": { "denom": "rig", "amount": "500000" }, "exact_standard_amt": "1", "min_liquidity": "1", "deadline": "1655205453", "sender": "iaa1zq26lwpvc2q74kkhsy3s3cl77cpl5typrgqmfr"}|
|MsgRemoveLiquidity|MsgRemoveLiquidity|{ "type": "MsgRemoveLiquidity", "withdraw_liquidity": { "denom": "lpt-11", "amount": "1000" }, "min_token": "1000", "min_standard_amt": "1000", "deadline": "1655202403", "sender": "iaa1zq26lwpvc2q74kkhsy3s3cl77cpl5typrgqmfr"}|
|MsgSwapOrder|MsgSwapOrder|{ "type": "MsgSwapOrder", "input": { "address": "iaa1zq26lwpvc2q74kkhsy3s3cl77cpl5typrgqmfr", "coin": { "denom": "unyan", "amount": "100000" } }, "output": { "address": "iaa1zq26lwpvc2q74kkhsy3s3cl77cpl5typrgqmfr", "coin": { "denom": "rig", "amount": "90000" } }, "deadline": "1655203526", "is_buy_order": true}|
|MsgIssueDenom|MsgIssueDenom|{ "type": "MsgIssueDenom", "id": "amber129c73c3017e2b0b884afb7d4cc9df069d", "name": "Botanical Garden coupon", "sender": "iaa1rezgxhee2atzg6v2la0j8jsmflamzj9ypux8lg", "schema": "{"type":"/uptick.coupon"}", "symbol": "", "mint_restricted": false, "update_restricted": false, "description": "create coupon category", "uri": "", "uri_hash": "", "data": ""}|
|MsgTransferDenom|MsgTransferDenom|{ "type": "MsgTransferDenom", "id": "cardamber0550b544fd5d9ea7618ff4ce671e8881", "sender": "iaa1zq26lwpvc2q74kkhsy3s3cl77cpl5typrgqmfr", "recipient": "iaa1lekqsx4eh42grqey7hk6w74jpfkn36kfpcedgv"}|
|MsgMintNFT|MsgMintNFT|{ "type": "MsgMintNFT", "sender": "iaa1rezgxhee2atzg6v2la0j8jsmflamzj9ypux8lg", "recipient": "iaa1rezgxhee2atzg6v2la0j8jsmflamzj9ypux8lg", "denom_id": "amber129c73c3017e2b0b884afb7d4cc9df069d", "id": "amberpeter8cb924i4ldbj3ldwrpgc21dgk", "uri": "[https://zsvideo.86itn.cn/20220621172212836300572.png](https://zsvideo.86itn.cn/20220621172212836300572.png)", "data": "[https://zsvideo.86itn.cn/20220621172212836300572.png](https://zsvideo.86itn.cn/20220621172212836300572.png)", "name": "conceptual landscape illustration","uri_hash": "1136e58ca7a04f6988a1f592f8b94a62f43385e6b91fd14a6df25e9257a21f0c"}|
|MsgEditNFT|MsgEditNFT|{ "type": "MsgEditNFT", "sender": "iaa1rezgxhee2atzg6v2la0j8jsmflamzj9ypux8lg", "denom_id": "amber129c73c3017e2b0b884afb7d4cc9df069d", "id": "amberpeter8cb924i4ldbj3ldwrpgc21dgk", "uri": "[https://zsvideo.86itn.cn/20220621172212836300572.png](https://zsvideo.86itn.cn/20220621172212836300572.png)", "data": "[https://zsvideo.86itn.cn/20220621172212836300572.png](https://zsvideo.86itn.cn/20220621172212836300572.png)", "name": "conceptual landscape illustration2", "uri_hash": "1136e58ca7a04f6988a1f592f8b94a62f43385e6b91fd14a6df25e9257a21f0c" }|
|MsgTransferNFT|MsgTransferNFT|{ "type": "MsgTransferNFT", "sender": "iaa1rezgxhee2atzg6v2la0j8jsmflamzj9ypux8lg", "recipient": "iaa1n2rxxskf4ns4jqwce5maqrnygmkl5w70k83e47", "denom_id": "amber129c73c3017e2b0b884afb7d4cc9df069d", "id": "amberpeter8cb924i4ldbj3ldwrpgc21dgk", "uri": "[https://zsvideo.86itn.cn/20220621172212836300572.png](https://zsvideo.86itn.cn/20220621172212836300572.png)", "data": "[https://zsvideo.86itn.cn/20220621172212836300572.png](https://zsvideo.86itn.cn/20220621172212836300572.png)", "name": "conceptual landscape illustration","uri_hash": "1136e58ca7a04f6988a1f592f8b94a62f43385e6b91fd14a6df25e9257a21f0c"}|
|MsgBurnNFT|MsgBurnNFT|{ "type": "MsgBurnNFT", "id": "amberpeter8cb924i4ldbj3ldwrpgc21dgk ", "denom_id": "amber129c73c3017e2b0b884afb7d4cc9df069d", "sender": "iaa1rezgxhee2atzg6v2la0j8jsmflamzj9ypux8lg"}|


3.5 Cosmos Policy Settings

There are two types of policy settings in Cosmos: transfer policy and message policy. The transfer policy is used to set the number of signatures required for different amounts of currency, as in previous versions. The message policy sets a fixed number of signature shields for a certain type of message on a certain platform, as shown in the figure below:

![image](https://user-images.githubusercontent.com/31681438/218011598-ee9a5b3b-95ea-4a83-aede-94fda6482f47.png)

The name of the message type is the last part, which is displayed on the signature shield. The message type is exactly the same as the "@type" in the message. For example, "/irismod.token.MsgBurnToken". In the settings shown in the figure, message policy 0 represents 0 signatures for this type.

6. Cosmos Token Management
​Add Token Address: If the token is for IBC transfer, the field represents the denomination of the on-chain token, for example, "ibc/96F569ACD14BC391AF3A9A81C1ACB03E8BF3F8FDF7A97CAC05EBAAE2095397B6". It is recommended to name the token "ibc-xxx" to distinguish assets. For non-IBC transfer tokens, the address is the same as the token name. When conducting token transfers, or IBC transfers, the "symbol" parameter is the token name, as shown in the figure.

![image](https://user-images.githubusercontent.com/31681438/218011659-2aeec91f-f72c-41a3-a884-ef9551054f5d.png)

