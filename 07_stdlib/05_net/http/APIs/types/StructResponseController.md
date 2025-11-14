# http.ResponseController

`http.ResponseController` 是一个控制器，允许更安全、更统一地访问 ResponseWriter 的扩展功能 （比如 Flush、Hijack、Push、Trailer
等）。

在 Go 1.19 及以前，要使用这些扩展功能，需要写大量类型断言：

```
if f, ok := w.(http.Flusher); ok {
    f.Flush()
}
if hj, ok := w.(http.Hijacker); ok {
    conn, rw, _ := hj.Hijack()
}
```

这种写法既繁琐又容易出错。
所以从 Go 1.20 开始，官方引入了 ResponseController 来统一管理这些底层控制。

## 核心接口

```
type ResponseController struct {
    rw ResponseWriter
}
```

## 创建函数

使用 `http.NewResponseController` 进行创建

```
func NewResponseController(rw ResponseWriter) *ResponseController

ctrl := http.NewResponseController(w)
```

## 相关方法 / 函数

| 方法                                                                           | 说明                                                     |
|------------------------------------------------------------------------------|--------------------------------------------------------|
| `func (c *ResponseController) Flush() error`                                 | 立即把缓冲区中的响应数据发送到客户端（替代 Flusher）                         |
| `func (c *ResponseController) Hijack() (net.Conn, *bufio.ReadWriter, error)` | 接管底层连接（如 WebSocket、原始 TCP）                             |
| `Push(target string, opts *PushOptions)`                                     | http.Pusher 接口，发起 HTTP/2 Server Push                   |
| `func (c *ResponseController) SetReadDeadline(deadline time.Time) error`     | 设置读超时（针对连接）                                            |
| `func (c *ResponseController) SetWriteDeadline(deadline time.Time) error`    | 设置写超时                                                  |
| `func (c *ResponseController) EnableFullDuplex() error`                      | 启用全双工模式（Go 1.22 + ）                                    |
| `Trailer()`                                                                  | 用 ResponseWriter.Header() 控制返回响应的 Trailer 头（延迟 Header） |
