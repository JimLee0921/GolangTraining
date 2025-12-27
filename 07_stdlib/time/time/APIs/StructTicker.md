# time.Ticker

`time.Ticker` 是按照固定间隔，重复周期性地触发事件直到遇到 `Stop`

```
type Ticker struct {
	C <-chan Time // The channel on which the ticks are delivered.
	// contains filtered or unexported fields
}
```

一个 `time.Ticker` 也包含三部分：

1. 一个固定的 Duration 周期
2. 一个 channel：`Ticker.C`
3. 一个无限循环触发机制

每到一个周期：当前时间被送入 C，如果没有人接收则只保留最近一次

## 创建方式

使用 `time.NewTicker` 进行创建

### time.NewTicker

用于创建一个每隔 d 时间就触发一次的 Ticker

```
func NewTicker(d Duration) *Ticker
```

每次 tick runtime 都会向 `t.C` 发送一个 `time.Time`

**注意事项**
如果 Ticker 的 d 设为 1 秒但是每个任务需要 5 秒才能处理一次，那么 Ticker 原则上尽量跟上真实时间，而不是忠实回放历史，因此可能会丢弃
tick，跳过中间的时间点，只给最新的那个

在 Go 1.23 之前，只要 Ticker 还在运行且没调用 Stop，即使已经没有任何引用 GC 也不会回收它，必须这么写：

```
ticker := time.NewTicker(time.Second)
defer ticker.Stop
```

Go1.23 后进行了修复，Ticker 成为普通对象，Stop 不是必须调用的，但仍然可以在需要的逻辑上使用 Stop 停止 tick 实例

## 核心方法

和 Timer 一样，两个方法，Stop 和 Reset

### Stop()

Stop 用于关闭（停止）Ticker，在 Stop 之后，不会再有任何 tick 被发送

```
func (t *Ticker) Stop()
```

Stop 不会关闭 ticker 关联的 channel，目的是防止并发读取该 channel 的 goroutine 误以为收到了一个错误的 tick

### Reset()

Reset 会先停止当前的 Ticker，然后把它的周期重设为新的 duration

```
func (t *Ticker) Reset(d Duration)
```

> d 必须大于 0，否则 Reset 会 panic

下一次 tick 一定是在新的周期 d 过去之后才到来：

- Reset 不会立刻触发
- Reset 不会沿用旧的剩余时间
- Reset 的时间起点是调用 Reset 的这一刻