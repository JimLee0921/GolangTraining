# Mutex 解决资源竞争

## 什么是互斥锁

资源竞争 = 并发读写同一份可变数据且缺少同步，而锁的作用是同一时刻只允许一个 goroutine 修改（或 RWMutex 允许并发读、独占写）

互斥锁 (sync.Mutex)，适合逻辑较复杂情况：`mu := sync.Mutex`
互斥量 (mutual exclusion)保证同一时刻只有一个 goroutine 能进入临界区（共享资源的访问代码块）

Go 的锁是进程内同步原语：sync.Mutex、sync.RWMutex 的零值可用，在解锁时也可以使用 `defer` 关键字确保解锁成功

- sync.Mutex：最常用的互斥锁
- sync.RWMutex：读多写少的场景

```go
mu := sync.Mutex
mu.Lock() // 上锁
// 临界区：安全地访问或修改共享变量
mu.Unlock() // 解锁

```

## 注意事项

- 复制带锁的 struct：把含 Mutex/RWMutex 的结构体按值传参/赋值会出事，锁是结构体的状态，方法用 *T 接收者，不要复制带锁的结构体（值拷贝会复制
  Mutex，导致未定义行为或 panic）
- 重复 Unlock / 未持锁 Unlock：都会 panic
- 递归调用导致“重入”：Go 的 Mutex 不可重入，同一 goroutine 再次 Lock 会死锁
- 加锁顺序不一致：多把锁交叉时一定统一顺序（如 A->B），否则容易死锁
- Mutex 不支持取消：拿不到锁就阻塞到拿到为止，若需要“可取消的等待”，考虑channel/条件变量或设计成消息传递（actor）
- 锁的粒度：锁的范围不要太大，否则会降低并发度，粒度太小：保护不住数据，还是有竞争，粒度太大：整个程序都被串行化，失去并发意义
