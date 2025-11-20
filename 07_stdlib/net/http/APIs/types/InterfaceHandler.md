## http.Handler

Handler 是一个可以处理 HTTP 请求的对象。它代表了 Web 服务中的请求处理逻辑。
在 Go 的设计中，所有能响应请求的东西，都是一个 Handler。

主要用于编写路由器、中间件、框架。

```
type Handler interface {
    ServeHTTP(ResponseWriter, *Request)
}
```

这个接口只有一个方法 ServeHTTP，但这个接口决定了：

- 谁可以成为处理请求的对象
- Server 如何把请求交给代码处理
- 凡是实现了 ServeHTTP 的对象，都可以当作 Handler

### ServeHTTP

```
ServeHTTP(w http.ResponseWriter, r *http.Request)
```

| 参数  | 类型                    | 含义                        | 做什么           |
|-----|-----------------------|---------------------------|---------------|
| `w` | `http.ResponseWriter` | 用于写响应（状态码、头、body）         | 向客户端返回内容      |
| `r` | `*http.Request`       | 请求数据对象（方法、URL、头、Body、表单等） | 从这里读到客户端传来的东西 |


