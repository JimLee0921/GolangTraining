# sql.Row

`Row` 是 `sql.DB/Conn/Stmt/Tx` 的 QueryRow 系列方法的返回值，用来表示期望至多返回一行结果，主要有两个特点：

```
type Row struct {
	// contains filtered or unexported fields
}
```

## 特点

### 至多返回一行

如果 QueryRow 系列获取多行结果只取第一行结果，如果获取到了 0 行则返回 0 行

### 惰性结果

QueryRow 返回的 `*Row` 结果是惰性的，错误信息也是延迟获取

`row := db.QueryRow("SELECT ...")` 此时数据库查询并没有真正完成，真正的生命周期或者说流程为：

1. 执行 `QueryRow()`
2. 返回一个 Row（轻量占位）
3. 调用 `Row.Scan()` 方法才真正读取数据 / 获取错误信息

## 核心方法

Row 一共两个方法：`Err` 和 `Scan`

### Scan()

Row 真正的执行入口，实现 `sql.Scanner` 接口

```
func (r *Row) Scan(dest ...any) error
```

在调用 `err := row.Scan(&a, &b, ...)` 会做三件事：

1. 真正触发查询结果的读取
2. 只取第一行，舍弃剩余行（如果返回多行）
3. 如果有错误返回错误信息

| 查询结果 | `Scan` 行为          |
|------|--------------------|
| 0 行  | 返回 `sql.ErrNoRows` |
| 1 行  | 正常 Scan            |
| >1 行 | 只取第一行，其余直接丢弃       |

> QueryRow 并不会检查唯一性

### Err()

用于包装场景，在不调用 Scan 的情况下检查如果调用 Scan 是否有错误，主要给框架/在基层使用，业务代码基本不用

```
func (r *Row) Err() error
```

- `Err()` 不会读取行数据
- 只返回查询结果已知的错误
- 直接使用 `err := row.Scan(...)` 就行了

## Row 对比 Rows

| 对比点  | Row      | Rows               |
|------|----------|--------------------|
| 返回行数 | 至多 1 行   | 0-N 行              |
| 错误暴露 | 延迟到 Scan | Query / Next / Err |
| 资源释放 | 自动       | 必须 `rows.Close()`  |
| 使用场景 | 精确查询     | 列表查询               |

> 使用 QueryRow 系列不用关心 `Close()`