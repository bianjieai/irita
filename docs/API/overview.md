<!--
order: 1
-->

# 概览

本文档所列 API 为 IRITA REST API（IRITA 轻客户端接口。

::: tip
**提示**：在节点配置文件 `config/app.toml` 中设置 `API.enable` 为 `true`，节点启动时将开启 `API` 服务，即轻客户端。

如同时设置 `api.swagger = true`，则可通过 `127.0.0.1:1317/swagger/` 访问支持的接口列表，并进行交互。
:::

IRITA 轻客户端作为应用系统与 IRITA 链的中间层，屏蔽了链相关的逻辑，方便开发者将传统业务融入 IRITA 网络，以现有应用兼容的方式与 IRITA 平台进行交互。

>**_提示：_** 由于核心模块（积分、数字资产建模、存证、iService）交易相关的 REST API 当前仅返回 _未签名交易_，因此并未在此文档中列出。API 请求示例使用 [curl](https://curl.haxx.se/) 命令发送请求，并用 [jq](https://stedolan.github.io/jq/) 解析工具格式化返回结果。

IRITA REST API 主要包括如下几类：

1. [节点与链](./node.md)
2. [积分](./token.md)
3. [数字资产建模](./nft.md)
4. [存证](./record.md)
5. [iService](./iservice.md)
6. [身份](./identity.md)
