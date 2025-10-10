# Go 资源竞争

在并发程序里，多个 goroutine / 线程 同时访问和修改同一个资源（变量、文件、网络连接等），而访问顺序和执行时间不确定，最终结果也就变得不可预测，
当两个或以上的 goroutine 在没有正确同步的情况下，同时读写同一份可变内存，就会触发数据竞争（data race）。
症状：结果偶尔错、时好时坏、难复现；严重时直接 panic（比如并发写 map）

## 定位资源竞争

可以使用 go run 命令查看是否存在资源竞争（Go 的 race detector 底层依赖 C 代码（需要 cgo 支持）默认情况下，Windows 下安装的
Go 有时会把 CGO_ENABLED=0，导致 -race 用不了）

- go run main.go：
  编译并运行 main.go
  不会检查资源竞争
  程序结果可能是错的，但不会提示
- go run -race main.go
  这是带竞态检测（Race Detector）的运行方式：
  -race 会在编译时插入额外的检查逻辑
  程序运行时会监控内存访问，发现 多个 goroutine 并发访问同一个变量且至少有一个写操作，就会报 DATA RACE 错误

## 修复资源竞争

Go 中主要有下面四种方法来修复资源竞争问题

1. 互斥锁：`sync.Mutex` 或 `sync.RWMutex`
2. 使用 `sync/atomic` 原子操作

原子操作就是 不可分割的最小操作单元，要么全部完成，要么完全不做
atomic = 原子操作，atomic 提供了 轻量级、锁无关 的并发安全手段
在并发环境下，原子操作不会被打断，因此可以避免资源竞争
Go 里，sync/atomic 包提供了一些常用的原子操作函数，用于在多个 goroutine 并发访问共享变量时保证安全
以 int32 / int64 为例：

```go
读取 & 写入
atomic.LoadInt64(&x)   // 原子读取
atomic.StoreInt64(&x, 100) // 原子写入

// 加减
atomic.AddInt64(&x, 1)   // 原子地加 1
atomic.AddInt64(&x, -1) // 原子地减 1

// 交换
atomic.SwapInt64(&x, 200) // 把 x 设置为 200，并返回旧值

/*
   CAS（Compare And Swap）
   如果 x 当前的值等于 old，就把它改成 new，并返回 true；
   否则返回 false（说明有别的 goroutine 抢先修改过了）
   CAS 是实现很多并发安全算法的核心
*/
atomic.CompareAndSwapInt64(&x, old, new)
```

- 使用 channel 管道
- 只初始化一次：sync.Once


```

## 总结

| 场景               | 首选                     |
|------------------|------------------------|
| 简单计数/标志          | `sync/atomic`          |
| 结构化共享状态（多个字段要一致） | `sync.Mutex`/`RWMutex` |
| 队列/状态机、串行化处理     | **channel**（actor 模式）  |
| 只初始化一次的单例        | `sync.Once`            |
| 并发读多、偶尔写、键只增不删   | `sync.Map`（或外层加锁）      |

- 先设计共享边界：能不共享就不共享；能只读就只读；能复制就复制
- 统一通道或统一锁：一份数据只认一种同步方式，避免半锁半通道的混搭
- 所有写都在同一 goroutine：其他人通过 channel 请求（actor 化）
- 经常跑 -race 对热路径做基准测试（go test -bench .）
- 重要结构封装成类型，在方法里统一加锁/解锁，减少忘了加锁的人为失误
- 初始化与发布用 sync.Once 或先构建好再发布的不可变对象
