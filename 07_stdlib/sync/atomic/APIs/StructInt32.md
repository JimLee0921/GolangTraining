# atomic.Int32

`atomic.Int32` 用于在并发环境下，对一个 int32 变量进行无锁、原子、可见的读写与更新，零值是 0，首次使用后不能被拷贝

**典型场景**

- 计数器（QPS、请求数、失败数）
- 状态码（0=init，1=running，2=stopped）
- 位标志（bitmask）
- 乐观并发更新（CAS）

```
type Int32 struct {
	// contains filtered or unexported fields
}
```

## 核心方法

### 1. Load()

原子读取：原子地读取当前值，并保证读取到的是某次 `Store / Add / Swap / CAS` 之后的完整结果

```
func (x *Int32) Load() int32
```

- 无锁
- 不阻塞
- 有 happens-before 保证
- 一旦使用 atomic，所有读都必须用 Load

---

### 2. Store()

原子写入：原子地把值设置为 `val`，并对后续 `Load()` 可见

```
func (x *Int32) Store(val int32)
```

覆盖式写，不返回旧值

---

### 3. Add()

原子加减：原子地执行：`x = x + delta`，并返回更新后的新值

```
func (x *Int32) Add(delta int32) (new int32)
```

**常用于**

- 并发计数器
- reference count
- 累加统计

> `Add(-1)` 是合法的

### 4. Swap()

原子交换：原子地把值设为 `new`，并返回旧值

```
func (x *Int32) Swap(new int32) (old int32)
```

这是只允许一次成功的无锁写法

### 5. CompareAndSwap()

CAS 操作：只有当当前值 == `old` 时，才原子地更新为 `new`，返回 bool 表示是否更新成功

```
func (x *Int32) CompareAndSwap(old, new int32) bool
```

### 6. And()

原子按位与：原子地执行 `old = x` + `x = x & mask` 并返回旧值

```
func (x *Int32) And(mask int32) (old int32)
```

用于清除 bit flag 或并发状态位管理

### 7. Or()

原子按位或：原子地执行 `old = x` + `x = x | mask` 并返回旧值

```
func (x *Int32) Or(mask int32) (old int32)
```

主要用于设置 bit flag 或并发 feature 标记

