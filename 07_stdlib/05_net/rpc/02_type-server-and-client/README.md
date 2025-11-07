## 自定义 net/rpc 服务端

使用顶层函数进行创建 Server 用的是 默认服务器：rpc.DefaultServer（内部的全局变量）。

type Server 就是这个默认服务器的原型，它管理：

| 组件         | 用途                            |
|------------|-------------------------------|
| **服务注册表**  | 存放 `服务名 -> 方法` 映射             |
| **请求派发**   | 根据 `Service.Method` 找到方法并反射调用 |
| **连接处理逻辑** | 从连接中读取请求、执行方法、写回响应            |

所以当调用 `rpc.Register(&Cal{})` 等同于 `rpc.DefaultServer.Register(&Cal{})`

服务端的核心职责是：把本地的方法暴露出去，让客户端可以远程调用。

为了做到这件事，RPC 框架需要知道：

- 哪些对象提供服务（结构体实例）
- 哪些方法可以被调用
- 通过什么方式（TCP/HTTP/Unix socket）对外暴露
- 如何接受客户端连接并处理请求

### 创建 Server

使用 `func NewServer() *Server` 进行创建：

```
server := rpc.NewServer()
```

所有操作都可以通过自定义的 `server` 的方法进行调用，而不是使用 `rpc` 的顶层函数

### 接收连接

```
func (server *Server) Accept(lis net.Listener)
```

`rpc.Accept(lis)` = `rpc.DefaultServer.Accept(lis)`

使用 `NewServer` 创建实例后可以直接使用实例方法：`server.Accept(lis)`。

- 循环从 lis.Accept() 接受新连接
- 对每个连接自动启动 goroutine
- 并在该 goroutine 中调用 server.ServeConn(conn)
- 不需要自己写 `for + go ServeConn`

Accept 函数接受监听器上的连接，并处理每个传入连接的请求。Accept 函数会阻塞，直到监听器返回非空错误。
调用者通常在 go 语句中调用 Accept 函数。

### 使用 HTTP 连接

```
func (server *Server) HandleHTTP(rpcPath, debugPath string)
```

- `rpc.HandleHTTP()` = `rpc.DefaultServer.HandleHTTP(DefaultRPCPath, DefaultDebugPath)`

可以自定义路径并且不影响其他 server

功能等同于 `rpc.HandleHttp()` 但是还可以自定义路径：

- 把 server 挂载到 HTTP Handler
- RPC 通过 HTTP 使用 CONNECT 方法建立持久连接
- 调试面板挂载在 debugPath
- 客户端必须使用 `DialHTTPPath()` 并传入相同的 rpcPath 路径

### 注册服务

`server.Register / RegisterName`  是用来把要对外提供的服务对象注册进 RPC Server 的。
也就是让 RPC 系统知道有哪些方法可以被远程调用。

```
func (server *Server) Register(rcvr any) error
func (server *Server) RegisterName(name string, rcvr any) error
```

这两个方法等同于 `rpc.Register` 和 `rpc.RegisterName` 两个顶层函数。

Register 自动使用结构体的 类型名（大写开头）作为 服务名。
比如 `type Cal struct{}` 注册后客户端的调用方式就是：`client.Call("Cal.Add", ...)`，
也就是 `ServiceName=Cal MethodName=Add`

RegisterName 可以自定义服务名，不使用结构体名。
比如 `srv.RegisterName("Calculator", new(Cal))`，客户端调用方式变为：`client.Call("Calculator.Add", ...)`

### `Server.ServeConn`

```
func (server *Server) ServeConn(conn io.ReadWriteCloser)
```

`ServeConn` 方法等同于 `rpc.ServeConn` 函数，用于处理单条连接

- 每个连接单独处理
- 默认使用 Gob 编码
- 一般用于原生 TCP RPC

### `Server.ServeCodec`

```
func (server *Server) ServeCodec(codec ServerCodec)
```

`ServeCodec` `与ServeConn` 类似，但可以使用指定的编解码器来解码请求和编码响应。

### `Server.ServeRequest`

```
func (server *Server) ServeRequest(codec ServerCodec) error
```

用于手动控制每个请求。
`ServeRequest` 与 `ServeCodec` 类似，但是同步处理单个请求的，并且不会在请求完成后关闭编解码器。
适用于：限流、鉴权、中间件、熔断、监控、优雅关闭。

## `type Client`

`type Client` 也就是 RPC 客户端对象本体。

```
client, _ := rpc.Dial("tcp", "localhost:1234")
client.Call("Cal.Square", args, &reply)
```

client 就是 `*rpc.Client` 类型。

Client 控制 请求发送、响应接收、序列号映射、并发管理、Call 队列。

### 主要作用

| 职责        | 描述                                                       |
|-----------|----------------------------------------------------------|
| 连接管理      | 它内部持有一个 `io.ReadWriteCloser`（TCP、HTTP 隧道、WebSocket 等都可以） |
| 请求发送      | 对每次请求生成序列号（`Seq`），序列化并发送到连接                              |
| 响应匹配      | 根据返回包的序列号匹配到对应的请求（支持乱序返回）                                |
| 并发支持      | 多个 goroutine 可同时并发调用 `client.Call`                       |
| 错误与连接关闭处理 | 出错后会标记为 shutdown，所有未完成请求立刻得到错误回调                         |

### 创建方式

| 创建方式                             | 说明                               |
|----------------------------------|----------------------------------|
| `rpc.Dial(network, address)`     | 最常用，用默认 Gob 协议                   |
| `rpc.DialHTTP(network, address)` | 用 HTTP 方式承载 RPC（配合 `HandleHTTP`） |
| `rpc.NewClient(conn)`            | 如果自己构建了连接（例如 TLS、WebSocket 等）    |
| `rpc.NewClientWithCodec(codec)`  | 如果要自定义协议（如 JSON-RPC、Protobuf 等）  |

### 方法

| 方法                                            | 用途                              | 同步/异步  |
|-----------------------------------------------|---------------------------------|--------|
| `client.Call(serviceMethod, args, reply)`     | 阻塞等待结果                          | **同步** |
| `client.Go(serviceMethod, args, reply, done)` | 返回 `*Call`，后台执行，结果通过 channel 通知 | **异步** |
| `client.Close()`                              | 关闭连接，终止所有未完成调用                  | -      |
