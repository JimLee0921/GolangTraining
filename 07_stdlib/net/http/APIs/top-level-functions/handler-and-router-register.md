## Handler & 路由注册

这一部分是理解请求如何被分发到代码

在 Go 的 net/http 里：

路由器 = 实现了 `http.Handler` 接口的东西
最常用的路由器是 `http.ServeMux`（也叫多路复用器、路由分发器）。

而 http 包内部提供了一个全局的默认路由器：`http.DefaultServeMux`，
当使用 `http.ListenAndServe(":8080", nil)` 传入的 handler = nil 时，
实际上：`handler = http.DefaultServeMux`

所以用 `http.Handle(...)` 注册路由，就是往 DefaultServeMux 里注册。

### `http.Handle`

```
func Handle(pattern string, handler Handler)
```

手动注册 一个实现了 `http.Handler` 接口的对象。

| 参数        | 类型                | 含义                    |
|-----------|-------------------|-----------------------|
| `pattern` | `string`          | 路由匹配规则（例如 `/user/`）   |
| `handler` | `http.Handler` 接口 | 任何实现了 `ServeHTTP` 的对象 |

```
type HelloHandler struct{}

func (HelloHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintln(w, "hello")
}

func main() {
    http.Handle("/hello", HelloHandler{}) // 向默认路由器注册
    http.ListenAndServe(":8080", nil)
}
```

所有路由最终都得变成 ServeHTTP 调用。

### `http.HandleFunc`

```
func HandleFunc(pattern string, handler func(ResponseWriter, *Request))
```

| 参数        | 类型                                             | 含义        |
|-----------|------------------------------------------------|-----------|
| `pattern` | `string`                                       | 路由匹配规则    |
| `handler` | `func(w http.ResponseWriter, r *http.Request)` | 普通函数，不是接口 |
