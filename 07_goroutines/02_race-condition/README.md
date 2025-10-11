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
  程序运行时会监控内存访问，发现多个 goroutine 并发访问同一个变量且至少有一个写操作，就会报 DATA RACE 错误

## 修复资源竞争

Go 中主要有下面四种方法来修复资源竞争问题

1. 互斥锁：`sync.Mutex` 或 `sync.RWMutex`
2. 使用 `sync/atomic` 原子操作
3. 使用 channel 管道（见 channel 章节）
4. 只初始化一次：sync.Once


