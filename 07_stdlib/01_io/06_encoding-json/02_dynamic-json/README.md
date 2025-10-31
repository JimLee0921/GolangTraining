## 解析动态 JSON

有些 JSON 字段不固定，比如：

```
{
  "id": 1001,
  "name": "phone",
  "extra": {
    "color": "black",
    "weight": 120,
    "tags": ["hot", "new"]
  }
}
```

如果不知道 extra 里面会出现什么字段，用结构体很麻烦，这时候就用：

```
map[string]interface{}
```

**解析规则**

- 对象 -> map[string]any
- 数组 -> []any
- 默认数字 -> 因为 JSON 里的数字没有类型，Go 默认当做 float64 来处理

适用于完全不知道结构或字段。