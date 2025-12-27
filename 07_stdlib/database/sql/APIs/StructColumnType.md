# sql.ColumnType

用来描述 SQL 结果集中每一列的元信息（metadata），不是给日常 CRUD 使用的，而是给需要理解结果集结构本身的代码用的

```
type ColumnType struct {
	// contains filtered or unexported fields
}
```

> 开发几乎不会使用

## 创建方式

不会手动 new 一个 `sql.ColumnType`，只会从 `*sql.Rows` 中获取：

```
rows, _ := db.Query("SELECT ...")
columnTypes, err := rows.ColumnTypes()
```

- `ColumnType` 是结果集级别的概念，而不是某一行的概念
- 只有 `Rows` 有，`Row` 单行没有

## 元信息

通过 ColumnType 的一些方法获取列信息

1. `ct.Name() string` 列名
2. `ct.DatabaseTypeName() string` 数据库类型名
3. `ct ScanType() reflect.Type` Go 扫描类型
4. `ct.Nullable` 是否可以为 NULL，判断是否需要 `sql.Null[T]` 的依据之一
5. `Length() / DecimalSize()` 长度/精度相关

## 方法

都是获取数据库列的相关信息

### DatabaseTypeName()

获取数据库原生类型名，返回数据库系统中的类型名字符串（不含长度/精度）

```
func (ci *ColumnType) DatabaseTypeName() string
```

例如：

| 列定义             | 返回值         |
|-----------------|-------------|
| `VARCHAR(50)`   | `"VARCHAR"` |
| `BIGINT`        | `"BIGINT"`  |
| `DECIMAL(10,2)` | `"DECIMAL"` |
| `ENUM('a','b')` | `"ENUM"`    |

- 不同数据库、不同驱动，字符串可能不同
- 不包含 `(50)`、`(10,2)` 这些规格
- 可能返回 `""`（驱动不支持）

### DecimalSize()

获取数据列的数值精度与小数位

```
func (ci *ColumnType) DecimalSize() (precision, scale int64, ok bool)
```

仅用于十进制数值表示（如 `DECIMAL` / `NUMERIC`）

- precision：总位数
- scale：小数位数
- ok：是否适用 / 是否支持，如果列不是 decimal 类型返回 false

`DECIMAL(10,2)` 返回：

- precision = 10
- scale = 2
- ok = true

### Length()

获取可变长度列的长度，只对变长类型有意义，`math.MaxInt64` 表示理论不限制长度，不代表数据库实际限制（数据库可能有上限）

```
func (ci *ColumnType) Length() (length int64, ok bool)
```

- `VARCHAR(50)` -> `length = 50`
- `TEXT / BLOB` -> `length = math.MaxInt64`
- `INT` -> `ok = false`

### Name()

获取结果集中的列名或别名

```
func (ci *ColumnType) Name() string
```

比如：`SELECT id AS student_id FROM student` 调用 `ci.Name() == "student_id"`

### Nullable()

该列是否允许为 NULL

```
func (ci *ColumnType) Nullable() (nullable, ok bool)
```

- `nullable`：这一列是否可能为 NULL
- `ok`：驱动是否能确定，为 false 很常见

| nullable | ok    | 含义       |
|----------|-------|----------|
| true     | true  | 允许 NULL  |
| false    | true  | NOT NULL |
| _        | false | 驱动不知道    |

主要用于：

- 判断是否需要 `sql.Null[T]`
- ORM 自动类型选择
- schema 检查

### ScanType()

返回一个适合用于 `Rows.Scan` 的 Go 类型

> 如果不知道这列该用什么 Go 类型接收，driver 建议用什么

```
func (ci *ColumnType) ScanType() reflect.Type
```

- 如果驱动不支持则返回 any 也就是 `interface{}`

**场景映射**

| SQL 类型     | ScanType          |
|------------|-------------------|
| VARCHAR    | `string`          |
| BIGINT     | `int64`           |
| BOOL       | `bool`            |
| DATETIME   | `time.Time`       |
| NULL / 不支持 | `interface{}/any` |


