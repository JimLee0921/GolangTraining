## Fanin and Fanout

Fan-Out：把同一批任务分发给多个 goroutine 并行处理

Fan-In：把多路 goroutine 的结果汇聚成一条通道供下游消费

两者常成对出现：源 ->（Fan-Out）worker 池 ->（Fan-In）合并 -> 下游

### 扇入（Fan in）

- 指的是将多路通道聚合到一条通道中处理，从多条输入 channel 读数据，写入一个输出 channel
- golang中最简单的扇入就是使用select聚合多条通道服务，当生产者的速度很慢时，需要使用扇入技术聚合多个生产者满足消费者
- 需要一个收尾协作：当所有输入都完结时，关闭输出 channel，让消费者优雅退出
- 通常用 sync.WaitGroup 或 计数器+done chan 统计合并者的存活数
- 取消/超时靠 context.Context：所有阻塞读写都应 select { case <-ctx.Done(): return }

### 扇出（Fan out）

- 指的是将一条通道发散到多条通道中处理，多个 goroutine 共同从一个输入 channel 读（或上游将任务轮询/广播到多个输入 channel）
- 在golang中的具体实现就是使用 go 关键字启动多个 goroutine 并发处理，当消费者的速度很慢时，需要使用扇出技术来并发处理请求
- 并发度 = worker 数
- channel 的缓冲区提供削峰填谷，下游慢时，写操作阻塞形成背压

> [博客地址](https://go.dev/blog/pipelines)
