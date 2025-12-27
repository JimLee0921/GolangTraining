# sql.Scanner

主要用于 `Rows.Scan` 实现 `sql.Scanner` 接口

```
type Scanner interface {
    Scan(src any) error
}
```

> 只要把一个实现了 Scanner 接口的类型传给 `rows.Scan`，`database/sql` 就会把如何接收数据库值的控制权交给你

## 流程

```
数据库驱动 -> driver.Value(原始值) -> database/sql -> Rows.Scan(...) -> 如果目标变量实现了 Scanner -> 调用 Scan(src)
                                                                  -> 否则                     -> 使用内建转换规则
```

## 调用时机

当 `rows.Scan(&dest)` 发生时：

1. 对当前列的每一列取一个值
2. 对应到传入的每一个 `&dest`
3. 如果 `dest` 实现了 `Scanner`：`dest.Scan(src)`
4. 在 `Scan` 中决定：
    - 是否接受这个值
    - 如何进行存储
    - 是否报错

> Scan 是数据库值 -> Go 内存表示的最后一道关卡

## `src any`

在 Scan 中，面对的是数据库语义类型，而不是 Go 业务类型，Go 文档明确规定了 src 的合法类型集合：

```
int64
float64
bool
[]byte
string
time.Time
nil   // 表示 SQL NULL
```

> 永远不会在 Scan 中拿到 `int32` / `uint` / `struct`，这是 `database/sql` 对 driver 的标准化输出

如果 `src == nil`，表示数据库中这一列是 NULL，并不是 空字符串 / 0 / false，而是无值，
这也就是为什么直接用 `string / int` 接 NULL 会出错，以及为什么有 NullString 这类类型

## 内存与生命周期

当 src 为 `[]byte` 时，这块内存是 driver 的，下一次 `Scan()` 可能直接复用或释放，如果保存引用或把它存入 struct 不拷贝，必定产生遮蔽
bug，正确示例：

```
func (s *MyType) Scan(src any) error {
    b, ok := src.([]byte)
    if ok {
        s.Data = append([]byte(nil), b...) // 拷贝
        return nil
    }
    ...
}
```

## Scan 错误原因

在下面情况应当报错：

- src 是 float64，但想存到 int
- src 是 string，但格式非法
- src 是 `time.Time`，但期望特定时区/格式

不应默默吞掉的情况：

- NULL 被当作 0
- 超出范围数值被截断
- 字符串被强制裁剪

> Scan 是数据正确性的最后防线

## 自定义 Scanner 示例

```
type MyInt struct {
    V     int64
    Valid bool
}

func (m *MyInt) Scan(src any) error {
    if src == nil {
        m.Valid = false
        m.V = 0
        return nil
    }

    switch v := src.(type) {
    case int64:
        m.V = v
        m.Valid = true
        return nil
    default:
        return fmt.Errorf("cannot scan %T into MyInt", src)
    }
}
```

## Scanner 与 `sql.Null*` 的关系

比如 NullString 本质就是：

```
type NullString struct {
    String string
    Valid  bool
}
```
- 实现了 Scan
- 严格处理 `nil`、`string`、`[]byte`