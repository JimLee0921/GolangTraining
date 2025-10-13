## 什么是 `panic`

在 Go 中：

> `panic` 会 **立即终止当前函数的正常执行流程**
> 并沿调用栈逐层向上传递，直到：
>
> * 被 `recover()` 捕获
> * 或者传播到 `main()` 导致程序崩溃

它相当于其他语言的：

* `throw`（Java）
* `raise`（Python）
* `throw Exception`（C++）

但 Go 的哲学是：
panic 只用于 **不可恢复的错误（fatal error）**，而非普通业务错误

---

## 的基本语法

```go
panic("错误原因描述")
```

* 参数可以是任意类型（常见为 `string` 或 `error`）
* 一旦触发，当前函数立即停止执行
* Go 会执行当前函数的所有 `defer`
* 然后把 panic 向上层函数传播

---

## 简单示例

```go
package main

import "fmt"

func main() {
	fmt.Println("before panic")
	panic("something went wrong")
	fmt.Println("after panic") // 不会执行
}
```

输出：

```
before panic
panic: something went wrong
```

解释：

1. 执行到 `panic()` -> 程序进入“恐慌”状态
2. main 函数没有 `recover()`
3. Go 打印堆栈并退出程序。

---

## panic 的传播机制

假设调用栈如下：

```
main -> A -> B -> C
```

如果 `C` 发生 panic：

1. Go 会执行 `C` 的所有 defer
2. 若未 recover，则退出 `C` 并将 panic 传给 `B`
3. 执行 `B` 的 defer
4. 若 `B` 也未 recover，则继续向上
5. 直到传到 `main`。若仍没人 recover -> 程序崩溃

---

### 调用栈传播

```go
func main() {
fmt.Println("main start")
A()
fmt.Println("main end") // 不会执行
}

func A() {
defer fmt.Println("A defer")
B()
}

func B() {
defer fmt.Println("B defer")
panic("something wrong in B")
}
```

输出：

```
main start
B defer
A defer
panic: something wrong in B
```

顺序说明：

* panic 在 `B` 发生
* 先执行 `B` 的 defer
* 再执行 `A` 的 defer
* 最后崩溃打印堆栈

## panic 与 defer 的协作

即使 panic 发生，`defer` 依然会被调用。

```go
func main() {
defer fmt.Println("defer 1")
defer fmt.Println("defer 2")
panic("boom")
}
```

输出：

```
defer 2
defer 1
panic: boom
```

**panic 不会跳过 defer** —— 所有已注册的 defer 一定执行。
这就是为什么可以在 defer 里用 `recover()` 来拦截 panic

---

## panic 典型触发场景

1. 手动触发严重错误
2. 系统自动触发的 panic

这些是 Go runtime 自动引发的：

| 错误类型     | 示例代码                                   |
|----------|----------------------------------------|
| 数组越界     | `arr := []int{1}; fmt.Println(arr[2])` |
| nil 指针调用 | `var p *int; fmt.Println(*p)`          |
| 类型断言错误   | `v := any(123); s := v.(string)`       |
| 并发错误     | 修改已关闭的 channel、重复关闭 channel            |

---

## panic 与普通错误 (`error`) 的区别

| 对比项    | `panic`           | `error`         |
|--------|-------------------|-----------------|
| 用途     | 致命、不可恢复错误         | 可预期、可恢复的业务错误    |
| 处理方式   | `recover()` 或程序崩溃 | `if err != nil` |
| 影响范围   | 打断当前函数并向上传播       | 仅当前逻辑分支         |
| 是否推荐滥用 | 不推荐               | 推荐作为常规错误处理      |

一般业务逻辑中不要用 `panic`，而应返回 `error`。

---


## panic 的执行流程总结

```
调用 panic
↓
执行当前函数 defer（若有）
↓
逐层向上传播
↓
若 defer 中有 recover -> 捕获并停止传播
↓
若无人 recover -> 打印堆栈 -> 程序崩溃
```


