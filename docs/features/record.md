# Record

## 简介

`Record`主要是利用了区块链不可篡改的特性来实现数据的存证功能。用户在链下通过hash算法对元数据进行数据摘要，将摘要信息存储在`irita`系统上，系统生成一个唯一`record_id`返回给用户。当用户需要取证时，可以根据`record_id`来检索元数据的hash，链下进行数据的比对。

## 操作

`Record`模块的相关命令请参考[文档](../cli-client/record.md)
