# http.Handler

`http.Handler` 是整个 Go HTTP 体系的 核心抽象。
写的所有业务逻辑、路由、中间件，本质上都是在实现或包装 `http.Handler`。
Server -> Handler -> ServeHTTP 调用链。

## 核心定义

```
type Handler interface {
	ServeHTTP(ResponseWriter, *Request)
}
```

只要一个类型实现了 ServeHTTP 方法，它就是一个 Handler。
这意味着：

- 可以用 结构体 + 方法 组织逻辑
- 可以用函数来充当 Handler（借助 `http.HandlerFunc`）
