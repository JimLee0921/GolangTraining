# Go Context 笔记

`context` 用于在多个 goroutine 之间传递 **取消信号**、**超时控制** 和 **请求相关的上下文信息**。
它是 Go 并发编程中的关键组件，能够有效地管理不同任务之间的协作和资源释放。

---

## 📌 Context 的接口定义

`context.Context` 是一个接口，定义了四个方法：

* `Deadline()`：返回 context 被取消的时间（若设置了超时或截止时间）。
* `Done()`：返回一个 channel，接收到信号时表示 context 已被取消。
* `Err()`：返回取消的原因（`context.Canceled` 或 `context.DeadlineExceeded`）。
* `Value(key)`：在调用链中传递请求范围内的元数据。

---

## 📌 核心功能

1. **取消（cancel）**

    * 上游调用 `cancel()`，所有持有该 context 的 goroutine 都会收到 `<-ctx.Done()` 信号，并自行退出。

2. **超时（timeout）**

    * 自动在指定时间后触发取消信号。

3. **截止时间（deadline）**

    * 和超时类似，但使用的是一个绝对时间点。

4. **传值（value）**

    * 可以在请求范围内传递少量公共数据（如 request id、trace id）。
    * **注意**：不建议用来传业务参数或大对象。

---

## 创建 Context 的方式

### 1. 根 Context

* 所有 context 都应从 `context.Background()` 开始。
* `context.Background()`：根节点，常用于 `main`、初始化、顶层，**永远不会被取消**。
* `context.TODO()`：当还不确定用哪种 context 时，先用它占位。

```go
ctx := context.Background()
ctx := context.TODO()
```

---

### 2. 可取消的 Context

使用 `context.WithCancel(parent)`：

* 返回一个新的子 `Context` 和一个 `cancel` 函数。
* 调用 `cancel()` 后，所有持有该 context 的 goroutine 都会收到 `<-ctx.Done()` 信号。

```go
ctx, cancel := context.WithCancel(context.Background())
defer cancel()
```

---

### 3. 带超时 / 截止时间的 Context

* `context.WithTimeout(parent, duration)`
  自动在一段时间后触发取消（相当于 `WithDeadline(parent, time.Now().Add(duration))`）。

* `context.WithDeadline(parent, deadline)`
  在指定的绝对时间点自动取消。

```go
// 使用 Deadline
ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(5*time.Second))
defer cancel()

// 使用 Timeout
ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
defer cancel()
```

---

### 4. 带值的 Context

* 使用 `context.WithValue(parent, key, value)`。
* 用于在调用链上传递少量、与请求相关的元信息（如 requestID、用户ID）。
* **注意**：不推荐放大对象或业务数据。

```go
ctx := context.WithValue(context.Background(), "requestID", "12345")
```

---

### 总结

* 起点：`Background()` / `TODO()`
* 取消：`WithCancel()`
* 超时：`WithTimeout()` 相对时间 / 截止时间：`WithDeadline()` 绝对时间
* 传值：`WithValue()`（仅元信息）

---

## context 取值

context.WithValue 可以多次调用，形成一棵嵌套树，每次都加一个键值对

```go
package main

import (
	"context"
	"fmt"
)

type ctxKey string

func main() {
	ctx := context.Background()

	// 依次添加多个值
	ctx = context.WithValue(ctx, ctxKey("requestID"), "abc-123")
	ctx = context.WithValue(ctx, ctxKey("userID"), 42)
	ctx = context.WithValue(ctx, ctxKey("role"), "admin")

	// 读取不同的值
	fmt.Println("requestID:", ctx.Value(ctxKey("requestID")))
	fmt.Println("userID:", ctx.Value(ctxKey("userID")))
	fmt.Println("role:", ctx.Value(ctxKey("role")))
}
```

- 每次调用 WithValue 都会返回一个新的 Context，里面多存了一个 key -> value
- 读取时会先从当前 Context 查找，如果没有就往父 Context 查找，直到 Background()

---

### Value方法

取值时通过Value(key interface{})方法，在需要访问context值的位置，调用Value(key)方法，并传入相应的键返回的是一个interface{}类型，需要进行类型断言才能使用

```go
requestID, ok := ctx.Value("requestID").(string)
if ok {
fmt.Printf("Request ID: %s\n", requestID)
}
```

---

### 注意事项

1. 可以存多个 key-value，但每次都要调用一次 WithValue
2. Key 要唯一：推荐用自定义类型 type ctxKey string，避免不同包用同样的字符串冲突
3. 不要滥用： 只适合存少量、跟请求上下文强相关的“元信息”（traceID、用户ID），不要用来传大对象、业务数据或必填参数

---

## 进阶使用

---

### 链式调用

context可以形成链条结构，每个子context继承自父context，并添加额外的值或取消操作
本质是从一个父 context 派生出子 context，取消/超时会从上游向下游级联传播

```go
parent := context.Background() // 根
ctx1, cancel1 := context.WithCancel(parent) // 可手动取消
ctx2, cancel2 := context.WithTimeout(ctx1, 2*time.Second) // 带超时（隐含可取消）
ctx3 := context.WithValue(ctx2, userKey{}, 42) // 传值（不影响取消）
```

- 取消传播：cancel1() 会依次取消 ctx1、ctx2、ctx3
- 超时传播：ctx2 到时，会取消 ctx2 和其子孙（ctx3），但不会反向取消 ctx1
- 取值链：ctx3.Value(k) 沿父链向上查找，最近的键优先

> 规范：在创建处 defer cancelX()（即使暂时用不到），避免资源泄露

---

在Context链中，子context会继承父context的值，同时也可以有自己的值和取消操作

### 多个Context选择

在多个context同时存在时，通常需要使用select语句来处理多个Done()信号

```go
select {
case <-ctx1.Done():
handleCancel(ctx1)
case <-ctx2.Done():
handleCancel(ctx2)
default:
// 进行其他操作
}
```

---

## 使用规范

1. 避免作为结构体的字段：context不应该作为结构体的字段，而是应该通过函数参数传递
2. 不应长时间持有context：context是用于短期的取消和超时控制，不应长时间持有，特别是在函数之间传递
3. 避免将context存储在全局变量中：全局变量会导致context的生命周期难以控制，增加资源泄漏的风险
4. 使用context管理资源：利用context的Done()信号，释放不再需要的资源，如文件句柄、网络连接等
