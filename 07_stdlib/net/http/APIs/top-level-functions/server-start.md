## Server Start Functions

这些是把服务跑起来的顶层函数

### `http.ListenAndServe`

最常用的一行式，监听 TCP 并开始处理 HTTP 请求。

ListenAndServe 监听 TCP 网络地址 addr，然后调用 Serve并传入处理程序来处理传入连接的请求。已接受的连接配置为启用 TCP 保活机制。

处理程序 handler 通常为 nil，在这种情况下使用 DefaultServeMux 。

```
func ListenAndServe(addr string, handler Handler) error
```

**参数**

- addr：":8080" 表示本机所有网卡 8080 端口；"127.0.0.1:8080" 仅本机回环
- handler 为 nil：走 DefaultServeMux（配合 http.Handle/HandleFunc 注册路由）

**返回值**

阻塞直到出错，常见返回：

- nil：极少见（通常只有自己关 listener）
- `http.ErrServerClosed`：调用了优雅关闭（Server.Shutdown）后返回。生产里应识别此错误不当成致命错误
- 其它错误：端口占用、权限问题等

```
// 最短用法：nil => 使用默认路由器 DefaultServeMux
log.Fatal(http.ListenAndServe(":8080", nil))
```

开发期可以 log.Fatal(err)；生产代码更常见：

```
if err := http.ListenAndServe(":8080", nil); err != nil && err != http.ErrServerClosed {
    log.Fatalf("listen: %v", err)
}
```

### `http.ListenAndServeTLS`

一行式 HTTPS 启动（会自动启用 HTTP/2，除非禁用）。

```
func ListenAndServeTLS(addr, certFile, keyFile string, handler Handler) error
```

```
// certFile / keyFile 必须是 PEM 文件
log.Fatal(http.ListenAndServeTLS(":8443", "server.crt", "server.key", nil))
```

- 证书要求：服务端证书（链）+ 私钥（PEM）
- 若需要更复杂的 TLS（SNI、多证书、强制 TLS1.3、ClientAuth 等），更推荐用 `http.Server` + TLSConfig

### `http.Serve`

当自己先创建好 listener 监听端口后，用 `http.Serve` 把 listener 交给 HTTP 服务处理

```
func Serve(l net.Listener, handler Handler) error
```

**适用场景**

* 手动管理监听端口
* 自定义监听方式（重用端口、SO_REUSEPORT、优雅重启）
* 限制最大连接数
* 接入自定义 `net.Listener`（比如 TLS、Unix Socket）

```
ln, err := net.Listen("tcp", ":8080")
if err != nil { log.Fatal(err) }


// 用顶层函数
log.Fatal(http.Serve(ln, mux))
// 对应方法版是 (&http.Server{Handler: mux}).Serve(ln)
```

**底层原理**

默认的写法 `http.ListenAndServe(":8080", handler)` 内部其实做了两件事：

```
ln, _ := net.Listen("tcp", ":8080")
http.Serve(ln)
```

也就是说 `ListenAndServe` 封装了 `Listen + Serve`

而 `net.Listener` 是一个接口：

```
type Listener interface {
    Accept() (Conn, error)
    Close() error
    Addr() Addr
}
```

也就是说可以使用任何实现了 Listener 的东西代替 TCP socket，比如：

| Listener 类型                           | 意义                       |
|---------------------------------------|--------------------------|
| `net.Listen("tcp", ...)`              | 普通 TCP 端口                |
| `net.Listen("unix", "/tmp/app.sock")` | Unix 域套接字（Nginx 反代高性能场景） |
| `tls.NewListener(ln, cfg)`            | 给 HTTP Server 增加 TLS 支持  |
| `netutil.LimitListener(ln, n)`        | 限制最大连接数                  |
| 自己包装 Listener                         | 加速、限流、审计、追踪、IP 黑白名单      |

### `http.ServeTLS`

```
func ServeTLS(l net.Listener, handler Handler, certFile, keyFile string) error
```

与 http.Serve 类似，但直接在已有 listener 上以 HTTPS 服务

```
ln, err := net.Listen("tcp", ":8443")
if err != nil { log.Fatal(err) }
mux := http.NewServeMux()
mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintln(w, "hello https")
})

log.Fatal(http.ServeTLS(ln, mux, "server.crt", "server.key"))
```

如果需要 复杂 TLS 行为（SNI、多证书、动态证书、mTLS），更常见的是：

- 自己创建 tls.Config
- 用 `http.Server{TLSConfig: ...}` + `Server.ServeTLS(ln, "")`（方法形态）
- 或者先把 ln 包成 `tls.NewListener(ln, tlsConfig)` 再走 `Server.Serve`

## 补充

### 和 `http.Server` 的关系

顶层函数是“快捷方式”，内部其实就是创建一个 `http.Server` 然后调用对应方法：

- ListenAndServe = `(&http.Server{Addr: addr, Handler: handler}).ListenAndServe()`
- Serve = `(&http.Server{Handler: handler}).Serve(l)`
- ListenAndServeTLS / ServeTLS 同理（加上证书参数）

### 生产更推荐专业用法

顶层函数够用但不够可控。生产里通常需要：

- 超时：ReadTimeout、WriteTimeout、IdleTimeout
- 报文限制：MaxHeaderBytes
- 优雅关闭：Shutdown(ctx)（让连接把在途请求处理完再退出）
- TLSConfig：SNI、多证书、mTLS、强制最低版本等

> 优雅关闭后 `Serve/ListenAndServe` 会返回 `http.ErrServerClosed`，不要当成致命错误

### handler == nil

- 顶层函数 / `http.Server` 的 Handler 字段为 nil 时，等价于使用 `http.DefaultServeMux`
- 因此用 `http.Handle/HandleFunc` 注册的路由就会生效；但如果传了自定义 Handler，默认的 mux 不会被使用