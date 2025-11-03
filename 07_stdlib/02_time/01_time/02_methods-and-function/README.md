# `time.Time`

time.Time 包中和时间等有关的方法和函数

## 基础方法

1. `func (t Time) Year() int`：公历年份
2. `func (t Time) Month() Month`：公历月份（time.Month 枚举）
3. `func (t Time) Day() int`：当月第几日（1–31）
4. `func (t Time) YearDay() int`：当年第几日（1–366）
5. `func (t Time) Weekday() Weekday`：星期几（time.Weekday，周日=0）
6. `func (t Time) ISOWeek() (year int, week int)`：ISO 8601 获取年和周编号（周一为一周起点）
7. `func (t Time) Hour() int`：小时（0–23）
8. `func (t Time) Minute() int`：分钟（0–59）
9. `func (t Time) Second() int`：秒（0–59）
10. `func (t Time) Nanosecond() int`：纳秒（0–999,999,999）
11. `func (t Time) Date() (year int, month Month, day int)`：一次返回年/月/日
12. `func (t Time) Clock() (hour, min, sec int)`：一次返回时/分/秒

## 时区有关方法

### 时区和 Location 获取

| 方法签名                                             | 含义                            |
|--------------------------------------------------|-------------------------------|
| `func (t Time) Location() *Location`             | 返回该 `Time` 所属的时区（Location 对象） |
| `func (t Time) Zone() (name string, offset int)` | 返回时区名称和相对 UTC 的秒级偏移量          |

### 时区转换

| 方法签名                                   | 含义                              |
|----------------------------------------|---------------------------------|
| `func (t Time) In(loc *Location) Time` | 将时间转换到指定 `Location`，返回新 `Time`。 |
| `func (t Time) UTC() Time`             | 将时间转换为 UTC 时区，返回新 `Time`。       |
| `func (t Time) Local() Time`           | 将时间转换为本地系统时区，返回新 `Time`。        |

## 格式化 / 解析相关

> layout 格式见 create-time 章节，两个 Parse 不属于 t 的方法，而是 time 模块工具方法

Format 是输出相关，Parse 是构造相关

| 方法签名                                                                      | 含义                                                                |
|---------------------------------------------------------------------------|-------------------------------------------------------------------|
| `func (t Time) Format(layout string) string`                              | 使用指定 layout 返回格式化后的字符串。Go 的 layout 使用固定模板 `"2006-01-02 15:04:05"` |
| `func (t Time) AppendFormat(b []byte, layout string) []byte`              | 将按 layout 格式化后的时间追加到切片 `b` 后面，返回新切片（常用于高性能字符串构建）                  |
| `func Parse(layout, value string) (Time, error)`                          | 按 layout 解析字符串，默认按 **UTC** 解释                                     |
| `func ParseInLocation(layout, value string, loc *Location) (Time, error)` | 在指定时区下解析字符串（不会默认转 UTC）                                            |

## timestamp 时间戳有关方法

Unix 时间戳与时区无关，只是一个绝对时间点。
`time.Unix()` 系列构造出的结果 默认在 UTC。
想以某个时区显示必须配合 `t.In(loc)` 使用。

**时间戳精度关系**

| 方法 / 单位       | 表示精度 | 典型用途         |
|---------------|------|--------------|
| `Unix()`      | 秒    | 存储、日志、跨系统兼容  |
| `UnixMilli()` | 毫秒   | Web 请求、前端时间戳 |
| `UnixMicro()` | 微秒   | 高频监控、速度分析    |
| `UnixNano()`  | 纳秒   | 精细性能 & 事件顺序  |

### Time 转 Unix 时间戳

时间戳的数值不受时区影响，因为它们始终基于 UTC 的固定起点

| 方法签名                              | 返回类型 | 含义                                          |
|-----------------------------------|------|---------------------------------------------|
| `func (t Time) Unix() int64`      | 秒级   | 返回距 Unix Epoch（1970-01-01 00:00:00 UTC）的秒数。 |
| `func (t Time) UnixMilli() int64` | 毫秒级  | 返回距 Epoch 的毫秒数（Go1.17+）。                    |
| `func (t Time) UnixMicro() int64` | 微秒级  | 返回距 Epoch 的微秒数（Go1.17+）。                    |
| `func (t Time) UnixNano() int64`  | 纳秒级  | 返回距 Epoch 的纳秒数。                             |

### Unix 时间戳 转 Time

这些函数构造出的时间总是 UTC，如果要显示为本地或其他时区，需要再用 `In()` 或 `Local()` 转换时区

| 函数签名                                        | 含义                        |
|---------------------------------------------|---------------------------|
| `func Unix(sec int64, nsec int64) Time`     | 根据秒和纳秒构造 `Time`，时区默认为 UTC |
| `func UnixMilli(ms int64) Time` *(Go1.17+)* | 用毫秒构造 `Time`，时区默认为 UTC    |
| `func UnixMicro(us int64) Time` *(Go1.17+)* | 用微秒构造 `Time`，时区默认为 UTC    |

## 时间比较与判断方法

### 比较丰富

三者比较的是时间点（即 Unix 时间），而不是格式化后的字符串，也不是时区。

| 方法签名                                | 含义                                |
|-------------------------------------|-----------------------------------|
| `func (t Time) Before(u Time) bool` | 判断 `t` 是否早于（在 `u` 之前）             |
| `func (t Time) After(u Time) bool`  | 判断 `t` 是否晚于（在 `u` 之后）             |
| `func (t Time) Equal(u Time) bool`  | 判断 `t` 与 `u` 是否表示同一时刻点（与是否同一时区无关） |

### 时间零值

零时间 (zero time) 和 Unix 时间戳 0是两回事。
`time.Unix(0,0)` 代表：`1970-01-01 00:00:00 +0000 UTC`，也就是 Unix 纪元开始时间（Unix Epoch）。
而 `IsZero()` 判定的零值时间是：`0001-01-01 00:00:00 +0000 UTC`，也就是 Go 中 time.Time 的 默认零值。

| 方法签名                          | 含义                                             |
|-------------------------------|------------------------------------------------|
| `func (t Time) IsZero() bool` | 判断 `t` 是否为零值 (`0001-01-01 00:00:00 +0000 UTC`) |

## 时间运算

`time.Time` 的加减与时间运算相关的方法

### 加法与差值

| 方法签名                                 | 含义                                     |
|--------------------------------------|----------------------------------------|
| `func (t Time) Add(d Duration) Time` | 在 `t` 的基础上增加或减少一个时间段 `d`（可正可负），返回新时间   |
| `func (t Time) Sub(u Time) Duration` | 返回 `t - u` 的时间差，结果是 `time.Duration` 类型 |

> Duration 本质上是纳秒数（int64）。 Add() 和 Sub() 并不会改变原来的 t，而是返回新值

### 基于日历的加法

不依赖 Duration 的加法运算

| 方法签名                                                          | 含义                                          |
|---------------------------------------------------------------|---------------------------------------------|
| `func (t Time) AddDate(years int, months int, days int) Time` | 在 `t` 基础上按日历逻辑增加/减少 年/月/日，并自动处理月份溢出、月天数等规则。 |

> 适用于下一月、明年今天这类日历操作。与 `Add(time.Hour*24*30)` 的语义完全不同（天数所需不是固定常量）

### 与当前时间的便利比较

包级功能函数，不属于 Time 对象

| 函数签名                          | 含义                                     |
|-------------------------------|----------------------------------------|
| `func Since(t Time) Duration` | 等价于 `time.Now().Sub(t)`，表示从 t 到现在的时间差。 |
| `func Until(t Time) Duration` | 等价于 `t.Sub(time.Now())`，表示从现在到 t 的时间差。 |

> Since() -> 过去了多久
> Until() -> 还剩多久

## 时区相关

| 方法                                                  | 作用       | 说明                  |
|-----------------------------------------------------|----------|---------------------|
| `func LoadLocation(name string) (*Location, error)` | 加载指定时区   | 如 `"Asia/Shanghai"` |
| `func FixedZone(name string, offset int) *Location` | 创建固定偏移时区 | 例如 `+8` 不考虑 DST     |

## 延时

```
func Sleep(d Duration)
```

Sleep 会暂停当前 goroutine 至少 d 的持续时间。d 为负数或零时，Sleep 会立即恢复。

## Month / Weekend

这两个是枚举类型而不是 int

- 可读性强
- 避免魔法数字
- 匹配 Date 和 Weekday() 返回类型

### Weekend

`type Weekday`：Weekday 表示一周中的某一天（星期几）

```
type Weekday int

const (
	Sunday Weekday = iota // 0
	Monday                // 1
	Tuesday               // 2
	Wednesday             // 3
	Thursday              // 4
	Friday                // 5
	Saturday              // 6
)
```

> Sunday 是 0，不是 Monday

### Month

`type Month`：Month 表示一年中的月份

```
type Month int

const (
	January Month = 1
	February
	March
	April
	May
	June
	July
	August
	September
	October
	November
	December
)
```

