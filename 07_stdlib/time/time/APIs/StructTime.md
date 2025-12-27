# time.Time

`time.Time` 是 Go 用来表示某一个具体时刻的结构体，包含一个时间点（纳秒级别）和 一个时区（Location）

使用时间的程序通常应该将时间存储和传递为值，而不是指针，也就是说，时间变量和结构体字段应该是 `time.Time` 类型，而不是
`*time.Time` 类型

多个 goroutine 可以同时使用 Time 值，但`Time.GobDecode`、`Time.UnmarshalBinary`、`Time.UnmarshalJSON`等一些 time 方法不是并发安全的

> time.Time 永远代表一个真实发生的时刻，即使格式看起来不同


下面先按“学习路径 + 工程用途”把你列出的 `time.Time` 相关函数/方法做一个**分层分类**
。这样后面我们可以按模块逐个攻克，每一块都能形成闭环（构造 → 解析 → 转换 → 运算 → 比较 → 输出 → 序列化）。

# 相关方法和函数

关于 `time.Time` 相关的函数和方法分为下面这些

## 1. 构造与获取当前时间（Construct / Acquire）

用于创建 `Time` 值或拿到现在的 `Time` 值

- `func Date(year int, month Month, day, hour, min, sec, nsec int, loc *Location) Time`
- `func Now() Time`
- `func Unix(sec int64, nsec int64) Time`
- `func UnixMilli(msec int64) Time`
- `func UnixMicro(usec int64) Time`

### time.Date()

最底层、最核心、最常用的 Time 构造器，根据年月日时分秒纳秒 + 时区，创建一个 `Time`，所有其它 Time 构建方式都开业等价为一次
Date

```
func Date(year int, month Month, day, hour, min, sec, nsec int, loc *Location) Time
```

**参数**

- `year int`：公历年份，支持负数（公元前），常规业务使用正数
- `month Month`：类型是 `time.Month`，不是 int，推荐使用 `time.January`，`time.Month(1)` 也可以但不推荐
- `day int`：月中第几天，可以溢出但是不合法，会自动进位
- `hour, min, sec int`：时(0-23)分(0-59)秒(0-59)
- `nsec int`：纳秒， `0 <= nsec < 1e9`，超出也会进位
- `loc *Location`：决定这是哪个失去下的当前时间，不能为 nil

> Date 会自动规范化，结果是一个完全合法的时间，永远不报错，只会进位

### time.Now()

获取当前时间，返回当前系统时间，精度默认是纳秒级别

```
func Now() Time
```

默认使用系统时钟，返回的 Time 包含时区，为 `time.Local` 可以通过 `time.Now().UTC()` 再转为 UTC 时区

### time.Unix()

时间范围从 Unix Epoch（1970-01-01 00:00:00 UTC）开始计算构造时间对象

```
func Unix(sec int64, nsec int64) Time
```

**参数**

- sec：自 Epoch 起的秒数
- nsec：纳秒偏移，`[0, 999999999]`

**返回值**

- 返回的 Time 对象的 Location 是 UTC，与本地时区无关

### time.UnixMilli() / time.UnixMicro()

更符合现代系统（DB / JS / Kafka / 日志系统）的时间戳单位来构造 Time 对象，和 Unix 一样时间范围从 `1970-01-01 00:00:00 UTC`
开始计算

```
func UnixMilli(msec int64) Time
func UnixMicro(usec int64) Time
```

- UnixMilli 参数为毫秒，返回的 Time 对象默认 Location 为 UTC
- UnixMicro 参数为微秒，返回的 Time 对象默认 Location 为 UTC

## 2. 解析（字符串 -> Time）（Parse / Interpret）

把文本按 layout 解释成 `Time`

- `func Parse(layout, value string) (Time, error)`
- `func ParseInLocation(layout, value string, loc *Location) (Time, error)`

### time.Parse

按 layout 解析字符串 value，返回对应的 `time.Time` 对象，当字符串中没有提供时区信息时，Parse 默认按 UTC 解释

```
func Parse(layout, value string) (Time, error)
```

**参数**

- `layout string`：格式样例（用参考时间写出来）
- `value string`：待解析的时间文本

**返回值**

- `Time`：解析成功的时间
- `error`：layout 与 value 不匹配、值非法等导致失败

> 推荐优先使用 `time.ParseInLocation` 函数进行时区处理

### time.ParseInLocation

和 `time.Parse` 类似用于将 layout 字符串解析为 `time.Time` 对象，但是需要显示指定时区

```
func ParseInLocation(layout, value string, loc *Location) (Time, error)
```

**参数**

- `layout string`：格式样例
- `value string`：待解析文本
- `loc *Location`：用于“无时区输入”的默认时区（不能为 nil）

**返回值**

同 `time.Parse`

**注意事项**

1. 如果 value 中本身已经带了 Location 信息，则时区信息按照输入的 value 为主，loc 只负责补充默认时区，而不是强制覆盖
2. 解析本地时间一定优先使用 `time.ParseInLocation`

## 3. 时区与位置（Time Zone / Location）

把同一时刻“映射成不同地区的表示”，或查询 DST/时区信息。

- `func (t Time) In(loc *Location) Time`
- `func (t Time) Local() Time`
- `func (t Time) UTC() Time`
- `func (t Time) Location() *Location`
- `func (t Time) Zone() (name string, offset int)`
- `func (t Time) ZoneBounds() (start, end Time)`
- `func (t Time) IsDST() bool`

### In()

把同一个时间点转换为指定时区下的表示形式（不修改原值，返回新的 Time 对象需要接收）

```
func (t Time) In(loc *Location) Time
```

**参数**

- `loc *Location`：不能为 nil，通常来自 `time.LoadLocation` / `time.UTC` / `time.Local`

**返回值**

- 新的 Time 值，Location 已经更新

### Local()

把时间表示切换为系统本地时区，等价于 `t.In(time.Local)`

```
func (t Time) Local() Time
```

> `time.Local` 取决于： 操作系统/容器配置/TZ 环境变量，在容器 / 服务器中非常不稳定

### UTC()

把时间表示切换为 UTC 时区，等价于 `t.In(time.UTC)`

```
func (t Time) UTC() Time
```

**使用场景**

- 存储 / 日志 / 跨系统通信：优先用 UTC
- 展示给用户：再转回本地或指定时区

### Location()

返回该 Time 当前使用的 Location

```
func (t Time) Location() *Location
```

### Zone()

返回当前时间点在该 Location 下的时区名称和相对于 UTC 时区的秒偏移量

```
func (t Time) Zone() (name string, offset int)
```

**注意事项**

1. name 可能不是 IANA 名：可能是 `CST` / `EST` / `EDT` 等缩写，不可以用于逻辑判断

### ZoneBounds()

返回当前时间点所在的时区规则区间的起止时间，也就是 offset / DST 状态从什么时候开始，到什么时候结束

```
func (t Time) ZoneBounds() (start, end Time)
```

**返回值**

- start：当前 offset 生效的起点
- end：当前 offset 失效的时间（DST 切换点）

### IsDST()

判断当前时间点在该 Location 下是否处于夏令时（DST）

```
func (t Time) IsDST() bool
```

- 取决于 Location 和当前时间点
- UTC 永远返回 false

## 4. 加减与对齐（Arithmetic / Alignment）

用于时间推进、日期推进、取整/截断。

- `func (t Time) Add(d Duration) Time`
- `func (t Time) AddDate(years int, months int, days int) Time`
- `func (t Time) Sub(u Time) Duration`
- `func (t Time) Round(d Duration) Time`
- `func (t Time) Truncate(d Duration) Time`

### Add()

把时间点 t 按固定时长推进 d（`time.Duration` 对象，纳秒精度，基于物理时长），返回新时间

```
func (t Time) Add(d Duration) Time
```

**参数**

- `d Duration`：推进的时长，可正可负

**返回值**

- 新的推进后的 `time.Time` 对象

### AddDate()

按照日历单位进行推进：年/月/日

```
func (t Time) AddDate(years int, months int, days int) Time
```

**参数**

均可正可负

- years：年增量
- months：月增量
- days：日增量

**返回值**

- 新的 `time.Time` 对象

> 如果溢出，会进行顺延进位

### Sub()

返回 `t-u` 的时间差，单位是 `time.Duration` 纳秒精度

```
func (t Time) Sub(u Time) Duration
```

**参数**

- `u time.Time`：被减的时间对象（相对于 t 可早可晚）

**返回值**

- `time.Duration`：
    - 正：t 在 u 之后
    - 负：t 在 u 之前

> 如果结果超过 Duration 类型可以存储的最大/最小值则返回最大或最小持续时间，比较的是绝对时间点，与 Location 无关

### Round()

返回把 t 四舍五入进行取整到最接近的 d 的倍数的结果，对于半数向上取整（不是银行家舍入round half to even而是始终向上）

```
func (t Time) Round(d Duration) Time
```

**参数**

- `d Duration`：舍入粒度，如果 `d <= 0` 返回的时间会移除 monotonic clock

### Truncate()

会把时间 t 截断到不大于当前时间的最近一个 d 的整数倍，也就是向下取整，等价于数学上的 floor

```
func (t Time) Truncate(d Duration) Time
```

**参数**

- `d Duration`：取整粒度，如果 `d <= 0` 返回的时间会移除 monotonic clock

**返回值**

- 向下取整后的新的 Time

## 5. 比较与判定（Comparison / Predicates）

判断先后、相等、零值等。

- `func (t Time) After(u Time) bool`
- `func (t Time) Before(u Time) bool`
- `func (t Time) Equal(u Time) bool`
- `func (t Time) Compare(u Time) int`
- `func (t Time) IsZero() bool`

### After()

判断时间 t 是否晚于时间 u，很重要的一个使用场景就是用 `time.After` 给 select 添加一个超时分支

```
func (t Time) After(u Time) bool
```

**参数**

- `u time.Time`：比较时间对象

**返回值**

- `bool`：t 早于 u 为 true

> 比较的是绝对时间点，与 Location 时区无关，同一时刻在 Shanghai/UTC 表示不同，但是 After/Before 结果一致

### Before()

判断时间 t 是否早于时间 u

```
func (t Time) Before(u Time) bool
```

**参数**

- `u time.Time`：比较时间对象

**返回值**

- `bool`：t 晚于 u 为 true

> 和 After 一样不受时区影响，即使两个时间点位于不同的时区也可能相等，比如 `6:00 +0200` 和 `4:00 UTC` 是相等的

### Equal()

判断 t 和 u 是否表示同一个绝对时间点

```
func (t Time) Equal(u Time) bool
```

- 只看时刻，不看时区表示
- Go 虽然允许 `time.Time` 使用 `==` 进行比较，但是比较的是结构体全部字段可能受一些细节影响
- 比较是否同一时刻一律用 `Equal` 不要使用 `==`

### Compare()

三路比较：将时间点 t 与 u 进行比较

```
func (t Time) Compare(u Time) int
```

**返回值**

- int 类型：
    - `-1`：`t < u`
    - `0`：`t == u`（同一时刻）
    - `+1`：`t > u`

**适用场景**

1. 排序
2. 构建区间比较逻辑时更直观
3. 需要三态输出时减少 `if/else` 链

### IsZero()

判断 t 是否为 `time.Time` 的零值

```
func (t Time) IsZero() bool
```

`time.Time{}` 零值对应的时间点是：

- 年：1
- 月：January
- 日：1
- 00:00:00 UTC

## 6. 拆解字段与日历信息（Field Accessors / Calendar）

把 `Time` 拆成年月日时分秒等，或周/年日等派生信息。

- `func (t Time) Date() (year int, month Month, day int)`
- `func (t Time) Year() int`
- `func (t Time) Month() Month`
- `func (t Time) Day() int`
- `func (t Time) Clock() (hour, min, sec int)`
- `func (t Time) Hour() int`
- `func (t Time) Minute() int`
- `func (t Time) Second() int`
- `func (t Time) Nanosecond() int`
- `func (t Time) Weekday() Weekday`
- `func (t Time) YearDay() int`
- `func (t Time) ISOWeek() (year, week int)`

### Date() / Year() / Month() / Day()

- Date：一次性返回年/月/日
- Year：返回年
- Month：返回月（time.Month，不是 int）
- Day：返回日（月中的第几天，不是年中的第几天）
-

```
func (t Time) Date() (year int, month Month, day int)
func (t Time) Year() int
func (t Time) Month() Month
func (t Time) Day() int
```

### Clock() / Hour() / Minute() / Second() / Nanosecond()

- Clock：一次性返回 小时 / 分钟 /秒（不含纳秒）
- Hour：返回小时
- Minute：返回分钟
- Second：返回秒
- Nanosecond：返回秒中的纳秒（0-999,999,999）

```
func (t Time) Clock() (hour, min, sec int)
func (t Time) Hour() int
func (t Time) Minute() int
func (t Time) Second() int
func (t Time) Nanosecond() int
```

### Weekday()

返回周几（`time.Weekday`）

```
func (t Time) Weekday() Weekday
```

- Sunday = 0
- 基于 Location

### YearDay()

返回一年中的第几天

```
func (t Time) YearDay() int
```

> 完全基于日历，平年为 1-365 闰年为 1-366

### ISOWeek()

获取 ISO 8601 周编号

```
func (t Time) ISOWeek() (year, week int)
```

**ISO 8601 规定**

1. 周一是每周的第一天
2. 一年中的第一周是包含该年第一个周四的那一周
    - 年初的几天可能属于上一年的最后一周
    - 年末的几天可能属于下一年的第一周

**注意事项**

- ISOWeek 的 `year != Year`，两个在年初/年末可能不同
- 只在明确要求 ISO 周时才用

## 7. 输出与格式化（Formatting / Stringification）

把 `Time` 变成可读字符串或追加到 buffer。

- `func (t Time) Format(layout string) string`
- `func (t Time) String() string`
- `func (t Time) GoString() string`
- `func (t Time) AppendFormat(b []byte, layout string) []byte`

### Format()

按照给定的 layout 字符串把 Time 格式化为字符串（Parse 的反向操作，使用同一套 layout 规则）

```
func (t Time) Format(layout string) string
```

**参数**

- `layout string`：时间字符串格式样例
- 不会返回 error

**返回值**

- `string`：格式化后的结果

**注意事项**

1. Format 基于当前 Location
2. layout 如果不符合规范不会报错，只会原样输出

### String()

返回一个可读性较好的字符串表示

```
func (t Time) String string
```

**注意事项**

1. 格式不是稳定保证，不同版本和系统可能细节不一样
2. 包含时区缩写信息

### GoString()

返回 Go 语法风格的字符串表示，用于 `%#v` 输出，主要用于 Debug 调试才会使用

```
func (t Time) GoString() string
```

### AppendFormat()

高级功能，把格式化后的时间追加到已有 `[]byte`，而不是新分配字符串

```
func (t Time) AppendFormat(b []byte, layout string) []byte
```

**参数**

- `b []byte`：目标 buffer
- `layout string`：格式样例

**返回值**

- `[]byte`：拓展后的 `[]byte`

> 主要用于高性能日志系统或自定义序列化

## 8. Unix 时间戳导出（Epoch / Unix Conversions）

把 `Time` 转为 epoch 计数（秒/毫秒/微秒/纳秒）这些时间戳都只表示绝对时间点，与 Location 无关

- `func (t Time) Unix() int64`
- `func (t Time) UnixMilli() int64`
- `func (t Time) UnixMicro() int64`
- `func (t Time) UnixNano() int64`

> Unix 时间戳的起点是：`1970-01-01 00:00:00 UTC` 所有 `Unix*` 方法返回的都是相对于这个时代的偏移量

### Unix()

返回 t 对于 1970年1月1日 UTC 开始这个时刻的秒级别 Unix 时间戳偏移量

```
func (t Time) Unix() int64
```

### UnixMilli()

返回 t 对于 1970年1月1日 UTC 开始这个时刻的毫秒级别 Unix 时间戳偏移量

```
func (t Time) UnixMilli() int64
```

### UnixMicro()

返回 t 对于 1970年1月1日 UTC 开始这个时刻的微秒级别 Unix 时间戳偏移量

```
func (t Time) UnixMicro() int64
```

### UnixNano()

```
func (t Time) UnixNano() int64
```

返回 t 对于 1970年1月1日 UTC 开始这个时刻的纳秒级别 Unix 时间戳偏移量

## 9. 编解码与序列化（Serialization / Encoding）

偏底层，使用用于二进制、文本、JSON、gob 等序列化协议。

- `func (t Time) MarshalBinary() ([]byte, error)`
- `func (t *Time) UnmarshalBinary(data []byte) error`
- `func (t Time) AppendBinary(b []byte) ([]byte, error)`
- `func (t Time) MarshalText() ([]byte, error)`
- `func (t *Time) UnmarshalText(data []byte) error`
- `func (t Time) AppendText(b []byte) ([]byte, error)`
- `func (t Time) MarshalJSON() ([]byte, error)`
- `func (t *Time) UnmarshalJSON(data []byte) error`
- `func (t Time) GobEncode() ([]byte, error)`
- `func (t *Time) GobDecode(data []byte) error`

### MarshalBinary() / UnmarshalBinary() / AppendBinary()

Binary 二进制形式编码解码，表示的是 绝对时间点 + Location 信息，单调时钟信息不会被编码，编码结果不适合人工阅读

```
func (t Time) MarshalBinary() ([]byte, error)
func (t *Time) UnmarshalBinary(data []byte) error
func (t Time) AppendBinary(b []byte) ([]byte, error)
```

### MarshalText() / UnmarshalText() / AppendText()

文本形式编码解码，与 String 方法不同，这是规范化格式

```
func (t Time) MarshalText() ([]byte, error)
func (t *Time) UnmarshalText(data []byte) error
func (t Time) AppendText(b []byte) ([]byte, error)
```

### MarshalJSON() / UnmarshalJSON()

Json 编码解码，最为重要，格式为 RFC3339 / RFC3339Nano 格式:

```
"2006-01-02T15:04:05Z07:00"
"2006-01-02T15:04:05.999999999Z07:00"
```

```
func (t Time) MarshalJSON() ([]byte, error)
func (t *Time) UnmarshalJSON(data []byte) error
```

- MarshalJSON 实现 `encoding/json.Marshaler` 接口，如果时间格式不符合规范会返回 error
- UnmarshalJSON 实现 `encoding/json.UnmarshalJSON`，时间格式同样必须是符合 RFC3339 格式的带引号的字符串

### GobEncode() / GobDecode()

Go 专用序列化协议进行 Gob 编码解码，类型安全，比 JSON 紧凑，比 Binary 友好，但是不可跨语言，实现 `gob.GobDecoder` 和
`gob.GobEncoder` 接口

```
func (t Time) GobEncode() ([]byte, error)
func (t *Time) GobDecode(data []byte) error
```


