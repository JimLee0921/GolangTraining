# `type ResponseWriter`

ResponseWriter 是一个接口，定义如下：

```
type ResponseWriter interface {
    Header() Header
    Write([]byte) (int, error)
    WriteHeader(statusCode int)
}
```

服务器端响应的核心控制器，ResponseWriter 是在 handler 中手里拿到的东西，所有设置 Header、设置状态码、写 Body 都是通过它完成。

## 功能行为

ResponseWriter 本质上提供了三个功能：

| 方法                  | 作用           | 内容阶段    |
|---------------------|--------------|---------|
| `Header()`          | 获取并设置响应头     | 头部阶段    |
| `WriteHeader(code)` | 发送状态码 + 发送头部 | 头部阶段结束  |
| `Write(data)`       | 向客户端写响应体     | Body 阶段 |

### 隐式调用

如果没有手动调用 WriteHeader，那么第一次 Write 会自动调用 WriteHeader(200)。
也就是说：写 Body = 隐式地锁死 Header。

### 工作顺顺序

1. 先通过 `Header()` 设置头部
2. `WriteHeader(code)` 设置状态码（可选，见隐式调用）
3. 最后 `Write()` 写响应体

### 接口设计理念

因为` http.Server` 内部可以根据协议不同（例如 `HTTP/1.1`、`HTTP/2`、TLS、反向代理情形）替换具体实现

> 永远不需要自己实现 ResponseWriter，只需要使用它

## ResponseWriter 拓展接口

这些接口不是每个响应都会用，但它们代表了在 HTTP 协议层进行更高级的能力。

ResponseWriter 是一个接口，而实际运行时的实现（如 `*http.response`）可能同时实现下列扩展接口：

### 1. `http.Hijacker`

Hijacker 接口由 ResponseWriter 实现它允许 HTTP 处理程序接管连接。

```
type Hijacker interface {
	Hijack() (net.Conn, *bufio.ReadWriter, error)
}
```

也就是劫持底层 TCP 连接，从 HTTP 模式切换为原始套接字通信。

因为有些协议不是 HTTP 文本请求/响应模型，比如：

* WebSocket（升级协议后就不再按 HTTP 格式通信）
* 自定义长连接协议
* 代理服务器（隧道，比如 CONNECT）

> 一旦 Hijack 成功，HTTP 就退出舞台，自己直接 `Read/Write` socket
>
> HTTP/1.x 连接的默认ResponseWriter支持 Hijacker，但 HTTP/2 连接有意地不支持。ResponseWriter 的包装器也可能不支持 Hijacker

### 2. `http.Flusher`

Flusher 接口由 ResponseWriter 实现，允许 HTTP 处理程序将缓冲的数据刷新到客户端。

```
type Flusher interface {
	// Flush sends any buffered data to the client.
	Flush()
}
```

立即把缓冲区的数据发送给客户端，不等待响应结束。

因为 HTTP 为了性能会缓冲输出，而某些场景要求边生成内容边发送：

* 流式响应
* 服务端推送日志（SSE，Server-Sent Events）
* 实时进度输出（例如浏览器看到后台执行进度条）

> 调用 `Flush()` 表示：不要等，立刻把已有数据发出去

### 3. `http.CloseNotifier`（已废弃）

用于检测客户端是否断开连接

因为它和某些新 HTTP / HTTP2 / Proxy 环境无法可靠保障。 现在使用 `Context` 替代：

```
<-r.Context().Done()
```

只需知道它存在过，现在不用它

### 4. `http.Pusher`（HTTP/2）

Pusher 是 ResponseWriter 实现的接口，支持 HTTP/2 服务器推送

```
type Pusher interface {
	Push(target string, opts *PushOptions) error
}
```

在 HTTP/2 中服务器主动推送资源，无需浏览器先请求。

例如客户端请求 `index.html`，服务端提前推：

```
style.css
main.js
logo.png
```

也就是 Client 要 A，Server 说：你肯定也会要 B C D，我提前给你推过来了。
主要为了减少网络往返，提升页面加载速度。

* 只在HTTP/2 协议中有用
* 需要客户端也允许推送

### 总结对比

| 扩展接口                | 功能          | 常见用途              | 是否常用  |
|---------------------|-------------|-------------------|-------|
| `Hijacker`          | 获取底层 TCP 连接 | WebSocket、隧道代理    | 需要时必用 |
| `Flusher`           | 立即发送缓冲数据    | 流式响应 / SSE / 实时输出 | 很有用   |
| `CloseNotifier`（废弃） | 监听客户端断开     | 被 `Context` 替代    | 不用    |
| `Pusher`（HTTP/2）    | 服务端推送资源     | 高性能网页服务           | 专业场景用 |

* Hijacker：我要跳出 HTTP，自己玩 TCP
* Flusher：我现在就要把数据发出去，别缓冲
* CloseNotifier：旧时代监听断开，现在用 Context
* Pusher：HTTP/2 才有，我提前把你一定要的资源推给你
