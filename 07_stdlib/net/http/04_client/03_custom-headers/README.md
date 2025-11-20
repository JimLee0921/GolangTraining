## 自定义请求头

在 Go 里要给请求加自定义 Header，其实就是操作 `*http.Request.Header`。
每个请求 (http.Request) 都有一个字段：

```
Header http.Header // 类型是 map[string][]string
```

### Header 常用方法

| 方法                           | 说明                 |
|------------------------------|--------------------|
| `req.Header.Set(key, value)` | 设置（若已存在则覆盖）        |
| `req.Header.Add(key, value)` | 添加（可叠加多个值）         |
| `req.Header.Del(key)`        | 删除某个 Header        |
| `req.Header.Get(key)`        | 读取 Header 值（区分大小写） |

### 批量设置

如果有以组 Header 可以直接赋值整个 map：

```
req.Header = http.Header{
	"User-Agent":    {"Go-HttpClient/2.0"},
	"Authorization": {"Bearer my-token"},
	"Accept":        {"application/json"},
}
```

### 不允许用户设置的 Header

出于安全和协议一致性考虑，Go 的 net/http 会禁止手动设置或覆盖某些头：

| Header              | 原因                |
|---------------------|-------------------|
| `Host`              | 用 `req.Host` 字段设置 |
| `Content-Length`    | 由 `net/http` 自动计算 |
| `Transfer-Encoding` | 自动处理 chunked 传输   |
| `Trailer`           | 仅在支持的流式传输中使用      |
