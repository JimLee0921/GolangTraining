# time.Timer

`time.Timer` 表示单个时间，用于在未来某个时间点只触发一次一次性动作，也就是一次性定时器

```
type Timer struct {
	C <-chan Time
	// contains filtered or unexported fields
}
```

> `time.Timer` 为一次性，`time.Ticker` 为周期性

## 底层原理

一个 Timer 主要包含三样东西

- 一个到期时间点
- 一个通知机制
- 一次性触发语义

当现在时间 >= 到期时间时有两件事情可以做：

1. 要么往一个 channel 中发送一个时间值
2. 要么调用指定的 `AfterFunc` 函数

> 官方文档给的解释：当 Timer 到期时，当前时间会被发送到其通道 C 上，除非该 Timer 是通过 AfterFunc 创建的。
> Timer 必须通过 NewTimer 或 AfterFunc 创建。

## 创建方式

time 包有两种方式创建 `time.Timer`，分别是 `time.NewTimer` 和 `time.AfterFunc` 函数

### time.NewTimer

`time.NewTimer` 用于创建一个定时器，在 d 持续时间后通过向 channel 发送 Timer 对象进行通知

```
func NewTimer(d Duration) *Timer
```

- Timer 有一个 channel：`t.C`
- 在到时间后会自动向 `t.C` 种发送一次当前时间
- 发送完成就直接结束（不会再触发）

**历史问题**

在 Go1.23 之前，如果一个 Timer 还没有到期（未被触发）且没有被 `Stop()`，即使不再持有它的引用，GC 也不会自动回收，所以老版本写法是”

```
t := time.NewTimer(d)
defer t.Stop()
```

Go1.23 后进行了修复，Timer 变成了真正的普通对象，只要不再持有，即使没到期且没有 `Stop()` 也会被 GC 自动回收

第二个问题是在 Go1.23 之前，Timer 触发时必须有人接收，如果使用了 `Stop`/`Reset` 成功就不能再接收旧值，也就是必须写：

```
if !t.Stop() {
    <-t.C
}
```

Go1.23 后把 `Timer.C` 变为了同步通道，在新语义下这段代码是多余的

### time.AfterFunc

创建一个定时器，在 d 时间之后直接调用 `f()，不需要等待 channel 的快捷版 Timer

```
func AfterFunc(d Duration, f func()) *Timer
```

行为和 `NewTimer()` 一样等待至少 d，和 `NewTimer` 不一样的是在到期后：

- runtime 新创建一个 goroutine 并调用 `f()`
- 不依赖 channel / select / 调用方是否在等待

返回值是 `*time.Timer`，这个 Timer 不是为了收时间值，而是为了在到期前调用 `Stop()` 阻止 `f()` 被执行：

```
t := time.AfterFunc(d, f)
t.Stop()   // 如果还没到期，那么 f() 永远不会被执行
```

返回值实际上：`t.C == nil`，也就是说不能使用 `<-t.C` 操作，不能拿去配合 select

> AfterFunc 只是一个延迟执行器，而不是一个时间信号源

## 核心方法

`time.Timer` 一共两个方法：`Stop` 和 `Reset`

### Stop()

`Stop` 方法用于尝试停止这个 Timer 触发，NewTimer 创建的就是阻止向 `t.C` 发送时间值，对于 AfterFunc 创建的就是阻止调用
`f()`

```
func (t *Timer) Stop() bool
```

**返回值**

- 如果这次调用成功阻止了定时器返回 true
- 如果定时器已经到期/之前已经被停止则返回 false

> Stop 的返回值是一个状态判断信号，而不是结果确认

注意对于通过 `AfterFunc(d, f)` 创建的定时器：

- 如果 `t.Stop()` 返回 false 说明定时器已经到期且函数 f 已经在它自己的 goroutine 中启动
- Stop 不会等待 f 执行完成，而是立刻返回

对于使用 `NewTimer(d)` 创建的、基于 channel 的定时器：

从 Go 1.23 开始，在 Stop 返回之后：

- 对 `t.C` 的任何接收操作保证会阻塞
- 不会再收到 Stop 之前残留的时间值
- Go 1.23 前需要使用 `if !Stop{ <-t.C }`

### Reset()

把一个已有的 Timer 重新设置为从现在开始，再过 d 时间触发

```
func (t *Timer) Reset(d Duration) bool
```

**返回值**

- true：Reset 之前，Timer 依然是活跃的（未被触发）
- false：Reset 之前，Timer 已经被触发或已被 Stop

对于通过 `AfterFunc(d, f)` 创建的定时器，Reset 有两种可能：

1. 重新安排 f 的执行时间
    - 原本还没执行
    - Reset 返回 true
2. 再安排一次 f 的执行
    - 原本已经执行过（或正在执行）
    - Reset 返回 false

> 可能同时会有多个 f 并发执行。`Reset(false) != 失败`，而是再启动一次 f，需要注意 Reset 不会保证之前的 f 执行完成也不会保证新一轮执行
> f 的 goroutine 不会和上一轮执行的 f 并发执行



对于通过 `NewTimer(d)` 创建的、基于 channel 的定时器，从 Go 1.23 开始，在 Reset 返回之后：

- 从 `t.C` 接收 
- 保证不会收到旧设置（Reset 之前）的时间值