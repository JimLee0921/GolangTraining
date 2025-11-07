# RPC

远程过程调用（英语：Remote Procedure Call，缩写为 RPC）是一个计算机通信协议。
该协议允许运行于一台计算机的程序调用另一个地址空间（通常为一个开放网络的一台计算机）的子程序，而程序员就像调用本地程序一样，无需额外地为这个交互作用编程（无需关注细节）。
RPC是一种服务器-客户端（Client/Server）模式，经典实现是一个通过发送请求-接受回应进行信息交互的系统。

> 允许像调用本地函数一样调用另一台机器上的函数

# net/rpc

net/rpc 是 Go 标准库中提供的 RPC（远程过程调用）框架。

> 文档地址：https://pkg.go.dev/net/rpc

## 具体作用

net/rpc 负责：

- 将 函数调用 -> 序列化成数据发送到网络
- 将 响应数据 -> 反序列化回来变成本地函数返回值
- 自动处理 TCP/HTTP 传输 和 连接管理
- 只需要关心调用函数，不用自己写 socket、协议、序列化等底层通信逻辑

## 使用约束

在 Go 中要让函数能被 RPC 调用，函数必须符合以下规则：

```
func (t *Type) MethodName(arg ArgType, reply *ReplyType) error
```

1. 方法类型（T）是导出的（首字母大写）
2. 方法名（MethodName）是导出的
3. 方法有2个参数(`argType T1`, `replyType *T2`)，均为导出/内置类型
4. 方法的第2个参数一个指针(`replyType *T2`)
5. 方法的返回值类型是 error

net/rpc 对参数个数的限制比较严格，仅能有2个，第一个参数是调用者提供的请求参数，第二个参数是返回给调用者的响应参数。
服务端需要将计算结果写在第二个参数中。如果调用过程中发生错误，会返回 error 给调用者。

## 前置条件

在 Go 标准库 net/rpc 中，一个方法要能被 RPC 调用，必须满足固定的签名和类型要求。

| 角色        | 例子                                                     | 为什么需要                    |
|-----------|--------------------------------------------------------|--------------------------|
| 服务类型（接收者） | `type Cal struct{}`                                    | 用作 RPC 命名空间，RPC 通过它找到方法  |
| 参数类型      | `type Args struct{ Num int }`                          | 用来装客户端传给服务端的输入数据（必须可序列化） |
| 返回值类型（指针） | `type Result struct{ Num, Ans int }`                   | 服务端通过它把结果写回客户端（必须指针）     |
| 方法        | `func (c *Cal) Square(args Args, reply *Result) error` | 必须精确符合 RPC 方法签名          |

### 服务类型

因为 RPC 是通过服务名 + 方法名调用的：`client.Call("Cal.Square", ...)`。
这里的 Cal 也就是 `type Cal struct{}` 的类型名。

如果写成：`type Cal int`，它也能当服务（没问题），但是：

- 必须 register 时 用指针：`server.Register(new(Cal))`，也就是只能使用 `new(Cal)` 进行传入
- 因为只有 `(*Cal)` 方法集符合 RPC 方法签名要求

所以一般写：`type Cal struct{}`，这样注册的使用可以使用传入 `new(Cal)` 也可以传入 `&Cal{}`

### 参数类型

> 这个不是必须的，但是建议进行创建使用

RPC 需要 明确的输入参数结构，例如：

```
type Args struct {
    Num int
}
```

- 字段必须首字母大写（可导出），否则 Gob/JSON 无法序列化
- 客户端与服务端必须用 相同类型定义（或者至少包路径 + 字段一致）

### 返回值

返回值 Result 必须是结构体且是指针，因为 RPC 要 把结果写进去。

```
func (cal *Cal) Square(args Args, reply *Result) error {
    reply.Num = args.Num
    reply.Ans = args.Num * args.Num
    return nil
}
```

客户端调用时也是传入指针：

```
var reply Result
client.Call("Cal.Square", Args{Num: 9}, &reply)
//                           ^^^^^ 这里必须传 &reply
```

> 如果传值，服务端改不到对象 -> RPC 会直接报错
