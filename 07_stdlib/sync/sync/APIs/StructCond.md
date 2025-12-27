# sync.Cond

`sync.Cond` 用来让 goroutine 在某个条件不满足时等待，在条件满足时唤醒

```
type Cond struct {

	// L is held while observing or changing the condition
	L Locker
	// contains filtered or unexported fields
}
```

## 解决的问题

Mutex 解决的时不能同时做的问题，而不是解决什么时候开始做的问题

Cond 解决的是如何在持有锁的前提下：

- 如何让 goroutine 安全地进入等待状态
- 如何让 goroutine 在条件变化时被唤醒
- 如何让 goroutine 重新检查并继续执行

Cond = Mutex + 条件队列 + 通知机制，内部会根据传入的锁做三件事

1. 原子地释放锁
2. 把当前 goroutine 放入等待队列
3. 被唤醒后重新加锁

## 应用场景

1. 生产者/消费者
    - 队列为空->消费者等待
    - 生产者释放数据->通知消费者
2. 资源池：连接池 / worker 池/ token bucket
3. 状态机等待：等待 ready / 等待 leader / 等待配置加载完成

**对比 channel**

- channel 是值推动：没值就等，有值就走
- Cond 是条件驱动：条件不满足就 Wait，条件满足就继续

> Cond 符合复杂条件，channel 适合数据流，能用 channel 就优先使用 channel

## 构造函数 `sync.NewCond`

用于 Cond 初始化，创建一个 Cond 条件变量，并把它绑定到一个互斥锁 Locker 上，Cond 自己没有锁，需要自行传入：

```
func NewCond(l Locker) *Cond
```

**使用示例**

```
mu := sync.Mutex{}
cond := sync.NewCond(&mu)

// RWMutex 创建，等待时用读锁语义，允许多个 Waiter 并发等待
var rw sync.RWMutex
cond := sync.NewCond(rw.RLocker())
```

## 核心方法

一共三个核心方法

### 1. Wait()

用于等待条件成立（最核心方法）

```
func (c *Cond) Wait()
```

**完整语义**

1. 要求调用方已经持有锁
2. 原子地释放锁
3. 把当前 goroutine 挂起（进入等待队列）
4. 被 Signal / Broadcast 唤醒
5. 重新获取锁
6. 返回给调用方

> Wait 返回时，锁一定时被重新持有的

Wait 必须放在 for 循环中，这是 Cond 第一铁律：

```
mu.Lock()

for !condition {
    cond.Wait()
}

doSomething()
mu.Unlock()
```

- 防止虚假唤醒：Cond 不保证每次唤醒条件都真的成立
- 防止广播唤醒多个 goroutine： `cond.Broadcast()` 多个 goroutine 被同时唤醒，只有一个能真正满足条件
- 防止条件被其他 goroutine 抢先改变：即使被唤醒，条件也可能已经再次不成立，所以必须重新检查条件

### 2. Signal

从等待队列中，唤醒至少一个等待者 goroutine，不保证唤醒的哪个等待着，不保证顺序，不保证立即运行

```
func (c *Cond) Signal()
```

Signal 把一个等待中的 goroutine 标记为可运行，不释放锁，不修改条件，不保证唤醒后条件成立，推荐在持锁状态下进行调用：

```
mu.Lock()
condition = true
cond.Signal()
mu.Unlock()
```

可以保证条件修改，通知，在同一临界区内完成，避免竞态

### 3. Broadcast()

唤醒当前等待队列中的所有 goroutine 等待者

```
func (c *Cond) Broadcast()
```

主要用于状态发生全局变化，所有等待者都应该重新检查条件，但是会唤醒大量 goroutine，性能成本较高，原则上能使用 Signal 就不要使用 Broadcast 

