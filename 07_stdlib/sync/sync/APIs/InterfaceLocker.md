# sync.Locker

`sync.Locker` 是 sync 包中最小抽象，是把锁抽象为一个统一接口，让不同的同步原语能够进行协作

## 定义

```
type Locker interface {
	Lock()
	Unlock()
}
```

`sync.Locker` 是对互斥性为最小的抽象，一共定义两个方法：Lock（能加锁）/ Unlock（能解锁）

## 类型实现

标准库中有下面这些 `sync.Locker` 的实现：

- `*sync.Mutex`
- `*sync.RWMutex`（写锁语义）
- `rw.RLocker()`（读锁语义的适配器）

主要是使用到 `sync.Cond`，它内部要求是一个 `sync.Locker` 接口的实现，

`sync.Cond` 的构造函数 `func NewCond(l Locker) *Cond` 的语义是并不关心用什么锁，只要求：

- Wait 前能 Lock
- Wait 中能 Unlock
- 被唤醒后能重新 Lock