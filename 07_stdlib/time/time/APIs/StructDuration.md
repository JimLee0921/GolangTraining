# `time.Duration`

time.Duration 用来表示时间长度，也就是两个时刻之间的时间间隔，单位位纳秒 int64 纳秒（ns），所以限制了可表示的持续时间最长约为290年

> 可以为负数

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

```
const (
	Nanosecond  Duration = 1                    // 纳秒
	Microsecond          = 1000 * Nanosecond    // 微秒
	Millisecond          = 1000 * Microsecond   // 毫秒
	Second               = 1000 * Millisecond   // 秒
	Minute               = 60 * Second          // 分中
	Hour                 = 60 * Minute          // 小时
)
```

# 常用方法/函数

## 构造来源函数

- `time.ParseDuration()`：把字符串解析为 Duration
- `time.Since()`：等价于 `time.Now().Sub(t)`，从 t 到现在过去了多久
- `time.Until()`：等价于 `t.Sub(time.Now)`，距离 t 还有多久

### time.ParseDuration

将指定字符串解析为 Duration 对象，主要用于：

- 配置文件
- 环境变量
- 命令行参数
- 动态规则

```
func ParseDuration(s string) (Duration, error)
```

**注意事项**

参数 s 必须符合规定的格式：

- 支持多段组合
- 支持小数
- 支持负号

> `"300ms" / "1.5h" / "2h45m" / "-10s"` 这些是合法的，`"1hour" / "10 sec"` 是非法的

### time.Since

获取从时间点到限制的时间间隔 Duration

```
func Since(t Time) Duration
```

- 语义上等价于：`time.Since(t) == time.Now().Sub(t)`
- 主要用于统计耗时：`start := time.Now()` + `do something...` + `elapsed := time.Since(start)`

### time.Until

计算从现在到目标时间点的时间间隔 Duration

```
func Until(t Time) Duration
```

- 语义上等价于：`time.Until(t) == t.Sub(time.Now())`
- 主要用于超时剩余时间计算与调度判断：`deadline := time.Now().Add(5 * time.Second)` + `remaining := time.Until(deadline)`

## 数值读取

- 浮点输出
    - `Hours()`
    - `Minutes()`
    - `Seconds()`

- 整数输出
    - `Nanoseconds()`
    - `Microseconds()`
    - `Milliseconds()`

### Hours() / Minutes() / Seconds()

- Minutes 返回持续时间以分钟为单位的浮点数表示
- Hours 返回持续时间以小时为单位的浮点数表示
- Seconds 返回持续时间以秒为单位的浮点数表示

```
func (d Duration) Hours() float64
func (d Duration) Minutes() float64
func (d Duration) Seconds() float64
```

**注意事项**

1. 负数不做截断、取整，保留符号与比例
2. 都是 float64 浮点数，可能存在浮点误差，不适合做精确边界判断

### Milliseconds() / Microseconds() / Nanoseconds()

- Milliseconds：返回持续时间以毫秒为单位的整数计数
- Microseconds：返回持续时间以微秒为单位的整数计数表示
- Nanoseconds：返回持续时间以纳秒为单位的整数计数表示

```
func (d Duration) Milliseconds() int64
func (d Duration) Microseconds() int64
func (d Duration) Nanoseconds() int64
```

**注意事项**

这些都是把 Duration 转为指定单位后直接向 0 进行截断（直接丢弃小数部份）

## 数值计算

- `Abs()`
- `Round()`
- `Truncate()`

### Ads()

计算 d 的绝对值，Duration 是由符号的，很多场景只关心大小，不关心方向

```
func (d Duration) Abs() Duration
```

### Round()

返回将 d 四舍五入到最接近的 m 的倍数的结果

```
func (d Duration) Round(m Duration) Duration
```

- 如果 `m<=0`，Round 返回的 d 值不变
- 如果结果超过了 Duration 类型可以存储的最大值或最小值就返回 Duration 的极值也就是最大/最小持续时间

### Truncate

返回将 d 向零取整到 m 的倍数的结果。如果 `m <= 0`，则 Truncate 函数返回 d 不变

```
func (d Duration) Truncate(m Duration) Duration
```

## 字符串表示

- `String()`

### String()

把 Duration 转成人类可读的规范字符串

```
func (d Duration) String() string
```

- 自动选择最合适的单位组合
- 自动处理符号
- 小数只在必要时出现
- 输出格式可被 ParseDuration 再解析：`d, _ := time.ParseDuration(d.String())`