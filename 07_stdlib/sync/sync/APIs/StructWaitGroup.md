# sync.WaitGroup

`sync.WaitGroup` 是一个计数信号量（并发安全的计数器+阻塞点），用于等待一组 goroutine 全部执行完毕

```
type WaitGroup struct {
	// contains filtered or unexported fields
}
```

## 使用场景

主要适合以下几个典型场景

1. fan-in/fan-out 模型：一个任务拆成 N 个并发子任务，等待所有任务完成后再进行汇总
2. 并发批处理
3. 程序退出前的收尾工作

正常情况下：

```
func main() {
    go doWork()
}
```

在 main goroutine 退出后，整个进程立即退出，所有 goroutine 会被强制终止，main goroutine 不会等待 goroutine

WaitGroup 用来把 main goroutine 和子 goroutines 的生命周期进行绑定，也就是提供了一种计数式同步机制：

- 知道要等待多少个任务
- 每个任务完成后进行记录
- main goroutine 等到所有任务全部完成

## 核心方法

WaitGroup 本质上是一个并发安全的计数器 + 阻塞点，所有方法，都是围绕这个计数器展开的

### 1. Add()

将内部计数器加上 delta（可以为正也可以为负），必须在 goroutine 启动之前调用

```
func (wg *WaitGroup) Add(delta int)
```

- `Add(1)`：增加一个待完成任务
- `Add(-1)`：减少一个待完成任务
- `Add(n)`：一次性增加 n 个

> 如果计数器变成负数或在 `Wait()` 期间并发调用 Add 会触发 panic

### 2. Done()

用于标记一个任务完成，将计数器 -1，等价于 `Add(-1)`

```
func (wg *WaitGroup) Done()
```

**标准用法**

使用 defer 可以防止 return 提前退出或防止 panic 导致计数器不归零

```
wg.Add(1)
go func() {
    defer wg.Done()
    work()
}()
```

> 不要在一个任务中多次调用 Done 或忘记调用 Done

### 3. Wait()

阻塞当前 goroutine 直到内部计数器变为 0 也就是任务全部完成

```
func (wg *WaitGroup) Wait()
```

- 如果一开始就是 0 会立即返回，否则挂起当前 goroutine
- 可以被多个 goroutine 同时调用，所有调用者会被一起唤醒
- WaitGroup 唤醒条件是计数器归零，不是具体哪个 goroutine 完成

> 只要 Add/Done 对齐，不关心 goroutine 和先后顺序

### 4. Go()

Go1.25+ 新推出的语法糖，在一个新的 goroutine 中调用 `f()`，可以减少模板代码，降低忘记调用 `Done` 的概率，本质上是一个安全封装

```
func (wg *WaitGroup) Go(f func())
```

`wg.Go(f)` 等价于：

```
wg.Add(1)

go func() {
    defer wg.Done()
    f()
}
```

> 适合简单并发任务