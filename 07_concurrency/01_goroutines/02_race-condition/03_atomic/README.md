# 原子操作解决资源竞争

原子操作就是对单个变量进行的不可分割读/写/加/交换操作，避免并发下的中间状态被别的 goroutine 看见。
适合简单计数/标志/指针交换这类“单值同步”，不适合多个字段的一致性（需要用锁）

## Go 1.18-

旧风格（Go 1.18- 及早期代码常见）：函数操作基础类型指针

```go
import "sync/atomic"

var total int64
atomic.AddInt64(&total, 1)
v := atomic.LoadInt64(&total)
atomic.StoreInt64(&total, 0)
ok := atomic.CompareAndSwapInt64(&total, 5, 6)
```

## Go 1.19+

新风格（Go 1.19+ 推荐）：类型化原子，有 atomic.Int64 / Uint64 / Bool / Pointer[T] ... 等结构体，方法如
Add/Load/Store/Swap/CompareAndSwap

```go
import "sync/atomic"

var cnt atomic.Int64

// ++
cnt.Add(1)

// 读
n := cnt.Load()

// 写
cnt.Store(0)

// CAS（期待旧值为 5，换成 6）
ok := cnt.CompareAndSwap(5, 6)
```

布尔标志

```go
var ready atomic.Bool
ready.Store(true)
if ready.Load() { /* ... */ }
```

指针/句柄热更新（零拷贝切换）

```go
type Config struct{ Limit int }
var cfg atomic.Pointer[Config]

// 发布新配置
cfg.Store(&Config{Limit: 100})

// 任意协程读取一致快照
c := cfg.Load()
do(c.Limit)
```

> 还有 atomic.Value（非泛型）：可存任意类型的不可变对象，Store/Load 间保持一致性（适合发布只读配置）

## CAS

CAS 是 Compare-And-Swap（比较并交换）的缩写，是一种原子更新内存中单个变量的指令/原语，用来在并发环境下实现无锁（lock-free）修改。

### 工作原理

1. 读取变量当前值 old
2. 比较：如果当前内存里的值仍然是 old（期间没人改过）
3. 交换：把它原子地替换成新值 new

> 如果比较失败（说明被别人改过），CAS 返回失败，不做修改；常见做法是重试（CAS 循环）