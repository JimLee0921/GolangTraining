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

### 定义使用

实际开发中最常用使用 sync.WaitGroup 来等待 goroutines 结束

```go
var wg sync.WaitGroup

wg.Add(1) // 期待 1 个 goroutine 完成
go func () {
defer wg.Done() // 完成时减 1
fmt.Println("work")
}()

wg.Wait() // 等待计数归零
fmt.Println("all done")
```

### 多个 goroutine

```go
var wg sync.WaitGroup
urls := []string{"a", "b", "c"}

wg.Add(len(urls))           // 先一次性 Add
for _, u := range urls {
u := u // (Go<1.22) 防闭包捕获；1.22+ 可省略
go func () {
defer wg.Done()
// fetch(u) ...
}()
}

wg.Wait()
```

- 先 Add，再启动 goroutine（或在启动前立刻 Add(1)）
- 在 goroutine 里 defer wg.Done()，写错忘记调 Done() 会导致永远堵塞

---

## 通信方式

**用通信共享内存**，少用“共享内存做通信”。

```go
ch := make(chan int) // 无缓冲：发送与接收必须配对同时发生
go func () { ch <- 10 }() // 发送
fmt.Println(<-ch)                 // 接收

buf := make(chan int, 2)          // 有缓冲：允许先存入 N 个元素
buf <- 1; buf <- 2
select {                          // 多路复用
case v := <-buf:
fmt.Println(v)
default:
fmt.Println("no data")
}
```

---

# 4) 取消与超时：context（防泄露的关键）

任何**可能阻塞**的 goroutine 都应该能被取消。

```go
ctx, cancel := context.WithTimeout(context.Background(), 500*time.Millisecond)
defer cancel()

go func (ctx context.Context) {
select {
case <-time.After(time.Second):
fmt.Println("done")
case <-ctx.Done(): // 超时/取消
return
}
}(ctx)
```

在你的循环/读写处都加：

```go
select {
case <-ctx.Done(): return
case v := <-in: /* ... */
}
```

---

## 常见并发模式

- **Worker Pool（fan-out）** 扇出
- **Fan-in（多路合并）** 扇入
- **Pipeline（分阶段处理）**

---

# 6) 错误汇总与限并发：`errgroup`

更“生产化”的并发执行：任何一个失败就取消其它任务。

```go
g, ctx := errgroup.WithContext(context.Background())
g.SetLimit(8) // 最大并发
for _, t := range tasks {
t := t
g.Go(func () error {
select { case <-ctx.Done(): return ctx.Err()
default: return do(t) }
})
}
if err := g.Wait(); err != nil { /* 处理错误 */ }
```

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

---

# 8) 性能与资源管理

* **别无限制创建 goroutine**：用 worker 池/限流（`SetLimit`、信号量 chan）。
* **I/O 密集**：并发数可远大于核数；**CPU 密集**：并发≈`GOMAXPROCS()`。
* **避免忙等**：不要在循环里空转 `default`；使用阻塞 I/O + `select`。
* **缓冲大小**：从 64/128 起压测，不要过大（占内存/隐藏延迟）。

---



---

# 10) 常见坑（避坑清单）

* **main 退出太快**：记得 `WaitGroup` 或消费通道直到关闭。
* **通道关闭方错误**：**谁生产谁关闭**；消费者不要关输入通道。
* **泄露**：阻塞读写没有 `ctx.Done()` 分支；外层取消后子协程仍挂着。
* **滥用 `default`**：为了“非阻塞”导致忙等或丢数据。
* **在循环里捕获变量**（闭包坑）：

  ```go
  for i := 0; i < 3; i++ {
      i := i // 重新绑定
      go func(){ fmt.Println(i) }()
  }
  ```

---

## 最小完整示例（可运行）

```go
package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 600*time.Millisecond)
	defer cancel()

	in := make(chan int, 100)
	go func() {
		defer close(in)
		for i := 0; i < 20; i++ {
			select {
			case <-ctx.Done():
				return
			case in <- i:
			}
		}
	}()

	out := startWorkers(ctx, 3, in) // fan-out 例子
	for v := range out {
		fmt.Println("result:", v)
	}
}

func startWorkers(ctx context.Context, n int, in <-chan int) <-chan int {
	out := make(chan int, 128)
	var wg sync.WaitGroup
	wg.Add(n)
	for i := 0; i < n; i++ {
		go func() {
			defer wg.Done()
			for {
				select {
				case <-ctx.Done():
					return
				case v, ok := <-in:
					if !ok {
						return
					}
					time.Sleep(100 * time.Millisecond)
					select {
					case out <- v * v:
					case <-ctx.Done():
						return
					}
				}
			}
		}()
	}
	go func() { wg.Wait(); close(out) }()
	return out
}
```

---

想继续的话，我可以按你“循序渐进”的节奏，接着讲 **调度器工作原理（G/P/M）**、**抢占/系统调用对调度的影响**、以及*
*如何用 `pprof/trace` 定位 goroutine 泄露与阻塞点**。
