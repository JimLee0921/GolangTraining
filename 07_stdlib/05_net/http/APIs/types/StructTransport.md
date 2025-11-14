# http.Transport 结构体

```
type Transport struct {...}
```

Transport 是 RoundTripper 的 默认实现，负责真正的网络通信，包括：

| 能力                     | 意义                      |
|------------------------|-------------------------|
| 连接池                    | 复用 TCP 连接，提高性能          |
| Keep-Alive             | 避免重复握手                  |
| TLS 握手与缓存              | HTTPS                   |
| 代理支持                   | HTTP proxy, PAC, 环境变量代理 |
| HTTP/1.1 / HTTP/2 自动切换 | 自动协商 ALPN               |
| 超时控制                   | 建连 / TLS / 等响应头         |
| 连接数量限制                 | 控制流量 / 防止打爆上游           |

这是 Go HTTP 客户端性能、并发、连接池、代理、HTTP/2、TLS、安全 等能力的核心所在。
可以理解为：

- Client.Do() ->  控制请求
- Transport -> 控制网络和连接池

## 创建方式

使用 `&http.Transport{}` 进行创建

```
tr := &http.Transport{}
client := &http.Client{Transport: tr}
// 如果不提供 tr
client.Transport = http.DefaultTransport
```

`http.DefaultTransport` 是一个预设较保守的 Transport

## 关键字段

按照功能进行划分

### 连接池 / 并发控制

```
MaxIdleConns          int // 全局空闲连接上限（默认 100）
MaxIdleConnsPerHost   int // 每主机空闲连接上限（默认 2）
MaxConnsPerHost       int // 每主机最大连接数 (默认不限 = 0)
IdleConnTimeout       time.Duration // 空闲连接在池中可存活多久
```

| 字段                    | 控制什么              | 意义            |
|-----------------------|-------------------|---------------|
| `MaxIdleConns`        | 整个进程池可保留多少空闲连接    | 并发量大时要调高      |
| `MaxIdleConnsPerHost` | 单个目标主机最多留多少连接等待复用 | 非常关键：默认 2，太小  |
| `MaxConnsPerHost`     | 限制总并发连接数          | 防止打爆上游服务      |
| `IdleConnTimeout`     | 空闲连接多久清掉          | 保守值通常 60~120s |

### 建连与 TLS 超时

```
DialContext: (&net.Dialer{
    Timeout:   5 * time.Second,   // 建立 TCP 连接的上限
    KeepAlive: 30 * time.Second,
}).DialContext,

TLSHandshakeTimeout: 5 * time.Second // TLS 握手超时
```

| 字段                    | 解决的问题         |
|-----------------------|---------------|
| `DialContext.Timeout` | 网络卡住或 IP 不通   |
| `TLSHandshakeTimeout` | 证书问题、SSL 协商卡住 |

### 等待服务器响应头

```
ResponseHeaderTimeout: 10 * time.Second
```

请求发出去后，等待第一字节响应头最多多久。用于防止上游慢到连响应头都不返回导致请求一直挂着

### 100-Continue 机制

```
ExpectContinueTimeout: 1 * time.Second
```

用于大 Body 上传优化。若服务器不允许，会避免提前上传大内容。

### HTTP 代理

```
Proxy: http.ProxyFromEnvironment,
// 或自定义
Proxy: func(req *http.Request) (*url.URL, error) { ... }
```

> 支持 HTTP_PROXY / HTTPS_PROXY / NO_PROXY

### 是否禁用 KeepAlive

```
DisableKeepAlives: bool
```

一般不要关闭，除非在写测试或在写调试代理需要故意避免连接复用

> 关闭 Keep-Alive 会导致每次请求都重新建连接导致性能变差

## transport 创建推荐模板

```
tr := &http.Transport{
    Proxy: http.ProxyFromEnvironment,

    DialContext: (&net.Dialer{
        Timeout:   5 * time.Second,
        KeepAlive: 30 * time.Second,
    }).DialContext,

    TLSHandshakeTimeout:   5 * time.Second,
    ResponseHeaderTimeout: 10 * time.Second,
    ExpectContinueTimeout: 1 * time.Second,

    MaxIdleConns:          200,
    MaxIdleConnsPerHost:   50,
    IdleConnTimeout:       90 * time.Second,
}
client := &http.Client{
    Transport: tr,
    Timeout:   15 * time.Second,
}
```

## http.RoundTripper 接口

```
type RoundTripper interface {
    RoundTrip(*Request) (*Response, error)
}
```

RoundTripper 是一个接口，表示执行单个 HTTP 事务的能力，即获取给定请求的响应。
RoundTripper 必须能够安全地供多个 goroutine 并发使用。
只一个方法：RoundTrip，负责：把 Request 变成 Response。

在 Go 的 HTTP 客户端中：`Client.Do(req)  <-  RoundTripper.RoundTrip(req)`，
也就是说：RoundTripper 是真正执行一次 HTTP 请求的接口。

### 实现

`http.Transport` 是 RoundTripper 接口最重要的一个实现，Transport 是 RoundTripper 的默认/标准实现。

而 Transport 通常是用来配置 连接池 / 代理 / TLS / HTTP/2 / 超时策略 / 重用连接 的地方，所以
`Client.Transport (RoundTripper) -> 通常是 Transport`，如果不写 `client := &http.Client{}`，
那么默认就是 `client.Transport = http.DefaultTransport`。

### 行为规范

RoundTrip 必须返回一个带可读 Body 的 Response，并且调用者必须关闭它。也就是说：

```
resp, err := transport.RoundTrip(req)   // 只有建立流通道
defer resp.Body.Close()                 // 消费方必须关闭
```

而 `Client.Do()` 会进一步处理 重定向策略 / CookieJar / Timeout / Context 取消

1. Client.Do() 并不直接发请求，真正发请求的是 `RoundTripper.RoundTrip(req)`
2. Transport 是默认的 RoundTripper
3. Transport 负责连接池 / TLS / 代理 / HTTP/2 / 超时