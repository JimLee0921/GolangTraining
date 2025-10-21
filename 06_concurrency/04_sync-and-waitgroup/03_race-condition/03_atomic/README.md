# 原子操作解决资源竞争

原子操作就是对单个变量进行的不可分割读/写/加/交换操作，避免并发下的中间状态被别的 goroutine 看见。
适合简单计数/标志/指针交换这类“单值同步”，不适合多个字段的一致性（需要用锁）

## Go 1.18-

旧风格（Go 1.18- 及早期代码常见）：函数操作基础类型指针

```
import "sync/atomic"

var total int64
atomic.AddInt64(&total, 1)
v := atomic.LoadInt64(&total)
atomic.StoreInt64(&total, 0)
ok := atomic.CompareAndSwapInt64(&total, 5, 6)
```

### atomic 常见操作

| 类型           | 函数                                                                                         |
|--------------|--------------------------------------------------------------------------------------------|
| 整型操作         | `AddInt32` / `AddInt64` / `LoadInt64` / `StoreInt64` / `SwapInt64` / `CompareAndSwapInt64` |
| 无符号整型        | `AddUint32` / `AddUint64` 等                                                                |
| 指针操作         | `LoadPointer` / `StorePointer` / `CompareAndSwapPointer`                                   |
| 布尔值（Go1.19+） | `Bool` 类型（`atomic.Bool`）                                                                   |

### 解释

| 函数                                  | 含义                            | 示例                                     |
|-------------------------------------|-------------------------------|----------------------------------------|
| `AddInt64(&x, n)`                   | 原子加 n 并返回新值                   | `atomic.AddInt64(&counter, 1)`         |
| `LoadInt64(&x)`                     | 原子读取当前值                       | `v := atomic.LoadInt64(&counter)`      |
| `StoreInt64(&x, v)`                 | 原子设置值                         | `atomic.StoreInt64(&flag, 1)`          |
| `SwapInt64(&x, v)`                  | 原子替换并返回旧值                     | `old := atomic.SwapInt64(&counter, 0)` |
| `CompareAndSwapInt64(&x, old, new)` | CAS 操作：只有当当前值等于 old 时才替换成 new | `atomic.CompareAndSwapInt64(&v, 0, 1)` |

## Go 1.19+

新风格（Go 1.19+ 推荐）：类型化原子，有 atomic.Int64 / Uint64 / Bool / Pointer[T] ... 等结构体，方法如
Add/Load/Store/Swap/CompareAndSwap

### 整体方法

| 新类型                                            | 作用                    | 旧版等价函数族                                                                   |
|------------------------------------------------|-----------------------|---------------------------------------------------------------------------|
| `atomic.Int32` / `Int64` / `Uint32` / `Uint64` | 原子整型计数器               | `AddInt64`, `LoadInt64`, `StoreInt64`, `SwapInt64`, `CompareAndSwapInt64` |
| `atomic.Bool`                                  | 原子布尔值                 | 以前需用 `int32` + `LoadInt32`/`StoreInt32`                                   |
| `atomic.Pointer[T]`                            | 原子指针（泛型）              | `LoadPointer`, `StorePointer`, `CompareAndSwapPointer`                    |
| `atomic.Value`                                 | 原子存取任意对象（interface{}) | （早已有，继续保留）                                                                |

### atomic.Int64

其它整数类型方法一致

| 方法                                    | 返回值  | 说明                        |
|---------------------------------------|------|---------------------------|
| `Add(n int64) int64`                  | 新值   | 原子地加 n 并返回结果              |
| `Load() int64`                        | 当前值  | 原子读取                      |
| `Store(v int64)`                      | –    | 原子写入                      |
| `Swap(new int64) int64`               | 旧值   | 原子替换                      |
| `CompareAndSwap(old, new int64) bool` | 是否成功 | CAS 操作，当前值等于 old 才替换为 new |

```
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

### atomic.Bool

布尔标志

| 方法                                   | 返回值  | 说明   |
|--------------------------------------|------|------|
| `Load() bool`                        | 当前值  | 读取   |
| `Store(v bool)`                      | –    | 写入   |
| `Swap(new bool) bool`                | 旧值   | 原子替换 |
| `CompareAndSwap(old, new bool) bool` | 是否成功 | CAS  |

```
var ready atomic.Bool
ready.Store(true)
if ready.Load() { /* ... */ }
```

### atomic.Pointer[T]

| 方法                                 | 返回值  | 说明   |
|------------------------------------|------|------|
| `Load() *T`                        | 当前指针 | 读取   |
| `Store(p *T)`                      | –    | 写入   |
| `Swap(new *T) *T`                  | 旧指针  | 原子替换 |
| `CompareAndSwap(old, new *T) bool` | 是否成功 | CAS  |

```
var p atomic.Pointer[int]
x := 10
p.Store(&x)
fmt.Println(*p.Load())       // 10
y := 20
p.CompareAndSwap(&x, &y)

```

### atomic.Value

| 方法             | 返回值  | 说明          |
|----------------|------|-------------|
| `Store(v any)` | –    | 原子写入任意类型    |
| `Load() any`   | 任意类型 | 原子读取（需类型断言） |

```
var v atomic.Value
v.Store(map[string]int{"a": 1})
m := v.Load().(map[string]int)
fmt.Println(m["a"]) // 1
```


## CAS

CAS 是 Compare-And-Swap（比较并交换）的缩写，是一种原子更新内存中单个变量的指令/原语，用来在并发环境下实现无锁（lock-free）修改。

### 工作原理

1. 读取变量当前值 old
2. 比较：如果当前内存里的值仍然是 old（期间没人改过）
3. 交换：把它原子地替换成新值 new

> 如果比较失败（说明被别人改过），CAS 返回失败，不做修改；常见做法是重试（CAS 循环）


