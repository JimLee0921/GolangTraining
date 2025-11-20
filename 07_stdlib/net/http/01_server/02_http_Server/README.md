## 手动构建 http.Server

更多的在真实项目和生产环境中必须使用的方式。

对比一行式 `http.ListenAndServe`

- 可以配置超时防止慢请求攻击
- 可以进行优雅关闭（Shutdown）
- 自定义 `net.Listener`（限速/Unix socket/SO_REUSEPORT）
- HTTP/HTTPS 自动切换、证书更新
- 多服务共享端口/网关整合

> 自定义路由器 / 中间件栈 是都可以的

### 定义路由（Handler）

可以使用 http.DefaultServeMux，也可以自己创建 router：

```
mux := http.NewServeMux()
mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("Hello from http.Server!"))
})
```

### 创建 Server

推荐配置模板

```
srv := &http.Server{
    Addr:              ":8080",
    Handler:           myMux,
    ReadTimeout:        5 * time.Second,
    WriteTimeout:       10 * time.Second,
    IdleTimeout:        120 * time.Second,
    ReadHeaderTimeout:  2 * time.Second,
    MaxHeaderBytes:     1 << 20, // 1MB
}
```

如果 Handler 传入 nil 或者不传默认使用的就是 `http.DefaultServeMux`，推荐自定义路由传入。

srv 必须取地址 (&http.Server)

- ListenAndServe 会改变 Server 结构中 内部字段（比如连接状态）
- 需要传递 *Server（指针），这样修改才能生效

> 用结构体值会复制一份，不会影响原对象，所以必须用指针

### 启动服务

```
if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
    // 如果错误不是“服务器正常关闭”
    log.Fatalf("listen error: %v\n", err)
}
```

ListenAndServe 会返回一个 error，不是在 goroutine 里而是直接返回，因此需要正确处理。

当优雅关闭服务器时（比如调用 `srv.Shutdown(ctx)`）
ListenAndServe 会返回 `http.ErrServerClosed`，这不是错误，是正常行为，所以不应该把它当成 panic 或 FATAL。

### 设置超时

- 防慢速攻击：故意不发完请求头，拖住 goroutine
- 防客户端网速太慢写不完，服务器资源被耗尽

正确的生产参数通常：

```
ReadTimeout: 5 * time.Second
WriteTimeout: 10 * time.Second
IdleTimeout: 60 * time.Second
MaxHeaderBytes: 1 << 20 // 1MB
```