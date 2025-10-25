## TCP 连接

TCP 通信基本分为两部分

| 角色         | 主要职责                                       |
|------------|--------------------------------------------|
| **Server** | 开一个监听端口，等客户端来连（Listen -> Accept -> Handle） |
| **Client** | 主动发起连接，请求通信（Dial -> Read/Write）            |

### Server 端

1. `net.Listen("tcp", ":4000")`
    - 在本机启动一个 TCP 监听器，等待客户端连接。
    - 参数：
        - "tcp"：指定协议类型，还可以是 "udp"、"unix" 等
        - ":4000"：监听所有网卡上的 4000 端口，如果只想监听本地，可以写 `127.0.0.1:4000`
    - 返回值：`listener net.Listener`，是一个接口
2. `listener.Accept()`
    - 作用：阻塞等待客户端连接请求，一旦有客户端连上，就返回一个 net.Conn 对象
    - 特点：
        - 是 阻塞调用：没有客户端连接时会一直卡在这里
        - 每个客户端连上后都要交给一个独立的 goroutine 去处理，否则会堵塞后续连接

3. `conn.Read() / bufio.NewReader(conn).ReadString('\n')`
    - 从客户端读取数据
    - TCP 是流式协议，没有消息边界，所以用 `bufio.NewReader().ReadString('\n')` 按行读取是一个好办法

4. `conn.Write([]byte("reply"))`
    - 作用：向客户端发送数据（响应），是双向通信的核心之一。
    - Write 是 阻塞 的，TCP 没有分包机制，写多少字节对方就收到多少，但可能分多次传输，所以双方最好约定数据格式（比如以 `\n`
      结尾）
5. conn.Close()
    - 关闭与客户端的连接，释放资源
    - 常配合 defer 使用，保证函数结束后一定关闭

### Client 端

1. `net.Dial("tcp", "127.0.0.1:4000")`
    - 作用：主动连接服务器
    - 第一个参数 "tcp" 表示使用 TCP 协议
    - 第二个参数 "127.0.0.1:4000" 是服务器地址和端口
    - 成功后返回 conn net.Conn 对象，之后可以直接 `Read / Write`

2. `conn.Write([]byte("hello\n"))`
    - 向服务器发送消息
    - 因为 TCP 是字节流，服务端可能一次收不完，所以常用 `\n` 或固定长度来划分消息边界

3. `conn.Read(buf)`
    - 从服务器读取响应
    - Read() 同样是阻塞的，如果服务器没有发数据，它会一直等
    - 返回 n 是实际读取的字节数，当 n == 0 且 err == io.EOF 时表示连接关闭

4. `defer conn.Close()`
    - 客户端退出时一定要关闭连接，否则服务器端的 goroutine 还会卡着

## TCP 通信生命周期

| 阶段   | Server             | Client         | 说明       |
|------|--------------------|----------------|----------|
| 建立连接 | `Listen -> Accept` | `Dial`         | TCP 三次握手 |
| 发送数据 | `conn.Write()`     | `conn.Write()` | 双方都可发    |
| 接收数据 | `conn.Read()`      | `conn.Read()`  | 双向流式通信   |
| 断开连接 | `conn.Close()`     | `conn.Close()` | TCP 四次挥手 |

## TCP 粘包拆包

非常棒的问题 👏！
“**粘包 / 拆包（Packet sticking / splitting）**” 不是 Go 独有的现象，
而是 **所有基于 TCP 协议的语言都会遇到的网络层现象**。

比如用 Python、Java、C++、Node.js 都会碰到。
只要你用的是 TCP（而不是 UDP），它就有可能出现。

---

### TCP 原理

TCP 是“字节流”，不是“消息流”，这是理解粘包的关键。

- UDP：你发一包，我收一包。每个 `send()` 对应一个 `recv()` -> 每个包边界是固定的，不会粘在一起

- TCP：你发的是一串连续的字节流，对方只看到一堆字节 -> 它不关心你发了几次，只管把字节拼在一起按顺序送达

举个比喻👇

> UDP：一瓶一瓶发果汁（每瓶独立）
> TCP：把果汁倒进一根管子，对方那边接到的只是连续流。
> 管道没告诉这瓶果汁到哪结束。

---

### 粘包（sticking）

> 多条消息在接收端“粘”成一条。

例如客户端连续发：

```
hello
world
```

但服务端只读了一次，就可能得到：

```
helloworld
```

因为 TCP 认为它们是连续的字节流，没有分界。

---

### 拆包（splitting）

> 一条消息被拆成了多次接收。

例如客户端发：

```
Hello world, this is a long message...
```

但服务端第一次只收到：

```
Hello world, this i
```

第二次又收到剩下的：

```
s a long message...
```

因为 TCP 会根据缓冲区、窗口大小、网络拥塞情况自动分割传输。

---

### 重点

**这些现象在 TCP 协议层是正常行为，不是 bug**

TCP 的职责是：

> 保证字节顺序正确、全部送达，
> 但它不保证发的消息边界。

---

TCP 是“流式”协议，数据是被底层按块缓存的。

假设客户端代码：

```
conn.Write([]byte("msg1"))
conn.Write([]byte("msg2"))
```

服务器端：

```
buf := make([]byte, 1024)
n, _ := conn.Read(buf)
fmt.Println(string(buf[:n]))
```

> 可能打印：
>
> ```
> msg1msg2
> ```
>
> （粘包）

或者：

> ```
> msg
> ```
>
> （拆包）

取决于：

* 操作系统的 TCP 缓冲区
* 网络延迟
* 接收端读取时机
* 数据包大小

---

### 解决

需要在应用层自己定义消息边界。
这叫做自定义协议封包 / 解包（packet framing）。

常见的几种方案👇

| 方案                | 示例                    | 优缺点                  |             |
|-------------------|-----------------------|----------------------|-------------|
| **1. 特殊分隔符**      | 用 `\n`、`              | `、`\r\n` 分隔消息        | 简单好用，适合文本消息 |
| **2. 固定长度消息**     | 每条消息固定 128 字节         | 简单但浪费带宽              |             |
| **3. 消息头 + 长度字段** | `[4字节长度][消息体]`        | 通用、高效，是最常用方案（如 gRPC） |             |
| **4. 使用序列化协议**    | JSON、Protobuf 自带结构化解析 | 实际开发中最常见             |             |

---

### 举个简单例子：长度前缀方案

#### 客户端发送

```
msg := []byte("hello world")
length := len(msg)

header := make([]byte, 4)
binary.BigEndian.PutUint32(header, uint32(length))

conn.Write(header) // 先发长度
conn.Write(msg) // 再发内容
```

#### 服务器接收

```
// 先读 4 字节长度
header := make([]byte, 4)
conn.Read(header)
msgLen := binary.BigEndian.Uint32(header)

// 再读真正的消息体
msg := make([]byte, msgLen)
conn.Read(msg)
fmt.Println("收到:", string(msg))
```

这样就不会粘包或拆包了：每次都能准确知道一条消息的边界。

