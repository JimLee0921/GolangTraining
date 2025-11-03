## `time.After`

延迟触发一次（简化版 Timer）

```
<-time.After(3 * time.Second)
```

等价于

```
timer := time.NewTimer(3 * time.Second)
<-timer.C
```

### 特点

| 特性         | 说明        |
|------------|-----------|
| 创建匿名 Timer | 不需要保存变量   |
| 自动触发并释放资源  | 不用 Stop() |
| 只能执行一次     | 无法 Reset  |
| 无法取消       | 一旦调用就等着执行 |

## `time.Tick`

周期触发（简化版 Ticker）

```
for t := range time.Tick(1 * time.Second) {
    fmt.Println(t)
}
```

等价于

```
ticker := time.NewTicker(1 * time.Second)
for t := range ticker.C {
    fmt.Println(t)
}
```

### 特点

**不能停止**

`time.Tick()` 返回的是 一个只读 channel，没有 `ticker.Stop()` 的机会

也就是说：

- `time.Tick` 是永不停止的 Ticker，会造成资源泄漏
- 真实项目中不推荐用 `time.Tick`
- 推荐用 `time.NewTicker` + `Stop()`

## `time.AfterFunc`

在指定的时间过去后，自动执行一个函数（回调）。
也就是说，它不是像 `time.After` 那样 <-chan 等待，而是 到点自动执行一个函数。

```
func AfterFunc(d Duration , f func()) * Timer
```

AfterFunc 会等待指定的时间过去，然后在它自己的 goroutine 中调用 f。
返回一个Timer 对象，可以使用其 Stop 方法取消调用。
返回的 Timer 对象的 C 字段不会被使用，其值为 nil。

## 对比

| API                 | 本质        | 触发次数 | 能否 Stop() | 是否推荐   |
|---------------------|-----------|------|-----------|--------|
| `time.NewTimer(d)`  | 定时器       | 一次   | 可 Stop    | 推荐     |
| `time.After(d)`     | 匿名 Timer  | 一次   | 不可 Stop   | 用于简单延迟 |
| `time.NewTicker(d)` | 周期器       | 无限   | 可 Stop    | 推荐     |
| `time.Tick(d)`      | 匿名 Ticker | 无限   | 不可 Stop   | 不推荐    |
