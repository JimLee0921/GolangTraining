# `type Request`

Request 是服务器接收到的或客户端要发送的 HTTP 请求，表示 一次 HTTP 请求的全部信息。
`type Request` 也就是一条 HTTP 请求在 Go 中的结构化表示。

- 服务端：表示客户端发来的请求
- 客户端：表示即将发出去的请求

> 你看请求和写请求用的都是同一类型

其中包括：

- 请求行（方法、URL、协议）
- 请求头
- 请求体（Body）
- 客户端信息、上下文等

## 核心字段

这里只展示最重要的一些字段

```
type Request struct {
    // 最核心的四个字段
    Method string       // HTTP 方法，如 GET / POST / PUT / DELETE ...
    URL    *url.URL     // 包含路径、查询参数、域名、scheme等，使用见 http/url 
    Header Header       // 请求头（map[string][]string），使用见 Header
    Body   io.ReadCloser // 请求体，使用见 io 操作
       
    // 其余字段
    Host   string       // Host头 或 URL中的host
    
    ContentLength int64
    TransferEncoding []string
    
    Close bool

    RemoteAddr string   // 客户端 IP:端口
    RequestURI string
    
    TLS *tls.ConnectionState

    
    Form url.Values     // 表单数据（解析后）
    PostForm url.Values // POST表单数据（解析后）
    MultipartForm *multipart.Form // 文件上传解析后
    Cookie()            // 获取 Cookie
    Context() context.Context // 取消/超时 的上下文
}
```

## 构造 / 解析 Request 函数

### 1. `http.NewRequest()`

手动构造一个客户端的 HTTP 请求对象 `*Request`。

```
func NewRequest(method, url string, body io.Reader) (*Request, error)
```

| 参数       | 含义                              |
|----------|---------------------------------|
| `method` | 请求方法，例如 `"GET"`, `"POST"`       |
| `url`    | 完整或相对 URL                       |
| `body`   | 请求体来源（可为 `nil`，也可以是字符串、文件、缓冲区等） |

在客户端要发请求时一般会这样用：

```
req := NewRequest(...)
client.Do(req)
```

- 不会自动设置 `Content-Typ`e（需要手动设置 Header）
- body 必须是 `io.Reader`，不是 `[]byte`，因为请求体要按流方式发送
- 最常用、最标准的构造 Request 的方式

### 2. `http.NewRequestWithContext()`

与 NewRequest 相同，但额外允许为请求附带一个 Context

```
func NewRequestWithContext(ctx context.Context, method, url string, body io.Reader) (*Request, error)
```

传入的 Context 主要用于：

- 超时：请求执行太久就中断
- 取消：用户点击取消或后台退出
- 链路追踪：分布式系统内部跟踪请求路径，在链路中传递 trace_id / user_id / session 等

```
ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
req := NewRequestWithContext(ctx, "GET", url, nil)
client.Do(req)
```

如果超时到了，请求会自动被取消，连接会关闭。

> NewRequestWithContext = NewRequest + 取消/超时控制能力。现代 Go 代码中，推荐始终使用这个版本。

### 3. `http.ReadRequest()`

从原始网络流中读取并解析一个 HTTP 请求，返回 `*Request`。也就是把 TCP 字节解析为 Request 结构体。用于底层或代理，不是日常开发使用。

在服务端代码中 Request 一般不是自行创建的，而是 `net/http` 自动解析传递的。所有 Handler 的签名都是：

```
func(w http.ResponseWriter, r *http.Request)
```

这里的 r 也就是请求信息。可以在 Handler 中直接使用：

```

func hello(w http.ResponseWriter, r *http.Request) {
    fmt.Println(r.Method)         // GET
    fmt.Println(r.URL.Path)       // /hello
    fmt.Println(r.Header)         // 所有请求头
}
```

> ReadRequest 属于协议层 API 只有在自己实现 HTTP 服务器、或写中间代理服务器时才会直接调用它

## Request 常用方法

### 1. AddCookie()

把一个 Cookie 添加到这条请求里（设置到 Cookie 请求头）。
常用于在客户端发起请求（或在代理里转发下游请求）时，需要携带某些 cookie。

```
func (r *Request) AddCookie(c *Cookie)
```

> 这是请求侧的 Cookie 添加；如果在服务端返回 cookie，应使用 `http.SetCookie(w, c)` 写入 响应头 `Set-Cookie`

### 2. BasicAuth()

从请求头 `Authorization: Basic ...` 中解析出用户名和密码。

```
func (r *Request) BasicAuth() (username, password string, ok bool)
```

如果返回 `ok=false` 表示没有 Basic 认证头或格式不对，否则给出解析后的 `username/password`。

> 只支持 Basic 方案（base64 的 user:pass），其他认证方式（如 Bearer/JWT）不在此方法处理范围

### 3. SetBasicAuth()

在请求头中设置： `Authorization: Basic base64(username:password)`

```
func (r *Request) SetBasicAuth(username, password string)
```

**使用场景**

- 客户端请求服务器时进行 Basic 身份认证
- 或代理认证、服务间内部测试接口认证

与 r.BasicAuth() 是对偶关系：

- `SetBasicAuth` 用于客户端设置 Authorization 头
- `BasicAuth` 用于服务端解析 Authorization 头

### 4. Cookies()

解析并返回该请求携带的 所有 Cookie（来自 Cookie 请求头，可有多条）

```
func (r *Request) Cookies() []*Cookie
```

常用于需要遍历所有 cookie，或做统一记录/过滤时，用它最直接。

### 5. Cookie(name)

按名称来取单个 Cookie ，找不到时返回 ErrNoCookie，若同名有多条，只返回其中一条（可能不是想要的那条）。

```
func (r *Request) Cookie(name string) (*Cookie, error)
```

> 若关心所有同名值，用上面的 CookiesNamed(name)

### 6. CookiesNamed()

Go 1.23+ 支持，返回所有名称为 name 的 cookie 列表（有的浏览器/场景可能会出现同名多条）
对比 Cookie(name) 更精确完整，不会丢弃同名的其他条目。用于补齐同名多条场景的易用性。

```
func (r *Request) CookiesNamed(name string) []*Cookie
```

### 7. FormValue()

获取请求参数（自动在 Query + POST Form 中查找）。

```
func (r *Request) FormValue(key string) string
```

等价于：

```
r.URL.Query().Get(key)  
// OR  
r.PostForm.Get(key)
```

### 8. PostFormValue()

只从 POST 表单数据 中查，不查 Query。用于明确要求参数只能从 `POST body` 来 的场景。

```
func (r *Request) PostFormValue(key string) string
```

等价于： `r.PostForm.Get(key)`

### 9. ParseForm()

ParseForm 会填充 `r.Form` 和 `r.PostForm`。

```
func (r *Request) ParseForm() error
```

- 对于所有请求，ParseForm 会解析 URL 中的原始查询并更新 `r.Form`。
- 对于 POST、PUT 和 PATCH 请求，它还会读取请求体，将其解析为表单，并将结果放入 `r.PostForm` 和 `r.Form` 中。请求体参数优先于
  `r.Form` 中的 URL 查询字符串值
- 当想使用 `r.Form.Get(...)` 和 `r.PostForm.Get(...)` 时，必须先执行 `ParseForm()`（但 `FormValue()` 不需要）

> 如果请求正文的大小尚未被 MaxBytesReader 限制，则大小上限为 10MB
>
> 对于其他 HTTP 方法，或者当 `Content-Type` 不是 `application/x-www-form-urlencoded` 时，不会读取请求正文，并且
`r.PostForm` 会被初始化为非 nil 的空值

`Request.ParseMultipartForm` 方法会自动调用 `ParseForm`

### 10. ParseMultipartForm()

解析 multipart/form-data（即 文件上传表单），maxMemory 指定允许将多少数据暂存到内存里，超过部分会写入临时文件。

```
func (r *Request) ParseMultipartForm(maxMemory int64) error

```

- 如果有必要 ParseMultipartForm 会调用 `Request.ParseForm` 方法来解析 Form 表单
- 如果 ParseForm 返回错误，ParseMultipartForm 也会返回该错误，但会继续解析请求体
- ParseMultipartForm 方法调用一次后，后续调用将不再生效
- 文件上传必须用 ParseMultipartForm，不能用 ParseForm 替代

```
maxMemory = 32 << 20   // 32MB

// 解析后可以使用
r.MultipartForm.File
r.MultipartForm.Value
```

### 11. FormFile()

从 multipart 表单（文件上传）中提取第一个上传文件

```
func (r *Request) FormFile(key string) (multipart.File, *multipart.FileHeader, error)
```

| 返回值                     | 含义                   |
|-------------------------|----------------------|
| `multipart.File`        | 可读取的文件内容             |
| `*multipart.FileHeader` | 包含文件名 / 大小 / MIME 信息 |
| `error`                 | 未找到或解析失败             |

```
表单：<input type="file" name="avatar">
方法：file, header, err := r.FormFile("avatar")
```

> 如有必要，FormFile 会自动调用 `Request.ParseMultipartForm` 和 `Request.ParseForm`，无需手动调用

### 12. MultipartReader()

获取 multipart/form-data 的底层 分段解析器。通常用于大文件上传、云对象存储直传、断点续传、流式处理（不一次性读入内存）

```
func (r *Request) MultipartReader() (*multipart.Reader, error)
```

**与 FormFile 的关系**

| 方法                  | 用途              | 抽象级别        |
|---------------------|-----------------|-------------|
| `FormFile()`        | 高级封装 -> 直接拿文件即可 | 简单业务推荐      |
| `MultipartReader()` | 底层分段读取 -> 可处理大流 | 上传大文件/流式处理用 |

### 13. Context()

返回与该请求关联的 Context

```
func (r *Request) Context() context.Context
```

- 要更改上下文，需要使用 `Request.Clone` 或 `Request.WithContext`
- 返回的上下文始终不为 nil，默认为后台上下文
- 对于发出的客户端请求，上下文控制取消操作
- 对于传入的服务器请求，当客户端连接关闭、请求被取消（使用 HTTP/2）或 ServeHTTP 方法返回时，上下文将被取消

### 14. WithContext()

返回一个带新 Context 的 Request 副本也就是浅拷贝（但共享 Header / Body / Method / URL）

```
func (r *Request) WithContext(ctx context.Context) *Request
```

**常用于**

- 中间件或拦截器想给请求添加：超时、deadline、trace-id、user-id
- 但是它不会复制整个 Request，只替换 Context 字段

> WithContext = 原样请求 + 替换Context，要创建带有新上下文的请求的深拷贝，需要使用 `Request.Clone`

### 15. Clone()

深拷贝整个 Request，包括 Header / URL / Trailer 等，同时替换 Context，主要用于 中间件 / 反向代理 / API Gateway。

```
func (r *Request) Clone(ctx context.Context) *Request
```

**与 `WithContext()` 区别**

| 方法            | 复制程度                   | 场景                 |
|---------------|------------------------|--------------------|
| `WithContext` | 只换 Context，不复制 Request | 本地修改 Context       |
| `Clone`       | 完整复制 Request           | 要转发 / 代理 / 重试请求时使用 |

> 克隆只会对 body 字段进行浅层复制

### 16. PathValue()

获取 URL 路径参数（Go 1.22+ 新的标准路由中可用）

```
func (r *Request) PathValue(name string) string
```

```
// 路由
/users/{id}/orders/{oid}

// 请求 URL
/users/42/orders/7

// 结果
r.PathValue("id")   → "42"
r.PathValue("oid")  → "7"
```

> 为新语法提供：只有使用 Go 1.22+ 的新 ServeMux 路由模式时才生效，老 `net/http` 默认 ServeMux 不支持 path 参数匹配

### 17. SetPathValue()

为该请求写入一个 URL 路径参数（Go 1.22+ 新路由功能）

```
func (r *Request) SetPathValue(name, value string)
```

与 `PathValue()` 方法是相反操作：

- `PathValue(name)` 用于从路由匹配中 读取 name
- `SetPathValue(name, value)` 用于人工 设置 / 覆盖 路径参数

**使用场景**

- 中间件修改路径参数
- 重写 URL 时保持参数语义一致
- 手动构建路由匹配后的 Request（例如手写 ServeMux 逻辑时）

### 18. Referer()

返回请求头中的 Referer 值

```
func (r *Request) Referer() string
```

- Referer 表示当前请求从哪里跳转来的页面
- 常用于来源分析、埋点统计、防盗链
- 如果请求头中没有该值，返回 ""

> 浏览器可能不会发送 Referer，不能强依赖

### 19. UserAgent()

返回请求头中的 `User-Agent`

```
func (r *Request) UserAgent() string
```

- 标识客户端类型，例如浏览器、手机 app、爬虫、命令行工具
- 常用于：判断来源设备类型 / 识别爬虫 / 日志分析
- 若不存在 `User-Agent`，返回 ""

### 20. Write()

将请求原始 HTTP 报文正常请求格式写入流

```
func (r *Request) Write(w io.Writer) error
```

会输出类似：

```
GET /path HTTP/1.1
Header-A: xxx
Header-B: yyy

<Body 数据>
```

**适用场景**

| 场景             | 说明                 |
|----------------|--------------------|
| 反向代理实现         | 将请求原样转发到下游服务器      |
| HTTP 请求录制/调试   | 输出完整原始请求内容         |
| 手工构造 HTTP 原始通讯 | 比如写原生 TCP HTTP 客户端 |

> 专业级 / 中间件 / 代理服务器 会用到的方法

### 21. WriteProxy()

输出适用于 HTTP 代理模式 的请求格式

```
func (r *Request) WriteProxy(w io.Writer) error
```

- 在代理模式中，请求行必须包含 完整 URL：`GET https://example.com/path HTTP/1.1`，而普通请求只需要：`GET /path HTTP/1.1`
- Forward Proxy / HTTP Proxy 实现时必须用它，例如在写 CONNECT / 代理 / 抓包工具 / 透明代理

### 22. ProtoAtLeast()

判断请求所使用的 HTTP 协议版本是否 ≥ 目标版本

```
func (r *Request) ProtoAtLeast(major, minor int) bool
```

| 请求版本     | `r.ProtoAtLeast(1,1)` | 含义                    |
|----------|-----------------------|-----------------------|
| HTTP/1.0 | false                 | 不支持 keep-alive 默认持久连接 |
| HTTP/1.1 | true                  | 支持持久连接、Host 必填        |
| HTTP/2   | true                  | 支持多路复用、二进制帧           |

> 常用于协议特性判断、底层服务器、代理、连接优化策略

