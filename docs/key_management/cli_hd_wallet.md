<!--
order: 1
-->

# HD 钱包（命令行）

HD 钱包即分层确定性钱包，是基于 [BIP44]() 的秘钥管理方式。HD 钱包为多资产、多秘钥的管理提供了极大的便利。HD 钱包的优势在于其可以通过一个主密钥推导出无数的分层（树状结构）子私钥。

主密钥由一个随机的种子产生。在 IRITA 中，该种子由24个有序单词构成的 `助记词` 生成。

当秘钥丢失时，可由此助记词恢复出所有的秘钥。

## 钱包初始化

```bash
iritawallet init
```

>**_提示_**：务必将生成的助记词安全存储。

## 修改密码

通过密码或助记词修改密码

```bash
iritawallet update [flags]
```

## 恢复钱包

通过助记词恢复钱包

```bash
iritawallet init --recover
```

## 创建密钥

```bash
iritawallet keys create <name> [flags]
```

## 查询密钥

通过名称或地址查询密钥信息

```bash
iritawallet keys show <name|address> [flags]
```

## 查询所有密钥

```bash
iritawallet keys list
```

## 导出密钥

通过名称或地址导出密钥，导出后的密钥文件可通过`irita keys import`导入到业务系统中使用。

```bash
iritawallet keys export <name|address> [flags]
```
