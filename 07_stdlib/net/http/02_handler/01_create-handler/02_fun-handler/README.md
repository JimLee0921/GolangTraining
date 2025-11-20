## 函数创建 `http.Handler`

只要某个东西有 ServeHTTP 方法那它就是 Handler。
而 Go 提供了一个适配器：`http.HandlerFunc`

```
type HandlerFunc func(ResponseWriter, *Request)
```

本身是一个函数类型，但关键在于它实现了 ServeHTTP：

```
func (f HandlerFunc) ServeHTTP(w ResponseWriter, r *Request) {
    f(w, r)  // 直接调用函数本体
}
```

> 只要函数符合 (w http.ResponseWriter, r *http.Request) 签名，它就能被当作 Handler 使用

### 底层原理

`http.HandleFunc` 本质就是包装。

```
http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello!")
})
```

底层其实已经做了：

```
http.Handle("/hello", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello!")
}))
```

> 也就是：普通函数 -> 通过 HandlerFunc -> 变成实现了 ServeHTTP 的 Handler

### 设计理念

- 把函数当成 Handler 使用
- 把 Handler 包装成另一个 Handler -> 中间件
- 框架（Gin/Echo/Fiber）都能复用 Handler 链
- 不需要写一堆 struct 实现接口

> HandlerFunc 让 Go 的 HTTP 编程既简单，又具备组合能力