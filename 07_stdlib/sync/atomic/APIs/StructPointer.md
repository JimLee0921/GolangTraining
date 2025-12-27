# atomic.Pointer

`atomic.Pointer[T]` 用于在并发环境下，对指向某个对象的指针进行无锁、原子、可见的替换与读取

```
type Pointer[T any] struct {
// contains filtered or unexported fields
}
```

## 核心动机

一个典型需求：

- 多个 goroutine 只读某个配置结构，
- 少量 goroutine 偶尔整体更新配置。

如果用 Mutex：

- 读也要加锁
- 高并发读下有锁竞争

如果用 `atomic.Pointer`：

- 无锁读取
- 写入是一次原子指针替换
- 读到的对象是完整一致的快照

## 能力边界

只能保证指针本身的原子性，不能保证指针指向对象内部的并发安全

正确使用方式是：

- 写：构造一个新对象 -> Store/Swap/CAS
- 读：Load -> 只读使用
- 要原地修改已发布的对象

## 核心方法

### 1. Load()

原子读取指针：原子地读取当前指针值，并保证内存可见性

```
func (x *Pointer[T]) Load() *T
```

- 无锁、不阻塞
- 返回的是某次成功 Store/Swap/CAS 后的指针
- 可能返回 `nil`（初始状态）

### 2. Store()

原子发布指针：原子地把指针设置为 `val`，并对后续 `Load()` 可见

```
func (x *Pointer[T]) Store(val *T)
```

**典型用途**

* 初始化
* 全量替换配置
* 发布新快照

> Store 后不要再修改 `val` 指向的对象

### 3. Swap()

原子替换并返回旧指针：原子地把指针替换为 `new`，并返回之前的旧指针

```
func (x *Pointer[T]) Swap(new *T) (old *T)
```

**典型用途**

- 更新并同时拿到旧值做清理
- 日志/监控里记录版本切换

### 4. CompareAndSwap()

CAS（条件替换）：只有当当前指针恰好等于 `old`时，才原子地替换为 `new`

```
func (x *Pointer[T]) CompareAndSwap(old, new *T) (swapped bool)
```