# Go 资源竞争

资源竞争（race condition） 指的是：
当多个 goroutine 同时访问同一份数据，并且至少有一个在修改，
而又没有同步机制来保证访问顺序时，就会发生资源竞争。

## 定位资源竞争

可以使用 go run 命令查看是否存在资源竞争（Go 的 race detector 底层依赖 C 代码（需要 cgo 支持）默认情况下，Windows 下安装的
Go 有时会把 CGO_ENABLED=0，导致 -race 用不了）

- go run main.go：
    - 编译并运行 main.go
    - 不会检查资源竞争
    - 程序结果可能是错的，但不会提示
- go run -race main.go
    - 这是带竞态检测（Race Detector）的运行方式：
    - `-race` 会在编译时插入额外的检查逻辑
    - 程序运行时会监控内存访问，发现多个 goroutine 并发访问同一个变量且至少有一个写操作，就会报 DATA RACE 错误

## 修复资源竞争

Go 中主要有下面几种方法来修复资源竞争问题

1. 互斥锁：`sync.Mutex` 或 `sync.RWMutex`，保证每次只有一个 goroutine 能进入临界区（对共享资源的访问操作）
2. 使用 `sync/atomic` 原子操作，保证读写不可分割，效率比锁更高，适合简单的整数累加等场景
3. 使用 channel 管道：Go 的推荐做法，不要让多个 goroutine 共享数据，而是通过 channel 传递数据
4. 只初始化一次：sync.Once


