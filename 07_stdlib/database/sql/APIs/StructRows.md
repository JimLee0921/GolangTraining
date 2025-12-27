# sql.Rows

`sql.Rows` 表示一个结果集的游标，用于逐行读取 0-N 行数据，不是数据本身，而是：

- 指向数据库结果集的游标
- 一次只返回一行真正的数据

```
type Rows struct {
	// contains filtered or unexported fields
}
```

## 获取/生命周期

通过 `sql.DB/Conn/Stmt/Tx` 获取 `sql.Rows`

1. `db.Query(...)`：返回 `*Rows`（游标在第一行之前）
2. 使用 `rows.Next()` + `rows.Scan` 不断读取每一行
3. 最后必须使用 `rows.Close` 进行资源释放

> Query 之后游标在第一行之前，Next 才是真正向数据库获取数据，`Close` 必须调用，错误不一定在 Query 时发生

# 主要方法

方法不多，简单分类

## 游标推进类

Cursor navigation，用于让游标往前走，决定你当前指向哪一行/哪一组结果集

- `Next() bool`：前进到下一行，返回 true 表示当前有一行可读，false 表示结束或出错（出错要看 `Err()`）。
- `NextResultSet() bool`：前进到下一组结果集（multi result sets）。常见于存储过程或一次执行返回多组 SELECT 的场景

### `Next()`

把游标从当前行之前/当前行推送到下一行，并准备好给 `rs.Scan()` 读取

```
func (rs *Rows) Next() bool
```

**返回值**

- 如果返回 true：当前至少有一行可读
- 如果返回 false：没有更多行，或发生了某些错误

> false 不等于没有错误，是否出错需要看 `Rows.Err()`

**注意事项**

每次 Scan 之前必须先 Next：

```
for rows.Next() {
    rows.Scan(...)  // 读取当前游标所指向的行
}

if err := rows.Err(); err != nil {
    return err
}
```

- Rows 的游标初始化是在第一行之前
- `Next()` 是把游标移动到某一行的唯一方式
- `Scan()` 只负责读取当前行，不负责移动
- 后续需要使用 `rows.Err` 判断是哪种方式结束的

### `NextResultSet()`

结果集级别推进，用于一次数据库调用，返回多组结果集（multiple result sets），每一组结果集本身又包含 0-N 行

结果集主要用于：

- 存储过程（CALL）
- 一次执行多条 SELECT
- 某些数据库的特殊用法

```
func (rs *Rows) NextResultSet() bool
```

**返回值**

- true：成功切换到下一组结果集
- false：没有更多结果集或发生了错误（需要使用 `Rows.Err()` 进行判断）

**注意事项**

在切换到新的结果集之后必须重新从 `Next()` 开始读行，新的结果集可能一行都没有

## 执行获取数据类

Data fetch / scan，用于把当前行的列值复制到你的目标变量中

- `Scan(dest ...any) error`：把当前行各列写入 `dest` 指向的变量；必须在 `Next()` 返回 `true` 后调用。若 `dest` 某个参数实现了
  `sql.Scanner`，会回调它的 `Scan`

### Scan()

用于把当前行的每一列值复制到提供的目标变量（必须为指针）中

```
func (rs *Rows) Scan(dest ...any) error
```

**注意事项**

- Rows 的游标初始在第一行之前，`Next()` 是唯一能把游标移动到某一行的方法，`Scan()` 不负责移动游标
- `len(dest)` 必须等于查询返回的列数，否则直接返回 error，不能做部分字段 Scan

**参数**
`dest ...any`，这里传入的需要是指针，Scan 做的是赋值

支持的目标类型有：

1. 基础 Go 类型指针：
    - `*string`
    - `*int / *int64 / *uint64`
    - `*bool`
    - `*float32 / *floate64`
    - `*time.Time`
2. `*interface{}`
    - Scan 会被 driver 返回的原始值放进去
    - `[]byte` 会被复制
    - 常用于动态查询
3. `*[]byte` 对比 `*RawBytes`
    - `*[]byte` 会进行复制，更安全，推荐使用
    - `*RawBytes` 不进行复制，可以性能优化，但是极不安全，业务代码不要使用
4. 可以是实现了 Scanner 的类型：
    - `sql.NullXxxx`
    - `sql.Null[T]`
    - 自定义实现的 Scanner
    - 如果 dest 实现了 Scanner，Scan 会调用它的 `Scan(value)`

## 元数据类

Metadata / schema of result set，用于了解这一组结果集有哪些列、列名是什么、列类型是什么，不读取行值

- `Columns() ([]string, error)`：返回列名列表（或别名）。最轻量、最常用
- `ColumnTypes() ([]*ColumnType, error)`：返回更完整的列元信息（类型名、可空、长度、ScanType 等）。常用于导出工具、ORM、动态扫描

### Columns()

获取 SQL 查询结果中的列名

```
func (rs *Rows) Columns() ([]string, error)
```

**返回值**

- `[]string`：一个列名切片，顺序与 `Scan()` 的列顺序一致，列明时 SQL 结果集中的名字或者别名
- `error`：错误信息

**注意事项**

- 在 `Query()` 之后即可调用，不需要 `Next()`
- 不能再 `rows.Close` 之后再进行调用（会报错）

### ColumnTypes()

返回完整的列元信息

```
func (rs *Rows) ColumnTypes() ([]*ColumnType, error)
```

**返回值**

- `[]*ColumnType`：参考 [StructColumnType.md](StructColumnType.md)，每个元素描述一列的元信息（顺序与 Columns / Scan 完全一致）
- `error`：错误信息

## 收尾与错误类

Lifecycle / error handling，用于资源释放与错误汇总（尤其是迭代过程中发生的错误）

- `Close() error`：释放结果集资源并归还连接/游标相关资源。凡是 `Query()` 拿到 `Rows`，都应 `defer rows.Close()`
- `Err() error`：返回迭代过程中遇到的错误（网络/驱动/解码等）。标准模式是循环结束后检查 `rows.Err()`

### Close()

释放资源，只要用了 `db.Query()`，就必须 `Close()`

```
func (rs *Rows) Close() error
```

**注意事项**

- 当完整遍历到末尾，且没有多结果集时内部会自动触发关闭
- 但是工程上不应依赖这个默认行为：
    - 可能中途退出循环
    - 可能没有遍历完
    - 可能有多结果集
- 最佳实践是永远显示使用 `defer rows.Close()`
- 调用多次不会出错，关闭前后 `Err()` 的结果是一致的

### Err()

在行迭代过程中发生的错误，例如：

- 网络中断
- 驱动解码错误
- 数据读取失败

> 这些错误不一定发生在 `Query()`，也不一定发生在 `Scan()` 可能发生在第 N 次 `Next()`

```
func (rs *Rows) Err() error
```

必须在循环后检查 `Err()`：

```
for rows.Next() {
    rows.Scan(...)
}
```

- 当 `Next()` 返回 false 时可能是正常结束，也可能是发生了错误
- 无法仅靠 `Next()` 判断
