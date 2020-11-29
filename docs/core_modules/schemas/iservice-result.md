<!--
order: false
-->

### iService 响应结果 Schema 示例

```json
{
    "$schema": "http://json-schema.org/draft-04/schema#",
    "title": "iservice-result",
    "description": "iService Result Schema",
    "type": "object",
    "properties": {
      "code": {
        "description": "result code",
        "type": "integer",
        "enum": [200, 400, 500]
      },
      "message": {
        "description": "result message",
        "type": "string"
      }
    },
    "additionalProperties": false,
    "required": [
      "code",
      "message"
    ]
}
```
