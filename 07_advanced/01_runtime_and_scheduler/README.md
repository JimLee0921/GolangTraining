> 参考博客地址：https://www.cnblogs.com/yubo-guan/p/18748830

## Runtime 库的作用

Runtime 库是 Go 语言运行时系统的实现，它在程序启动时自动加载，并为 Go 程序提供以下核心功能：

* Goroutine 调度：管理 Goroutine 的创建、调度和销毁。
* 内存管理：负责内存分配和垃圾回收（GC）。
* 网络 I/O：实现非阻塞 I/O 操作。
* 系统调用：封装操作系统调用，提供跨平台支持。
* 类型系统：支持反射、接口和类型断言等高级特性。
* panic 和 recover：实现异常处理机制。

---

## Runtime 库的核心组件

### （1）Goroutine 调度器（Scheduler）

* M:N 调度模型：Go 运行时采用 M:N 调度模型，将大量的 Goroutine (N) 调度到少量的操作系统线程 (M)
  上运行。这种模型既保证了并发性能，又减少了线程切换的开销。
* G-M-P 模型：Go 调度器的实现基于 G-M-P 模型：

    * G（Goroutine）：表示一个 Goroutine，包含栈、程序计数器、状态等信息。
    * M（Machine）：表示一个操作系统线程，负责执行 Goroutine。
    * P（Processor）：表示一个逻辑处理器，负责管理 Goroutine 队列和调度。
* 工作 stealing（任务窃取）：当一个 P 的本地 Goroutine 队列为空时，它会从其他 P 的队列中“窃取” Goroutine
  来执行，以实现负载均衡。

### （2）内存管理

* 内存分配器：Go 运行时使用高效的内存分配器来管理堆内存。内存分配器基于 TCMalloc（Thread-Caching
  Malloc）设计，将内存划分为多个大小类别，以减少内存碎片和提高分配效率。
* 垃圾回收（GC）：Go 的垃圾回收器采用并发标记-清除算法（Concurrent Mark-Sweep），具有以下特点：

    * 低延迟：GC 的大部分工作是与用户程序并发执行的，减少了 STW（Stop-The-World）的时间。
    * 分代回收：Go 的 GC 会优先回收年轻代对象，以提高回收效率。
    * 写屏障（Write Barrier）：在并发标记阶段，写屏障用于确保对象引用的正确性。

### （3）网络 I/O 与系统调用

* 非阻塞 I/O：Go 运行时使用非阻塞 I/O 模型（如 epoll、kqueue 等）来实现高效的网络通信。Goroutine 在等待 I/O
  操作时会被挂起，而不会阻塞操作系统线程。
* 系统调用封装：Runtime 库封装了操作系统的系统调用，提供了跨平台的抽象接口。例如，文件操作、网络通信等都会通过 Runtime
  库转换为相应的系统调用。

### （4）类型系统与反射

* 类型信息：Runtime
  库维护了所有类型的元信息（如大小、对齐方式、方法集等），这些信息用于支持接口、类型断言和反射等特性。
* 反射：Go 的反射功能（`reflect` 包）依赖于 Runtime
  库提供的类型信息。通过反射，程序可以在运行时动态地获取和操作类型信息。

### （5）异常处理

* panic 和 recover：Go 的异常处理机制基于 `panic` 和 `recover`。当发生 `panic` 时，Runtime 库会展开调用栈，并执行延迟函数（
  `defer`）。`recover` 可以捕获 `panic` 并恢复程序的正常执行。

---

## Runtime 库的工作机制

### （1）程序启动

* 当 Go 程序启动时，Runtime 库会初始化以下组件：

    * 内存分配器和垃圾回收器。
    * Goroutine 调度器。
    * 网络 I/O 和系统调用模块。
* 然后，Runtime 库会创建主 Goroutine，并开始执行 `main` 函数。

### （2）Goroutine 调度

* 调度器会周期性地检查 Goroutine 的状态，并在多个 Goroutine 之间进行上下文切换。
* 当 Goroutine 阻塞时（如等待 channel 或 I/O 操作），调度器会将其挂起，并调度其他 Goroutine 运行。

### （3）垃圾回收

* 垃圾回收器会周期性地扫描堆内存，标记存活对象，并回收未使用的内存。
* 在 GC 期间，Runtime 库会暂停所有 Goroutine（STW 阶段），以确保标记的正确性。

### （4）程序退出

* 当 `main` 函数执行完毕时，Runtime 库会等待所有 Goroutine 退出，然后释放资源并结束程序。

---

## Runtime 库的调试与调优

### （1）调试工具

* `GODEBUG`：通过设置 `GODEBUG` 环境变量，可以启用 Runtime 库的调试信息。例如：

  ```bash
  GODEBUG=gctrace=1 ./myprogram
  ```

  这会输出垃圾回收的详细信息。
* `pprof`：Go 提供了 `pprof` 工具，用于分析程序的 CPU、内存和 Goroutine 使用情况。

### （2）调优参数

* `GOGC`：通过设置 `GOGC` 环境变量，可以调整垃圾回收的触发频率。例如：

  ```bash
  GOGC=100 ./myprogram
  ```

  这会将 GC 的触发阈值设置为 100%。
* `GOMAXPROCS`：通过设置 `GOMAXPROCS` 环境变量，可以控制运行时使用的 CPU 核心数。例如：

  ```bash
  GOMAXPROCS=4 ./myprogram
  ```

  这会限制程序最多使用 4 个 CPU 核心。

---

## 总结

Go 语言的 Runtime 库是 Go 程序运行时的核心组件，它负责管理 Goroutine 调度、内存分配、垃圾回收、网络 I/O
等底层操作。通过高效的调度器和内存管理机制，Runtime 库为 Go 语言提供了强大的并发能力和高性能。理解 Runtime
库的工作原理，有助于编写高效、可靠的 Go 程序，并在需要时进行调试和调优。

> Do not communicate by sharing memory; instead, share memory by communicating.

## `runtime` 包常用函数

### `runtime.NumCPU()`

**作用：**
返回当前系统的逻辑 CPU 数量（例如你的机器 8 核，就返回 8）。

```
fmt.Println(runtime.NumCPU()) // 输出：8
```

> 常用于决定并发度，例如在爬虫、worker pool、并行计算中动态设置并发数量。

---

### `runtime.NumCgoCall()`

**作用：**
返回当前进程中已执行过的 cgo 调用次数（即 Go 调用 C 函数的次数）。

```
fmt.Println("初始 cgo 调用数:", runtime.NumCgoCall())
```

> 可以通过 import "C" 调用 C 语言函数，调用一次 C 函数，Go 的 runtime 会进行一次 cgo 调度过程（涉及线程切换、栈切换、锁定等）

---

### `runtime.GOMAXPROCS(n int)`

**作用：**
设置同时执行用户级 Go 代码的最大操作系统线程数（并行度）
如果参数为 0，则返回当前值

```
fmt.Println(runtime.GOMAXPROCS(0)) // 查看当前设置
runtime.GOMAXPROCS(4)              // 让调度器最多并行 4 个线程
```

> 注意：
>
> * 不要盲目设太大，一般设为 `runtime.NumCPU()`
> * Go 1.5 之后默认值就是 CPU 核心数

---

### `runtime.NumGoroutine()`

**作用：**
返回当前活跃的 goroutine 数量。

```
fmt.Println("当前 goroutine 数:", runtime.NumGoroutine())
```

> 常用于调试 goroutine 泄漏或死锁。

---

### `runtime.Gosched()`

**作用：**
让出 CPU 使用权，调度器可以切换去运行其他 goroutine。
当前 goroutine 会暂时挂起，稍后再被恢复。

```
go func() {
	for i := 0; i < 5; i++ {
		fmt.Println("A:", i)
		runtime.Gosched()
	}
}()
for i := 0; i < 5; i++ {
	fmt.Println("B:", i)
}
```

> 用于实验或协作式调度（一般不必手动调用）。

---

### `runtime.Goexit()`

**作用：**

- 立即终止当前 goroutine 的执行
- 不会影响其他 goroutine，也不会引发 panic
- 但会执行当前 goroutine 的 `defer` 语句。

```
go func() {
	defer fmt.Println("defer 被执行")
	fmt.Println("即将退出")
	runtime.Goexit()
	fmt.Println("不会执行到这里")
}()
```

> 常用于服务 goroutine 主动退出的场景，比如 worker 中检测到停止信号。

---

### `runtime.LockOSThread()` / `UnlockOSThread()`

**作用：**
将当前 goroutine 固定（或解除固定）在当前操作系统线程上。

```
runtime.LockOSThread()
// 此 goroutine 永远运行在同一 OS 线程
defer runtime.UnlockOSThread()
```

> 典型用途：
>
> * 调用依赖线程上下文的 C 代码（通过 cgo）
> * GUI 程序中必须在主线程执行的操作
> * 某些需要线程局部变量的特殊场景

---

### `runtime.Caller(skip int)`

**作用：**
返回调用栈信息，用于日志或错误追踪。

```
func trace() {
	pc, file, line, ok := runtime.Caller(1)
	if ok {
		fmt.Printf("调用来自：%s:%d [%v]\n", file, line, runtime.FuncForPC(pc).Name())
	}
}
```

> `skip=0` 表示当前函数，`skip=1` 表示调用者。

---

### `runtime.Stack(buf []byte, all bool)`

**作用：**
将当前（或全部）goroutine 的栈信息写入 `buf`。
返回写入的字节数。

```
buf := make([]byte, 1024)
n := runtime.Stack(buf, false)
fmt.Println(string(buf[:n]))
```

> 用于调试死锁或 panic 时的堆栈打印。
> `all=true` 会打印所有 goroutine 的堆栈。

---

### `runtime.ReadMemStats(m *runtime.MemStats)`

**作用：**
读取当前内存使用与 GC 统计信息。

```
var m runtime.MemStats
runtime.ReadMemStats(&m)
fmt.Printf("Alloc = %v KB\n", m.Alloc/1024)
fmt.Printf("TotalAlloc = %v KB\n", m.TotalAlloc/1024)
fmt.Printf("NumGC = %v\n", m.NumGC)
```

> 用于性能分析与内存监控。

---

### `runtime.GC()`

**作用：**
立即触发一次垃圾回收（GC）。

```
runtime.GC()
```

> 一般仅用于测试或性能实验；生产环境无需手动调用。

---

### `runtime/debug.FreeOSMemory()`

**作用：**
强制释放当前空闲的堆内存给操作系统。

```
import "runtime/debug"

debug.FreeOSMemory()
```

> 通常 GC 不会立刻归还空闲内存，这个函数会强制归还。

---

### `runtime/debug.SetGCPercent(p int)`

**作用：**
设置 GC 触发频率（百分比）。
默认是 `100`，表示当堆大小增长到上次 GC 后的 2 倍时触发下一次 GC。

```
debug.SetGCPercent(50) // 提高 GC 频率
```

> 在低延迟系统中可用来控制 GC 频率和停顿时间。

---

### 调度与 GC 调试环境变量

| 变量                         | 说明          |
|----------------------------|-------------|
| `GODEBUG=schedtrace=1000`  | 每秒打印一次调度器状态 |
| `GODEBUG=scheddetail=1`    | 打印调度器详细日志   |
| `GODEBUG=gctrace=1`        | 打印垃圾回收信息    |
| `GODEBUG=allocfreetrace=1` | 打印内存分配与释放信息 |

示例：

```bash
GODEBUG=schedtrace=1000,scheddetail=1 go run main.go
```

---

### 小结

| 分类       | 主要函数                                        | 说明                        |
|----------|---------------------------------------------|---------------------------|
| **调度控制** | `GOMAXPROCS`, `Gosched`, `Goexit`           | 控制并行度、让出 CPU、退出 goroutine |
| **信息查询** | `NumCPU`, `NumGoroutine`, `Caller`, `Stack` | 查询系统与运行状态                 |
| **内存管理** | `ReadMemStats`, `GC`, `debug.FreeOSMemory`  | 手动或监控内存使用                 |
| **线程绑定** | `LockOSThread`, `UnlockOSThread`            | 控制 goroutine 与线程绑定        |
| **调试分析** | `GODEBUG` 环境变量, `pprof`, `trace`            | 调度与性能分析工具支持               |




