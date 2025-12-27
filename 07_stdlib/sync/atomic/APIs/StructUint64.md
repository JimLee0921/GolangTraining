# atomic.Uint64

`atomic.Uint64` 用于在并发环境下，对一个 `uint64` 变量进行无锁、原子、可见的读写与更新

```
type Uint64 struct {
	// contains filtered or unexported fields
}
```

使用前提不变：只涉及一个变量、逻辑简单。比 `Uint32` 更适合这些场景：

- 大计数（QPS、累计请求数、累计字节数）
- 时间戳 / 序列号（不需要负数）
- 64 位位标志（flags）
- 跨较长时间运行的统计，避免溢出

如果不需要负数，且数值可能很大，优先 `Uint64`

## 核心方法

### 1. Load()

原子读取：原子地读取当前值，并保证可见性（happens-before）

```
func (x *Uint64) Load() uint64
```

- 无锁、不阻塞
- 一旦用了 atomic，所有读都必须 `Load()`

### 2. Store()

原子写入：原子地把值设置为 `val`

```
func (x *Uint64) Store(val uint64)
```

### 3. Add()

原子加法：原子地执行：`x = x + delta`，并返回新值

```
func (x *Uint64) Add(delta uint64) (new uint64)
```

- 无符号加法
- 溢出会按 `uint64` 回绕（极少见，但要心里有数）

### 4. Swap()

原子交换：原子地把值设为 `new`，并返回旧值

```
func (x *Uint64) Swap(new uint64) (old uint64)
```

### 5. CompareAndSwap()

CAS 操作：仅当当前值等于 `old` 时，才原子地更新为 `new`

```
func (x *Uint64) CompareAndSwap(old, new uint64) (swapped bool)
```

### 6. Or()

原子按位或（置位）：原子地执行 `old = x` + `x = x | mask`

```
func (x *Uint64) Or(mask uint64) (old uint64)
```

主要用于设置某个状态位或打开功能标志

### 7. And()

原子按位与（清位）：原子地执行 `old = x` + `x = x & mask`

```
func (x *Uint64) And(mask uint64) (old uint64)
```

主要用于清除某个状态位或回收标志
