# `time.Duration`

time.Duration 用来表示时间长度，也就是两个时刻之间的时间间隔，
单位位纳秒 int64 纳秒（ns），所以限制了可表示的持续时间最长约为290年。

```
type Duration int64
```

| 时间单位 | 所代表的数值           |
|------|------------------|
| 1 纳秒 | 1                |
| 1 微秒 | 1,000 纳秒         |
| 1 毫秒 | 1,000,000 纳秒     |
| 1 秒  | 1,000,000,000 纳秒 |

## 常见 Duration 变量

Go 中已经定义好了常用时间单位

| 常量                 | 含义 |
|--------------------|----|
| `time.Nanosecond`  | 纳秒 |
| `time.Microsecond` | 微秒 |
| `time.Millisecond` | 毫秒 |
| `time.Second`      | 秒  |
| `time.Minute`      | 分  |
| `time.Hour`        | 小时 |
