# sync.Once 解决资源竞争

使用 `sync.Once` 保证初始化只执行一次，且其他 goroutine 看到的是已完成的结果。
once.Do(f) 能保证 f 在多 goroutine 并发下只执行一次（且只在 f 正常返回后才算执行过）

## 注意事项

1. Do 方法接收的函数必须是无参数、无返回值的函数。如果需要传递参数，可以通过闭包或提前定义
2. Panic 处理：如果 f 中发生 panic，sync.Once 会标记 done 为 1，但不会捕获 panic。后续调用 Do 不会再执行 f
3. 不可复制：sync.Once 对象不能被复制，否则可能导致意外行为

```go
var once sync.Once
var conn *DB

func GetConn() *DB {
once.Do(func () {
conn = Dial(...)
})
return conn
}
```

## 其它方法

1. sync.OnceFunc：把“只执行一次”的逻辑封装成可复用函数
    - 把只执行一次的逻辑封装成一个可反复调用的函数
    - 签名：func OnceFunc(f func()) func()
    - 适合场景：一次性初始化（加载配置、预热缓存、建立连接），但希望在多处/多次调用同一个入口
2. sync.OnceValue：懒加载并缓存单个返回值
    - 作用：懒加载并缓存一个返回值；后续调用直接返回缓存
    - 签名：func OnceValue[T any](f func() T) func() T
    - 适合场景：只需计算一次的纯值（如版本号、时间戳快照、重型配置对象等）
3. sync.OnceValues：懒加载并缓存多个返回值（常用于 value + error）
    - 作用：懒加载并缓存多个返回值（最常见是“值 + error”）
    - 签名：`func OnceValues[T1, T2 any](f func() (T1, T2)) func() (T1, T2)`
    - 适合场景：一次性初始化但需要携带错误或多个返回值的场景（例如加载配置/密钥、建连接并返回句柄+err）

> 与经典 sync.Once 相比：这些 API 更易组合（直接得到函数/取值器），避免显式保存 Once 与结果的样板代码

## 总结

- sync.Once 适合 “只初始化一次” 的场景（配置、连接、缓存预热等）
- 不要把 sync.Once 复制或通过值传递（会破坏内部状态）；通常定义为包级变量或结构体字段
- 如果初始化可能失败并需要返回错误，可以在 Do 里设置一个包级变量 initErr，外层调用后检查它
- Go 1.21+ 还提供了 sync.OnceFunc / sync.OnceValue（把函数包装成只执行一次、或只计算一次返回值的函数/工厂），但基础的
  sync.Once 已足够覆盖大多数需求