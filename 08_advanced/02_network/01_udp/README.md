使用终端在本目录下运行

```shell
go run server.go
go run client.go
```

### UDP 特点

| 特点                     | 含义                       |
|------------------------|--------------------------|
| 无连接（connectionless）    | 不需要握手/断开，直接发就行           |
| 不可靠（unreliable）        | 丢包、乱序、重复都可能发生            |
| 面向报文（message-oriented） | 每个包是独立的，不像 TCP 那样是连续的数据流 |

- 不需要建立连接
- 报文边界清晰
- 代码结构最直观
- 能立刻看到客户端->服务端->响应的全流程

## Go 创建 UDP

### UDP Server

1. `net.ResolveUDPAddr("udp", ":3000")`
    - 作用：解析一个 UDP 地址（字符串 -> `*net.UDPAddr` 对象）
    - 返回类型：`*net.UDPAddr`，它包含 IP 和端口
    - 参数含义：
        - `udp`：表示协议类型，也可以写 `udp4` 或 udp6`
        - `:3000`：表示监听所有网卡上的 3000 端口O（相当于`0.0.0.0:3000`）

2. `net.ListenUDP("udp", addr)`
    - 作用：启动一个 UDP socket 来监听指定的地址
    - 返回值：*net.UDPConn（UDP 连接对象）
    - 做的事情相当于： `socket()` + `bind()`。UDP 没有“accept” 阶段，也不维护连接状态

3. `conn.ReadFromUDP(buf)`
    - 作用：读取一个完整的 UDP 报文（包）。
    - 返回：
        - n：实际读取的字节数
        - clientAddr：远程客户端的地址（*net.UDPAddr）
        - err：错误（如超时）
    - > UDP 是“面向报文”的，所以每次 ReadFromUDP 就对应一个完整的“包”

4. `conn.WriteToUDP(reply, clientAddr)`
    - 作用：发送一个 UDP 包到指定的客户端地址
    - 参数：
        - reply：要发送的字节切片
        - clientAddr：目标地址
    - 没有连接状态，每发一次都要指定目标

5. `defer conn.Close()`
    - 关闭 socket，即释放操作系统的网络资源

### UDP Client

1. `net.ResolveUDPAddr("udp", "127.0.0.1:3000")`
    - 和服务端一样，用来解析目标地址
    - 区别只是这次绑定的是目标而不是本地绑定

2. `net.DialUDP("udp", nil, addr)`
    - 作用：创建一个 UDP 连接对象，用于发送和接收数据
    - 参数：
        - "udp"：协议
        - nil：本地地址（让系统自动分配一个端口），也可以手动指定端口，但这样多次运行客户端会因为端口占用而报错
        - addr：远程服务器的地址
    - 返回：*net.UDPConn
    - > 虽然名字叫 Dial，但注意：在 UDP 里并不会真正“建立连接”，只是让 Go runtime 知道目标地址是谁，以后可以直接发数据

3. `conn.Write([]byte("hello UDP server"))`
    - 作用：向之前 DialUDP 指定的远程地址发送数据
    - 因为已经假连接，所以不需要每次写都传目标地址
4. `conn.ReadFromUDP(buf)`
    - 作用：接收服务器回发的数据
    - 一次读取一个完整的包
    - 返回 (n, addr, err)，其中 addr 是服务端的地址

### 关系图

```
┌──────────────┐              ┌──────────────┐
│  UDP Client  │              │  UDP Server  │
│              │              │              │
│ Write()──────┼─────────────►│ ReadFromUDP()│
│              │              │              │
│ ReadFromUDP()│◄─────────────┼──────WriteToUDP()
│              │              │              │
└──────────────┘              └──────────────┘
```

| 方法                   | 类型   | 作用                 |
|----------------------|------|--------------------|
| `net.ResolveUDPAddr` | 工具函数 | 把字符串地址转成 `UDPAddr` |
| `net.ListenUDP`      | 服务端  | 绑定端口并开始接收 UDP 包    |
| `net.DialUDP`        | 客户端  | 建立“伪连接”，简化通信       |
| `ReadFromUDP`        | 双方   | 接收一个完整包            |
| `WriteToUDP`         | 双方   | 发送一个完整包            |
| `Close`              | 双方   | 关闭 socket          |

回环地址和本机地址

| 写法                 | 说明               | IP 部分         | 含义              |
|--------------------|------------------|---------------|-----------------|
| `"127.0.0.1:8080"` | **显式绑定本地回环地址**   | 指定为 127.0.0.1 | 只接受来自本机的 UDP 请求 |
| `":3000"`          | **省略 IP，监听所有网卡** | 空（即 0.0.0.0）  | 接受来自任何网卡/IP 的请求 |