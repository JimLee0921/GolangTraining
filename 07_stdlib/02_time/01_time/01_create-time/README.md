## 创建 Time 对象

`time.Time` 的获取方式大致分成 4 类：当前、构造、解析、时间戳

### 获取当前时间

* `func Now() Time`
  返回当前本地时间

```
t := Time.Now()
fmt.Println(t)
```

类似输出：`2025-10-31 09:18:25.123456 +0800 CST`

| 部分                         | 说明           |
|----------------------------|--------------|
| 2025-10-31 09:18:25.123456 | 日期 + 时间 + 微秒 |
| +0800                      | 时区偏移量（对 UTC） |
| CST                        | 时区名称         |

> 版本和平台不同，输出也可能不同，内部仍记录为UTC时间点，只是显示为本地时区

### 直接构造

* `func Date(year int, month Month, day, hour, min, sec, nsec int, loc *Location) Time`
    - 根据年月日时分秒纳秒 + 时区，创建一个 `Time`
    - month 可以传 int 也可以使用 `time.Month` 枚举

* `func Unix(sec int64, nsec int64) Time`
  根据 Unix 时间戳（秒 + 纳秒）构造时间（使用 UTC）

* `func UnixMilli(ms int64) Time`
  根据毫秒时间戳构造时间（使用 UTC）

* `func UnixMicro(us int64) Time`
  根据微秒时间戳构造时间（使用 UTC）

* `func UnixNano(ns int64) Time`
  根据纳秒时间戳构造时间（使用 UTC）

> Date = 人类可读输入；Unix = 数字时间戳输入

---

### 从字符串解析

* `func Parse(layout, value string) (Time, error)`
  按 layout 对字符串解析，解析结果为 UTC

* `func ParseInLocation(layout, value string, loc *Location) (Time, error)`
  按 layout 解析，并在指定时区解释该字符串

> ParseInLocation 用来避免时区歧义问题，是日常解析更常用的方式

---

### 特殊情形构造

* `func (t Time) In(loc *Location) Time`
  将已存在的时间转换为另一时区

* `func (t Time) Local() Time`
  转为本地时区

* `func (t Time) UTC() Time`
  转为 UTC

> 已有时间对象 -> 修改时区

## layout 补充

layout 是 Go 语言 独有的时间格式模板，它不是用 `%Y-%m-%d` 这种格式，而是用一个固定的参考时间来表示格式。

Go 的设计者认为与其用 `%Y`, `%m`, `%d` 这种抽象符号，不如直接用一个真实日期作为模板，
写什么格式，就写成想要时间输出的样子，只是其中的数字要用这个固定参考时间来替代

这个参考时间是且只能是：`2006-01-02 15:04:05`也可以写成：`Mon Jan 2 15:04:05 MST 2006` 等其它格式

### layout 格式

记忆口诀：2006 1 2 3 4 5

| 数字       | 表示含义   |
|----------|--------|
| **2006** | 年（4位）  |
| **06**   | 年（2位）  |
| **01**   | 月      |
| **02**   | 日      |
| **15**   | 24 小时制 |
| **03**   | 12 小时制 |
| **04**   | 分钟     |
| **05**   | 秒      |

下面这些都是合法的

```
"2006/01/02"
"02-Jan-2006"
"15:04"
"03:04 PM"
"Mon Jan 2 15:04:05"
"2006-01-02T15:04:05Z07:00"   // RFC3339
```

### 常用预定义格式

| 常量                          | 含义                              |
|-----------------------------|---------------------------------|
| `time.RFC3339`              | `2006-01-02T15:04:05Z07:00`     |
| `time.RFC1123`              | `Mon, 02 Jan 2006 15:04:05 MST` |
| `time.DateTime` *(Go1.20+)* | `2006-01-02 15:04:05`           |
| `time.DateOnly` *(Go1.20+)* | `2006-01-02`                    |
| `time.TimeOnly` *(Go1.20+)* | `15:04:05`                      |



