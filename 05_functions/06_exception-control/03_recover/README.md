## 简介

> `recover` 是 Go 提供的一个**内建函数**，用于 **捕获 panic**，防止程序崩溃

* 只能在 **`defer` 函数中调用**
* 如果当前 goroutine 处于 panic 状态，`recover()` 会：
    * 拦截 panic
    * 返回 panic 的值
    * 让程序从恐慌中“恢复”
* 如果当前没有 panic，`recover()` 返回 `nil`

---

## 基本语法

```go
func recover() any
```

使用示例：

```go
defer func () {
if r := recover(); r != nil {
fmt.Println("Recovered:", r)
}
}()
```

> 只有在 defer 中调用 `recover()` 才能生效，
> 直接在普通函数中调用不会捕获 panic

---

## 捕获 panic

```go
package main

import "fmt"

func main() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered from:", r)
		}
	}()
	fmt.Println("Before panic")
	panic("something went wrong")
	fmt.Println("After panic") // 不会执行
}
```

输出：

```
Before panic
Recovered from: something went wrong
```

程序没有崩溃，`recover` 捕获到了 panic 并“恢复执行”

---

## recover 的运行机制

### 执行流程图：

```
panic("boom")
↓
执行当前函数 defer
↓
defer 中调用 recover() -> 捕获 panic
↓
停止 panic 向上传播
↓
继续执行 defer 之后的代码（如果有）
↓
程序正常退出
```

## recover 只能在 defer 中生效

下面的写法不会起作用：

```go
func main() {
recover() // 无效，程序仍然崩溃
panic("boom")
}
```

必须写成：

```go
defer func () {
recover() // 有效
}()
panic("boom")
```

---

## recover 与返回值

如果 panic 被捕获，程序会从 defer 执行完毕后正常返回

```go
func f() (result int) {
defer func () {
if r := recover(); r != nil {
fmt.Println("Recovered:", r)
result = -1 // 修改返回值
}
}()
panic("failed")
return 1
}

func main() {
fmt.Println("Result:", f())
}
```

输出：

```
Recovered: failed
Result: -1
```

说明：

* panic 发生 -> recover 捕获
* defer 执行完后函数继续返回
* 可以在 defer 中修改命名返回值（如 `result`）

---

## goroutine 中使用 recover

panic 不会跨 goroutine 传播，
如果在 goroutine 里 panic，没有 recover，会直接崩溃整个程序

因此要在 goroutine 内部包一层：

```go
func safeGo(fn func ()) {
go func () {
defer func () {
if r := recover(); r != nil {
fmt.Println("goroutine recovered:", r)
}
}()
fn()
}()
}

func main() {
safeGo(func () {
panic("worker panic!")
})
time.Sleep(time.Second)
fmt.Println("main still alive")
}
```

输出：

```
goroutine recovered: worker panic!
main still alive
```

panic 只在当前 goroutine 中传播，因此要在每个 goroutine 中单独 recover。

---

## 常见使用场景

| 场景            | 示例                        |
|---------------|---------------------------|
| 程序最外层保护（防止崩溃） | Web 服务主循环                 |
| goroutine 保护  | 并发任务                      |
| 框架内部          | 捕获用户逻辑中的 panic 并转化为 error |
| 日志系统          | 打印 panic 堆栈后继续运行          |

---

## 错误与 panic 的区分

| 比较项  | `error`         | `panic` + `recover`  |
|------|-----------------|----------------------|
| 处理方式 | `if err != nil` | `defer + recover()`  |
| 用途   | 可预期的业务错误        | 程序级严重异常              |
| 可恢复性 | 可以              | 一般不可恢复，但 recover 可救场 |
| 推荐使用 | 一般业务逻辑          | 底层框架、容错保护            |

---

## 小结

| 关键点             | 说明                |
|-----------------|-------------------|
| 只能在 `defer` 中调用 | 否则无效              |
| 返回 panic 值      | 如果当前在 panic 状态    |
| 捕获后终止传播         | 程序恢复正常执行          |
| 常用于防止崩溃         | 保护 goroutine、主循环等 |


