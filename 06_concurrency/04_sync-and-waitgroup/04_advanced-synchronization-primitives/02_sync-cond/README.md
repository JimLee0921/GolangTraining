# sync.Cond

sync.Cond（Condition）是一个条件变量（Condition Variable），
条件变量是用于多个goroutine之间进行阻塞等待的一种机制。
sync.Cond可以用于等待和通知goroutine，以便它们可以在特定条件下等待或继续执行。

> 是实现等待 -> 通知机制的工具，类似于：
> - 操作系统里的条件变量
> - Python 的 threading.Condition
> - Java 的 wait()/notify()

## 应用场景

| 场景        | 示例                   |
|-----------|----------------------|
| 生产者-消费者模型 | 消费者等待数据到来            |
| 等待资源可用    | 等待池中有空闲对象            |
| 控制并发      | 等待某状态切换（如任务完成、标志位改变） |
| 事件驱动      | 一组 worker 等待某个事件触发   |

## 结构与核心方法

### 结构

```
type Cond struct {
    // L 在观察或改变条件时保持
	L Locker        // 一般为 *sync.Mutex 或 *sync.RWMutex
	// 包含已过滤或未导出的字段
}
```

> L 是依赖的锁（Locker 接口，通常用 *sync.Mutex）

### 核心方法

| 方法                       | 说明                                                   |
|--------------------------|------------------------------------------------------|
| `Wait()`                 | 使当前线程进入阻塞状态，等待其他协程唤醒                                 |
| `Signal()`               | 唤醒一个等待该条件变量的 goroutine ，如果没有 goroutine 在等待，则该方法会立即返回 |
| `Broadcast()`            | 唤醒所有等待该条件变量的 goroutine，如果没有线程在等待，则该方法会立即返回           |
| `sync.NewCond(l Locker)` | 创建 `Cond` 实例的方法，`l` 通常是 `&sync.Mutex{}`              |

### 工作机制

```
+-------------+       +-------------+
| goroutine A |       | goroutine B |
|-------------|       |-------------|
| Wait() ↓    |       | Signal() ↑  |
| (等待条件)  | <---- | (条件满足)  |
+-------------+       +-------------+
```

```
var (
    // 1. 定义一个互斥锁
    mu    sync.Mutex
    cond  *sync.Cond
    count int
)
func init() {
    // 2.将互斥锁和sync.Cond进行关联
    cond = sync.NewCond(&mu)
}
go func(){
    // 3. 在需要等待的地方,获取互斥锁，调用Wait方法等待被通知
    mu.Lock()
    // 这里会不断循环判断 是否满足条件
    for !condition() {
       cond.Wait() // 等待任务
    }
    mu.Unlock()
}

go func(){
     // 执行业务逻辑
     // 4. 满足条件，此时调用Broadcast唤醒处于等待状态的协程
     cond.Broadcast() 
}
```

- 定义一个互斥锁，用于保护共享数据
- 创建一个sync.Cond对象，关联这个互斥锁
- 在需要等待条件变量的地方，获取这个互斥锁，并使用Wait方法等待条件变量被通知
- 在需要通知等待的协程时，使用Signal或Broadcast方法通知等待的协程
- 最后，释放这个互斥锁
