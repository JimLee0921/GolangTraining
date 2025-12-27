# sql.Null

Go1.22 加入，`sql.Null` 是一个泛型版本的 Nullable 容器，用于在 Go 中安全表达 SQL 的 NULL

## 定义

与 `NullString` / `NullInt64` / `NullTime` 在语义上完全一致，只是变成了泛型

```
type Null[T any] struct {
    V     T
    Valid bool
}
```

Valid 决定有没有值，V 决定值是什么

### Valid

`Valid bool` 是语义位，表示是否存在值

`Valid == false`：表示在数据库中这一列为 NULL（值不存在），而不是 0 / 空字符串 / false / 零时间，此时 V 的内容是未定义的，不能使用

`Valid == true`：值在 v 中

### V

`V T` 载荷位也就是真正的值，只有在 `Valid == true` 时才有意义

当 `Valid == true` 时表示：数据库里确实有一个非 NULL 的值，它被安全地存放在 V 中

## 方法

### Scan()

用于从数据库读取值

```
func (n *Null[T]) Scan(value any) error
```

当使用：

```
var s sql.Null[string]
row.Scan(&s)
```

`database/sql` 会发现：`&s` 实现了 `sql.Scanner` 接口，于是调用 `s.Scan(value)`，这里的 value 就是 driver 返回的标准值之一：

```
int64 / float64 / bool / []byte / string / time.Time / nil
```

Scan 主要用于：

1. 判断是否为 SQL NULL
2. 能转换就进行转换
3. 不能无损转换则报错

如果 `value == nil`：

```
if value == nil {
    n.Valid = false
    var zero T
    n.V = zero
    return nil
}
```

- 数据库为 NULL
- Valid = false
- V 被设置为零值（不可使用，只是占位）
- 不会报错（合法情况，`NULL != error`）

如果 `value != nil` 则尝试把 value 转为 `T`

- 如果转换成功则 `n.V = 转换后的值`、`n.Valid = true`
- 如果转换失败则返回 error

### Value()

用于写回数据库

```
func (n Null[T]) Value() (driver.Value, error)
```

在调用 `db.Exec("INSERT INTO ...", ...)` 时 `database/sql` 会检查参数是否实现了 `driver.Valuer`，内部逻辑很简单：

```
if !n.Valid {
    return nil, nil
}
return n.V, nil
```

> 返回值必须为：int64 / float64 / bool / []byte / string / time.Time / nil，nil 是 driver 语义上的 NULL

## 新增方式

在 Go 没有泛型之前，早期只能使用 `database/sql` 定义好的：

```
type NullString struct { String string; Valid bool }
type NullInt64  struct { Int64 int64; Valid bool }
```

- 提供一堆重复但明确的类型
- 而不能提供一个通用方案

在 Go 1.18 提出泛型后，标准库的规则就是：

- 观察真实需求
- 避免破换既有 API
- 避免泛型滥用

`sql.Null[T]` 对比那些老的 `sql.Null*` 非常明确，几乎没有歧义，与现有的 API 完全兼容

## 实现 Scanner 接口

`Null[T]` 实现了 Scanner 接口，也就是：

```
var s sql.Null[string]
row.Scan(&s)
```

1. `Rows.Scan` 发现 `&s` 实现了 Scanner
2. 调用 `s.Scan(src)`
3. 在 Scan 里：
    - `src == nil` -> `Valid = false`
    - 否则会尝试把 src 转为 `T`，并写入 `V`

## `T` 的严格约束

`T` 不能是任意类型，只能是：

```
int64
float64
bool
string
[]byte
time.Time
```

以及这些类型的合理别名/包装

> database/sql 只认识 driver.Value 那一套类型系统

## 与旧 `Null*` 的关系

1. 功能层面上是等价的：

    ```
    sql.NullString   ≈ sql.Null[string]
    sql.NullInt64    ≈ sql.Null[int64]
    sql.NullTime     ≈ sql.Null[time.Time]
    ```

2. 旧的 API 依旧保留，为了
    - 向后兼容
    - 老代码无需重构
    - 避免破坏公共 API

3. Go 1.22 +和新项目还是建议使用 `sql.Null[T]`，代码更统一，便于维护

## 适用场景

### Scan 数据库行

在从数据库读数据时需要使用判断数据是否为 NULL

```
type StudentRow struct {
    ID     int64
    Name   string
    Gender string
    Phone  sql.Null[string]
}
```

- Phone 允许为 NULL
- 直接使用 string 可能导致 Scan 读取时报错
- `sql.Null` 是标准解决方案

### 插入更新可选字段

```
db.Exec(
    "UPDATE student SET phone=? WHERE id=?",
    sql.Null[string]{Valid: false},
    id,
)
```

- 明确写 NULL
- 不依赖 SQL 拼接
- 不依赖魔法值（空字符串）

### 不能用于业务模型

不推荐用于业务层面，业务层面不应该关注 SQL NULL，会污染领域模型

```
type Student struct {
    Phone sql.Null[string]
}
```

### 不应该直接返回 API / JSON

主要面向数据库层面，不要使用交互

```
return studentRow
```