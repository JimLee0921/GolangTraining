# sync.Pool

`sync.Pool` 是一个用于临时对象复用的并发安全对象池，目的是减少内存分配和 GC 压力，可以供多个 goroutine 同时使用

```
type Pool struct {

	// New optionally specifies a function to generate
	// a value when Get would otherwise return nil.
	// It may not be changed concurrently with calls to Get.
	New func() any
	// contains filtered or unexported fields
}
```

## 并发性能问题

比如一段高并发程序中有类似这段代码：

```
for i := 0; i < N; i++ {
    buf := make([]byte, 4096)
    use(buf)
}
```

这段代码可能会影响性能：

- 每次循环都使用 make 创建
- 产生大量短生命周期对象
- GC 频繁扫描，回收
- 吞吐下降，延迟抖动

如果使用传统缓存对象，手写一个全局缓存，需要 加锁/控制大小/防止内存泄露/与 GC 协作，很容易写错

`sync.Pool` 解决的本质问题就是如何在并发环境下安全地复用短生命周期、可重置的临时对象，并且让 GC 在需要时随时清空这些临时对象

## 核心思想

可以把 `sync.Pool` 视为 GC 友好的临时对象缓存，核心思想是程序尽量复用对象，GC 可以随时丢弃池中对象，不保证对象一定来自于
Pool 池，Put 进去的下次也不一定会被 Get 到

`sync.Pool` 主要适合创建成本高、使用后可 Reset、声明周期短的对象：

- `[]byte / bytes.Buffer`
- `bytes.Reader`
- 编解码中间对象
- 临时 struct（request context, parser state）

不适合长期持有的对象/带状态且不可重置/需要`close/release`的对象/很小的对象（int、小 struct）

**注意事项**

1. Pool 里的对象可能随时消失，当 GC 回收时 Pool 中的对象可能会被全部清空，所有不能假设 Put 的对象下一次一定可以 Get 到
2. Pool 不是容量可控的缓存，没有 size，没有上限也没用淘汰策略
3. Pool 时并发安全的，多 goroutine 可以同时 Get/Put，内部已经做了并发优化

## 核心方法

一共就两个方法：`Get` / `Put`

### 1. Get()

Get 方法用于从 Pool 中返回一个可用对象，不承诺从哪里来，如果 Pool 中没有可用对象：

1. 如果设置了 `New()`，则通过 `p.New()` 进行创建
2. 否则返回 nil

```
func (p *Pool) Get() any
```

使用时必须假设返回的对象是新建的或者是被别人用过的，所以一定要重置，也就是 Get 只会必须做的一件事就是把对象当成脏的，立即
Reset
> 并发安全 / 不堵塞 / 不保证公平 / 不保证命中

### 2. Put()

归还一个对象，把对象 x 放回 pool，作为未来可能被复用的候选对象

```
func (p *Pool) Put(x any)
```

**行为边界**

1. 并发安全
2. 不阻塞
3. GC 可能随时清空 Pool
4. Put 之后不能再使用 x

Put 之前必须做的一件事是把对象恢复到干净可复用的状态，这是 Pool 能安全工作的前提