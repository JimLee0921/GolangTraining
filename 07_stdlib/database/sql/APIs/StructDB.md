# sql.DB

`sql.DB` 是一个数据库句柄，代表着 0 个或多个底层连接组成的池子，相当于数据资源管理中心

`sql.DB` = 数据库句柄（database handle） + 连接池（connection pool） + 并发协调器

## 底层原理

### 懒加载

`*seq.DB` 不是连接而是一个连接池管理其，刚 `sql.Open` 时是 0个连接连接，只有在第一次 `Query / Exec / Ping` 时才会创建新的连接

```
db, _ := sql.Open("mysql", dsb) 

// 此时没有任何 TCP 连接
```

### 并发安全

`*sql.DB` 是并发安全的，可以被多个 goroutine / 多个 HTTP 请求 / 多个后台 worker 同时使用

> 一个进程中至英国有一个 `*sql.DB` 实例

这就是为什么 `sql.DB` 通常：

- 全局初始化
- 注入到 service / repository
- 生命周期 = 进程生命周期

### 创建和关闭

不需要手动进行创建/关闭连接，`database/sql` 会：

- 自动创建连接
- 自动回收连接
- 维护一个空闲连接池（idle pool）

**机制**

- Query 需要连接就从池子中获取，使用完毕放回 idea pool
- idle pool 中连接太多会被回收
- 超过 MaxLifetime 会被关闭

所以只需要配置策略：

```
db.SetMaxIdleConns(10)
db.SetMaxOpenConns(100)
db.SetConnMaxLifetime(time.Hour)
...
```

### 连接级状态

per-connection state 指的是绑定在某一个连接上的状态：

- Mysql
    - `SET SESSION ...`
    - 临时表
    - `LAST_INSERT_ID()`
- PostgreSQL
    - `SET LOCAL`
- SQLite
    - PRAGMA

如果需要可靠地观察或依赖连接级状态，必须使用 `TX`(事务) 或 `Conn`(显式绑定连接)

> 不能指望 `db.Query()` 每次用的是同一个连接

### 事务

事务 = 独占一个连接：

- 从 `DB.Begin()` 开始：`TX` 事务中的所有 `Exec / Query` 都在同一条物理连接上
- 保证了 隔离级别 / 会话变量 / 临时表 / 锁
- 使用 `Tx.Commit` 或 `Tx.Rollback` 是为了归还事务，将连接归还到连接池中
- 如果忘了 `Commit` 或 `Rollback` 永远不会归还连接池，最终可能导致连接池耗尽

标准做法：

```
tx, er := db.Begin()
if err != nil { ... }

defer tx.Rollback() // 完全兜底

// ... do work ...
err = tx.Commit()
```

## 创建 `sql.DB` 方式

一共有两种创建方式： `Open` / `OpenDB`，但是使用上大大不同

| 方法           | 面向对象          | 是否常用 | 核心定位               |
|--------------|---------------|------|--------------------|
| `sql.Open`   | 99% 应用开发者     | 是    | DSN -> DB（标准入口）    |
| `sql.OpenDB` | driver / 框架作者 | 基本不用 | 已有 Connector -> DB |

### `sql.Open`

连接一个数据库，创建一个 `*sql.DB` 数据库句柄，`database/sql` 不会自己连接任何数据库，必须借助数据库驱动（Go
标准库不能内置所有协议，所以如何连接数据库是交给了第三方 driver 进行处理）

引入 driver：`import _ "github.com/go-sql-driver/mysql"` 只有一个作用就是触发 `init()` 把 driver 注册到 `database/sql`

> https://golang.org/s/sqldrivers 维护的有 driver 列表

```
func Open(driverName, dataSourceName string) (*DB, error)
```

#### 参数

- `driverName string`：
    - 数据库驱动名字，也就是在引入的 `import _ "xxx/driver"` 后注册的名字
    - 比如：`go-sql-driver/mysql` 就是 `"mysql"`
    - 比如：`lib/pq` 就是 `"postgres"`
    - 比如：`mattn/go-sqlite3` 就是 `"sqlite3"`
- `dataSourceName string`：DSN 是完全由 driver 解析的字符串，标准库不解析 DSN
    - 比如：MYSQL：`user:pass@tcp(host:3306)/db`
    - Postgres：`postgres://user:pass@host/db`

#### 注意事项

- `sql.Open` 不保证已连接数据库 / 不保证数据库存在 / 不保证用户名账号密码正确 / 只是 driver 存在且可能 DSN 形式合法
- 在 Open 建议的是使用 `Ping()` 进行测试：
    - 会触发真实连接创建
    - 从连接池拿一个连接
    - 执行 driver 的健康检查
    - 验证 网络 / 认证 / 数据库是否存在
- 正确使用是将 DB 注册为 全局或单例：`var DB *sql.DB`
- `db.Close()` 是禁止新连接、等待已有连接归还并关闭整个连接池
    - 只有在程序即将推出或测试用例结束时才会 Close
    - Web 服务 / 常驻进程通常不需要主动 Close

主要是为了：

- 避免服务启动时因为 DB 不可用导致直接 crash
- 支持懒加载、延迟连接
- 增强应用健壮性

#### 返回值

返回的 `*sql.DB` 是 `goroutine-safe`的：

- 内部做了完整的同步控制
- 可以在任意 goroutine 中使用同一个 DB
    - `go db.Query(...)`、`go db.Exec(...)`
- `sql.DB` = 连接池管理器，其中维护了：
    - idle connections
    - open connections
    - in-use connections
    - 所有的连接池参数都挂在 DB 上可以手动设置

### `sql.OpenDB`

`sql.OpenDB` 是用代码对象代替 DSN 字符串来创建 DB，需要传入 `driver.Connector`

```
func OpenDB(c driver.Connector) *DB
```

因为 DSN 只能表达 用户名 / 密码 / host / 参数，不能表达较为复杂的

- 动态 Token（每次连接不同）
- IAM / 云数据库临时凭证
- 自定义连接生命周期
- 复杂初始化逻辑

使用 Connector 就是把怎么连接数据库完全交给代码控制，使用这个 API 的目标用户是：
driver 作者 / 框架作者 / 中间件作者

和 Open 一样也是懒连接：

- 不保证已连接
- 不保证数据库可用
- 只创建 DB 句柄 + 连接池

# 主要方法

`sql.DB` 的方法主要分为下面这些

## 事务控制

主要用于开启事务 Transaction，返回 `*sql.Tx` 事务绑定单连接

- `Begin()`
- `BeginTx()`

### Begin() / BeginTx()

> 具体事务是干嘛的在 [StructTx.md](StructTx.md) 中讲解

Begin / BeginTx 用于开启一个数据库事务。返回 `*sql.Tx`，一个事务等同于：

- 绑定一条物理连接
- 再 `Commit` / `Rollback` 之前
- 所有操作都在这条连接上完成

```
func (db *DB) Begin() (*Tx, error)
func (db *DB) BeginTx(ctx context.Context, opts *TxOptions) (*Tx, error)
```

**Begin 对比 BeginTx**

- Begin
    - 内部等价于 `BeginTx(context.Background(), nil)`
    - 不可取消、不可超时
    - 使用默认事务选项
- BeginTx
    - 支持 context
    - 事务级别隔离
    - 配合部分 driver 可以进行只读事务

> 生产环境优先使用 BeginTx


**参数**

- `ctx context.Context`：这个 context 的生命周期等于整个事务的生命周期
    - 从 BeginTx 开始直到 `Commit` 或者 `Rollback`
    - ctx 超时或者取消，`database/sql` 会自动执行 Rollback
    - 并且 `Tx.Commit()` 会返回 error
- `opts *TxOptions`：可选，参考 [StructTxOptions.md](StructTxOptions.md)
    - 因为默认的隔离环境不是 Go 决定的，而是数据库 / driver 决定的
    - Go 不支持模拟隔离级别，如果 driver 不支持会直接报错

**返回值**

- `*TX`：`*sql.Tx`，一个绑定单连接的事务上下文
    - 使用它的 `Exec` / `Query` 等方法时不要在使用 `db.Exec` 等方法，逻辑上比较混乱
- `error`：通常遇到下面这些情况会返回错误
    - 无法从连接池获取连接
    - driver 不支持指定的隔离级别
    - context 已经被 cancel
    - 数据库拒绝开启事务

**注意事项**

1. 事务等于独占一条连接：
    - Begin 成功后从连接处拿走一条连接标记为 InUse
    - 在 Commit / Rollback 之前不会归还连接池
2. Commit / Rollback 是释放练级的唯一方式
3. 不调用 Commit / Rollback 可能导致：
    - 连接泄露
    - 连接池耗尽
    - 后续 Begin 堵塞

## 连接获取与驱动信息

Connection / Driver，不常用，了解就行

- `Conn()`：从池里租用一条底层连接，返回 `*sql.Conn`（需要手动 `Close()` 归还池）。
- `Driver()`：拿到当前 DB 使用的 driver（多用于框架/诊断）

### Driver()

返回当前 `*sql.DB` 背后使用的 `driver.Driver` 实例

```
func (db *DB) Driver() driver.Driver
```

- `database/sql` 是中间层
- 真正和数据库打交道的是 driver
- `Driver()` 只是看一眼 driver 信息，而不是操作 driver，几乎不会使用

### Conn()

Conn 用于显式占用一条物理连接，并在这条连接上执行一系列操作，和普通的 `db.Exec()` 本质不同

```
func (db *DB) Conn(ctx context.Context) (*Conn, error)
```

调用 `db.Conn()`：

1. 从连接池中去除一条 idle 连接或新建一条连接
2. 标记为 InUse
3. 在调用 Close 之前，这条连接不会被其它人使用

**参数**
`ctx context.Context`：

- 如果没有 idle 连接且已达到 MaxOpenConns
- 会阻塞等待直到有链接归还或有 ctx 被 cancel / 超时

**返回值**

`*sql.Conn`：注意 `Conn.Close()` 的真正含义不是关闭数据库，而是归还连接池

- 如果连接未过期变成 idle
- 如果已经过期才是真正的 Close

## Sql 执行与查询

Execution / Query，最常见的 CRUD 入口，Context 版本是支持超时和取消

1. 写操作 / 无结果集

    - `Exec()`
    - `ExecContext()`

2. 读操作 / 多行结果集

    - `Query()`
    - `QueryContext()`

3. 读操作 / 单行结果（延迟报错）

    - `QueryRow()`
    - `QueryRowContext()`

### Exec() / ExecContext()

Exec 系列用于执行不需要返回结果集合的 SQL 语句，一定不要用于 SELECT 主要使用场景有：

- INSERT / UPDATE / DELETE 数据操作
- CREATE / DROP / ALTER 表操作
- TRUNCATE
- 调用不返回行的存储过程

Exec 和 ExecContext的差异在于：

**Exec：**

- 内部等价于 `ExecContext(context.Background(), ...)`
- 不可设置取消和超时

**ExecContext：**

- 由调用方控制 context 上下文
- 支持 超时 / 取消 / 链路传递

> 生产代码优先使用 ExecContext

```
func (db *DB) Exec(query string, args ...any) (Result, error)
func (db *DB) ExecContext(ctx context.Context, query string, args ...any) (Result, error)
```

**参数**

- `query string`：
    - 要执行的 SQL 语句
    - 占位符由 driver 决定
- `args ...any`：query 中占位符的绑定参数
    - 按照顺序进行填充占位符，args 个数必须和占位符数量匹配
    - 会被 driver 做转义 / 编码
    - 主要用于避免 SQL 注入
    - 类型需要是 driver 支持的基础类型（int / string / time.Time 等）
- `ctx context.Context`：ExecContext 支持的上下文

**返回值**

- `error`：如果执行失败返回失败信息，可能为：
    - SQL 语法错误
    - 参数数量/类型不匹配
    - 连接失败
    - 约束冲突（唯一键、外键）
    - 上下文取消（ExecContext）
- `Result`：参考 `sql.Result` 接口
    - RowsAffected 可以获取为 UPDATE / DELETE 等受影响行数
    - LastInsertId 可以获取最后一次插入生成的自增 ID（不是所有数据库都支持）

**注意事项**

1. Exec 系列不返回 Rows，用于不要用于查询
2. 不要拼接 SQL 字符串，而是使用占位符
3. Result 的方法可能返回 error
4. ExecContext 可能因为 ctx 失败
5. Exec 不等于单条连接，每次 Exec 都是从连接池获取执行并返回给池子，不能保证两次 Exec 在同一连接使用 `Tx` 或 `Conn` 获取同一连接

### Query() / QueryContext

Query 系列用于执行多行结果集的 SQL，通常是 SELECT：

- 列表查询
- 聚合查询（多行）
- Join 查询

**Query vs QueryContext**

- Query 内部使用 `context.Background()`，不可取消、不可超时
- QueryContext：由调用方控制 context

> 生产代码应优先使用 QueryContext

```
func (db *DB) Query(query string, args ...any) (*Rows, error)
func (db *DB) QueryContext(ctx context.Context, query string, args ...any) (*Rows, error)
```

**参数**

- `query string`：SQL 语句，通常是 SELECT 语句，占位符由 driver 决定
- `args ...any`：
    - 按照顺序进行填充占位符，args 个数必须和占位符数量匹配
    - 会被 driver 做转义 / 编码
    - 主要用于避免 SQL 注入
    - 类型需要是 driver 支持的基础类型（int / string / time.Time 等）
- `ctx context.Context`：QueryContext 支持的上下文

**返回值**

- `*Rows`：参考 `DB.Rows`，一个游标式的结果集合迭代器
    - 可能返回 0 行，但不是错误
    - 需要检查 `rows.Err()`
- `error`：错误信息

### QueryRow() / QueryRowContext()

QueryRow 系列用于最多返回一行的查询：

- 主键查询
- 唯一索引查询
- `SELECT COUNT(*)`
- `SELECT MAX(...)`

QueryRowContext 可用通过传入自定义上下文进行控制超时/取消等行为

```
func (db *DB) QueryRow(query string, args ...any) *Row
func (db *DB) QueryRowContext(ctx context.Context, query string, args ...any) *Row
```

**参数**

- `query string`：SQL 语句，通常是 SELECT 语句，占位符由 driver 决定
- `args ...any`：
    - 按照顺序进行填充占位符，args 个数必须和占位符数量匹配
    - 会被 driver 做转义 / 编码
    - 主要用于避免 SQL 注入
    - 类型需要是 driver 支持的基础类型（int / string / time.Time 等）
- `ctx context.Context`：QueryRowContext 支持的上下文

**返回值**

`QueryRow / QueryRowContext` 永远返回非 nil 的 `*Row`，真正的错误比如 SQL 错误 / 连接错误 / 参数错误 全部延迟到 `Scan()`：

```
err := row.Scan(&v)
if err == sql.ErrNoRows {
    // 没有数据
}
```

> ErrNoRows 不是异常是业务分支条件

如果 SQL 实际返回是多行，则只会读取第一行，其它的直接丢弃

## 预编译语句

Prepared Statements，用于创建 `*sql.Stmt` 以复用 SQL（可减少重复解析/编译成本；是否真正受益取决于 driver/数据库），Context
版本是支持超时和取消

- `Prepare()`
- `PrepareContext()`

### Prepare() / PrepareContext()

Prepare 用于提前准备一条 SQL 模板，然后在之后反复、高并发地执行这条 SQL

```
func (db *DB) Prepare(query string) (*Stmt, error)
func (db *DB) PrepareContext(ctx context.Context, query string) (*Stmt, error)
```

**Prepare 对比 PrepareContext**

- Prepare：
    - 内部使用 `context.Background()`
    - 准备阶段不可取消
- PrepareContext：
    - 使用自定义提供的 ctx 可以用于超时/取消
    - 只影响准备阶段，不会影响后续执行（只会控制创建 Stmt 这一步，而不会自动作用到 `stmt.Exec` / `stmt.Query`）

**用途**

- 同一条 SQL 会被多次执行
- 需要不同参数但是相同的结构
- 需要减少重复解析/编译成本
- 需要更清晰的 SQL 生命周期管理

**参数**

- `query string`
    - 一条 SQL 模板
    - 包含占位符（由 driver 决定）
    - 不能是多条 SQL

**返回值**

- `error`：发生错误时返回
- `*sql.Stmt`：一个可并发使用的 SQL 执行器，绑定了 SQL 结构但不绑定具体参数，参考 [StructStmt.md](StructStmt.md)
    - 同一个 `*Stmt` 可以被多个 goroutine 同时使用
    - 内部自行处理并发安全问题
    - 执行上和 DB 一样使用 `Exec` / `Query` / `QueryRow` 等方法

**注意事项**

- Stmt 是一种有生命周期的资源，使用完毕必须调用 Close，如果忘记调用会导致服务器 prepared statement 不释放、driver 内部泄露等问题
- DB.Prepare 返回的 Stmt 是 DB 级别的
    - 不是绑定单一连接，而是在每条需要的连接上按需创建对应的 prepared statement
    - 由 `database/sql` 进行管理，所以 `*Stmt` 可以并发使用
- Prepare 不一定提升性能，取决于数据库和 driver

**适用场景**

- 高频、固定 SQL
- 热路径
- 框架 / DAO 层
- 批量处理
- 不适用于一次性 SQL / SQL 动态拼接 / 简单 CRUD

## 健康检查与生命周期

Health / Lifecycle

- `Ping()/PingContext()`：触发实际连通性验证（以及必要时建连）
- `Close()`：关闭整个 DB（连接池），通常只在进程退出/测试结束时调用

### Ping() / PingContext()

Ping 系列用于确认数据库现在能不能连接上，能不能用，会做两件事：

1. 如果还没有建立连接：创建一个真实的数据库连接
2. 如果已经有空闲连接：复用一条连接并做存活性检查

```
func (db *DB) Ping() error
func (db *DB) PingContext(ctx context.Context) error
```

- 主动触发真实连接，验证 网络 / 认证 / 数据库可用性
- 不是业务查询，不会返回任何数据

**Ping 对比 PingContext**

- Ping：
    - 内部使用 `context.Background()`
    - 不可取消，不可设置超时
- PingContext：
    - 可以传入自定义的 context
    - 支持超时/取消/服务优雅关闭

> 生产环境推进使用 PingContext


**使用场景**

1. 应用启动时调用确保 DSN 正确 / 网络可达 / 账号密码凭证有效
2. 健康检查接口

**注意事项**

- Ping 只验证当前时刻
- 不保证下一秒连接还是可用状态
- 网络 / DB 都可能随时变化

### Close

用于关闭 DB 句柄，主要做三件事：

1. 阻止新的查询：新的 `Query / Exec` 会直接报错
2. 等待已经开始的查询完成：不会强制关闭正在执行的 SQL
3. 关闭连接池中所有连接

```
func (db *DB) Close() error
```

**注意事项**

- `*sql.DB` 的设计目标是：进程级别、长期存在、全局共享
- Web服务、后台 Worker、长期运行进程通常不需要主动调用 `Close()`

**适用场景**

- 程序即将结束：main 主 goroutine 返回前
- 测试用例结束
- 短生命周期工具：CLI、一次性脚本

## 连接池配置与统计

Pool Tuning / Observability，控制连接池资源与稳定性，Stats 用于监控、压测与排障

1. 配置
    - `SetMaxOpenConns(n int)`：最大打开连接数上限
    - `SetMaxIdleConns(n int)`：最大空闲连接数
    - `SetConnMaxLifetime(d time.Duration)`：单连接最大存活时长
    - `SetConnMaxIdleTime(d time.Duration)`：单连接最大空闲时长
2. 统计
    - `Stats() DBStats`

### SetConnMaxIdleTime()

空闲连接的时间维度控制，设置连接空闲多久会被回收

```
func (db *DB) SetConnMaxIdleTime(d time.Duration)
```

- 设置的这个空闲时间 idle time 是指没有被任何 goroutine 使用的时间
- 超过这个时间的连接不会立刻被杀死，而是在下次要使用之前被懒回收
- 如果设置的 `d <= 0`，则连接不会因为空闲时间而关闭

### SetConnMaxLifetime()

连接的时间维度控制，一个连接最长能用多久，lifttime 指从连接创建开始计时，无论是否在工作，超过这个世界后都会被回收

```
func (db *DB) SetConnMaxLifetime(d time.Duration)
```

**IdleTime 对比 LifeTime**

两个可以同时设置

|        | IdleTime | Lifetime |
|--------|----------|----------|
| 关注点    | 闲着多久     | 活了多久     |
| 是否要求空闲 | 是        | 否        |
| 适用场景   | 减少闲置     | 防止长时间存活  |

### SetMaxIdleConns()

空闲连接的数量维度控制，控制最多保留多少空闲连接，控制用完之后不关闭，留着等待复用的连接数，默认只有两条

```
func (db *DB) SetMaxIdleConns(n int)
```

如果 `n <= 0` 表示每次用完都关闭不留缓存，主要原因非常低频访问或数据库连接十分昂贵的情况

### SetMaxOpenConns()

连接的数量维度控制，最多允许打开多少连接，整个连接池的硬上限，包括正在使用的和空闲中的

```
func (db *DB) SetMaxOpenConns(n int)
```

- 默认值 `n = 0` 表示无上限（很危险）
- 生产环境一定要设置

### Stats()

用于连接池观察与诊断，返回当前连接池的实时状态快照

```
func (db *DB) Stats() DBStats
```

返回的字段参考 [StructDBStats.md](StructDBStats.md)

- `WaitCount` 不断增长说明 MaxOpenConns 太小
- `Idle` 长期很大说明 MaxIdleConns 太大
- `OpenConnections` 接近 DB 上限说明有风险

OpenConnections 接近 DB 上限

**官方对于` Idle / Open` 的规则**

- `if MaxOpenConns > 0 but < MaxIdleConns`：MaxIdleConns 会被自动压缩
- `if MaxIdleConns > 0 but new MaxOpenConns < MaxIdleConns`：MaxIdleConns 会被同步降低

Idle 空闲连接数量永远不能超过 Open 连接数量

**生产级别设置参考**

实际值必须结合：

- 数据库最大连接数
- 服务实例数量
- 并发请求峰值

```
db.SetMaxOpenConns(50)
db.SetMaxIdleConns(10)
db.SetConnMaxLifetime(1 * time.Hour)
db.SetConnMaxIdleTime(30 * time.Minute)
```

