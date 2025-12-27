# atomic.Bool

`atomic.Bool` 用于在并发环境下对一个布尔值进行无锁、原子、可见的读写与切换，零值是 false，首次使用后不能被复制

```
type bool struct {
    // contains filtered or unexported fields
}
```

## 核心方法

### 1. Load()

原子读取：原子地读取当前布尔值并保证读取到的是某个 Store/Swap/CAS 完成后的值

```
func (x *Bool) Load() bool
```

无锁 / 不堵塞 / 有内存可见性（happens-before）

> 一旦使用了 atomic.Bool，所有读都必须用 Load()

### 2. Store()

原子写入：原子地把布尔值设置为 val，并对后续 `Load()` 可见

```
func (x *Bool) Store(val bool)
```

- 覆盖式写入
- 不返回旧值
- 无条件成功

### 3. Swap()

原子交换：原子地把值设为 new，并返回之前的旧值

```
func (x *Bool) Swap(new bool) (old bool)
```

只执行一次，防止重复启动/重复关闭，只允许一个 goroutine 成功切换状态（新值换为旧值）

### 4. CompareAndSwap()

CAS 操作：只有当当前的值等于 old 时，才原子地把旧值替换为 new 新值

```
func (x *Bool) CompareAndSwap(old, new bool) (swapped bool)
```