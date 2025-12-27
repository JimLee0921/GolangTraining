# atomic.Int64

`atomic.Int64` 用于在并发环境下，对单个 int64 变量进行无锁、原子、可见的读写与更新，int返回更大，更不容易溢出

```
type Int64 struct {
	// contains filtered or unexported fields
}
```

典型用途：

* 请求计数、错误计数（更不容易溢出）
* 累计字节数、累计金额（注意单位）
* 时间戳（UnixNano）
* 状态值（也可以，但通常 Int32 足够）

**边界不变：只适合一个变量的简单操作。**

## 核心方法

### 1. Load()

原子读取：原子读取当前值，对 `Store/Add/Swap/CAS` 的结果有可见性保证，一旦用了 atomic，所有读都必须 Load

```
func (x *Int64) Load() int64
```

### 2. Store()

原子写入：原子设置为 `val`，覆盖写，不返回旧值，常用于发布状态或初始化数值

```
func (x *Int64) Store(val int64)
```

### 3. Add()

原子加减：原子执行 `x = x + delta`，返回更新后的新值

```
func (x *Int64) Add(delta int64) (new int64)
```

### 4. Swap()

原子交换：原子把值设为 `new`，返回旧值，常用于取出并清零或只让第一个成功

```
func (x *Int64) Swap(new int64) (old int64)
```

### 5. CompareAndSwap()

CAS 操作：只有当当前值 == `old` 时才更新为 `new`，返回是否成功，常用于状态跃迁、乐观并发

```
func (x *Int64) CompareAndSwap(old, new int64) (swapped bool)
```

### 6. And()

原子按位与：语义等价于：

```
old = x
x = x & mask
return old
```

```
func (x *Int64) And(mask int64) (old int64)
```

主要用于清除某些 bit flag（把某些位变 0）

### 7. Or()

原子按位或：语义等价于：

```
old = x
x = x | mask
return old
```

```
func (x *Int64) Or(mask int64) (old int64)
```

主要用于设置某些 bit flag（把某些位置 1）