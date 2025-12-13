# `testing.BenchmarkResult`

benchmark 基准测试运行后产生的结果类型

- 可以用于手动运行 benchmark （不使用 `go test -bench` 也能跑）
- 自己处理 benchmark 输出（做自动化分析，性能回归对比）
- 可以更深入了解 Go 内部如何表示 benchmark 数据

## 类型定义

Go 在运行 benchmark 会返回 BenchmarkResult 结构体

```
type BenchmarkResult struct {
	N         int           
	T         time.Duration // 总耗时
	Bytes     int64         // 每次操作处理的字节数（b.SetBytes 设置的）
	MemAllocs uint64        // 总分配次数（需要）
	MemBytes  uint64        // The total number of bytes allocated.

	// Extra records additional metrics reported by ReportMetric.
	Extra map[string]float64
}
```

字段解释如下：

### N

最终执行次数，也就是 `B.N`，由 Go runtime 自动生成

### T

总耗时，只包含计时窗口的时间：

- 不包括 ResetTimer 之前的
- 不包括 StopTimer 到 StartTimer 之间的

### Bytes

用户通过 `b.SetBytes(n)` 设置的用于每次处理的数据大小，主要用于计算吞吐量

### MemAllocs

benchmark 期间发生的总分配次数，需要 `b.ReportAllocs` 才会展示，对应 CLI 输出的 `allocs/op`

### MemBytes

benchmark 期间分配的总分配字节数，同样需要 `b.ReportAllocs` 才会展示，对应 CLI 输出的 `B/op`

## 主要方法

BenchmarkResult 主要有下面几个方法

```
func (r BenchmarkResult) AllocedBytesPerOp() int64
func (r BenchmarkResult) AllocsPerOp() int64
func (r BenchmarkResult) MemString() string
func (r BenchmarkResult) NsPerOp() int64
func (r BenchmarkResult) String() string
```

### 1. NsPerOp

返回每次操作平均耗时（ns/op），计算公式为 总耗时 / 操作次数，也就是 CLI 中的 `xxx ns/op`

### 2. MemString

内存分配信息的格式化展示，返回`r.AllocedBytesPerOp` 和 `r.AllocsPerOp`格式化结果，例如：`24 B/op   2 allocs/op`，和
`go test` 格式化结果相同

> 注意必须设置了 `b.ReportAllocs` 才有效

### 3. AllocsPerOp

返回每次操作平均发生多少次内存分配：`allocs/op = MemAllocs / N`，必须调用 `b.ReportAllocs` 才会展示，否则为 0，对应 CLI 输出的
`allocs/op`

### 4. AllocedBytesPerOp

返回每次操作平均分配多少字节，同样需要 `b.ReportAllocs` 才会展示，对应 CLI 输出的 `B/op`

### 5. String

返回基准测试结果的照耀字符串展示，注意不包含基准测试名称，额外指标会覆盖同名的内置指标

> 不包含 `allocs/op` 或 `B/op`，因为这些指标由 `BenchmarkResult.MemString` 报告

## 获取 BenchmarkResult

在进行基准测试时运行 `go test -bench=.` 并不会直接产生 BenchmarkResult，但是可以在代码中手动运行 benchmark 并获得结果。

需要使用 `testing.Benchmark` 顶层函数进行获取：

```
func Benchmark(f func(b *B)) BenchmarkResult
```

- 手动执行一个 benchmark 函数，并返回 BenchmarkResult
- 用于在不运行在 `go test` 命令的情况下在程序内部动态执行 benchmark
- 在 CI 中做性能回归测试
- 构建可控的 benchmark harness （性能分析根据）获取 BenchmarkResult 进行性能分析