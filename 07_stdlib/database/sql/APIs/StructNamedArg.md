# sql.NamedArg

`sql.NamedArg` 是用于给 SQL 语句传递命名参数（named parameters）的工具（需要看数据库和驱动支不支持命名占位符）

```
type NamedArg struct {

	// Name is the name of the parameter placeholder.
	//
	// If empty, the ordinal position in the argument list will be
	// used.
	//
	// Name must omit any symbol prefix.
	Name string

	// Value is the value of the parameter.
	// It may be assigned the same value types as the query
	// arguments.
	Value any
	// contains filtered or unexported fields
}
```

## 字段

### Name

参数名（不带参数），比如：`id`、`name`，不是 `@id`、`:id`

### Value

和普通 `Exec/Query` 传入的参数完全一样，支持

- 基础类型
- `sql.Null[T]`
- 实现了 `driver.Valuer` 的类型

## 解决问题

NameArg 可以用名字而不是位置来给 SQL 绑定参数

```
// 位置参数（传统）
db.Exec("UPDATE student SET name=?, phone=? WHERE id=?", name, phone, id)

// 命名参数
db.Exec(
    "UPDATE student SET name=:name, phone=:phone WHERE id=:id",
    sql.Named("name", name),
    sql.Named("phone", phone),
    sql.Named("id", id),
)
```

> `:id`、`@id`、`?`、`$1` 这种位置参数和命名参数占位符语法都是数据库+驱动决定的

1. 解决参数顺序错误的风险
2. 提升可读性
3. 方便复用和重构

## `sql.Named` 构建方式

`sql.Named` 是用来构建命名参数的便捷函数，用于把参数名+参数值绑定到 SQL 里的命名占位符

```
func Named(name string, value any) NamedArg
```

只是个语法糖，等价于手写 `NamedArg{}` 但是更短，更清晰：

```
sql.Named("id", 1)
// 等价于
sql.NamedArg(Name: "id", Value: 1)
```