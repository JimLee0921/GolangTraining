# Timer

Timer = 过一段时间后执行一次

可以把它理解成：

> “X 秒后提醒程序做一件事”

就像设定一个闹钟，只响一次。

## 主要方法和属性

Timer 是 一次性延迟触发事件。使用 `func NewTimer(d Duration) *Timer` 进行创建

```
timer := time.NewTimer(3 * time.Second)
```

### timer.C（通道）

C 是一个 只读 channel，在时间到达时发出一个信号

```
<-timer.C   // 阻塞等待 3 秒后继续执行
```

Timer.C 是 Timer 最核心的部分，依靠 channel 通知时间到达。

### Stop 停止计时

`func (t *Timer) Stop() bool`

```
ok := timer.Stop()
fmt.Println(ok)
```

- 如果 Timer 还没触发，Stop 会成功返回 true
- 如果 Timer 已经触发了，Stop 返回 false

> 实际开发中：在退出 goroutine 时，必须 Stop Timer，避免 goroutine 泄漏
> ```
> timer := time.NewTimer(5 * time.Second)
> go func() {
>    <-timer.C
>    fmt.Println("trigger")
> }()
>    timer.Stop() // 防止 goroutine forever blocked
> ```

### Reset 重新设置时间

`func (t *Timer) Reset(d Duration) bool`

```
timer.Reset(2 * time.Second)
```

- 可以用来 重新开始倒计时
- 常在需要“重新刷新超时时间”时使用（例如连接心跳超时重置）

> 注意：Reset 前应先 Stop（否则可能有竞态问题）

# Ticker

Ticker = 每隔固定时间执行一次

你可以把它理解成：

> “每 X 秒提醒一次”

就像手机上的整点报时，会持续不断地触发

## 主要方法和属性

Ticker 是 周期性触发事件，使用 `func NewTicker(d Duration) *Ticker` 进行创建

### ticker.C（通道）

```
ticker := time.NewTicker(1 * time.Second)
```

每隔 1 秒发送一次时间值，类似一个不断触发的闹钟

### Stop 停止周期触发

`func (t *Ticker) Stop()`

如果不 Stop：

- goroutine 会一直存在
- `ticker.C` 会一直产生事件
- 会导致内存泄漏

因此必须 `defer Stop()`

```
ticker := time.NewTicker(1 * time.Second)
defer ticker.Stop()
```

### Reset

`func (t *Ticker) Reset(d Duration)`

Reset 函数会停止当前计时器，并将其周期重置为指定的持续时间。下一个计时器将在新的持续时间结束后到来。
持续时间 d 必须大于零；否则，Reset 函数会引发 panic。

## 本质区别

| 特性   | Timer | Ticker     |
|------|-------|------------|
| 触发次数 | 只触发一次 | 无限次触发      |
| 功能   | 延迟执行  | 周期执行       |
| 常用场景 | 超时、延迟 | 轮询、心跳、定时任务 |

## 使用场景

| 场景                | 用 Timer 还是 Ticker | 原因      |
|-------------------|-------------------|---------|
| 数据库查询超时 5 秒自动取消   | Timer             | 只需要一次取消 |
| 网页爬虫每 1 秒请求一次     | Ticker            | 周期执行    |
| API 调用超过 3 秒就停止等待 | Timer + context   | 一次性超时   |
| 微服务之间发送心跳包        | Ticker            | 持续发送    |
