## Query（查询参数）拼接与解析

Query（查询字符串）就是 URL 中 ? 之后的部分：

```
https://api.example.com/users?name=JimLee&age=22&lang=go
```

**主要用于**

- GET 请求的参数传递
- 搜索过滤、分页、排序
- API 查询条件（REST、GraphQL-like 接口常见）

### url.Values

`url.Values` 是 Go 中管理 Query 的核心类型：

```
type Values map[string][]string
```

它本质就是一个 支持多值的 map。

**常用方法**

| 方法                | 说明                   |
|-------------------|----------------------|
| `Set(key, value)` | 设置单值（会覆盖旧值）          |
| `Add(key, value)` | 添加多值（不会覆盖）           |
| `Del(key)`        | 删除某个 key             |
| `Get(key)`        | 取第一个值                |
| `Encode()`        | 生成编码后的字符串（自动 URL 编码） |

### 总结

`url.Values` 是构建和解析查询字符串的标准方式，
它会自动编码、支持多值、与 `url.URL` 完美结合。
不要手动拼 `?key=value`。

| 操作       | 推荐做法                      |
|----------|---------------------------|
| 创建 query | `q := url.Values{}`       |
| 设置单值     | `q.Set(key, value)`       |
| 多值添加     | `q.Add(key, value)`       |
| 删除       | `q.Del(key)`              |
| 生成字符串    | `q.Encode()`              |
| 拼回 URL   | `u.RawQuery = q.Encode()` |
| 获取参数     | `u.Query().Get(key)`      |
