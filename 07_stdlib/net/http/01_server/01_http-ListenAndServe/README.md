## `http.ListenAndServe`

使用 `http.ListenAndServe` 创建 服务：

```
http.ListenAndServe(":8080", nil)       // 用默认路由器 DefaultServeMux
// 或
http.ListenAndServe(":8080", myHandler) // 显式传入 Handler
```

底层等价于：

```
func ListenAndServe(addr string, h Handler) error {
    srv := &http.Server{Addr: addr, Handler: h}
    return srv.ListenAndServe()
}
```

也就是：底层临时创建了一个 `http.Server` 并启动。这意味着：

- 不能配置超时/最大头大小/优雅关闭等细节（后面说限制）
- 返回值是 阻塞的 error。成功启动后只有出错或进程退出才返回

常见返回：

- 端口被占用：listen tcp :8080: bind: address already in use
- 如果用 Server.Shutdown（一行式没法用）关闭才会得到 http.ErrServerClosed

### handler 参数

1. nil：表示使用 全局默认路由器 `http.DefaultServeMux`
   可以用 `http.HandleFunc` 或 `http.Hadnle` 进行路由添加（具体在 Handle 中再进行讲解）
    ```
    http.HandleFunc("/hello", hello)
    http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))
    http.ListenAndServe(":8080", nil)
    ```

2. 也可以使用 `http.HandlerFunc` 创建 Handler（任何实现了 `http.Handler` 接口的对象）然后使用传入
    ```
    myHandler := http.HandlerFunc(){}
    http.ListenAndServe(":8080", myHandler)
    ```

> 两种风格都行，本质只是谁来接住每个请求

### 自定义 Handler

```
http.ListenAndServe(":8080", myHandler)
```

这里的 myHandler 必须实现 `http.Handler` 接口，可以用 `http.HandlerFunc` 进行创建

```
myHandler := http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
    switch {
    case request.URL.Path == "/":
        fmt.Fprintln(writer, "this is home")
    case strings.HasPrefix(request.URL.Path, "/api/"):
        fmt.Fprintln(writer, "api:", request.URL.Path)
    default:
        http.NotFound(writer, request) // 使用快捷函数返回 404
    }
})
```