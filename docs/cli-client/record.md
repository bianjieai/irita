# iritacli record

`record`模块可以轻松实现链上数据的存证功能。

## 可用命令

| 名称                                        | 描述         |
| ------------------------------------------- | ------------ |
| [record create](#iritacli-tx-record-create) | 创建一条记录 |
| [record record](#iritacli-tx-token-edit)    | 查询指定记录 |

## iritacli tx record create

创建一条记录

```bash
iritacli tx record create [digest] [digest-algo] --uri=<uri> --meta=<meta-data> --chain-id=<chain-id> --from=<key-name> --fees=0.3iris
```

**标识：**

| 名称，速记 | 类型   | 必须 | 默认 | 描述             |
| ---------- | ------ | ---- | ---- | ---------------- |
| --uri      | string | 否   | ""   | 元数据的网络地址 |
| --meta     | string | 否   | ""   | 元数据           |

## iritacli q record record

查询指定记录

```bash
iritacli q record record [record-id]
```
