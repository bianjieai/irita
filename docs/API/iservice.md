<!--
order: 6
-->

# iService

## 查询服务定义

查询指定的服务定义。

**API：**

```bash
GET /service/definitions/{service-name}
```

**参数：**

- service-name：string，服务名称

**返回值：**

- 服务定义查询结果：object

**请求示例：**

```bash
curl -X GET "http://localhost:1317/service/definitions/test" -H "accept: application/json" | jq
```

```json
{
  "height": "2098",
  "result": {
    "type": "irismod/service/ServiceDefinition",
    "value": {
      "name": "test",
      "description": "test",
      "tags": [
        "test"
      ],
      "author": "iaa12v3r0unp6nprp9zur0fn446n832dfag5l3w38e",
      "author_description": "test",
      "schemas": "{\"input\":{\"type\":\"object\",\"properties\":{\"id\":{\"type\":\"string\"}}},\"output\":{\"type\":\"object\",\"properties\":{\"name\":{\"type\":\"string\"}}}}"
    }
  }
}
```

## 查询服务绑定

查询指定的服务绑定。

**API：**

```bash
GET /service/bindings/{service-name}/{provider}
```

**参数：**

- service-name：string，服务名称

- provider：string，服务提供者地址

**返回值：**

- 服务绑定查询结果：object

**请求示例：**

```bash
curl -X GET "http://localhost:1317/service/bindings/test/iaa12v3r0unp6nprp9zur0fn446n832dfag5l3w38e" -H "accept: application/json" | jq
```

```json
{
  "height": "2111",
  "result": {
    "type": "irismod/service/ServiceBinding",
    "value": {
      "service_name": "test",
      "provider": "iaa12v3r0unp6nprp9zur0fn446n832dfag5l3w38e",
      "deposit": [
        {
          "denom": "point",
          "amount": "10000"
        }
      ],
      "pricing": "{\"price\":\"1point\"}",
      "qos": "50",
      "available": true,
      "disabled_time": "0001-01-01T00:00:00Z",
      "owner": "iaa12v3r0unp6nprp9zur0fn446n832dfag5l3w38e"
    }
  }
}
```

## 查询服务绑定列表

查询指定服务的所有绑定。

**API：**

```bash
GET /service/bindings/{service-name}
```

**参数：**

- service-name：string，服务名称

**返回值：**

- 服务绑定列表查询结果：object

**请求示例：**

```bash
curl -X GET "http://localhost:1317/service/bindings/test" -H "accept: application/json" | jq
```

```json
{
  "height": "2126",
  "result": [
    {
      "service_name": "test",
      "provider": "iaa12v3r0unp6nprp9zur0fn446n832dfag5l3w38e",
      "deposit": [
        {
          "denom": "point",
          "amount": "10000"
        }
      ],
      "pricing": "{\"price\":\"1point\"}",
      "qos": "50",
      "available": true,
      "disabled_time": "0001-01-01T00:00:00Z",
      "owner": "iaa12v3r0unp6nprp9zur0fn446n832dfag5l3w38e"
    }
  ]
}
```

## 查询服务请求

根据请求 ID 查询服务请求。

**API：**

```bash
GET /service/requests/{request-id}
```

**参数：**

- request-id：string，请求 ID

**返回值：**

- 服务请求查询结果：object

**请求示例：**

```bash
curl -X GET "http://localhost:1317/service/requests/3BBA63B4BD750DF87CC5DF996518861E11FFF161A606EC36A007886B6E4C05580000000000000000000000000000000100000000000008990000" -H "accept: application/json" | jq
```

```json
{
  "height": "2222",
  "result": {
    "type": "irismod/service/Request",
    "value": {
      "id": "3BBA63B4BD750DF87CC5DF996518861E11FFF161A606EC36A007886B6E4C05580000000000000000000000000000000100000000000008990000",
      "service_name": "test",
      "provider": "iaa12v3r0unp6nprp9zur0fn446n832dfag5l3w38e",
      "consumer": "iaa12v3r0unp6nprp9zur0fn446n832dfag5l3w38e",
      "input": "{\"id\":\"001\"}",
      "service_fee": [
        {
          "denom": "point",
          "amount": "1"
        }
      ],
      "request_height": "2201",
      "expiration_height": "2301",
      "request_context_id": "3BBA63B4BD750DF87CC5DF996518861E11FFF161A606EC36A007886B6E4C05580000000000000000",
      "request_context_batch_counter": "1"
    }
  }
}
```

## 查询服务绑定的请求列表

查询指定服务绑定的当前活跃请求。

**API：**

```bash
GET /service/requests/{service-name}/{provider}
```

**参数：**

- service-name：string，服务名称

- provider：string，服务提供者地址

**返回值：**

- 服务请求列表查询结果：object

**请求示例：**

```bash
curl -X GET "http://localhost:1317/service/requests/test/iaa12v3r0unp6nprp9zur0fn446n832dfag5l3w38e" -H "accept: application/json" | jq
```

```json
{
  "height": "2205",
  "result": [
    {
      "id": "3BBA63B4BD750DF87CC5DF996518861E11FFF161A606EC36A007886B6E4C05580000000000000000000000000000000100000000000008990000",
      "service_name": "test",
      "provider": "iaa12v3r0unp6nprp9zur0fn446n832dfag5l3w38e",
      "consumer": "iaa12v3r0unp6nprp9zur0fn446n832dfag5l3w38e",
      "input": "{\"id\":\"001\"}",
      "service_fee": [
        {
          "denom": "point",
          "amount": "1"
        }
      ],
      "request_height": "2201",
      "expiration_height": "2301",
      "request_context_id": "3BBA63B4BD750DF87CC5DF996518861E11FFF161A606EC36A007886B6E4C05580000000000000000",
      "request_context_batch_counter": "1"
    }
  ]
}
```

## 查询服务响应

根据请求 ID 查询服务响应。

**API：**

```bash
GET /service/responses/{request-id}
```

**参数：**

- request-id：string，请求 ID

**返回值：**

- 服务响应查询结果：array

**请求示例：**

```bash
curl -X GET "http://localhost:1317/service/responses/3BBA63B4BD750DF87CC5DF996518861E11FFF161A606EC36A007886B6E4C05580000000000000000000000000000000100000000000008990000" -H "accept: application/json" | jq
```

```json
{
  "height": "2293",
  "result": {
    "type": "irismod/service/Response",
    "value": {
      "provider": "iaa12v3r0unp6nprp9zur0fn446n832dfag5l3w38e",
      "consumer": "iaa12v3r0unp6nprp9zur0fn446n832dfag5l3w38e",
      "result": "{\"code\":200,\"message\":\"ok\"}",
      "output": "{\"name\":\"test\"}",
      "request_context_id": "3BBA63B4BD750DF87CC5DF996518861E11FFF161A606EC36A007886B6E4C05580000000000000000",
      "request_context_batch_counter": "1"
    }
  }
}
```

## 查询请求上下文

查询指定的请求上下文。

**API：**

```bash
GET /service/contexts/{request-context-id}
```

**参数：**

- request-context-id：string，请求上下文 ID

**返回值：**

- 请求上下文查询结果：object

**请求示例：**

```bash
curl -X GET "http://localhost:1317/service/contexts/3BBA63B4BD750DF87CC5DF996518861E11FFF161A606EC36A007886B6E4C05580000000000000000" -H "accept: application/json" | jq
```

```json
{
  "height": "0",
  "result": {
    "type": "irismod/service/RequestContext",
    "value": {
      "service_name": "test",
      "providers": [
        "iaa12v3r0unp6nprp9zur0fn446n832dfag5l3w38e"
      ],
      "consumer": "iaa12v3r0unp6nprp9zur0fn446n832dfag5l3w38e",
      "input": "{\"id\":\"001\"}",
      "service_fee_cap": [
        {
          "denom": "point",
          "amount": "1"
        }
      ],
      "timeout": "100",
      "repeated": true,
      "repeated_frequency": "150",
      "repeated_total": "100",
      "batch_counter": "1",
      "batch_request_count": 1
    }
  }
}
```

## 查询请求上下文的请求列表

查询请求上下文指定批次的活跃请求列表。

**API：**

```bash
GET /service/requests/{request-context-id}/{batch-counter}
```

**参数：**

- request-context-id：string，请求上下文 ID

- batch-counter：uint64，批次计数器

**返回值：**

- 服务请求列表：array

**请求示例：**

```bash
curl -X GET "http://localhost:1317/service/requests/3BBA63B4BD750DF87CC5DF996518861E11FFF161A606EC36A007886B6E4C05580000000000000000/1" -H "accept: application/json" | jq
```

```json
{
  "height": "2251",
  "result": [
    {
      "id": "3BBA63B4BD750DF87CC5DF996518861E11FFF161A606EC36A007886B6E4C05580000000000000000000000000000000100000000000008990000",
      "service_name": "test",
      "provider": "iaa12v3r0unp6nprp9zur0fn446n832dfag5l3w38e",
      "consumer": "iaa12v3r0unp6nprp9zur0fn446n832dfag5l3w38e",
      "input": "{\"id\":\"001\"}",
      "service_fee": [
        {
          "denom": "point",
          "amount": "1"
        }
      ],
      "request_height": "2201",
      "expiration_height": "2301",
      "request_context_id": "3BBA63B4BD750DF87CC5DF996518861E11FFF161A606EC36A007886B6E4C05580000000000000000",
      "request_context_batch_counter": "1"
    }
  ]
}
```

## 查询请求上下文的响应列表

查询请求上下文指定批次的活跃响应列表。

**API：**

```bash
GET /service/responses/{request-context-id}/{batch-counter}
```

**参数：**

- request-context-id：string，请求上下文 ID

- batch-counter：uint64，批次计数器

**返回值：**

- 服务响应列表查询结果：object

**请求示例：**

```bash
curl -X GET "http://localhost:1317/service/responses/3BBA63B4BD750DF87CC5DF996518861E11FFF161A606EC36A007886B6E4C05580000000000000000/1" -H "accept: application/json" | jq
```

```json
{
  "height": "2277",
  "result": [
    {
      "provider": "iaa12v3r0unp6nprp9zur0fn446n832dfag5l3w38e",
      "consumer": "iaa12v3r0unp6nprp9zur0fn446n832dfag5l3w38e",
      "result": "{\"code\":200,\"message\":\"ok\"}",
      "output": "{\"name\":\"test\"}",
      "request_context_id": "3BBA63B4BD750DF87CC5DF996518861E11FFF161A606EC36A007886B6E4C05580000000000000000",
      "request_context_batch_counter": "1"
    }
  ]
}
```

## 查询提取地址

查询指定所有者的服务费提取地址。

**API：**

```bash
GET /service/owners/{owner}/withdraw-address
```

**参数：**

- owner：string，所有者账户地址

**返回值：**

- 提取地址查询结果：object

**请求示例：**

```bash
curl -X GET "http://localhost:1317/service/owners/iaa12v3r0unp6nprp9zur0fn446n832dfag5l3w38e/withdraw-address" -H "accept: application/json" | jq
```

```json
{
  "height": "2308",
  "result": "iaa12v3r0unp6nprp9zur0fn446n832dfag5l3w38e"
}
```

## 查询服务费

查询指定服务提供者赚取的服务费。

**API：**

```bash
GET /service/fees/{provider}
```

**参数：**

- provider：string，服务提供者地址

**返回值：**

- 服务费查询结果：object

**请求示例：**

```bash
curl -X GET "http://localhost:1317/service/fees/iaa12v3r0unp6nprp9zur0fn446n832dfag5l3w38e" -H "accept: application/json" | jq
```

```json
{
  "height": "2322",
  "result": [
    {
      "denom": "point",
      "amount": "1"
    }
  ]
}
```
