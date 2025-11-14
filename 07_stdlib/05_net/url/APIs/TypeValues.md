## `type Values`

Values 映射将字符串键映射到值列表。它通常用于查询参数和表单值。与 `http.Header` 映射不同，Values 映射中的键区分大小写。

```
type Values map[string][]string
```

Query 参数允许同名键重复，因此用 []string 存值。每个查询参数 key 可以对应多个值。`?id=1&id=2&debug=true` 会变成：

```
map[string][]string{
    "id":   {"1", "2"},
    "debug": {"true"},
}
```

### 修改流程

修改 Query 参数的标准工作流程永远是：Query -> 修改 -> Encode -> 回写

步骤：

1. `q := u.Query()` 获取副本
2. 修改 q
3. `u.RawQuery = q.Encode()` 写回到 URL

> Query() 返回的是副本，不是引用

### 主要方法

| 方法                                     | 用途                     | 行为说明               |
|----------------------------------------|------------------------|--------------------|
| func (v Values) Get(key string) string | 获取该 key 的第一个值          | 不存在则返回空字符串         |
| func (v Values) Set(key, value string) | 覆盖该 key 的所有值           | 变成单值模式（`[]=value`） |
| func (v Values) Add(key, value string) | 追加一个值到该 key            | 支持同 key 多值（累积）     |
| func (v Values) Del(key string)        | 删除 key 对应的所有值          | 删除整项，而不是置空         |
| func (v Values) Has(key string) bool   | 判断 key 是否存在            | 不检查值，只看 key 存在与否   |
| func (v Values) Encode() string        | 把 Values 编码为 `a=1&b=2` | 用于写回 `u.RawQuery`  |





