# time 顶层函数

一共三个：After / Sleep / Tick

| 函数      | 一句话定位                                    |
|---------|------------------------------------------|
| `Sleep` | 让当前 goroutine 停一会儿，相当于 py 的 `time.sleep` |
| `After` | 把一次性延时变成 channel 事件                      |
| `Tick`  | 把“周期时间”变成 channel 事件（简化版 Ticker）         |

## time.Sleep

Sleep 会暂停当前 goroutine 至少 d 的持续时间。d 为负数或零时，Sleep 会立即恢复

```
func Sleep(d Duration)
```

> python 的 time.sleep()

## time.After

After 用于等待 d 时间后，向返回的 channel 发送当前时间，参考 `time.Timer`

```
func After(d Duration) <-chan Time
```

`After(d)` 是个顶层函数，快捷通道，语义上等价于 `time.NewTimer(d).C`，

```
ch := time.After(d)
<-ch // d 时间之后收到 time.Time
```

- 这是个一次性事件
- 只会发送一次
- 发送的是触发那一刻的时间

**版本讲解**

Go1.23 前 After 内部创建的 Timer 在到期之前不会被 GC 回收也就是说：

```
select {
case <-time.After(10 * time.Second):
case <-ctx.Done():
}
```

如果 `ctx.Done()` 先触发：

- `time.After` 创建的 Timer 仍然存在，仍然会在 10 秒后触发
- 在旧版本中 GC 不会提前回收它

所以老版本推进如果关心效率问题不要使用 `After`，而是使用 `NewTimer` 并在不需要时显示调用 `Stop()` 方法

Go1.23 后只要 Timer 没有任何引用，没有 `Stop()` 也会被 GC 自动回收，所以新版本只要 `After` 能满足需求就内必要可以使用
`time.NewTimer`：

```
select {
case <-workDone:
case <-time.After(5 * time.Second):
    return errors.New("timeout")
}
```

适合：

- 一次性超时
- 不需要取消
- 不需要 Reset
- 不需要复用

## time.Tick

`time.Tick` 是 NewTicker 的一个便捷封装，只暴露周期性 tick 的 channel

```
func Tick(d Duration) <-chan Time
```

等价于：

```
ticker := time.NewTicker(d)
return ticker.C
```

- Tick 内部一定创建了一个 Ticker
- 但拿不到 `*Ticker`，只能拿到 `Ticker.C`
- 如果 `d <= 0`，Tick 不会 panic 而是直接返回 nil，这一段与 `NewTicker` 不同（但是很危险）

**版本问题**

和 `After` 与 `NewTimer` 一样，Go1.23 前关心效率应该用 `NewTicker`并在不需要的适合调用 `Ticker.Stop`

Go1.23 后能用 `Tick` 就用 `Tick`