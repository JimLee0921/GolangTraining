# sql.DBStats

`sql.DBStats` 使用 `stats := db.Stats()` 获取，是 db 连接池的实时快照+累计计数器

- 一部分字段是当前状态
- 一部分字段是从进程启动到现在的连接累计值
- 主要用于判断连接池当前状态与检查潜在问题

## 字段定义

```
type DBStats struct {
	MaxOpenConnections int // Maximum number of open connections to the database.

	// Pool Status
	OpenConnections int // The number of established connections both in use and idle.
	InUse           int // The number of connections currently in use.
	Idle            int // The number of idle connections.

	// Counters
	WaitCount         int64         // The total number of connections waited for.
	WaitDuration      time.Duration // The total time blocked waiting for a new connection.
	MaxIdleClosed     int64         // The total number of connections closed due to SetMaxIdleConns.
	MaxIdleTimeClosed int64         // The total number of connections closed due to SetConnMaxIdleTime.
	MaxLifetimeClosed int64         // The total number of connections closed due to SetConnMaxLifetime.
}
```

### 配置上限

1. `MaxOpenConnections int`：当前 DB 允许的最大打开连接
    - 通过 `db.SetMaxOpenConns(n)` 进行设置，不设置默认值为 0 也就是无限制
    - 主要用于对照 OpenConnections，判断是否连接上限

### 连接池当前状态

1. `OpenConnections int`：当前已经建立的物理连接总数
    - 等价于 `InUse + Idle`
    - 包括正在用的和空闲的，不包括已经关闭的
2. `InUse`：当前正在被 goroutine 使用的连接数量
    - 正在执行 SQL 的连接活被 `Tx / Conn` 占用着的连接
    - 如果这个值长期接近于 `MaxOpenConnections` 说明连接池可能偏小或某些 SQL 命令执行世界过长
3. `Idle`：当前空闲、可以立即被取走复用的连接数量
    - 执行完 SQL、已经归还到连接池等待下一次使用
    - 如果这个值长期很大说明 `MaxIdleConns` 可能设置太大，存在资源浪费问题

> OpenConnections = InUse + Idle

### 等待与阻塞

这两个是性能信号相关的字段

1. `WaitCount int64`：表示有多少次需要连接但是拿不到，被迫等待
    - 每当连接池中没有 idle 连接且 `OpenConnections == MaxOpenConnections` 时
    - 来了新的请求就会被阻塞直到有了空闲连接进行分配，并且 `WaitCount + 1
    - WaitCount 不断增长表明连接池压力过大，需要调参或优化 SQL

2. `WaitDuration time.Duration`：所有等待连接的累计阻塞时间
    - 是阻塞时间的总和而不是最大值，通常和 `WaitCount` 一起观察判断
    - 平均等待时间 = `WaitDuration / WaitCount`，平均等待时间过大说明连接池明显不够用

### 连接被关闭原因统计字段

这三个都是累计计数器，只增不减

1. `MaxIdleClosed int64`：因为空闲连接数量超限而被关闭的连接数量
    - 对应的设置为 `db.SetMaxIdleConns(n)`
    - 这个值不断增长说明空闲时间总是被关掉，可能因为设置的 `MaxIdleConns` 等原因

2. `MaxIdleTimeClose int64`：因为空闲时间过长而被关闭的连接数
    - 对应的设置为：`db.SetConnMaxIdleTime(d)`
    - 这是一个健康信号，说明旧连接正在被周期性请理
    - 如果为 0 说明没有设置 IdleTime，也就是默认不关闭空闲时间长的连接或流量很频繁

3. `MaxLifetimeClosed int64`：因为连接寿命到期而被关闭的连接数量
    - 对应的设置为：`db.SetConnMaxLifetime(d)`
    - 重要健康指标，持续增长说明连接在进行轮换
    - 如果长期为 0 说明可能没有设置 Lifetime，存在一定风险

## 连接生命周期

在 Go 的 `database/sql` 中，连接是可复用的资源，而不是一次性对象，生命周期由连接池统一管理

1. 不存在（Not Created）：此时没有连接
    - 程序刚启动
    - `sql.Open()` 只创建了 DB 句柄
    - 没有任何真实连接

2. 创建（Created）：`[Conn#1 | Created -> InUse]`
    - 触发为下面条件之一
        - 第一次调用 `Exec` / `Query` / `Ping`
        - 连接池里没有 idle 空闲连接
        - `OpenConnections < MaxOpenConns`
    - 会建立 TCP 连接 -> 完成认证 -> 创建 `driver.Conn`

3. 使用中（InUse）：`[Conn#1 | InUse]`
    - 发生在执行 SQL 或等待数据库响应：
        - `db.Exec / Query`
        - `tx.Exec / Query`
        - `conn.Exec / Query`
    - 被一个 goroutine 独占，此时不能被复用

4. 释放占用（BackPool）：`[Conn#1 | Idle]`
    - 当第一次操作结束时进行释放，注意不是 Close 关闭连接，而是放回池子中等待下一次被调用
    - 普通 SQL：SQL 执行完毕
    - Query：`rows.Close()` 被调用
    - Tx：执行了 `Commit` / `Rollback`
    - Conn：`conn.Close()`

5. 空闲（Idle）：`[Conn#1 ]`
    - 这是被使用过的连接重新被放回池子后的状态
    - 连接仍然是可用状态，TCP 依然保持
    - 等待下一次被复用

6. 决定是否被关闭：每当连接准备被复用或放回池子时，连接池会检查：
    - 空闲实现 IdleTime 检查：`db.SetConnMaxIdleTime(d)`
        - 如果连续空闲 > d
        - 会被标记为已过期
        - 下次复用前会被关闭
    - 生命周期 Lifetime 检查：`db.SetConnMaxLifetime(d)`
        - 从连接的创建时间开始累计
        - 如果超过 d 会被标记为茨连接已过期
        - 如果是 idle 连接直接关闭，如果是 inuse 连接使用完再关闭
    - 数量限制 MaxIdleConns 检查：`db.SetMaxIdleConns(n)`
        - 如果 idle 连接多余设置的最大数，将多余的进行关闭
    - > `database/sql` 不会暴力中断正在工作的连接

7. 关闭（Closed）：`[Conn#1 | Closed]`
    - TCP 断开
    - 资源释放
    - 永远不可复用

### 三种连接占用方式

1. 普通 DB 操作（最常见）：`db.Exec / Query / QueryRow / ...`
    - 临时占用连接
    - 执行完毕自动归还
    - 生命周期最短
2. 事务 Tx 操作：`tx, _ := db.BeginTx(...)`
    - 独占一条连接
    - 多条 SQL 在同一连接进行处理
    - 直到 `Commit` / `Rollback`
3. 显式连接 Conn：`conn, _ := db.Conn(ctx)`
    - 显式占用
    - 直到 `conn.Close()`
    - 用于会话级需求

### 补充

1. 释放不等于关闭，释放是将连接放回池子变为 idle 状态
2. IdleTime 每次使用会进行重置，Lifetime 不会重置
3. 连接池只在边界点做关闭决策