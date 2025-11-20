## http.Server

是一个 HTTP 服务器的抽象，负责：

| 职责     | 描述                                   |
|--------|--------------------------------------|
| 监听端口   | `net.Listen`                         |
| 接收连接   | `Accept()` TCP 连接                    |
| 解析请求   | 解析 HTTP 请求为 `http.Request`           |
| 分发请求   | 交给 Handler 处理                        |
| 控制生命周期 | Shutdown、超时、KeepAlive、MaxHeaderBytes |

去掉不常用和内部字段

```
type Server struct {
    Addr           string
    Handler        Handler
    TLSConfig      *tls.Config

    ReadTimeout    time.Duration
    WriteTimeout   time.Duration
    IdleTimeout    time.Duration
    ReadHeaderTimeout time.Duration

    MaxHeaderBytes int

    ErrorLog       *log.Logger
}
```

> Go 原生 http server 默认不设置超时（为了兼容老系统）。
> 但是实际生产中一定要设置，否则容易被慢请求攻击（Slowloris Attack）

### 1. `Addr`

监听地址：`Addr: ":8080"`

表示服务器要监听的主机和端口。如果只写 `:8080`，表示监听所有网卡。

* 部署时区分内网/公网（如 `127.0.0.1:8080` 只本地访问）
* 多端口多服务时也需要进行配置

### 2. `Handler`

负责处理请求的处理器

```
Handler: myMux // 如果不写，则使用 DefaultServeMux
```

如果 Handler 为 nil 或者被不写使用 `http.DefaultServeMux`

* 大型项目会自定义一个 `ServeMux` 或使用框架（如 gin 的 `engine`）

### 3. `TLSConfig`

HTTPS 配置

用于：

* 配置证书
* 限制加密算法
* 打开 HTTP/2

一般配合：`srv.ListenAndServeTLS("server.crt", "server.key")` HTTPS 才会用到。

### 4. `ReadTimeout`

限制：从客户端发起请求到**读取完 body** 的最大时间。

```
ReadTimeout: 5 * time.Second
```

防止有人慢慢传 body 导致 Server 卡住。

### 5. `WriteTimeout`

限制：服务器开始写响应到写完的最大时间。

```
WriteTimeout: 10 * time.Second
```

防止客户端一直不收数据，导致 Server 卡住。

### 6. `ReadHeaderTimeout`

限制：只用来读取请求头的超时

```
ReadHeaderTimeout: 2 * time.Second
```

这个是上生产必设的，因为请求头小、应当快速送达。如果这个时间过长可能说明攻击/异常客户端

### 7. `IdleTimeout`

keep-alive 长连接空闲超时

```
IdleTimeout: 120 * time.Second
```

影响连接池与长连接性能

**Web 服务推荐合理设置**

* API 服务：10-30s
* 网关/长连接：120s+

### 8. `MaxHeaderBytes`

限制请求头大小，默认 1MB。

```
MaxHeaderBytes: 1 << 20 // 1MB
```

如果不设置，有人可以构造一个巨大的 header 直接搞崩程序。

### 9. `ErrorLog`

```
ErrorLog: log.New(os.Stderr, "HTTP ERROR: ", log.LstdFlags)
```

控制 `http.Server` 的日志输出目的。

在生产环境会把它接到统一的 logging 系统中（如 `zap/logrus`）

## 主要方法

### `ListenAndServe`

```
func (srv *Server) ListenAndServe() error
```

| 项       | 含义                                         |
|---------|--------------------------------------------|
| **接收者** | `srv *Server`，说明它是绑定在某个 Server 实例上的        |
| **参数**  | 无（但会用 `srv.Addr` 和 `srv.Handler`）          |
| **返回值** | 正常优雅关闭返回 `http.ErrServerClosed`，其他错误返回异常信息 |

**内部行为**

- 调用 net.Listen("tcp", srv.Addr) 创建 TCP 监听套接字
- 成功后调用 srv.Serve(listener) 开始接受和处理连接。
- 若发生错误（端口占用、关闭等）返回 error

### `Serve`

```
func (srv *Server) Serve(l net.Listener) error
```

| 项   | 含义                               |
|-----|----------------------------------|
| `l` | 一个已经创建好的 `net.Listener`（调用者自己提供） |
| 返回值 | `error`                          |

- 不会自己创建 listener，而是接受别人提供的。
- 从 `l.Accept()` 循环接收连接，每个连接交给一个 goroutine 并用 `srv.Handler` 处理

ListenAndServe() = 自动创建 listener + 调用 Serve()
Serve() = 接管已有的 listener

### ListenAndServeTLS

```
func (srv *Server) ListenAndServeTLS(certFile, keyFile string) error
```

| 参数         | 类型       | 含义         |
|------------|----------|------------|
| `certFile` | `string` | PEM 格式证书文件 |
| `keyFile`  | `string` | PEM 格式私钥文件 |

与 ListenAndServe 类似，但会根据证书/私钥构造 `tls.Config`，然后调用 `srv.ServeTLS(...)`

### ServeTLS

```
func (srv *Server) ServeTLS(l net.Listener, certFile, keyFile string) error
```

与 Serve 类似，但会把 l 包装成 `tls.Listener` 再处理连接

| 参数                 | 含义          |
|--------------------|-------------|
| `l`                | 现有 listener |
| `certFile/keyFile` | TLS 证书与私钥   |

### Shutdown

优雅关闭， ctx 用于控制等待在处理请求的 goroutine 完成的超时

```
func (srv *Server) Shutdown(ctx context.Context) error
```

**底层行为**

- 停止接收新连接
- 等待正在处理的请求完成
- 超时则强制中断（由 ctx 控制）

### Close

```
func (srv *Server) Close() error
```

不等待、直接关闭所有连接。请求可能立刻中断。如果关闭连接出现问题会返回 error。

> 一般不用于生产关闭流程，主要用于错误处理或快速停止