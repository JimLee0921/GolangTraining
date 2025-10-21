# sync.Once 解决资源竞争

sync.Once 是 Go 并发包 (sync) 中的一个同步原语，
它的作用是：保证某个操作只执行一次，无论有多少个 goroutine 同时调用它。

比如初始化数据库连接、加载配置文件、注册日志系统、启动后台任务等。

是实现单例模式（Singleton）或懒加载（Lazy Initialization）的首选工具。

## 基本使用

sync.Once 只有一个方法：

| 方法             | 说明                     |
|----------------|------------------------|
| `Do(f func())` | 仅第一次调用时执行函数 f，后续调用立即返回 |

1. Do 方法接收的函数必须是无参数、无返回值的函数。如果需要传递参数，可以通过闭包或提前定义
2. Panic 处理：如果 f 中发生 panic，sync.Once 会标记 done 为 1，但不会捕获 panic。后续调用 Do 不会再执行 f
3. 不可复制：sync.Once 对象不能被复制，否则可能导致意外行为

## 工作原理

### 内部结构

```
type Once struct {
    done uint32
    m    Mutex
}
```

- done：标记是否已执行（0 未执行，1 已执行）
- m：互斥锁，防止并发进入执行区域

### 执行流程伪代码

```
func (o *Once) Do(f func()) {
    if atomic.LoadUint32(&o.done) == 0 {
        o.doSlow(f)
    }
}

func (o *Once) doSlow(f func()) {
    o.m.Lock()
    defer o.m.Unlock()
    if o.done == 0 {
        defer atomic.StoreUint32(&o.done, 1)
        f()
    }
}
```

- 用原子变量判断是否已执行
- 若未执行，进入加锁路径
- 第一次执行成功后 done = 1
- 之后所有 goroutine 都直接返回，不再调用

## 其它方法

Go 1.21 引入了 OnceValue / OnceFunc，是对传统 sync.Once 的泛型化与函数封装升级版，
让只执行一次的逻辑更优雅、更实用。

传统 sync.Once 有两个明显限制：

- 只能调用函数，但不能返回值
- 如果想获取初始化结果（比如配置对象），需要额外声明全局变量

> 与经典 sync.Once 相比：这些 API 更易组合（直接得到函数/取值器），避免显式保存 Once 与结果的样板代码

### sync.OnceValue

带返回值的 Once，用来延迟执行一个函数，并缓存其d单个返回值。
后续调用会直接返回之前的结果，不会重复执行。
懒加载并缓存一个返回值；后续调用直接返回缓存

`func OnceValue[T any](fn func() T) func() T`

- 接收一个返回值的函数 fn
- 返回一个新的函数
- 该函数在第一次调用时执行 fn()，并缓存结果
- 之后所有调用都直接返回第一次的结果。

> 适合场景：只需计算一次的纯值（如版本号、时间戳快照、重型配置对象等）

### sync.OnceValues

sync.OnceValue 的变体，用于支持 返回多个值 的函数（即多返回值版本），
懒加载并缓存多个返回值（常用于 value + error）

`func OnceValues[T1, T2 ... any](f func() (T1, T2, ...)) func() (T1, T2, ...)`

> 适合场景：一次性初始化但需要携带错误或多个返回值的场景（例如加载配置/密钥、建连接并返回句柄+err）
>
> Go 官方在标准库中自动生成了 OnceValues 的 1~8 返回值版本

### sync.OnceFunc

把只执行一次的逻辑封装成一个可反复调用的函数，类似 sync.Once 的 .Do()，但更简洁，可直接传递

`func OnceFunc(f func()) func()`

- 接收一个普通函数
- 返回一个包装函数
- 第一次调用时执行 fn()
- 后续调用立即返回，不再执行

> 适合场景：一次性初始化（加载配置、预热缓存、建立连接），但希望在多处/多次调用同一个入口



## 总结

- sync.Once 适合只初始化一次的场景（配置、连接、缓存预热等）
- 不要把 sync.Once 复制或通过值传递（会破坏内部状态），通常定义为包级变量或结构体字段
- 如果初始化可能失败并需要返回错误，可以在 Do 里设置一个包级变量 initErr，外层调用后检查它
- Go 1.21+ 还提供了 sync.OnceFunc / sync.OnceValue（把函数包装成只执行一次、或只计算一次返回值的函数/工厂），但基础的
  sync.Once 已足够覆盖大多数需求