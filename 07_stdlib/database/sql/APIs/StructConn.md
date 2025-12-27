# sql.Conn

`sql.Conn` 是一个相对底层、偏高级的概念，不是为日常 CURD 使用的，而是为必须绑定到同一条物理连接的特殊场景提供能力

```
type Conn struct {
	// contains filtered or unexported fields
}
```

`sql.Conn` 表示的是数据库连接池中一条物理连接，在获取后接下来在这个连接上执行的所有操作一定发生在同一条底层连接上

## 获取连接

使用 `func (db *DB) Conn(ctx context.Context) (*Conn, error)` 进行获取：

```
conn, _ := db.Conn(ctx)
conn.Query(...)
conn.Exec(...)
```

- 显式占用连接池中的一条连接
- 所有操作强制走这一条连接
- 使用完毕后必须调用 `conn.Close`，否则可能导致连接泄露

## 特点

1. 除非明确需要连续使用同一条连接，否则永远不要使用 `sql.Conn`，这相当于绕过连接池，不安全
2. 可以使用 `sql.Tx`，也是单连接且保证事务性
3. 必须最后调用 `conn.Close`，这是硬性要求
4. 可以在另外一个 goroutine 中调用 `conn.Close`，这是并发安全的
5. 在调用 `conn.Close()` 后再使用 `conn.*` 必定返回 `sql.ErrConnDone`

## 适用场景

1. 必须绑定连接的数据库语义
2. 驱动级别状态/会话变量
3. 在实现数据库驱动或中间件

## 对比 Tx

`sql.Tx` 内部就是绑定了一条连接，是携带了事务语义的 Conn，能用 Tx 就不要用 Conn，Conn 是 Tx 的底层

| 对象     | 是否独占连接 | 是否有事务 |
|--------|--------|-------|
| `DB`   | 否      | 否     |
| `Conn` | 是      | 否     |
| `Tx`   | 是      | 是     |

## 核心方法

SQL 语义上和 `sql.DB` 的这些方法基本一样，差别在于：

- `db.*`：每次调用可能从池里拿不同连接（不保证同一条连接）
- `conn.*`：所有调用都固定走这一条连接（直到 `conn.Close()` 归还）

Conn 的执行类方法主要有：`ExecContext` / `PingContext` / `PrepareContext` / `QueryContext` / `QueryRowContext`
，因为较低层不咋用，都直接使用最安全版本，没有设置普通版本

### ExecContext()

执行不返回行的语句（INSERT/UPDATE/DELETE/DDL 等）

```
func (c *Conn) ExecContext(ctx context.Context, query string, args ...any) (Result, error)
```

### QueryContext()

执行返回多行的查询（典型 SELECT）

```
func (c *Conn) QueryContext(ctx context.Context, query string, args ...any) (*Rows, error)
```

返回 `*Rows`，需要 `defer rows.Close()`，逐行 `rows.Next()/Scan()`

### QueryRowContext

期望最多一行结果（或只关心第一行结果）

```
func (c *Conn) QueryRowContext(ctx context.Context, query string, args ...any) *Row
```

### PrepareContext()

创建预编译语句 `*Stmt`，之后可以多次执行

```
func (c *Conn) PrepareContext(ctx context.Context, query string) (*Stmt, error)
```

> PrepareContext 的 ctx 只用于准备阶段，不控制后续执行

### PingContext()

验证连接是否还活着

```
func (c *Conn) PingContext(ctx context.Context) error
```

- `db.PingContext`：更像验证数据库整体可用性/池可用性
- `conn.PingContext`：验证这条物理连接本身可用

### BeginTx()

在这条 Conn 所绑定的物理连接上，开启一个事务

- Tx 事务本身必然需要绑定到一个单连接上
- `Conn.BeginTx` 等价于已经手里拿着一条连接，现在在这条连接上开事务
- `db.BeginTx`等价于从连接池里挑一条连接，然后在那条连接上开事务

```
func (c *Conn) BeginTx(ctx context.Context, opts *TxOptions) (*Tx, error)
```

**参数**

- `ctx context.Context`：用于控制超时/取消，贯穿整个事务生命周期
    - 一旦 `ctx.Done`（超时/取消），`database/sql` 会自动执行 Rollback，`Tx.Commit` 返回错误
    - context 是事务的生命线
- `opts *TxOptions`：可以为 nil，就是使用默认事务配置

> Tx 结束后，连接仍然属于 Conn，必须再调用 `conn.Close()`

### Close()

归还连接给连接池，注意是归还，而不是关闭，是否关闭由连接池决定

```
func (c *Conn) Close() error
```

- 并发安全的，可以并发调用，会堵塞
- 直到正在进行的 `Query/Exec/Scan` 完成
- 不需要自己加锁，`database/sql` 会保证生命周期安全
- 调用 `Close()` 之后，所有 conn 后续操作都会立即失败，返回 `sql.ErrConnDone`

**最佳关闭顺序**

1. `ctx.cancel()`
2. `conn.Close()`

可以主动中断阻塞中的 SQL ，更快释放连接

### Raw()

驱动级别方法，开发基本不可能使用，临时暴露底层 driver 的原始连接对象，也就是停止抽象，直接拿到底层对象

```
func (c *Conn) Raw(f func(driverConn any) error) (err error)
```