# atomic.Uintptr

用于在并发环境下，对一个 `uintptr` 数值进行无锁、原子、可见的读写与更新，零值是 0

```
type Uintptr struct {
	// contains filtered or unexported fields
}
```

> API 形式和 `Uint32/Uint64` 很像，但语义与使用边界更敏感，因为 `uintptr` 常和指针/unsafe/地址运算关联；一旦用错，可能不是数据竞争，而是内存安全问题

注意：它是数值的原子操作；`uintptr` 只是机器字大小的无符号整数（32 位或 64 位，取决于架构）

---

## 核心概念

### `uintptr` 的本质

- `uintptr` 是整数，不是 `*T`
- 把指针转成 `uintptr` 再转回指针，涉及 unsafe 规则与 GC 规则

### 危险操作

如果把一个 Go 指针地址存进 `uintptr`，GC 不会把它当作指针根，可能导致：

* 对象被回收（因为 GC 看不到引用）
* 或后续转回指针指向无效内存

> `atomic.Uintptr` 适合存整数标记/计数/位标志/句柄值，不适合存Go 指针地址。

如果目标是原子指针，建议使用：`atomic.Pointer[T]`（推荐）

## 核心方法

### 1. Load()

原子读取

```
func (x *Uintptr) Load() uintptr
```

- 原子读当前值
- 有可见性保证
- 返回可能为 0（常用作“未设置”）

### 2. Store()

原子写入

```
func (x *Uintptr) Store(val uintptr)
```

- 原子设置为 `val`
- 覆盖式写，不返回旧值

### 3. Add()

原子加法

```
func (x *Uintptr) Add(delta uintptr) (new uintptr)
```

- 原子执行 `x = x + delta`
- 返回新值
- 溢出回绕按 `uintptr` 规则

### 4. Swap()

原子交换

```
func (x *Uintptr) Swap(new uintptr) (old uintptr)
```

- 原子替换为 `new`
- 返回旧值

### 5. CompareAndSwap()

CAS 操作

```
func (x *Uintptr) CompareAndSwap(old, new uintptr) (swapped bool)
```

当前值等于 `old` 才替换为 `new`

### 6. Or()

原子按位或（置位）

```
func (x *Uintptr) Or(mask uintptr) (old uintptr)
```

- 原子：`old = x; x = x | mask`
- 典型：flags 的设置位

### 7. And()

原子按位与（清位）

```
func (x *Uintptr) And(mask uintptr) (old uintptr)
```

- 原子：`old = x; x = x & mask`
- 典型：flags 的清除位