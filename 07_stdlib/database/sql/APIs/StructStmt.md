# sql.Stmt

`sql.Stmt` 是一条已经已经准备好的 SQL 模板。用来被反复、高并发地执行，本身不执行 SQL，而是执行器的蓝图

```
type Stmt struct {
	// contains filtered or unexported fields
}
```

## 特点

### prepared statement

prepared statement 不绑定某一次执行，而是绑定 SQL 结构，也就是说：

- SQL 的语句结构（字段、表、占位符）
- 被提前固定
- 参数在执行时再进行传入
- 不一定使用用 Stmt 就有性能收益，取决于数据库与 driver

### 并发安全

Stmt 是共享对象，同一个 `*sql.Stmt` 可以被多个 goroutine 同时使用：

```
stmt, _ := db.Prepare("INSERT INTO t(a) VALUES (?)")

go stmt.Exec(1)
go stmt.Exec(2)
go stmt.Exec(3)
```

- 不需要加锁
- 不会串行
- `database/sql` 内部保证安全

### 不同 stmt 构建与生命周期

Stmt 准备在哪儿决定了它绑定在哪儿，stmt 有三种构建绑定方式：

1. `stmt, _ := tx.Prepare(...)`
    - Stmt 永久绑定这条屋里连接
    - 不能被迁移
    - 不能跨连接使用
    - 如果在 `tx.Commit` / `tx.Rollback()` 之后再调用 `stmt.Exec(...)` 会报错
    - Stmt 的生命周期 <= Tx / Conn 的生命周期
2. `stmt, _ := conn.Prepare(...)`
    - Stmt 永久绑定这条屋里连接
    - 不能被迁移
    - 不能跨连接使用
    - 如果在 `conn.Close()` 之后再调用 `stmt.Exec(...)` 会报错
    - Stmt 的生命周期 <= Tx / Conn 的生命周期
3. `stmt, _ := db.Prepare(...)`：最常用
    - 连接池级别的 Stmt
    - 不绑定单一连接，生命周期与 `*sql.DB` 相同，可以长期复用
    - DB 级别的 Stmt 是逻辑上的预编译，不是只在一条连接上的预编译
    - 第一次执行在 `Conn#1` 上 prepare
        - 下一次执行如果复用 `Conn#1` 就直接使用
        - 如果还 `Conn#2` 会自动在 `Conn#2` 上 prepare
        - 由 `database/sql` 自动完成

# 核心方法

和 db 方法差不多，具体注意事项什么的参考 db

## 执行类方法

这几个在执行语义上，与 DB / Tx 上的同名方法完全一致，区别是：SQL 已经在 Prepare 阶段固定，不再传 SQL 字符串，主要传入参数什么

参数和返回值都参考 `sql.DB` 的这些同名方法

### Exec() / ExecContext()

执行一条不返回行的已预编译 SQL

- `Exec`：不可取消
- `ExecContext`：支持超时 / 取消

生产环境优先使用 `ExecContext`

```
func (s *Stmt) Exec(args ...any) (Result, error)
func (s *Stmt) ExecContext(ctx context.Context, args ...any) (Result, error)
```

典型 SQL： `INSERT` / `UPDATE`/ `DELETE` / `DDL`

### Query() / QueryContext()

多行查询，执行一条返回多行结果集的已预编译 SQL

- `QueryContext` 的 ctx：
    - 控制执行阶段
    - 可超时 / 可取消
- `Query` 内部使用 `context.Background()`

```
func (s *Stmt) Query(args ...any) (*Rows, error)
func (s *Stmt) QueryContext(ctx context.Context, args ...any) (*Rows, error)
```

典型 SQL：`SELECT * FROM ...` / 多行聚合查询

### QueryRow() / QueryRowContext()

单行查询，执行一条最多返回一行的已预编译查询

```
func (s *Stmt) QueryRow(args ...any) *Row
func (s *Stmt) QueryRowContext(ctx context.Context, args ...any) *Row
```

典型 SQL：主键查询 / 唯一索引查询 / `COUNT(*)`

## 状态管理

就一个 Close 方法

### Close()

用于关闭 Stmt ，生命周期与创建方式有关系

```
func (s *Stmt) Close() error
```

1. tx 得到的 stmt：
    - Stmt 绑定单条事务连接
    - 调用 `stmt.Close()`：释放该连接上的 prepared statement
    - 如果没调用 stmt.Close：
        - 在 `tx.Commit()` / `tx.Rollback()` 时
        - Stmt 也会随事务一起失效

2. conn 得到的 stmt：
    - Stmt 永久绑定该 Conn
    - 调用 `stmt.Close()`：释放该连接上的 prepared statement
    - `conn.Close()`：会让 Stmt 直接失效

3. db 得到的 stmt：
    - 标记该 Stmt 为不可用
    - 通知 database/sql：以后不再允许用这个 Stmt 执行 SQL
    - 在所有曾经使用过的底层连接上：释放（或安排释放）对应的 prepared statement
    - 断开 Stmt 与 DB 的绑定关系
    - 不会关闭 DB / 不会关闭连接池 / 不会影响其他 Stmt

