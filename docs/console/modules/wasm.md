<!--
order: 7
-->

# 智能合约

## Store

上传智能合约代码。

```bash
irita tx wasm store [wasm-file] --source [source] --builder [builder]
```

**参数：**

| 名称      | 类型   | 必须 | 默认 | 描述                       |
| --------- | ------ | ---- | ---- | -------------------------- |
| wasm-file | string | 是   |      | 智能合约编译后的二进制文件 |

**标志：**

| 名称，速记                 | 类型   | 必须 | 默认 | 描述                     |
| -------------------------- | ------ | ---- | ---- | ------------------------ |
| --source                   | string | 否   |      | 智能合约源码地址         |
| --builder                  | string | 否   |      | 合法的docker标签         |
| --instantiate-only-address | string | 否   |      | 只有该地址可以初始化合约 |
| --instantiate-everybody    | string | 否   |      | 任何人都可以初始化合约   |

### 上传合约示例

```bash
irita tx wasm store election.wasm --from node0 --chain-id=test --keyring-backend=file --home ./irita/node0/iritacli --fees 6point -b block --gas="auto"
```

**_注_：**合约代码上传完成后，系统会分配一个`code_id`给用户，用于下次初始化该合约。

## Instantiate

初始化合约的状态。

```bash
irita tx wasm instantiate [code_id] [json_encoded_init_args] --label [text] --admin [address] --amount [coins] [flags]
```

**参数：**

| 名称                   | 类型   | 必须 | 默认 | 描述                               |
| ---------------------- | ------ | ---- | ---- | ---------------------------------- |
| code_id                | int64  | 是   |      | 合约上传完成后系统分配的id         |
| json_encoded_init_args | string | 是   |      | 合约初始化函数需要的参数，json格式 |

**标志：**

| 名称，速记 | 类型   | 必须 | 默认 | 描述                            |
| ---------- | ------ | ---- | ---- | ------------------------------- |
| --label    | string | 是   |      | 合约的简要描述                  |
| --admin    | string | 否   |      | 可以执行`migrate`操作的用户地址 |
| --amount   | string | 否   |      | 发送指定金额代币到合约地址      |

### 初始化合约示例

```bash
CODE_ID=1
INIT='{"start":1,"end":100,"candidates":["iaa1qvty8x0c78am8c44zv2n7tgm6gfqt78j0verqa","iaa1zk2tse0pkk87p2v8tcsfs0ytfw3t88kejecye5"]}'
irita tx wasm instantiate "$CODE_ID" "$INIT" --from node0 --label "mint iris" --chain-id=test --keyring-backend=file --home ./irita/node0/iritacli --fees 6point --gas="auto" -b block
```

**_注_：**合约代码初始化完成后，系统会返回智能合约地址，用于智能合约的调用以及状态的查询等。

## Execute

执行智能合约中的方法。

```bash
irita tx wasm execute [contract_addr_bech32] [json_encoded_send_args] [flags]
```

**参数：**

| 名称                   | 类型   | 必须 | 默认 | 描述                               |
| ---------------------- | ------ | ---- | ---- | ---------------------------------- |
| contract_addr_bech32   | string | 是   |      | 智能合约地址                       |
| json_encoded_send_args | string | 是   |      | 调用智能合约方法的参数，json格式。 |

### 执行合约示例

下面是执行了合约中的`vote`方法，参数为`candidate:iaa1qvty8x0c78am8c44zv2n7tgm6gfqt78j0verqa`

```bash
Vote='{"vote":{"candidate":"iaa1qvty8x0c78am8c44zv2n7tgm6gfqt78j0verqa"}}'
irita tx wasm execute iaa18vd8fpwxzck93qlwghaj6arh4p7c5n89fqcgm9 "$Vote" --from node0 --chain-id=test --keyring-backend=file --home ./irita/node0/iritacli --fees 6point --gas="auto" -b block
```

## Migrate

将智能合约迁移到新的代码版本

```bash
irita tx wasm migrate [contract_addr_bech32] [new_code_id_int64] [json_encoded_migration_args] [flags]
```

**参数：**

| 名称                        | 类型   | 必须 | 默认 | 描述                                        |
| --------------------------- | ------ | ---- | ---- | ------------------------------------------- |
| contract_addr_bech32        | string | 是   |      | 智能合约地址                                |
| new_code_id_int64           | int64  | 是   |      | 迁移到的智能合约`code_id`。                 |
| json_encoded_migration_args | string | 是   |      | 执行合约迁移方法`migrate`的参数，json格式。 |

**_注_：**合约代码更新完成后，合约的地址依然采用原来的合约地址，但是合约的状态依据`migrate`的执行逻辑也可以保持不变。

## Set contract admin

变更智能合约的管理人，以便对合约执行`migrate`操作，必须由上个管理员指定。

```bash
irita tx wasm set-contract-admin [contract_addr_bech32] [new_admin_addr_bech32] [flags]
```

**参数：**

| 名称                  | 类型   | 必须 | 默认 | 描述           |
| --------------------- | ------ | ---- | ---- | -------------- |
| contract_addr_bech32  | string | 是   |      | 智能合约地址   |
| new_admin_addr_bech32 | string | 是   |      | 新的管理员地址 |

### 变更管理员示例

```bash
irita tx wasm set-contract-admin iaa18vd8fpwxzck93qlwghaj6arh4p7c5n89fqcgm9 iaa18lwh8r66wf2hc278ncu4mlgqcxh5slhudkuler --from node0  --chain-id=test --keyring-backend=file --home ./irita/node0/iritacli --fees 6point --gas="auto" -b block
```

## Clear contract admin

清空合约管理人权限，以住址合约的`migrate`操作。

```bash
irita tx wasm clear-contract-admin [contract_addr_bech32] [flags]
```

**参数：**

| 名称                 | 类型   | 必须 | 默认 | 描述         |
| -------------------- | ------ | ---- | ---- | ------------ |
| contract_addr_bech32 | string | 是   |      | 智能合约地址 |

### 清空合约管理人权限示例

```bash
irita tx wasm clear-contract-admin iaa18vd8fpwxzck93qlwghaj6arh4p7c5n89fqcgm9 --from node0  --chain-id=test --keyring-backend=file --home ./irita/node0/iritacli --fees 6point --gas="auto" -b block
```

## Download contract code

下载智能合约的二进制数据(上传的.wasm数据)。

```bash
irita query wasm code [code_id] [output filename] [flags]
```

**参数：**

| 名称    | 类型   | 必须 | 默认 | 描述                |
| ------- | ------ | ---- | ---- | ------------------- |
| code_id | int64  | 是   |      | 智能合约的`code_id` |
| output  | string | 是   |      | 合约code输出文件    |

### 下载智能合约示例

```bash
irita query wasm code 1 election.wasm
```

## Contract

查询合约的信息，包括合约地址、codeID，label等信息。

```bash
irita query wasm contract [bech32_address] [flags]
```

**参数：**

| 名称           | 类型   | 必须 | 默认 | 描述         |
| -------------- | ------ | ---- | ---- | ------------ |
| bech32_address | string | 是   |      | 智能合约地址 |

### 查询合约信息示例

```bash
irita query wasm contract iaa18vd8fpwxzck93qlwghaj6arh4p7c5n89fqcgm9
```

输出信息：

```text
address: iaa18vd8fpwxzck93qlwghaj6arh4p7c5n89fqcgm9
code_id: 1
creator: iaa1rgnu8grzt6mwnjg7jss7w0sfyjn67g4et0hzfz
label: test wasm
```

## Contract History

使用合约地址查询合约的`migrate`历史信息。

```bash
irita query wasm contract-history [bech32_address] [flags]
```

**参数：**

| 名称           | 类型   | 必须 | 默认 | 描述         |
| -------------- | ------ | ---- | ---- | ------------ |
| bech32_address | string | 是   |      | 智能合约地址 |

### 查询`migrate`历史信息示例

```bash
irita query wasm contract-history iaa18vd8fpwxzck93qlwghaj6arh4p7c5n89fqcgm9
```

## List code

查询链上所有的合约`code_id`。

```bash
irita query wasm list-code [flags]
```

### 查询链上所有的合约`code_id`示例

```bash
irita query wasm list-code
```

输出信息：

```text
code_infos:
- creator: iaa1rgnu8grzt6mwnjg7jss7w0sfyjn67g4et0hzfz
  data_hash: E5F29AA07C14DCA498680AFC5376284937FC158677475FD72DBD934B4E023174
  id: 1
pagination: {}
```

## List contract-by-code

使用`code_id`查询合约的基本信息。

```bash
irita query wasm list-contract-by-code [code_id] [flags]
```

**参数：**

| 名称    | 类型  | 必须 | 默认 | 描述                |
| ------- | ----- | ---- | ---- | ------------------- |
| code_id | int64 | 是   |      | 智能合约的`code_id` |

### 使用`code_id`查询合约信息示例

```bash
irita query wasm list-contract-by-code 1
```

输出信息：

```text
contract_infos:
- address: iaa18vd8fpwxzck93qlwghaj6arh4p7c5n89fqcgm9
  code_id: 1
  creator: iaa1rgnu8grzt6mwnjg7jss7w0sfyjn67g4et0hzfz
  label: test wasm
pagination: {}
```

## Contract state all

根据合约地址查询合约内部存储的所有状态信息，信息以编码后的格式输出，用户需自行解码。

```bash
irita query wasm contract-state all [bech32_address] [flags]
```

**参数：**

| 名称           | 类型   | 必须 | 默认 | 描述         |
| -------------- | ------ | ---- | ---- | ------------ |
| bech32_address | string | 是   |      | 智能合约地址 |

### 查询合约所有状态示例

```bash
irita query wasm contract-state all iaa18vd8fpwxzck93qlwghaj6arh4p7c5n89fqcgm9
```

输出信息：

```text
models:
- key: 0006636F6E666967
  value: eyJzdGFydCI6MSwiZW5kIjoxMDAsImNhbmRpZGF0ZXMiOlsiaWFhMXF2dHk4eDBjNzhhbThjNDR6djJuN3RnbTZnZnF0NzhqMHZlcnFhIiwiaWFhMXprMnRzZTBwa2s4N3Aydjh0Y3NmczB5dGZ3M3Q4OGtlamVjeWU1Il0sInZvdGVzIjpbeyJ2b3RlciI6ImlhYTFyZ251OGdyenQ2bXduamc3anNzN3cwc2Z5am42N2c0ZXQwaHpmeiIsImNhbmRpZGF0ZSI6ImlhYTFxdnR5OHgwYzc4YW04YzQ0enYybjd0Z202Z2ZxdDc4ajB2ZXJxYSJ9XX0=
pagination: {}
```

**_注_：**key的编码为hex编码，例如`0006636F6E666967`对应字符串为`config`，value编码为base64编码，例如

```text
eyJzdGFydCI6MSwiZW5kIjoxMDAsImNhbmRpZGF0ZXMiOlsiaWFhMXF2dHk4eDBjNzhhbThjNDR6djJuN3RnbTZnZnF0NzhqMHZlcnFhIiwiaWFhMXprMnRzZTBwa2s4N3Aydjh0Y3NmczB5dGZ3M3Q4OGtlamVjeWU1Il0sInZvdGVzIjpbeyJ2b3RlciI6ImlhYTFyZ251OGdyenQ2bXduamc3anNzN3cwc2Z5am42N2c0ZXQwaHpmeiIsImNhbmRpZGF0ZSI6ImlhYTFxdnR5OHgwYzc4YW04YzQ0enYybjd0Z202Z2ZxdDc4ajB2ZXJxYSJ9XX0=
```

对应的字符串为：

```json
{"start":1,"end":100,"candidates":["iaa1qvty8x0c78am8c44zv2n7tgm6gfqt78j0verqa","iaa1zk2tse0pkk87p2v8tcsfs0ytfw3t88kejecye5"],"votes":[{"voter":"iaa1rgnu8grzt6mwnjg7jss7w0sfyjn67g4et0hzfz","candidate":"iaa1qvty8x0c78am8c44zv2n7tgm6gfqt78j0verqa"}]}
```

## Contract state raw

根据用户指定的合约地址以及合约编写时指定的`key`，查询当前合约的状态信息，输出信息以base64编码。

```bash
irita query wasm contract-state raw [bech32_address] [key] [flags]
```

**参数：**

| 名称           | 类型   | 必须 | 默认 | 描述                                 |
| -------------- | ------ | ---- | ---- | ------------------------------------ |
| bech32_address | string | 是   |      | 智能合约地址                         |
| key            | string | 是   |      | 编写合约时指定的存储key(hex编码格式) |

### 查询指定的`key`的value信息示例

```bash
irita query wasm contract-state raw iaa18vd8fpwxzck93qlwghaj6arh4p7c5n89fqcgm9 0006636F6E666967
```

输出信息：

```text
data: eyJzdGFydCI6MSwiZW5kIjoxMDAsImNhbmRpZGF0ZXMiOlsiaWFhMXF2dHk4eDBjNzhhbThjNDR6djJuN3RnbTZnZnF0NzhqMHZlcnFhIiwiaWFhMXprMnRzZTBwa2s4N3Aydjh0Y3NmczB5dGZ3M3Q4OGtlamVjeWU1Il0sInZvdGVzIjpbeyJ2b3RlciI6ImlhYTFyZ251OGdyenQ2bXduamc3anNzN3cwc2Z5am42N2c0ZXQwaHpmeiIsImNhbmRpZGF0ZSI6ImlhYTFxdnR5OHgwYzc4YW04YzQ0enYybjd0Z202Z2ZxdDc4ajB2ZXJxYSJ9XX0=
````

**_注_：data解码格式为base64。

## Contract state smart

调用合约的查询方法并返回解码后的数据。

```bash
irita query wasm contract-state smart [bech32_address] [query] [flags]
```

**参数：**

| 名称           | 类型   | 必须 | 默认 | 描述                 |
| -------------- | ------ | ---- | ---- | -------------------- |
| bech32_address | string | 是   |      | 智能合约地址         |
| query          | string | 是   |      | 查询字符串，json格式 |

### 调用合约指定的查询方法示例

```bash
Query='{"get_vote_info":{}}'
irita query wasm contract-state smart iaa18vd8fpwxzck93qlwghaj6arh4p7c5n89fqcgm9 $Query
```

输出信息：

```text
data:
  end: 100
  start: 1
  votes:
  - candidate: iaa1qvty8x0c78am8c44zv2n7tgm6gfqt78j0verqa
    count: 1
```
