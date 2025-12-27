# atomic.Uint32

`atomic.Uint32` 用于在并发环境下，对一个 `uint32` 变量进行无锁、原子、可见的读写与更新

```
type Uint32 struct {
	// contains filtered or unexported fields
}
```

相比 `Int32`，`Uint32` 在这些场景更自然：

- 位标志（bitmask）
- 状态集合（多个 boolean 合并成一个变量）
- 计数但不需要负数
- 协议 / 标志字段（天然是无符号）

如果需要负数、差值判断，选 `Int32`
如果需要位运算 + 状态位，选 `Uint32`

## 核心方法

### 1. Load()

原子读取：原子地读取当前值，并保证可见性

```
func (x *Uint32) Load() uint32
```

- 无锁
- 不阻塞
- 一旦用了 atomic，所有读都必须用 `Load()`

### 2. Store()

原子写入：原子地把值设置为 `val`

```
func (x *Uint32) Store(val uint32)
```

### 3. Add()

原子加法：原子地执行：`x = x + delta`，并返回新值

```
func (x *Uint32) Add(delta uint32) (new uint32)
```

- 这是无符号加法
- 溢出会按 `uint32` 规则回绕（wrap around）

### 4. Swap()

原子交换：原子地设置为 `new`，并返回旧值

```
func (x *Uint32) Swap(new uint32) (old uint32)
```

### 5. CompareAndSwap()

CAS 操作：仅当当前值 == `old` 时，才原子地更新为 `new`

```
func (x *Uint32) CompareAndSwap(old, new uint32) (swapped bool)
```

### 6. Or()

原子按位或（设置位）：原子地执行 `old = x` + `x = x | mask`

```
func (x *Uint32) Or(mask uint32) (old uint32)
```

### 7. And()

原子按位与（清除位）：原子地执行 `old = x` + `x = x & mask`

```
func (x *Uint32) And(mask uint32) (old uint32)
```

