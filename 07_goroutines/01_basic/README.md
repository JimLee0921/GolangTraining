# goroutine

## 什么是 goroutine

goroutine 是 Go 的**轻量级线程**：创建成本很低（起始栈 ~2KB，按需增长），由 Go 运行时调度（M:N 调度），
在 Golang 中一个 goroutine 就是一个执行单元，而每个程序都应该有一个主函数main也就是主 goroutine。类似于 python 中的协程等概念，
和线程类似，共享堆，不共享栈，协程的切换一般由程序员在代码中显式控制。避免了上下文切换的额外耗费，兼顾了多线程的优点，简化了高并发程序的复杂。

---

## goroutine 创建

使用 go 关键字进行创建 goroutine，`go f()` 就是把函数 `f` 放到后台并发执行，**不阻塞**当前 goroutine

### 最小创建

```go
go func (x int) {
fmt.Println("run in background:", x)
}(42)
```

> **调度**：G（goroutine）在 P（逻辑处理器）上被 M（OS 线程）执行。`GOMAXPROCS`≈可并行运行的 G 数（默认=CPU核数）。

### 传参与返回值

goroutine 启动时值会被拷贝（按正常传参规则）

```go
go func (x int, s string) {
fmt.Println(x, s)
}(42, "hi")

```

> goroutine 没有直接返回值。若要把结果带回来，需要使用 channel 或 回调

---

### 闭包捕获

这个是老版本存在的问题，在 Go 1.22 后的版本被修复

错误写法（捕获同一循环变量）：

```go
for i := 0; i < 3; i++ {
go func () {           // i 被闭包捕获，三个协程可能都看到同一个最终值
fmt.Println(i)
}()
}
```

正确写法（重新绑定参数或变量）：

```go
for i := 0; i < 3; i++ {
i := i // 为当前迭代重新绑定
go func () {
fmt.Println(i)
}()
}

// 或者显式传参：
for i := 0; i < 3; i++ {
go func (v int) {
fmt.Println(v)
}(i)
}

```

## WaitGroup

`sync.WaitGroup` 用来等待一组 goroutine 结束。
它一个待完成数（Add(n)），每个 goroutine 完成时调用 Done()（等价于 Add(-1)），最后在需要阻塞等待的地方调用 Wait()，直到计数归零再继续

## 通信方式

channel 是 Go 语言中 goroutine 之间通信最通用、最核心的机制，它正是 Go 并发模型的基石之一，
Go 有句经典口号：Don’t communicate by sharing memory; share memory by communicating.
（不要通过共享内存来通信，而要通过通信来共享内存。）

- 不推荐像传统语言那样用锁去保护共享变量
- 而是用 channel 传递数据
- 每个 goroutine 各自独立，通过 channel 交换信息

---

## 取消与超时

任何**可能阻塞**的 goroutine 都应该能被取消，可以使用 context 包进行操作（防泄露的关键）

## 常见并发模式

- **Worker Pool（fan-out）** 扇出
- **Fan-in（多路合并）** 扇入
- **Pipeline（分阶段处理）**

---

## 错误汇总与限并发

---

# 7) Panic 与恢复（谨慎使用）

在 goroutine 边界做兜底，避免整个进程崩掉；但**不要吞错**。

```go
go func () {
defer func () {
if r := recover(); r != nil {
log.Printf("panic: %v\n%s", r, debug.Stack())
}
}()
// do risky work
}()
```

