# sync 包顶层函数

sync 包主要就三个顶层函数`OnceFunc / OnceValue / OnceValues`，都是 Go1.21+ 使用，这三个顶层函数主要都是对 `sync.Once.Do`
的函数式封装，把 `Once.Do` 变成一个函数闭包让 Once 的用法变得更安全、简介、可组合

> OnceFunc / OnceValue / OnceValues 把一次性执行的语义，直接封装进返回的函数里，不用再显式管理 `sync.Once` 和共享变量

## 1. `sync.OnceFunc`

把一个只需要执行一次的动作进行包装，包装成一个无参、无返回值、可安全并发调用的函数，返回一个只调用 `f()` 一次的函数，返回的函数可以并发调用

```
func OnceFunc(f func()) func()
```

等价于：

```
once := sync.Once{}

g := func() {
    once.Do(f)
}
```

> OnceFunc 把 once 和模板代码都藏了起来

## 2. OnceValue

把只计算一次的值封装成一个函数，第一次调用计算，后续直接返回缓存结果（并发安全），这是并发安全的 lazy value

```
func OnceValue[T any](f func() T) func() T
```

等价于：

```
var (
    once sync.Once
    v   T
)

g := func() T {
    once.Do(func() {
        v = f()
    })
    return v
}
```

## `sync.OnceValues`

和 OnceValue 一样，但支持多个返回值（最常见是 `value+error`）

```
OnceValues[T1, T2 any](f func() (T1, T2)) func() (T1, T2)
```

等价于：

```
var (
    once sync.Once
    v1   T1
    v2   T2
)

g := func() (T1, T2) {
    once.Do(func() {
        v1, v2 = f()
    })
    return v1, v2
}
```