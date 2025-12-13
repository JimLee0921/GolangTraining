# `testing.B`

`testing.B` 主要用于创建基准测试（Benchmark test），是用来测试性能（吞吐、延迟、分配情况）的专用类型。
基准测试是一种测量 Go 代码性能的工具，主要用于得到准确，可比较的耗时和资源消耗数据。

在 Go 中 benchmark 的形式为：

```
func BenchmarkXxx(b *testing.B){
    for i := 0; i < b.N; i++ {
        ...
    }
}
```

> 因为 `testing.B` 和 `testing.T` 一样实现了 `testing.TB` 接口，这里讲解方法依旧会跳过实现自 TB 接口的方法

## 定义

```
type B struct {
	N int
	// contains filtered or unexported fields
}
```

`b.N` 字段是 `testing.B` 最核心的机制，是 Go benchmark 的基础，表示这一轮基准测试中，框架要求把要测试的操作执行多少次：

- 一次 benchmark 测试不是只跑一遍，而是会跑很多论
- 每一轮 Go runtime 会根据时间自动调整 `b.N` 的大小
- 写的 benchmark 函数会被多次调用，每次的 `b.N` 可能不同（越来越大）
- 最终 Go 会根据多轮的总时间，`b.N` 等计算出：`ns/op`, `B/op`, `allocs/op`
  `b.N` 也就是这一轮要做多少次工作的配额，由 Go 自动调节

> Go 不会强制要求在 benchmark 基准测试必须使用 `b.N` for 循环进行测试，但是一般进行基准测试都会写
`for i:=0; i<b.N; i++{...}`（当然新版本更多使用 Loop 代替 `b.N` 进行循环）

## 计时统计与查询相关方法

Benchmark 的核心输出为 `ns/op` 也就是每次操作平均耗时（纳秒）这个值是

```
ns/op = 测试窗口内总耗时 / b.N
```

下面几个方法主要是为了解决：

- 测量窗口从什么时候开始算
- 什么时候停止
- 中间哪些代码算时间、哪些不算

**StopTimer/StartTimer 与 ResetTimer 的区别**

| 方法           | 是否清零时间 | 是否暂停/恢复 | 常用场景           |
|--------------|--------|---------|----------------|
| `ResetTimer` | 清零     | 不改变开关状态 | 排除一次性 setup    |
| `StopTimer`  | 不清零    | 暂停计时    | 排除循环内的准备阶段     |
| `StartTimer` | 不清零    | 恢复计时    | 与 StopTimer 配合 |

### 1. Elapsed

返回 benchmark 计时器当前累计记录的时间，也就是基准测试的测量运行时间，很少使用，主要用于调试观察 benchmark 运行到某一时刻耗费了多久。
与 `B.StartTimer`、`B.StopTimer` 和 `B.ResetTimer` 测量的持续时间一致。

```
func (b *B) Elapsed() time.Duration
```

- 只记录计时器打开期间的耗时
- 遇到 `StopTimer()` 计时器会暂停，不再累计
- 遇到 `StartTimer()` 计时器会重新继续累计

> 一般不建议在 benchmark 实现内依赖它控制逻辑，因为 benchmark 的目的是测量性能，而不是动态行为控制

### 2. ResetTimer

清零计时器，从此刻重新开始计时并清零分配统计，
因为 benchmark 测试通常包括 setup 准备阶段（进行数据加载，源数据链接等操作）和 run 测试阶段，如果只想做 run
的部份，setup 阶段会污染结果

```
func (b *B) ResetTimer()

```

- 把已累计的时间归零
- 把分配统计归零（配合 `ReportAllocs()` ）
- 不改变计时器的开关状态（如果之前 open 就继续 open）

```
func BenchmarkProcess(b *testing.B) {
    data := loadTestData()  // setup，不进入计时

    b.ResetTimer()  // 从这里开始才算时间

    for i := 0; i < b.N; i++ {
        Process(data)
    }
}
```

> `ResetTimer()` 总是用在 benchmark 中的准备工作之后，用于排除 setup 成本，这个方法是专业 benchmark 一定会用到的

### 3. StopTimer() / StartTimer()

这两个方法成对出现，用来控制计时窗口

- StopTimer 用于暂停计时器，暂停累积时间
- StartTimer 用于恢复计时器，继续累积时间

```
func (b *B) StartTimer()
func (b *B) StopTimer()
```

当 benchmark 每次操作都需要少量准备步骤，但这些步骤不能计入性能统计时，用它们来进行排除：

```
func BenchmarkSort(b *testing.B) {
    origin := makeData()

    for i := 0; i < b.N; i++ {
        b.StopTimer()
        data := clone(origin)    // 每次都要做，但不是想要测的内容
        b.StartTimer()

        sort(data)               // 只想测这段
    }
}
```

## 执行循环与并行控制相关方法

这一部分配合 `b.N` 基本就定义了 benchmark 如何执行的行为模型

### 1. Loop

这是 Go 1.22 引入的 `b.N` 的新写法，语义上就是在当前这轮 benchmark 中根据 `b.N` 决定循环次数，
每次 Loop 返回 true 表示本轮需要继续执行，如果返回 false 表示本轮结束。

```
func (b *B) Loop() bool
```

底层依然是使用 `b.N` 来控制总执行次数，只是：

- 不需要显示写 `i := 0; i < b.N; i++`
- Go 未来也可以在 Loop 中进行更复杂的批次/调度优化
- 新代码推荐使用 `for b.Loop() {...}`，两种是等价的

老版本写法：

```
for i := 0; i < b.N; i++ {
    work()
}
```

新版本写法：

```
for b.Loop(){
    work()
}
```

### 2. Run

创建子基准测试，benchmark 版本的 `t.Run()`，每个子 benchmark 有自己的：

- 名字（父名/子名）
- `b.N`
- 计时器
- 分配统计

```
func (b *B) Run(name string, f func(b *B)) bool
```

返回值为 bool，表示子 benchmark 是否通过测试，一般不关心，除非做一些 meta 的逻辑（例如某子 benchmark 挂了就不继续后面的子
benchmark）

**主要用途**

1. 对比不同实现版本（V1/V2/V3）
2. 对比不同输入规模（10/100/1000）
3. 对同类算法进行分组展示

### 3. RunParallel

`testing.B` 独特且重要的能力：并发压测，用于启动多个 goroutine，每个 goroutine 拿到一个 `*PB`(parallel benchmark)，通过
`pb.Next()` 控制循环次数，从而在多核下并发执行被测代码。

```
func (b *B) RunParallel(body func(*PB))
```

**简化示例**

```
b.RunParallel(func(pb *PB)){
    for pb.Next(){  // 类似 Loop()
        work()
    }
}
```

- pb.Next() 返回 true 时继续执行一次操作
- Go runtime 会根据 b.N 和并行度分配这些操作到多个 goroutine
- 不需要自己写 `for i<b.N` 或 `for b.Loop()`，完全交给 PB

**使用场景**

- 多 goroutine 同时访问某数据结构（map + lock / sync.Map / 自己的并发结构）
- 锁竞争性能（mutex / RWMutex / CAS 等）
- server 在多核环境下的 handler 性能
- 不适合单线程算法（直接使用普通 benchmark）

### 4. SetParallelism

配合 RunParallel 使用，控制并行度：把默认并行 goroutine 数量从 “约等于`GOMAXPROCS`” 调整为 “约等于`p * GOMAXPROCS`。


> 默认并行数大概是 GOMAXPROCS（比如 8 核就开 8 个 worker）

```
func (b *B) SetParallelism(p int)
```

**使用场景**

- 需要测试更高并发场景下的行为
- 研究锁、atomic、无锁结构在不同并发度下的性能曲线
- 做类似压测实验，观察性能随并行度的变化

## 性能指标与报告相关方法

这几个方法决定了 benchmark 测试输出中除了 `ns/op` 之外的那几列内容，是专业性能分析和调优的关键工具

### 1. ReportAllocs

用于开启内存分配统计，开启后，benchmark 测试的输出会多出两列：

- `B/op`：每次操作平均分配的字节数（Bytes per operation）
- `allocs/p[`：每次操作平均发生的分配次数（Allocations per operation）

```
func (b *B) ReportAllocs()
```

> 只需要在 benchmark 测试的开头调用一次 `b.ReportAllocs` 即可

比如输出结果为：`BenchmarkX-8   1000000   200 ns/op   24 B/op   2 allocs/op`

- 每次操作平均耗时 200ms
- 每次操作平局分配了 24 个字节
- 每次操作平均触发两次内存分配

> 只影响打印结果，不会改变程序行为

### 2. SetBytes

用于配置吞吐量单位 `bytes/op`，设置每次操作处理的数据量（单位为字节），这样 benchmark 就能额外报告每秒处理多少字节：例如
`500MB/s`

```
func (b *B) SetBytes(n int64)
```

**使用场景**

1. 网络 I/O：每次发送/接收 N 字节
2. 文件读写：每次读写块大小固定
3. 加密/压缩/哈希：每次处理一个数据块

> SetBytes 不会改变逻辑，仅用于输出统计，不会改变 benchmark 循环行为，只是告诉 testing 每次循环等价于处理了 n 字节数据

### 3. ReportMetric

自定义指标输出，高级用法，向 benchmark 输出中添加一列自定义指标

```
func (b *B) ReportMetric(n float64, unit string)
```

自定义指标列可以自己定义：

- 单位名（unit）：例如 "latency_ns"、"qps"、"hit_rate"、"errors_per_sec" 等
- 数值（n)：自己计算出的某个统计值

Go 会最终在 benchmark 输出中额外加上这一列
