# `type Response`

`type Response` 对应的是客户端收到的 HTTP 响应对象，即 Client 侧用的，不是服务器返回的 ResponseWriter。

## 核心字段

标准定义（部分省略）

```
type Response struct {
    Status     string       // 完整状态行，如 "200 OK"
    StatusCode int          // 仅状态码，如 200
    Proto      string       // 这三个是 HTTP 协议版本相关，如 HTTP/1.0
    ProtoMajor int          // 1
    ProtoMinor int          // 0

    Header     Header       // 响应头，与 ResponseWriter.Header() 同类型
    Body       io.ReadCloser// 响应体内容（需要读完并关闭）
    ContentLength int64     // 响应体长度（可能为 -1 表示未知）
    Close      bool         // 是否在响应后关闭连接
    Uncompressed bool       // 是否已自动解压

    Request    *Request     // 指向发起这个响应的请求
}
```

## 使用讲解

### Body 类型

因为 HTTP 响应体通常 是流 (stream)，可能非常大，所以 Body 是 `io.ReadCloser` 而不是 `[]byte`

- 视频、音频、文件下载
- 持续推送（SSE、长连接）
- chunked 分块传输

> 不能一次性全部读入内存，所以必须提供流式读取

### Status vs StatusCode

| 字段               | 示例                | 用途          |
|------------------|-------------------|-------------|
| `Status string`  | `"404 Not Found"` | 完整描述，展示给人看的 |
| `StatusCode int` | `404`             | 逻辑判断与代码分支   |

### 客户端读取响应头流程

> 类似于使用 python 的 requests 库发起请求后如何获取信息

使用 `resp, err := client.Do(req)` 接收 response 后。

1. 使用 resp
2. 读取 `resp.Body`
3. 必须关闭 `resp.Body`

如果忘记 `Close()`，连接不会归还给连接池，最终会导致 TCP 资源耗尽

### 响应头 Header

Response.Header 与之前学的 `type Header` 是完全一样的类型，可以直接使用 `Get(key)`、`Values(key)`、`Clone()` 等方法。

### Request 对象

这是原始请求对象，用于：

- 重试网络请求
- 传递 trace/context

简单理解：Response 与 Request 是成对出现的。也就相当于 python 中的 request 和 response

## 相关函数

文档中几个与客户端发起 HTTP 请求 & 构造 Response 的函数。

> 真正开发中更常用的是使用 `client.Do(req)`（来自 `*http.Client`）来执行请求获取 Response 对象
> 
> 这是执行自定义请求（最灵活的方式），见 Client 文档

### `http.Get()`

发起 HTTP GET 请求，是最常见的请求方式。相当于 python 中的 `request.get()`

```
func Get(url string) (resp *Response, err error)
```

- 内部会自动创建一个 `http.Client`
- 适用于快捷测试、简单请求
- 返回 `*Response`（必须手动 `Close Body`）

> 属于 快速调用函数（Convenience Function），不是专业用法

### `http.Head()`

发起一个 HTTP HEAD 请求。与 GET 类似，但 只返回响应头，不返回 Body。

```
func Head(url string) (resp *Response, err error)
```

**常用于**

- 获取资源是否存在
- 获取 `Content-Length `
- 检查缓存、检查服务器状态
- 不下载内容，节省流量

> 返回的 resp.Body 通常为空或长度为 0。仍然需要手动 `resp.Body.Close()`

### `http.Post()`

发起 HTTP POST 请求，请求体由 body 提供，相当于 python 中的 `requests.post()`

```
func Post(url, contentType string, body io.Reader) (resp *Response, err error)
```

| 参数            | 类型          | 含义                                       |
|---------------|-------------|------------------------------------------|
| `url`         | string      | 请求目标                                     |
| `contentType` | string      | `Content-Type` 头，例如 `"application/json"` |
| `body`        | `io.Reader` | 请求体来源（可以是字符串、文件、buffer 等）                |

> 也是快捷函数，适用于简单 POST

### `http.PostForm()`

发送一个表单类型的 POST 请求。

```
func PostForm(url string, data url.Values) (resp *Response, err error)
```

等价于（自动完成 URL 编码）：

```
Content-Type: application/x-www-form-urlencoded
Body: name=value&last=xxx...
```

**常用于**

- 网站表单提交
- OAuth 请求 token
- 老式 Web 系统交互

> Post() 的封装，简化了表单写法

### `http.ReadResponse()`

从底层流中解析 HTTP 响应，返回 `*Response`。

```
func ReadResponse(r *bufio.Reader, req *Request) (*Response, error)
```

要求已经有了一个：`r = bufio.NewReader(networkConnection)`

**常用于**

- 协议层 / 底层网络交互
- 实现 HTTP 代理服务器
- 测试 `HTTP raw` 报文
- 与非 `Go HTTP server` 通信

> 不是给普通开发者用的，是给写 HTTP 框架或代理服务器的人用的

## 常见方法

### 1. Cookies()

用于从 响应头中提取所有 `Set-Cookie` 字段，并解析为 `[]*http.Cookie` 结构体。

```
func (r *Response) Cookies() []*Cookie
```

HTTP 允许响应中返回多个 Cookie：

```
Set-Cookie: session_id=abc; Path=/; HttpOnly
Set-Cookie: theme=dark; Path=/; Max-Age=3600
```

`Cookies()` 会：

1. 扫描 `r.Header["Set-Cookie"]`
2. 逐条解析
3. 返回解析好的 Cookie 对象列表

> Cookies() 是对 Set-Cookie 响应头的结构化解析器

### 2. Location()

用于获取响应头中的 `Location URL`，通常用于重定向响应。

```
func (r *Response) Location() (*url.URL, error)
```

当服务器返回：

```
HTTP/1.1 302 Found
Location: https://example.com/login
```

Location 表示客户端应当跳转到另一个 URL。Location() 方法会：

1. 读取 `r.Header.Get("Location")`
2. 将字符串解析为 *url.URL
3. 返回结构化 URL

> 用于处理重定向目标地址

### 3. ProtoAtLeast()

判断响应的 HTTP 协议版本是否 >= 传入的版本（偏协议层）

```
func (r *Response) ProtoAtLeast(major, minor int) bool
```

比如：

```
ProtoMajor = 1
ProtoMinor = 1   
// 拼接后协议版本就是 HTTP/1.1
```

然后调用：

```
ProtoAtLeast(1, 0) -> true
ProtoAtLeast(2, 0) -> false
```

**常用于**

- 判断是否支持 `HTTP/1.1` 的 持久连接
- 判断是否支持 `HTTP/2 Server Push `
- 判断分块传输（chunked）等行为

### 4. Write()

将整个 Response 按 HTTP 原始报文格式写入 `io.Writer`。
也就是把 Response 原样输出为 HTTP 报文的函数，主要用于网关 / 代理 / 转发场景。

```
func (r *Response) Write(w io.Writer) error
```

也就是写出类似：

```
HTTP/1.1 200 OK
Content-Type: text/plain
Content-Length: 5

Hello
```

**常用于**

- 代理服务器转发响应
- 调试调试 / 生成 `raw HTTP` 报文
- 把 Response 回放或复制到别处

> 只写 头部和 Body，不会自动处理连接控制，某些 header 会按照 HTTP 规范进行格式化