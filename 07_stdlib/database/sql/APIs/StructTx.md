# sql.Tx

`sql.Tx` 表示一次正在进行中的数据库事务，是一个单独绑定单条物理连接的，具备 ACID：

- 原子性（Atomicity）
- 一致性（Consistency）
- 隔离性（Isolation）
- 持久性（Durability）

语义的执行上下文，注意 Tx 不是事务配置或者说事务描述，而是一个已经开始但尚未结束的事务实例

如果说 Conn 保证的是同一会话，Tx 就是在此基础上保证了 原子性 / 一致性 / 隔离性

> 一句话描述数据库事务：一组作为整体执行的数据库操作，要么全部成功提交，要么在发生错误时全部回滚，从而保证数据一致性和可靠性

## 特点

### 事务必须结束

事务等于独占连接，如果事务执行完毕后不进行结束：

- 连接永远不会归还给池子
- 后续请求可能会被堵塞
- 事务不是 GC 能够自动帮忙回收的资源

### Commit / Rollback

事务是一次性对象，一旦调用 `Commit` / `Rollback` 后：

- 事务状态变为 Done
- Tx 对象逻辑上已失效
- 再调用 tx 的任何方法都会 `return ErrTxDone`

### 和 Stmt 的关系

- 在 Tx 上 prepare 的 Stmt：生命周期 <= Tx
- 当事务结束：这些 Stmt 会被自动 Close
- 不需要、也不应该再手动 Close 它们

### 创建 tx 事务

有三种创建事务的方式，可以使用 Conn 或 DB：

- `func (c *Conn) BeginTx(ctx context.Context, opts *TxOptions) (*Tx, error)`
- `func (db *DB) BeginTx(ctx context.Context, opts *TxOptions) (*Tx, error)`
- `func (db *DB) Begin() (*Tx, error)`

Begin 使用的默认的 `context.Background` 和 默认的 `*TxOptions`，两个 BeginTx 可以自定义 ctx 和 `*TxOptions`
（参考 [StructTxOptions.md](StructTxOptions.md)）

在 `tx, err := db.BeginTx(...)` 之后：

1. 一条连接被独占
2. 必须保证一定会执行 `tx.Commit` 或 `tx.Rollback`
3. 这个对象只能用一次
4. 用完就丢，不能复用

# 核心方法

## 执行类方法

`Tx.Exec / Query / QueryRow` 在接口形态和返回值上，与 `DB.Exec / Query / QueryRow` 完全一致；
唯一、也是本质的区别是：它们一定在同一条事务连接上执行

注意这些执行方法的 Context 版本与 db 有些不一样：

- `Tx.ExecContext / QueryContext / QueryRowContext` 的 `ctx`
    - 不会只影响某一条 SQL
    - 如果 ctx 被 cancel：
        - 事务会被标记为失败
        - `Commit` 会返回 error
        - 通常应直接 Rollback

- Tx 中的 context 是事务级约束

### Exec() / ExecContext()

事务内写操作

- 执行不返回结果集的 SQL
- 必定属于当前事务
- 受事务隔离级别控制
- 提交前不可见（取决于隔离级别）

```
func (tx *Tx) Exec(query string, args ...any) (Result, error)
func (tx *Tx) ExecContext(ctx context.Context, query string, args ...any) (Result, error)
```

### Query() / QueryContext()

事务内多行查询

- 执行返回多行的查询
- 查询结果：
    - 受当前事务影响
    - 可见事务内尚未提交的变更（读己之写）

```
func (tx *Tx) Query(query string, args ...any) (*Rows, error)
func (tx *Tx) QueryContext(ctx context.Context, query string, args ...any) (*Rows, error)
```

**注意事项**

- 在事务中，如果用 `tx.Query()` / `tx.QueryContext()` 拿到了 `*Rows
- 在 Rows 被 `Close()` 之前
    - 这条连接仍然被认为正在使用中（InUse）
    - `Tx.Commit()` / `Tx.Rollback()` 无法真正完成连接释放。

# 四、Tx.QueryRow / Tx.QueryRowContext

事务内单行查询

```
func (tx *Tx) QueryRow(query string, args ...any) *Row
func (tx *Tx) QueryRowContext(ctx context.Context, query string, args ...any) *Row
```

- 永远返回非 nil `*Row`
- 错误延迟到 `Scan`
- 如果查询到 0 行返回`sql.ErrNoRows`
- 如果查询到多行只取第一行

> 读取的是事务视图里的数据

## 创建 Stmt 方法

一共有两种方法可以创建 Stmt ：`Tx.Prepare*` 和 `Tx.Stmt*`，但是有区别：

1. `Tx.Prepare`：在当前事务中创建一条新的 prepared statement
    - SQL 字符串自行提供
    - Stmt 绑定该事务的单连接
    - 事务结束后自动关闭
2. `Tx.Stmt*`：把一个已有的 Stmt 转换当前事务可用的 Stmt
    - 提供的 `*Stmt` 通常是由 `db.Prepare` 得到的
    - 返回一个事务专用的 `*Stmt`
    - 同样在事务接收后自动关闭

### Prepare() / PrepareContext()

在事务里 prepare 的 Stmt：

- 只能在该事务中使用
- 只能在该事务绑定的那条连接上执行
- 事务结束后（Commit/Rollback）后 Stmt 会自动 Close，再使用会报错

```
func (tx *Tx) Prepare(query string) (*Stmt, error)
func (tx *Tx) PrepareContext(ctx context.Context, query string) (*Stmt, error)
```

- PrepareContext 可以控制取消和超时
- 但是 PrepareContext 的 ctx 只控制创建 Stmt 的过程
- 后续执行仍然需要使用 `stmt.ExecContext/QueryContext` 才能控制执行的超时
- 但是执行时仍然处于事务上下文（同一连接、同一事务）

### Stmt() / StmtContext()

这两个把一个全局可用的 DB 级 Stmt，转换成当前事务专用的 Stmt，因为 DB 级别的 Stmt 不保证在事务连接上执行

- `db.Prepare` 得到的 stmt 是 DB 级别，可能在不同连接上执行
- 事务要求所有操作必须在事务那条连接上执行

```
func (tx *Tx) Stmt(stmt *Stmt) *Stmt
func (tx *Tx) StmtContext(ctx context.Context, stmt *Stmt) *Stmt
```

- `tx.StmtContext(ctx, stmt)` 的 ctx 仍然只用于适配/准备阶段而不是用于执行阶段
- 执行阶段要控制超时仍然需要用 `stmt.ExecContext/QueryContext`

> 这两个方法返回的事务专用 Stmt 不需要手动 Close（一般不需要），会在事务结束后自动失效

## 事务终止指令

`Tx.Commit` 和 `Tx.Rollback`，用于释放连接，结束事务，把连接放回连接池

注意这两个必须调用一个且只能调用一个且只能用一次：

- 必须调用其中一个
- 调用任意一个之后：
    - 再调用另外一个会返回 `ErrTxDone`
    - 所有的 `tx.*` 方法都会返回 `ErrTxDone`
- 因为 Tx 是一次性对象

### Commit()

提交事务，通知数据库永久提交事务中的所有修改

- 受数据库隔离级别与约束校验
- 释放事务独占连接
- 结束事务生命周期

```
func (tx *Tx) Commit() error
```

可能会失败：

- 违反约束（唯一键/外键）
- 死锁
- 网络中断
- context 被 cancel
- 数据库拒绝提交

### Rollback()

放弃事务，通知数据库撤销该事务中的所有修改：

- 释放事务独占的那条连接
- 结束事务生命周期
- Rollback 成功后会将连接归还连接池或关闭，`tx` 进入 Done 关闭状态
- 之后再使用 `tx.*` 方法返回 `ErrTxDone`

```
func (tx *Tx) Rollback() error
```

Rollback 是放弃操作，几乎不会失败，即使失败连接也会被标记不可用并回收

> 开发中推荐使用 `defer tx.Rollback()` 进行兜底