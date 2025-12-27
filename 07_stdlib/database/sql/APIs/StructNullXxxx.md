# sql.NullXxx

`sql.NullXxxx` 这些都是为了 Go 中安全、 无歧义地表达数据库列可能为 NULL 这一事实，因为直接使用 Go 基本类型的零值会丢失
NULL 语义

新版本主要使用 `sql.Null[T]` 泛型写法来代替这些写法，具体使用见 [StructNull.md](StructNull.md)

```
type NullBool struct {
	Bool  bool
	Valid bool // Valid is true if Bool is not NULL
}
type NullByte struct {
	Byte  byte
	Valid bool // Valid is true if Byte is not NULL
}
type NullFloat64 struct {
	Float64 float64
	Valid   bool // Valid is true if Float64 is not NULL
}
type NullInt16 struct {
	Int16 int16
	Valid bool // Valid is true if Int16 is not NULL
}
type NullInt32 struct {
	Int32 int32
	Valid bool // Valid is true if Int32 is not NULL
}
type NullInt64 struct {
	Int64 int64
	Valid bool // Valid is true if Int64 is not NULL
}
type NullString struct {
	String string
	Valid  bool // Valid is true if String is not NULL
}
type NullTime struct {
	Time  time.Time
	Valid bool // Valid is true if Time is not NULL
}
```

## 统一结构

内部结构本质上都是，Valid 永远是语义位，值字段永远是载荷位。

```
type NullXxx struct {
    Xxx T
    Valid bool
}
```

所有 NullXxx.Scan 都遵循同一套规则：

1. `value == nil`
    - 数据库是 NULL
    - `Valid = false`
    - 值字段设为零值
    - 不报错

2. `value != nil`
    - 尝试转换为目标类型
    - 成功：写入值字段，`Valid = true`
    - 失败：返回 error（防止信息丢失）

都有一个 Scan 和 Value 方法分别用于从数据库读取值和向数据库写入值，都是为了实现 `sql.Scanner` 和 `driver.Valuer` 接口

## 对比 `sql.Null[T]`

Go1.22 开始：

```
sql.NullString   ≈ sql.Null[string]
sql.NullInt64    ≈ sql.Null[int64]
sql.NullBool     ≈ sql.Null[bool]
sql.NullTime     ≈ sql.Null[time.Time]
```

功能等价、语义等价、实现原理等价