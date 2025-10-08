# 管道 Channel

> 具体代码见：[channel管道](../../../02_data-type/06_composite-types/06_channel)
>
> 更多使用见：[goroutines](../../../07_goroutines)

---

## 1 概念速览

* **channel 是一等公民类型**：`chan T`，用于在 goroutine 之间**安全传递数据**并实现同步。
* **无缓冲**：`make(chan T)` —— 发送与接收**必须同时就绪**（强同步点）。
* **有缓冲**：`make(chan T, n)` —— 允许暂存 `n` 个元素（生产/消费解耦）。

---

## 2 声明与类型形式

```go
var ch chan int // 仅声明，零值为 nil（不能直接用）
ch = make(chan int) // 无缓冲
ch2 := make(chan int, 2) // 有缓冲

// 单向化（常用于函数签名约束 API）
var sendOnly chan<- int // 只写
var recvOnly <-chan int // 只读
```

**方向约束规则**：

* 参数里常用：`in <-chan T`（只读输入）、`out chan<- T`（只写输出）
* 返回值里常用：`<-chan T`（只读输出）
* 双向 `chan T` 可以**收窄**为只读/只写再传递或返回；反之不行。

---

## 3 零值与初始化

* `var ch chan T` 的**零值是 `nil`**：对 `nil` channel 的发送/接收会**永久阻塞**。
* 必须用 `make` 初始化后才能使用。

示例：

```go
var ch chan int
// <- ch      // 会阻塞
// ch <- 1    // 会阻塞
ch = make(chan int) // 正确
```

---

## 4 基本操作符

* **发送**：`ch <- v`
* **接收**：`v := <-ch`
* **接收带 ok**：`v, ok := <-ch`（关闭后 `ok=false` 且 `v` 为零值）
* **关闭**：`close(ch)`（表示**不会再有新数据**；已在缓冲的数据仍可被读出）

示例（完整行为）：

```go
ch := make(chan int, 2)
ch <- 10; ch <- 20
fmt.Println(<-ch) // 10
fmt.Println(<-ch) // 20
close(ch)
v, ok := <-ch // 0, false（已关闭且读空）
```

---

## 5 无缓冲 vs 有缓冲（直观对比）

**无缓冲**（同步点）：

```go
c := make(chan int)
go func (){ c <- 1 }() // 发送阻塞到有接收者
fmt.Println(<-c) // 接收时配对，解除阻塞
```

**有缓冲**（解耦）：

```go
c := make(chan int, 2)
c <- 1 // 不阻塞（缓冲未满）
c <- 2
// c <- 3 // 若再发将阻塞，直到有人接收
fmt.Println(<-c, <-c)
```

---

## 6 for-range vs 单次 `<-`

* **单次接收**：`v, ok := <-ch` —— 适合“只取一次/不确定是否关闭”的场景。
* **for-range**：`for v := range ch { ... }` —— 适合“**把通道里所有数据都读完**”；

    * **前提**：发送方最终会 `close(ch)`，否则循环会一直阻塞等待。

---

## 7 关闭语义与“为什么 close 后还能继续读到数据”

* `close(ch)` **不会清空**缓冲：已写入的数据仍可被消费。
* `for range ch` 会在**通道关闭且读空**后**自动退出**。
* 规范：**谁发送谁关闭**（多生产者场景由**统一协调者**关闭，避免重复关闭）。

---

## 8 典型并发模式

### 8.1 Fan-in（N → 1，多生产者汇聚）

```go
func fanIn(n int) <-chan int {
out := make(chan int)
var wg sync.WaitGroup
wg.Add(n)
for i := 0; i < n; i++ {
go func (id int) {
defer wg.Done()
for j := 0; j < 5; j++ { out <- id*100 + j }
}(i)
}
go func (){ wg.Wait(); close(out) }()
return out
}

func main() {
for v := range fanIn(10) {
fmt.Println("consume:", v)
}
}
```

要点：**统一关闭者**在 `wg.Wait()` 后 `close(out)`。

### 8.2 Fan-out（1 → N，任务分发 / Worker Pool）

**无缓冲版本**（强同步；即产即消）：

```go
jobs := make(chan int) // 无缓冲
var wg sync.WaitGroup
numWorkers := 3
wg.Add(numWorkers)
for w := 1; w <= numWorkers; w++ {
go func (id int){
defer wg.Done()
for j := range jobs { /* 处理 j */ }
}(w)
}
go func (){             // 生产者
for j := 1; j <= 10; j++ { jobs <- j }
close(jobs)
}()
wg.Wait()
```

**有缓冲版本**（能堆任务）只需 `make(chan int, N)`。

### 8.3 Pipeline（分阶段加工）

```go
func gen(n int) <-chan int {
out := make(chan int)
go func(){
for i := 0; i < n; i++ { out <- i }
close(out)
}()
return out
}

func sum(in <-chan int) <-chan int {
out := make(chan int)
go func(){
s := 0
for v := range in { s += v }
out <- s; close(out)
}()
return out
}

func main() {
fmt.Println(<-sum(gen(10))) // 0..9 的和：45
}
```

**API 设计要点**：返回 `<-chan T`（只读）、参数用 `<-chan T`/`chan<- T` 限制方向。

---

## 9 WaitGroup vs 信号通道（done/quit/close）

**只需要“等所有 goroutine 结束”** → `sync.WaitGroup`：

```go
var wg sync.WaitGroup
wg.Add(n)
for i := 0; i < n; i++ { go func (){ defer wg.Done(); /* work */}() }
wg.Wait()
```

**需要“取消/超时/多路监听/广播”** → 信号通道或 `context.Context`：

```go
ctx, cancel := context.WithCancel(context.Background())
go func (){
select {
case v := <-in:
case <-ctx.Done(): return
}
}()
cancel() // 广播取消
```

**组合**：常见写法是 **WaitGroup 等收尾**，**close/ctx 负责停机信号**。

---

## 10 常见错误 & 调试

### 10.1 死锁（all goroutines are asleep）

* **单 goroutine** 用**无缓冲**先发后收：

  ```go
  c := make(chan int)
  c <- 1          // 阻塞，无接收者
  fmt.Println(<-c)
  ```

  **修复**：让另一方在 goroutine 中并发运行，或使用缓冲 `make(chan int, 1)`。

* `for range ch` 未关闭就等待：读完后继续等新数据 → **阻塞**。
  **修复**：发送方**最终要 `close(ch)`**；或按条数显式读取。

### 10.2 WaitGroup 竞态

* **不要在与 `Wait()` 并发的情况下调用 `Add()`**。应在启动 goroutine **之前**一次性 `Add(n)`。

### 10.3 多生产者重复关闭

* 只能由**唯一协调者** `close(ch)`；多处关闭会 **panic: close of closed channel**。

### 10.4 闭包捕获循环变量

* **Go 1.21 及更早**：`for` 里直接闭包捕获迭代变量会拿到同一实例，常见输出全是最后一个值。
* **Go 1.22+**：标准 `for range` 每轮**独立副本**，多数情形已安全，但**传参/变量 shadow**仍是好习惯：

```go
for _, v := range []string{"a", "b", "c"} {
v := v
go func (){ fmt.Println(v) }()
}
```

---

## 12 总结

* `chan T`：双向；`chan<- T`：只写；`<-chan T`：只读
* `make(chan T)`：无缓冲；`make(chan T, n)`：有缓冲
* 发送 `ch <- v`；接收 `v := <-ch`；关闭 `close(ch)`；`v, ok := <-ch`
* `for v := range ch`：**需 close** 才会退出
* **无缓冲**=强同步；**有缓冲**=可堆积
* **Fan-in**：`WaitGroup` 等生产者 → 协调者 `close(out)`
* **Fan-out/Worker Pool**：主线程 `close(jobs)`，workers `for range jobs` + `WaitGroup`
* **只等结束**用 `WaitGroup`；**取消/超时/广播**用 `context` 或信号通道
* 常见坑：死锁、重复关闭、`Add` 与 `Wait` 并发、（旧版本）闭包捕获循环变量

