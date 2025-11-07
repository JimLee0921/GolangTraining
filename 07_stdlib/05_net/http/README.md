# net/http

Go 的 net/http 标准库是 构建 Web 服务、HTTP 客户端、HTTP 服务器 的基础模块。

| 功能          | 描述                                                 |
|-------------|----------------------------------------------------|
| HTTP Server | 很轻松就能写一个 Web 服务（不用额外框架）                            |
| HTTP Client | 发起 HTTP 请求比 Requests/axios 更直接高效                   |
| 中间件机制       | 可以基于 `Handler` 组合实现路由 / 日志 / 鉴权                    |
| 性能强         | 标准库就内置了高性能服务器，不像 Python/JS 需要 Gunicorn/Nginx 之类的反代 |
| 并发内置        | 内部自带 goroutine 处理每个请求，天然并发                         |

## 学习文档

[https://pkg.go.dev/net/http](https://pkg.go.dev/net/http)：所有函数、类型、接口定义
[https://go.dev/src/net/http/](https://go.dev/src/net/http/)：最纯粹的官方源码示例

## 五大核心概念

| 概念           | 中文理解 | 比喻      | 在 Go 中的职责                   |
|--------------|------|---------|-----------------------------|
| **Server**   | 服务器  | 餐厅      | 接收客人（客户端）的请求并安排处理           |
| **Client**   | 客户端  | 顾客      | 发起请求，等待结果                   |
| **Handler**  | 处理器  | 厨房师傅    | 决定收到请求之后怎么处理、返回什么内容         |
| **Request**  | 请求   | 菜单 & 点单 | 顾客点了什么菜，携带参数、方法、URL         |
| **Response** | 响应   | 上菜      | 回传给客户端的内容，比如 HTML、JSON 二进制等 |

### Server（服务器）

Go 自带了一个 HTTP 服务器，不需要 Nginx、Tomcat、Flask、Django 的 Gunicorn 等。

最常见的启动方式是：

```
http.ListenAndServe(":8080", nil)
```

含义：

* 监听本机 8080 端口
* 第二个参数 `nil` 表示使用默认路由器（ServeMux）

服务器会自动：

* 监听 TCP 连接
* 拿到请求
* 交给对应的 **Handler** 来处理

**Server 本身不处理逻辑，它只负责分发请求**



## Client（客户端）

Go 中发 HTTP 请求也很简单：

```
resp, err := http.Get("https://example.com")
```

或更通用：

```
client := &http.Client{}
req, _ := http.NewRequest("GET", "https://example.com", nil)
resp, err := client.Do(req)
```

**Client 的任务：**

* 建立连接
* 发送 Request
* 等待 Response

---

## Handler（处理器）

可以理解为：

> **Handler = 收到请求后，应该做什么**

Handler 在 Go 中是一个接口：

```
type Handler interface {
ServeHTTP(ResponseWriter, *Request)
}
```

只需要实现 `ServeHTTP` 方法就行，所以可以自定义：

```
type MyHandler struct{}

func (h MyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
fmt.Fprintln(w, "Hello from MyHandler")
}
```

然后注册：

```
http.Handle("/", MyHandler{})
```

但更常用 `http.HandleFunc`（语法糖）：

```
http.HandleFunc("/", func (w http.ResponseWriter, r *http.Request) {
fmt.Fprintln(w, "Hello!")
})
```

---

## Request（请求）

这里包含客户端发来的所有信息：

```
func handler(w http.ResponseWriter, r *http.Request) {
fmt.Println(r.Method) // GET / POST
fmt.Println(r.URL) // URL 和路径
fmt.Println(r.Header) // 请求头
r.Body                // 请求体（POST 数据 / JSON）
}
```

典型读取 JSON：

```
body, _ := io.ReadAll(r.Body)
```

**Request = 客户端说“我想要什么”**

---

## Response（响应）

通过 `http.ResponseWriter` 写回内容：

```
func handler(w http.ResponseWriter, r *http.Request) {
w.WriteHeader(200) // 可选
w.Header().Set("Content-Type", "text/plain")
w.Write([]byte("Hello World"))
}
```

写 JSON：

```
w.Header().Set("Content-Type", "application/json")
json.NewEncoder(w).Encode(map[string]string{"msg": "ok"})
```

**Response = 服务端说“我给你什么”**
