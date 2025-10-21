# Worker Pool（工作池）

> 生产者-消费者模型解决的是：任务的生成与消费解耦，而 Worker Pool 模型进一步解决：如何高效地执行大量任务，同时限制并发数量

在 Go 中，创建 goroutine 虽然非常轻量，
但在高并发场景中，如果无限制地为每个任务开新协程，会带来：

* 过多内存分配
* 调度器负载剧增
* CPU 上下文切换频繁
* 系统崩溃（协程风暴）

所以可以使用 Worker Pool 工作池模式：创建一个工作池，其中包含固定数量的工作者（Worker），这些工作者从一个共享的队列中获取任务，并执行它们。
当一个工作者完成任务时，它会返回工作池，等待分配新的任务。

## 核心思想

`worker pool` 其实就是线程池 `thread pool`。只是对于 go 来说，直接使用的是 goroutine 而非线程

工作池模式的核心目标是用 N 个固定协程，消费 M 个任务（M ≥ N）

| 组件             | 作用                  |
|----------------|---------------------|
| **Job Queue**  | 缓冲任务的 Channel       |
| **Workers**    | 持续消费任务的 Goroutine   |
| **Dispatcher** | 分发任务、管理 worker 生命周期 |

每个 Worker 类似工人，Dispatcher 类似调度员。

## 核心组件说明

### Job（任务）

定义任务结构，用于传递要执行的数据。


### Worker（工人）

每个 worker 是一个独立的 goroutine，从 job channel 中取任务执行。


### Dispatcher（调度器）

负责创建 job channel、启动 workers 并发送任务。

## 运行机制

1. 初始化任务队列（`jobs := make(chan Job, N)`）
2. 启动固定数量 worker（每个 worker 在独立 goroutine 中运行）
3. 主线程分发任务（写入 `jobs`）
4. workers 从 `jobs` 中读取任务执行
5. 所有任务完成后关闭通道并等待退出

> 这是一种生产者–消费者模型（Producer–Consumer）的优化形式。

## 模式变体

| 模式                           | 特征                |
|------------------------------|-------------------|
| **固定工作池（Fixed Pool）**        | 固定 worker 数量      |
| **动态工作池（Dynamic Pool）**      | 可根据负载调整 worker 数量 |
| **限速工作池（Rate-limited Pool）** | 限制任务提交速率          |
| **优先级工作池（Priority Pool）**    | 任务按优先级调度          |
| **分区工作池（Sharded Pool）**      | 各 worker 处理独立数据分区 |

