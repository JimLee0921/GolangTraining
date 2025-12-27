# sql.TxOptions

`sql.TxOptions` 是 `sql.Tx` 也就是事务的启动参数包，当不想用默认事务行为时，用它显式声明规则。
只在开启事务的那一刻生效，用于告诉数据库这个事务的隔离级别和是否可读。

```
type TxOptions struct {
	// Isolation is the transaction isolation level.
	// If zero, the driver or database's default level is used.
	Isolation IsolationLevel
	ReadOnly  bool
}
```

在使用 `DB.BeginTx` 和 `Conn.BeginTx` 开启事务是传入，`DB.Begin` 使用的是默认事务设置

## 字段概述

### `Isolation IsolationLevel`

事务隔离级别，如果 `Isolation = 0` 则表示使用数据库/driver 默认隔离级别

参考 [TypeIsolationLevel.md](TypeIsolationLevel.md)

> `DB.BeginTx` 中驱动程序可能支持多种隔离级别。如果驱动程序不支持给定的隔离级别，则可能会返回错误

### `ReadOnly bool`

是否开启只读事务：

- `ReadOnly = true`：告诉数据库这个事务不会写数据，数据库可以做
    - 优化
    - 路由
    - 权限校验
- `ReadOnly = false`：默认值，普通读写事务
- 是否真的禁止写需要取决于数据库和 driver，有些强制拒绝写操作，有些只是作为优化提示

## 作用边界

`TxOptions` 只作用于：

- 事务开始时
- 数据库内部的事务模式

不影响：

- `tx.Exec / Query` 的用法
- `tx.Commit / Rollback` 的行为
- 事务生命周期
- 连接池策略