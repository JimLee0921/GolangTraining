# 基准测试 (Benchmark Test)

基准测试（Benchmark Test） 是用来测试代码性能的测试方法，用于测量函数的运行速度、耗时、内存分配等。

## 函数签名

```
func BenchmarkXxx(b *testing.B)
```

- 函数名以 Benchmark 开头
- 不返回任何值
- 运行基准测试：`go test -bench .` 或 `go test -bench=.`

## b *testing.B

Go 中基准测试函数的参数类型。
它是由测试框架自动创建的一个结构体指针，用来：

- 控制基准测试循环次数 (b.N)
- 记录耗时、内存分配
- 提供运行控制（如子基准测试、并行测试、清理函数等）

> 每一个基准测试函数都会自动接收到这个 b 实例

### 常用字段与方法总览

| 方法 / 字段                                      | 类型                  | 说明                               |
|----------------------------------------------|---------------------|----------------------------------|
| `b.N`                                        | `int`               | 基准循环次数，由框架自动控制                   |
| `b.ResetTimer()`                             | `func()`            | 重置计时器（在 setup 后调用，避免准备阶段被计时）     |
| `b.StopTimer()`                              | `func()`            | 暂停计时                             |
| `b.StartTimer()`                             | `func()`            | 重新开始计时                           |
| `b.ReportAllocs()`                           | `func()`            | 在输出中显示内存分配统计信息                   |
| `b.SetBytes(n int64)`                        | `func()`            | 告诉测试每次处理的数据量（如网络 I/O），用于 MB/s 显示 |
| `b.Run(name string, f func(b *testing.B))`   | `func()`            | 创建子基准测试                          |
| `b.RunParallel(f func(pb *testing.PB))`      | `func()`            | 并行执行基准测试                         |
| `b.Error`, `b.Errorf`, `b.Fatal`, `b.Fatalf` | 与 `t *testing.T` 相同 | 报告测试失败或错误                        |
| `b.Cleanup(f func())`                        | 注册清理函数，在基准测试结束后执行   |                                  |
| `b.Helper()`                                 | 标记辅助函数，使错误堆栈更干净     |                                  |
| `b.ReportMetric(float64, string)`            | 自定义性能指标输出           |                                  |
| `b.SetParallelism(n int)`                    | 设置并行度（默认 CPU 数 × n） |                                  |

### 重置定时器

如果在运行前基准测试需要一些耗时的配置，则可以使用 b.ResetTimer() 先重置定时器，例如：

```
func BenchmarkHello(b *testing.B) {
    ... // 耗时操作
    b.ResetTimer()
    for i := 0; i < b.N; i++ {
        fmt.Sprintf("hello")
    }
}
```

### 测试并发性能

使用 RunParallel 测试并发性能

```
func BenchmarkHelloRunParallel(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			fmt.Sprintf("Hello")
		}
	})
}
```

- `b.RunParallel(func(pb *testing.PB){...})`：Go 会启动多个 goroutine 并发执行测试体。并发数默认是 `GOMAXPROCS` 的倍数（可用 b.SetParallelism(n) 调整）
- `pb.Next()`：控制每个 goroutine 的循环次数，直到总运行次数达到基准测试要求。用它代替传统的 `for i := 0; i < b.N; i++`

## 报告对应值

基准测试报告每一列值对应的含义如下：

```
type BenchmarkResult struct {
    N         int           // 迭代次数
    T         time.Duration // 基准测试花费的时间
    Bytes     int64         // 一次迭代处理的字节数
    MemAllocs uint64        // 总的分配内存的次数
    MemBytes  uint64        // 总的分配内存的字节数
}
```

```
pkg: github.com/JimLee0921/GolangTraining/07_advanced/02_testing/02_benchmark-test
cpu: Intel(R) Core(TM) i3-10100 CPU @ 3.60GHz
BenchmarkHello-8        21597338                56.31 ns/op            5 B/op          1 allocs/op
```

| 字段                   | 含义                                     |
|----------------------|----------------------------------------|
| **pkg:**             | 当前测试包的导入路径。说明这是哪一个 Go 包的测试结果           |
| **cpu:**             | 测试时使用的 CPU 信息。说明环境配置，用于性能对比参考          |
| **BenchmarkHello-8** | 测试函数名 + 使用的逻辑 CPU 数（这里是 8 个线程）         |
| **21597338**         | 表示测试框架执行了 21,597,338 次循环（即 `b.N` 的最终值） |
| **56.31 ns/op**      | 平均每次操作耗时 56.31 纳秒。数值越小，性能越高            |
| **5 B/op**           | 每次操作分配的内存字节数（5 字节）。越少越好                |
| **1 allocs/op**      | 每次操作的内存分配次数。越少越好                       |





