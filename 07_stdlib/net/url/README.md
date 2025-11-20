# `net/url`

`net/url` 包主要用来解析、构建、修改和编码 URL 结构的。
HTTP 请求中的 URL，本质上都是一个 url.URL 结构体。

**主要用途**

- 解析 URL（从字符串 -> 结构体）
- 构建 URL（从结构体 -> 字符串）
- 读取和修改 URL 内的 Query 参数
- 对路径、参数进行 URL 编码 / 解码

```
req *http.Request
req.URL  // 这个字段就是 *url.URL
```

