## rpc 顶层函数

rpc 的几个顶层函数都主要用于服务端或客户端的创建。
而服务端和客户端也可以使用 `NewServer` 或 `NewClient` 进行创建。

- 顶层函数路径＝用默认全局 Server / Client + 若干便捷入口，是方便好用的全局默认
- NewServer / NewClient 路径＝自己持有并管理一个（或多个）Server / Client 实例，可以自行编排

| 侧别         | 作用             | 代表函数                                                                                          |
|------------|----------------|-----------------------------------------------------------------------------------------------|
| 服务端 Server | 接收连接、注册服务、处理请求 | `Register`, `RegisterName`, `Accept`, `HandleHTTP`, `ServeConn`, `ServeCodec`, `ServeRequest` |
| 客户端 Client | 建立连接、创建客户端对象   | `Dial`, `DialHTTP`, `DialHTTPPath`                                                            |

## 顶层函数 VS 实例方法

### Server

| 类型          | 顶层函数（使用默认全局 Server/Client）    | 等价的实例方法（使用 NewServer/NewClient）         | 说明                              |
|-------------|-------------------------------|-----------------------------------------|---------------------------------|
| 注册服务        | `rpc.Register(svc)`           | `server.Register(svc)`                  | 顶层函数操作的是 `rpc.DefaultServer`    |
| 自定义服务名      | `rpc.RegisterName(name, svc)` | `server.RegisterName(name, svc)`        | 完全同义，只是 Server 实例不同             |
| TCP 自动监听服务  | `rpc.Accept(lis)`             | `server.Accept(lis)`                    | 顶层函数调用的是 `DefaultServer.Accept` |
| HTTP 暴露 RPC | `rpc.HandleHTTP()`            | `server.HandleHTTP(rpcPath, debugPath)` | 顶层版本使用默认路径；实例可自定义路径             |
| 单连接 Gob 服务  | `rpc.ServeConn(conn)`         | `server.ServeConn(conn)`                | 完全对等                            |
| 单连接自定义协议    | `rpc.ServeCodec(codec)`       | `server.ServeCodec(codec)`              | 完全对等                            |
| 单次请求处理      | `rpc.ServeRequest(codec)`     | `server.ServeRequest(codec)`            | 完全对等                            |

### Client

| 操作          | 顶层函数版本                    | 实例版本                           | 场景                             |
|-------------|---------------------------|--------------------------------|--------------------------------|
| TCP 连接      | `rpc.Dial(network, addr)` | `NewClient(conn)`              | 日常最常用                          |
| HTTP RPC 连接 | `rpc.DialHTTP(...)`       | `NewClient(httpTransportConn)` | HTTP/反向代理情境                    |
| 自定义协议连接     | —                         | `NewClientWithCodec(codec)`    | JSON-RPC / Protobuf RPC / 自研协议 |

## 顶层 Server 函数

### Accept

```
func Accept(lis net.Listener)
```

在给定的 net.Listener 上持续接受进来的连接。
每来一个连接，就交给默认服务器去处理该连接上的所有 RPC 请求。

**行为特点**

- 阻塞：会一直阻塞在监听循环（通常放在主协程里或单独起一个 goroutine）
- 并发：每个新连接会在单独 goroutine 中服务，不同连接之间并发处理；单个连接内的请求是串行处理（默认 Gob/默认 Server）
- 使用前提：已调用 rpc.Register/RegisterName 注册好服务

**典型用法**

```
l, err := net.Listen("tcp", ":1234")
if err != nil { panic(err) }
rpc.Register(&Arith{})
rpc.Accept(l) // 阻塞，直到 listener 关闭或进程退出
```

- 不要在 Accept 之前忘了 Register，否则客户端调用会报找不到方法
- 相比 ServeConn：Accept 包办了接受新连接这一步；ServeConn 只处理已拿到的单条连接

> `rpc.Accept` 就等同于 `for + go ServeConn` 进行并发处理

### HandleHTTP

```
func HandleHTTP()
```

把默认服务器挂到 HTTP 上：注册两个 HTTP 路径（/rpc 和 /debug/rpc），从而可以通过 HTTP 承载 RPC。

> 仅仅是注册 HTTP handler，仍需自己 `http.ListenAndServe` 或 `http.Serve` 启动 HTTP 服务

**典型用法**

```
rpc.Register(&Arith{})
rpc.HandleHTTP()                         // 注册到默认路径
l, _ := net.Listen("tcp", ":1234")
go http.Serve(l, nil)                    // 或 http.ListenAndServe
// client 侧用 rpc.DialHTTP("tcp", "host:1234") 连接
```

- 客户端要用 rpc.DialHTTP（或 DialHTTPPath），不是普通的 Dial
- 不能再使用 Accept 并且需要使用 `http.Server` 或 `http.ListenAndServe` 启动服务
- 如果需要自定义路径，考虑 `Server.HandleHTTP(rpcPath, debugPath)`（实例方法），或客户端用 `DialHTTPPath`

### Register

```
func Register(rcvr any) error
```

把 `rcvr`（通常是指针接收者的实例）注册到默认服务器；其满足 RPC 签名的方法将被暴露为远程可调用方法。
客户端调用名默认为 `TypeName.Method`（可以用 RegisterName 自定义服务名）

> 导出方法满足规范见 rpc 下的 README

**典型用法**

```
type Arith struct{}
type Args struct{ A, B int }
func (a *Arith) Multiply(args Args, reply *int) error { *reply = args.A*args.B; return nil }

if err := rpc.Register(&Arith{}); err != nil { panic(err) }
```

### RegisterName

```
func RegisterName(name string, rcvr any) error
```

与 Register 相同，但服务名可以自定义指定为 name，而不是类型名。客户端调用名为 `name.Method`。

**适用场景**

- 版本化/别名：`OrderService`、`OrderServiceV2`
- 避免类型名暴露，或多个不同类型合并到同一个服务名下（不建议乱用，注意方法命名冲突）

**典型用法**

```
rpc.RegisterName("OrderService", &OrderV2{})
// 客户端：client.Call("OrderService.Get", args, &reply)
```

> 与默认类型名注册的服务重名/方法同名可能冲突

### ServeCodec

```
func ServeCodec(codec ServerCodec)
```

在当前 goroutine 中，使用自定义提供的 ServerCodec（自定义协议/编解码）来持续处理该底层连接上的多个请求，直到出错或关闭底层连接

**行为特点**

- 单连接上的循环处理：这更像是 ServeConn 的可插拔协议版。ServeConn 等价于 `ServeCodec(defaultGobCodec)`
- 需要提供一个满足 ServerCodec 接口的实现（读请求头体、写响应、关闭）
- 通常会在每接收到一个连接时 `go ServeCodec(jsonrpc.NewServerCodec(conn))` 之类地并发处理

**典型用法**

JSON-RPC 服务端

```
l, _ := net.Listen("tcp", ":1234")
for {
    conn, _ := l.Accept()
    go rpc.ServeCodec(jsonrpc.NewServerCodec(conn)) // 用 JSON-RPC 处理此连接
}
```

- 自定义 ServerCodec 实现要保证线程安全（在需要的地方）以及正确处理消息边界和错误
- 默认只处理一个连接；若要监听多个连接，自行 Accept 循环
- 若不需要自定义协议，直接 ServeConn 更简单。若只想处理单个请求（不是整条连接），用 ServeRequest。

### ServeConn

```
func ServeConn(conn io.ReadWriteCloser)
```

在当前 goroutine 上，使用默认编解码（Gob）在这条连接 conn 上持续处理多个请求，直到连接关闭或出错

**行为特点**

- 阻塞处理该连接；一般会 `go rpc.ServeConn(conn)` 让每条连接独立 goroutine 处理
- 多条连接并发需要在外层 `for { Accept(); go ServeConn(conn) }`

**典型用法**

```
l, _ := net.Listen("tcp", ":1234")
rpc.Register(&Arith{})
for {
    conn, _ := l.Accept()
    go rpc.ServeConn(conn) // 每个连接一个 goroutine
}
```

- 需要为每个连接起 goroutine；否则一个慢连接会拖死后续连接处理
- Accept = 自动 `Accept + ServeConn`
- ServeCodec = ServeConn 的可插拔协议版

### ServeRequest

```
func ServeRequest(codec ServerCodec) error
```

基于给定的 ServerCodec 仅处理一个请求（单次），返回后不自动关闭 codec；是否继续处理更多请求可以自行决定——可以再次调用
ServeRequest。

**行为特点**

- 适合需要精细控制请求生命周期的场景：比如自定义多路复用、单/双工协议、限流、每请求鉴权等
- 需要自己管理循环与错误处理

**典型用法**

```
conn, _ := l.Accept()
codec := jsonrpc.NewServerCodec(conn)
for {
    if err := rpc.ServeRequest(codec); err != nil {
        // 例如客户端断开/协议错误，决定是否 break
        break
    }
}
conn.Close()
```

- 不会自动关闭 codec/连接；需要自己收尾。
- ServeConn/ServeCodec 会在内部循环处理一个连接上的多个请求；ServeRequest 只处理一次请求，控制更细

## 顶层 Client 函数

### Dial

```
func Dial(network, address string) (*Client, error)
```

直接 TCP 连接服务器，并返回一个可调用 RPC 方法的 `*Client`。

| 参数        | 示例                 | 说明      |
|-----------|--------------------|---------|
| `network` | `"tcp"`            | 套接字类型   |
| `address` | `"localhost:1234"` | 服务端监听地址 |

**典型用法**

```
client, err := rpc.Dial("tcp", "localhost:1234")
var reply Result
client.Call("Cal.Square", Args{Num: 8}, &reply)
```

> 默认使用 Gob 编解码

### DialHTTP

```
func DialHTTP(network, address string) (*Client, error)
```

用于连接通过 HandleHTTP() 发布的 RPC 服务

**典型用法**

```
client, _ := rpc.DialHTTP("tcp", "localhost:1234")
```

要求服务端使用 `rpc.HandleHTTP()` 进行暴露

### DialHTTPPath

```
func DialHTTPPath(network, address, path string) (*Client, error)
```

与 DialHTTP 类似，但允许自定义 RPC HTTP 路径，需要配合自定义 server 并使用 HandleHTTP() 指定路径进行使用

**典型用法**
客户端：

```
client, _ := rpc.DialHTTPPath("tcp", "localhost:1234", "/myrpc")
```

服务端：

```
rpc.NewServer().HandleHTTP("/myrpc", "/debug")
```