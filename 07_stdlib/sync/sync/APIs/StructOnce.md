# sync.Once

`sync.Once` 用于保证某个函数在并发环境下只执行一次，Once 对象只会执行一次操作，首次使用后不能复制

```
type Once struct {
	// contains filtered or unexported fields
}
```

## 解决问题

在单线程中：

```
if !inited {
    initConfig()
    inited = true
}
```

这段代码是没有问题的，但是在并发场景下：

- 多个 goroutine 同时看到 `inited == false`
- `initConfig()` 被执行多次
- 导致状态被破坏

Once 解决的本质问题就是如何在并发条件下保证：

- 初始化逻辑只执行一次
- 其它 goroutine 要么等待初始化逻辑完成，要么看到已完成状态
- 不需要自己写锁、标志位、双重检查

## 原理

1. Once 保证严格只执行一次：
    - `once.Do(f)` 无论调用多少从，来自多少个 goroutine，`f()` 最多只会被执行一次
2. 完成可见性：`Do(f)` 返回时，`f()` 中的所有写操作对后续 goroutine 都是可见的
3. 自动并发安全：内部加锁，内部状态管理，不需要手动写 if/flag/mutex

```
var once sync.Once

func initConfig() {
    fmt.Println("init config")
}

func handler() {
    once.Do(initConfig)
}
```

无论有多少个 `handler()` 并发调用， `initConfig` 都只会执行一次，`init config` 只会被打印一次

## 使用场景

1. 单例初始化
2. 懒加载资源：配置文件/连接池/缓存/全局 map
3. 全局一次性动作：注册 handler/启动后台 goroutine/ 初始化 metrics
4. 不适合需要重置/可能失败、需要重试的初始化

## 核心方法

只有一个核心方法 `Do()`，once 入口方法，保证在并发环境下，`f()` 在整个进程生命周期中最多只执行一次

```
func (o *Once) Do(f func())
```

最多执行一次，第一次调用会有某个 goroutine 执行 `f()`，后续所有调用会直接返回，不再执行 `f()`

**panic 行为**

如果 `f()` 导致 panic，panic 也会被视为执行过了：

- panic 会向上传播
- Once 的状态会被标记为已执行
- 后续 Do 调用不会再执行 f

> 这是 Once 的一个强设计决策，所以 Do 时不可重试的，Once 没有 reset 能力
