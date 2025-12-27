# sync.Mutex

`sync.Mutex` 是用来保护共享状态的互斥锁，保证在同一时刻，只有一个 goroutine 能进入某段临界代码。主要用于保证：

```
type Mutex struct {
	// contains filtered or unexported fields
}
```

1. 互斥性：同一时间只有一个 goroutine 进入临界区
2. 顺序一致性：在 Unlock 前的写入操作，对下一个 Lock 成功的 goroutine 一定可见
3. 消除数据竞争：只要所有访问共享状态的路径都被同一个 Mutex 包裹，可能会出现数据竞争的地方是并发安全的
4. 满足以下三点之一必须考虑使用 Mutex：
    - 存在多个 goroutine
    - 访问同一份内存
    - 至少有一个写操作

## 注意事项

- Mutex 不关心业务逻辑
- Mutex 不保证公平性，不是严格的 FIFO（主要为了性能优先）
- Mutex 是非可重入的 `mu.Lock()  mu.Lock()` 同一个 goroutine 两次上锁会导致死锁，导致永久堵塞
- 在使用是应注意：锁和数据放在一起，不暴露裸数据，所有访问路径都一致加锁

## 最小示例

```
type Counter struct {
    mu sync.Mutex
    n  int
}

func (c *Counter) Inc() {
    c.mu.Lock()
    defer c.mu.Unlock()
    c.n++
}

func (c *Counter) Value() int {
    c.mu.Lock()
    defer c.mu.Unlock() // 推荐使用 defer 防止中途 return / panic 忘记 Unlock
    return c.n
}
```

## 主要方法

| 方法          | 是否阻塞 | 是否失败        | 典型用途        |
|-------------|------|-------------|-------------|
| `Lock()`    | 会阻塞  | 不会失败        | 进入临界区（标准做法） |
| `TryLock()` | 不阻塞  | 可能失败        | 试探性进入、避免等待  |
| `Unlock()`  | 不阻塞  | 不会失败（前提是合法） | 离开临界区       |

### 1. Lock()

尝试获得锁也就是锁定 m，如果锁已经被占用，则堵塞当前 goroutine 直到锁可用

```
func (m *Mutex) Lock()
```

当 `Lock()` 返回时，Go 会有三个明确保证：

1. 互斥性，在没有 `Unlock()` 之前，没有其它 goroutine 能成功 Lock 同一个 Mutex
2. 内存可见性：上一个锁持有者在 `Unlock()` 前的写入操作对当前锁持有者全部可见
3. 程序顺序保证：`Lock()` 之后的代码一定在临界区内顺序执行

### 2. Unlock()

释放当前持有的互斥锁也就是解锁 m，并唤醒一个等待该锁的 goroutine（如果存在）

```
func (m *Mutex) TryLock() bool
```

也就是将 Mutex 状态从占用转换为空闲，如果由等待者唤醒其中一个，不是全部，不保证其唤醒顺序

**注意事项**

1. 不能 Unlock 一个未 Lock 的 Mutex，如果在进入 Unlock 函数时 m 未被锁定，则会发生运行时错误
2. 不要跨 goroutine Unlock，原则上是谁 Lock 谁 Unlock

### 3. TryLock()

Go 1.18+ 支持，非阻塞尝试获取锁，如果锁当前是可用的，则获取锁并返回 true，否则立即返回 false，不阻塞，目的是 不等待/不睡眠/不自旋

```
func (m *Mutex) TryLock() bool
```

**适用场景**

1. 可选性临界区，失败了就放弃，不影响主流程
2. 避免死锁