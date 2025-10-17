## Pipeline 并发模型

在 Go 并发编程中，Pipeline（流水线）模式 是一种常见的设计思路
把一项工作拆成若干顺序阶段（stage），每个阶段用一个或多个 goroutine 工作，阶段与阶段之间用 channel 传递数据

1. 把一个大的任务拆解成多个阶段（stage）
2. 每个阶段由一个或多个 goroutine 处理，并通过 channel 连接
3. 数据像水流一样，从上游流向下游，每个阶段做自己的事

就像现实里的工厂流水线：

- 第一道工序负责把零件放到传送带上
- 第二道工序负责打磨
- 第三道工序负责上漆
- 最终产出成品

1. 每个 stage 都有输入和输出 channel
   输入：接收上游的数据
   输出：把处理结果交给下游
2. goroutine 独立执行
   每个 stage 都在 goroutine 中运行，互不阻塞
3. 数据流向单向
   只要上游继续发送数据，下游就能不断消费
4. 关闭 channel 表示结束
   当上游关闭 channel，下游的 range 会自然结束

```
source --> [stage1] --> [stage2] --> [stage3] --> sink
             |            |            |
           chanA        chanB        chanC
```

优势：结构清晰、易扩展；天然支持“背压”（下游慢会让上游在写 channel 处阻塞）、易做并发（某些阶段可开 worker 池）