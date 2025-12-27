# sync.RWMutex

`sync.RWMutex` 是读写互斥锁，允许多个 goroutine 并发进行读取，但是写必须单独 goroutine 独占，用于读多写少的场景

```
type RWMutex struct {
	// contains filtered or unexported fields
}
```

## 对比 Mutex

对于 Mutex 来说，无论是读还是写，同一时候都只能有一个 goroutine 进入临界区，也就是读、写、读+写都不能并发，这样也就导致了所有操作之间都会互相阻塞，在高并发下是纯性能浪费

而 RWMutex 的核心思想是读不改变状态，所以读与读之间不需要互斥访问，所以可以使用 `RLock / RUnlock` 进行读操作，
`Lock / Unlock` 进行写操作来进行性能优化

## 核心方法

```
Lock    = 写锁（独占）
RLock   = 读锁（共享）
```

### 1. Lock()

获取写操作互斥锁，如果有任何读锁或写锁存在，则阻塞直到所有锁释放

```
func (rw *RWMutex) Lock()
```

写操作是绝对独占锁，写操作会阻塞所有新的读/写操作，已存在的读必须全部退出

### 2. Unlock

释放写操作锁，并唤醒等待的读/写操作 goroutine（唤醒顺序不保证）

```
func (rw *RWMutex) Unlock()
```

- 必须和 `Lock()` 成对出现
- `Unlock()` 为加锁的 RWMutex 会导致 panic

### 3. Rlock

获取读操作锁，只要没有写锁，多个 goroutine 可以同时获取成功

```
func (rw *RWMutex) RLock()
```

> 读+读是并发的、读+写是阻塞的、写+读是阻塞的

### 4. RUnlock

实放一个读操作锁，当所有读操作锁都释放了，写操作锁才有机会获取

```
func (rw *RWMutex) RUnlock()
```

- 必须和 `Rlock` 成对出现
- RUnlock 多次会导致 panic

### 5. TryLock

尝试获取写锁（非阻塞），如果当前没有任何读锁或写锁，则获取写锁并返回 true，否则立即返回 false

```
func (rw *RWMutex) TryLock() bool
```

常用于后台维护/清理任务/可选更新

### 6. TruRLock

尝试获取读锁（非阻塞），如果当前没有写锁，则获取读锁并返回 true，否则立即返回 false

```
func (rw *RWMutex) TryRLock() bool
```

极少使用，一般读可以等待，多用于避免阻塞主 goroutine

### 7. RLocker

语义适配器，把读锁包装成 Locker，返回一个 Locker 接口，该接口通过调用 `rw.RLock` 和 `rw.RUnlock` 来实现 `Locker.Lock` 和
`Locker.Unlock` 方法，其 `Lock = RLock, Unlock = RUnlock`

```
func (rw *RWMutex) RLocker() Locker
```

因为狠多 API 只接受 `sync.Locker` 接口，而 Locker 接口不区分读/写操作，虽然 RWMutex 已经实现了 Locker
接口，但是实现的只是写操作锁语义，RLocker 用于实现 读锁语义