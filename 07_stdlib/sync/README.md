# sync 包

sync 是 Go 标准库中并发控制（concurrency control）的核心包之一，主要用于在多 goroutine 并发执行时，保证数据一致性、顺序性和安全性。解决的是多个
goroutine 同时访问共享资源会发生什么的问题。

## 设计哲学

Go 并发的核心里面是："Do not communicate by sharing memory; share memory by communicating."（优先使用 channel 通信，而不是共享内存）

在 Go 中：

- goroutine 是并发执行的
- goroutine 之间共享内存
- Go 不自动解决并发冲突

而这些在多个 goroutine 共同访问资源时可能会遇到：

- 数据竞争（data race）
- 脏读/写覆盖
- 状态不一致
- 死锁 / 活锁

> sync 的职责就是为并发 goroutine 提供低层次、可控、确定性的同步原语

## 核心组件

| 类型           | 解决什么问题          |
|--------------|-----------------|
| `Mutex`      | 互斥访问            |
| `RWMutex`    | 读写分离锁           |
| `WaitGroup`  | 等待 goroutine 完成 |
| `Once`       | 只执行一次           |
| `Cond`       | 条件等待/通知         |
| `Map`        | 并发安全 map        |
| `Pool`       | 对象复用            |
| `atomic`（子包） | 无锁原子操作          |


