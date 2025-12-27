# time.Month

`time.Month` 本质上是个 int，属于月份枚举类型，主要用于提高可读性

```
type Month int
```

## 枚举值

注意月份是从 1 开始的，而不是跟其它某些语言一样从 0 开始，主要为了跟人的思维认知一致，避免错位

```
const (
    January Month = 1 + iota
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

## String 方法

Month 只有一个 String 方法用于返回月份的英文名称

```
func (m Month) String() string
```