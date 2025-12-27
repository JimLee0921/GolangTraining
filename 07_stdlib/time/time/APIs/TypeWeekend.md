# time.Weekend

`time.Weekend` 和 `time.Month` 一样底层是 int，属于周的枚举类型，用于表示一周中的某一天，主要用于

- 日历逻辑
- 周统计
- 工作日/周末判断

```
type Weekday int
```

## Weekend 枚举值

```
const (
    Sunday Weekday = iota
    Monday
    Tuesday
    Wednesday
    Thursday
    Friday
    Saturday
)
```

这里和 Month 不一样的是 Weekend 是从 0 开始的 0-6，而 Month 是从 1 开始的 1-12

| Weekday   | 数值 |
|-----------|----|
| Sunday    | 0  |
| Monday    | 1  |
| Tuesday   | 2  |
| Wednesday | 3  |
| Thursday  | 4  |
| Friday    | 5  |
| Saturday  | 6  |

因为 Go 遵循的是：

- Unix / POSIX / IANA 日历传统
- 与 C / libc / tzdata 一致
- 与 tm_wday 对齐
- 更容易和系统时间交互，避免底层换算成本

## String 方法

Weekend 也只有一个 String 方法用于返回星期的英文名称

```
func (d time.Weekday) String() string
```