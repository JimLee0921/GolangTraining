## 基本定义

> `defer` 语句用于**延迟执行一个函数调用**，直到外层函数返回时才会被调用。

* `defer` 语句只能出现在函数内部
* 当所在函数**正常返回**或**发生 panic** 时，defer 都会被执行
* 常用于：**资源释放、日志记录、解锁、错误恢复**等
* `defer` 语句在函数退出前执行
* 多个 `defer` 是 **后进先出（LIFO）** 执行顺序，所以先注册的最后执行

---

## defer 的典型应用场景

- 场景 1：资源清理（关闭文件），无论函数正常返回还是 panic，`f.Close()` 都一定会执行。
- 场景 2：锁释放，defer 可以保证锁总能释放
- 场景 3：trace 追踪成对操作（开始/结束）

> 技巧：`defer trace("func")()`
> 先调用 `trace("func")`，返回一个匿名函数；`defer` 延迟执行它，实现优雅的进入/退出日志

## defer 的参数求值时机

> defer 的 **参数在注册时求值**，不是在执行时。

```go
func main() {
x := 10
defer fmt.Println("x =", x)
x = 20
}
```

输出：

```
x = 10
```

解释：

* 当执行 `defer fmt.Println("x =", x)` 时，`x` 的值就被拷贝成 10；
* 后续 `x = 20` 不会影响打印。

---

## defer 与 return 的执行顺序

```go
func f() (result int) {
defer func () {
result++
}()
return 10
}

func main() {
fmt.Println(f())
}
```

输出：

```
11
```

执行顺序：

1. `return 10` 先把返回值 `result = 10`
2. 调用 defer
3. defer 内把 `result` 加 1
4. 最终返回 11

> 如果返回值不是命名返回值（例如 `return 10` 没声明 `result`），`defer` 修改不了它

---

## defer + panic

即使函数中途 panic，defer 也一定会执行

```go
func main() {
defer fmt.Println("defer still runs")
panic("something went wrong")
}
```

输出：

```
defer still runs
panic: something went wrong
```

所以在资源释放、清理逻辑中用 defer 是非常安全的

---

## defer 与循环

defer 不能使用在循环中

```go
for i := 0; i < 3; i++ {
defer fmt.Println(i)
}
```

输出：

```
2
1
0
```

每次循环都会注册一个 defer，但都在函数退出时才执行（不是每轮循环执行），
想要立即执行的逻辑，不要用 defer 放在循环里

---

## 小结

| 特性         | 说明                   |
|------------|----------------------|
| 执行时机       | 函数返回前（正常返回或 panic）   |
| 执行顺序       | 后进先出（LIFO）           |
| 参数求值       | 在注册 defer 时求值        |
| 常见用途       | 资源清理、锁释放、日志记录、安全退出   |
| 与 panic 关系 | panic 触发时 defer 仍会执行 |

