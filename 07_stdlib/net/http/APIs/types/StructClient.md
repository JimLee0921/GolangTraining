# http.Client

Client 是 HTTP 客户端，用于发送 Request 并接收 Response。

| 用途        | 说明                  |
|-----------|---------------------|
| 建立连接      | TCP / TLS / HTTP2   |
| 维护连接池     | Keep-Alive、多请求复用    |
| 超时控制      | 避免卡住不返回             |
| 代理        | 支持 HTTP / HTTPS 代理  |
| 重定向策略     | 是否跳转，跳转几次           |
| Cookie 管理 | 可以通过 CookieJar 自动处理 |

> Request -> 由 Client 发送 -> 得到 Response

## 主要字段

```
type Client struct {
    Transport     RoundTripper   
    CheckRedirect func(req *Request, via []*Request) error
    Jar           CookieJar      
    Timeout       time.Duration  
}
```

> Client = 连接池 + 超时控制 + 代理 + Cookie + HTTP/1.1 & 2 + TLS

| 字段            | 负责什么                      | 用途            |
|---------------|---------------------------|---------------|
| Transport     | 低层连接、HTTP1.1/2、代理、TLS、连接池 | 高级优化点，决定性能    |
| CheckRedirect | 控制是否跟随 301/302 等跳转        | 登录 / 安全 / 防盗链 |
| Jar           | 自动管理 Cookie               | 会话维持、模拟浏览器    |
| Timeout       | 请求总超时时间（包括 DNS/TLS/读写）    | 防止卡死，非常重要     |

## 核心方法

### 1. Do()

Client 的核心方法，所有请求最终都走它。

```
func (c *Client) Do(req *Request) (*Response, error)
```

- 所有 Get / Post / PostForm / Head 最终都 调用 `Do()`
- 因为只有 Do() 能接受完整的自定义 Request（Headers、Context、Body、Cookie...）

**底层操作**

1. 根据 Request + Client 配置（Transport / Jar / Redirect）发送请求
2. 接收服务器返回
3. 返回 Response
4. 不会自动关闭 `resp.Body`（必须手动 Close）

> 在实际工程中：99% 的请求是 `client.Do(req)` 发送的

### 2. Get()

发起 GET 请求的快捷方式，等价于 `req, _ := http.NewRequest("GET", url, nil)` + `resp, err := client.Do(req)`

```
func (c *Client) Get(url string) (resp *Response, err error)
```

- 不适合复杂请求
- 没办法设置 Header、Cookie、Context、Body
- 适用于：简单脚本、快速测试、demo
- 在生产代码中更推荐 `Do()`

### 3. Head()

发起 HEAD 请求，只获取响应头，不获取 Body，也就是轻量版的 GET

```
func (c *Client) Head(url string) (resp *Response, err error)
```

**用途**

| 场景                     | 为什么用 HEAD |
|------------------------|-----------|
| 检查资源是否存在               | 不下载内容     |
| 获取 Content-Length/Type | 准备文件下载    |
| CDN / 缓存探测             | 低成本探测资源状态 |

### 4. Post()

```
func (c *Client) Post(url, contentType string, body io.Reader) (resp *Response, err error)
```

发起 POST 请求，带请求体，等价于:

```
req, _ := NewRequest("POST", url, body)
req.Header.Set("Content-Type", contentType)
client.Do(req)
```

- 常用于提交 JSON / 提交表单 / 上传小块数据
- 它只设置 `Content-Type`，不会自动做编码
- 大型/复杂 POST 仍建议使用 Do() 手动构造请求

### 5. PostForm

发送 表单格式 的 POST 请求：`Content-Type: application/x-www-form-urlencoded`

```
func (c *Client) PostForm(url string, data url.Values) (resp *Response, err error)
```

等价于：

```
body := strings.NewReader(data.Encode())
client.Post(url, "application/x-www-form-urlencoded", body)
```

- 常用于传统网站表单提交 / OAuth2 token 授权 / 登录接口（非 JSON）
- 会自动将 `url.Values` 进行 `key=value&...` URL 编码

### 6. CloseIdleConnections()

非常重要，和连接池相关，Client 默认使用 连接复用（Keep-Alive）：一个 Client = 一个连接池。

`CloseIdleConnections()` 的作用是关闭池中空闲但未使用的连接，不影响正在使用的连接。

```
func (c *Client) CloseIdleConnections()
```

**使用场景**

- 服务长时间运行，想释放多余连接
- 服务器端点变更（DNS 变更 / 微服务滚动更新）
- 优雅退出程序时，清理连接资源

> 不是必须调用，只在 资源回收/重启/切换服务器 时用

## 超时 timeout 配置

http.Client 的超时（timeout）配置。这一块直接决定请求是否卡死、是否能优雅失败。

**三层超时手段**

1. 整体超时：`Client.Timeout` 一刀切，覆盖 DNS、TCP 连接、TLS 握手、重定向、读取响应体全部阶段
2. 逐阶段超时：Transport 与 `net.Dialer` 上的细粒度字段，控制建连 / TLS / 等首包 / 继续 100-continue / 空闲连接等
3. 每次请求的上下文：Context（NewRequestWithContext），任意时刻可取消/设定 deadline，适合流式/长连接等不适合一刀切的场景

### Client.Timeout

```
client := &http.Client{
    Timeout: 10 * time.Second, // 整个请求的最长期限，包含读 Body
}
```

- 覆盖从发起到读 body 完成的全部流程（包括多次重定向）
- 到时未完成会自动取消请求，错误通常是 context deadline exceeded
- 适合常规 API 调用、下载/上传有上限的任务
- 不适合流式响应（SSE、gRPC 长轮询、长下载），会被总时长硬切断

> 默认的 `http.DefaultClient` 没有超时，生产环境一定要设

### Transport & Dialer

当需要更精细的控制（或做长流任务时不用整体超时），配置 Transport：

```
tr := &http.Transport{
    Proxy: http.ProxyFromEnvironment,

    // 1 建连（TCP）阶段
    DialContext: (&net.Dialer{
        Timeout:   5 * time.Second,  // TCP 连接超时
        KeepAlive: 30 * time.Second, // TCP keep-alive 心跳
    }).DialContext,

    // 2 TLS 阶段
    TLSHandshakeTimeout:   5 * time.Second, // TLS 握手

    // 3 等首包（响应头）阶段
    ResponseHeaderTimeout: 5 * time.Second, // 发送完请求后等待响应头的时间

    // 4 100-continue（大 body 上传前试探）
    ExpectContinueTimeout: 1 * time.Second, // 等服务器“100 Continue”的等待上限

    // 5 连接池相关（不直接是超时，但强相关）
    IdleConnTimeout:       90 * time.Second, // 空闲连接在池中存活多久
    MaxIdleConns:          100,              // 全局空闲连接上限
    MaxIdleConnsPerHost:   10,               // 每主机空闲连接上限
    MaxConnsPerHost:       0,                // 每主机最大连接数（0=不限制）
}
client := &http.Client{Transport: tr}
```

### Context 控制取消

每次请求的可取消/自定义时限

1. 创建 context：`ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second) // 每次三秒请求上限`
2. 注意使用 defer 关键字确保 context 的关闭：`defer cancel()`
3. 使用 `http.NewRequestWithContext` 创建 request ：`req, _ := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)`
   （如果传入的是 POST 等方法替换 nil 为对应的数据）
4. 使用 `client.Do()` 发起请求：`resp, err := client.Do(req)`

适合以下场景：

- 流式/长时间任务（SSE、长下载、chat 流）
- 需要随时中止（用户取消、UI 关闭）
- 需要不同请求不同超时策略

> Context 的取消会穿透到 Transport，优雅地关闭连接

## 补充

### 默认客户端 `http.DefaultClient`

Go 给了一个全局默认客户端：`http.Get()` 和 `http.Post()`

- 底层等价于 `http.DefaultClient.Do(req)`
- 默认 Client 没有超时，可能卡住不返回
- 生产中推荐使用 `&http.Client` 传入 Timeout 字段进行自定义客户端创建并使用

### 不要频繁创建

Client 是长生命对象，不要频繁 new

```

client := &http.Client{Timeout: 5 * time.Second} // 程序启动创建一次

for each request:
client := &http.Client{} // 这样会破坏连接池，性能变差

```

- Client 内部有 连接池，可以重用 -> 性能高
- 每次新建 Client -> 连接池无法复用 -> 性能差
- 一个程序 / 一个服务 / 一个模块 -> 通常一个 Client

