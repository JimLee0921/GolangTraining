## http.ServeMux

ServeMux 的全名是：Serve + Mux，本身实现了 ServeHTTP 方法，即：

- 服务 + Multiplexer（多路复用器）
- HTTP 请求路由分发器

> ServeMux 是用来根据请求的 URL 路径，选择应该调用哪个 Handler 的组件

```
请求来了 -> ServeMux 看请求路径 -> 决定调用哪个 Handler -> 执行 ServeHTTP
```

因为 HTTP Server 本身 不负责路径匹配。 `http.Server` 只做两件事：

1. 接受连接，解析请求
2. 把请求交给 一个 Handler

而一个 Web 程序往往有不止一个处理逻辑（`/login`, `/users`, `/api/v1/orders`），若果只有一个 Handler 是不够的。

**结构流程图**

```
http.Server
    |
    | 处理所有请求的 Handler
    v
ServeMux (路由器)
    |
    | 根据 URL 路径匹配
    v
具体 Handler（你写的）
```

### 本质定义

```
type ServeMux struct {
    mu    sync.RWMutex
    m     map[string]muxEntry // 路由表，pattern -> handler
}
```

- 它内部维护了一个 路由表（map）
- key 是 pattern（例如 /api/、/user）
- value 是对应的 handler

### 默认全局路由器

在使用 `http.Handle(...)` 或者 `http.HandleFunc(...)` 是本质上就是在往全局变量 `var DefaultServeMux = &ServeMux{}`
里面注册路由。

而使用 `http.ListenAndServe(":8080", nil)`时 nil 的含义就是 `Handler == nil`，所以使用 DefaultServeMux。

### 总结

ServeMux 是 Go HTTP 服务器的默认路由器，它通过匹配 URL 路径来选择要调用的 Handler。

| 项                                             | 理解内容                         |
|-----------------------------------------------|------------------------------|
| ServeMux 是路由器                                 | 用于根据 URL 路径选择 Handler        |
| ServeMux 自己就是一个 Handler                       | 因为它实现了 `ServeHTTP`           |
| DefaultServeMux 是全局默认路由器                      | `Handle` 和 `HandleFunc` 都注册它 |
| `ListenAndServe(..., nil)` 使用 DefaultServeMux | 真正的路由中心在这里                   |

### 创建 ServeMux

最常见的就是使用 `http.NewServeMux()` 进行创建：

```
mux := http.NewServeMux()
```

## 主要方法

### Handle

ServeMux 会把 (pattern -> handler) 存入内部的路由映射表。

handler 必须满足 `ServeHTTP(ResponseWriter, *Request)` 如果传入 nil，会在运行时触发 panic，这属于设计规则。

```
func (mux *ServeMux) Handle(pattern string, handler Handler)
```

| 参数        | 类型                  | 含义                 |
|-----------|---------------------|--------------------|
| `pattern` | `string`            | 路由匹配规则             |
| `handler` | `http.Handler` 接口对象 | 真正处理此路由请求的 Handler |

### HandleFunc

HandleFunc 只是 Handle 的一个更简洁语法糖。内部会把函数转换为 HandlerFunc，再调用 `mux.Handle`。

函数 -> HandlerFunc 类型 -> 实现 ServeHTTP -> 成为 Handler

```
func (mux *ServeMux) HandleFunc(pattern string, handler func(ResponseWriter, *Request))
```

| 参数        | 类型          | 含义        |
|-----------|-------------|-----------|
| `pattern` | `string`    | 路由匹配规则    |
| `handler` | 函数 `(w, r)` | 不要求自己实现接口 |

### ServeHTTP

调用匹配到的 Handler 的 ServeHTTP

```
func (mux *ServeMux) ServeHTTP(w ResponseWriter, r *Request)
```

| 参数  | 类型               | 含义        |
|-----|------------------|-----------|
| `w` | `ResponseWriter` | 用于写回响应    |
| `r` | `*Request`       | HTTP 请求对象 |

**内部语义**

1. ServeMux 根据 r.URL.Path 找到匹配的 handler 。匹配规则是 最长前缀匹配（Longest Match）
2. 找到后调用 `handler.ServeHTTP(w, r)`
3. 如果找不到则调用默认的 NotFoundHandler（默认响应为 404 Not Found）

> ServeMux 自己不处理请求，它只做把请求交给谁处理的决定

### Handler

这个方法不会直接处理请求。只是根据 `Request.URL.Path` 查路由表，返回匹配结果。

一般不会直接用，为内部方法。可以在请求被执行前，先知道它最终会匹配到谁。

```
func (mux *ServeMux) Handler(r *Request) (h Handler, pattern string)
```

| 返回值       | 类型             | 含义            |
|-----------|----------------|---------------|
| `h`       | `http.Handler` | 匹配到的处理器       |
| `pattern` | `string`       | 用于匹配该处理器的路由规则 |

