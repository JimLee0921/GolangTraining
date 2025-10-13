## 函数表达式（Function Expression）

函数表达式（Function Expression）是指：在表达式中直接定义函数值，可以赋给变量、传递、或立即执行，
也就是： 把函数当作值（value） 使用，而不是独立声明。

### 对比一下普通定义

| 类型                             | 示例                                         | 说明            |
|--------------------------------|--------------------------------------------|---------------|
| **函数声明（Function Declaration）** | `func add(a, b int) int { return a + b }`  | 普通命名函数，定义在包级别 |
| **函数表达式（Function Expression）** | `f := func(a, b int) int { return a + b }` | 匿名函数赋给变量或直接使用 |

表达式意味着它可以出现在任何可以用值的地方。

### 基本语法

```go
f := func (x int, y int) int {
return x + y
}
```

* `func(...) {...}` 是一个匿名函数表达式
* 没有函数名
* 返回一个函数值
* 可以赋值给变量、传参、或立即执行

### 使用方式

| 用法             | 示例                                           | 说明        |
|----------------|----------------------------------------------|-----------|
| **赋值给变量**      | `f := func() { fmt.Println("hi") }`          | 最常见       |
| **作为参数传递**     | `run(func() { fmt.Println("hello") })`       | 回调函数      |
| **立即执行（IIFE）** | `func() { fmt.Println("start") }()`          | 临时逻辑块     |
| **返回值为函数**     | `return func(x int) int { return x + base }` | 用于闭包与工厂函数 |

### 常见用途

| 场景        | 示例                                              | 说明     |
|-----------|-------------------------------------------------|--------|
| **回调函数**  | `operate(func(a,b int){...})`                   | 动态逻辑传入 |
| **闭包实现**  | `return func(){n++}`                            | 捕获外部变量 |
| **延迟执行**  | `defer func(){cleanup()}()`                     | 延迟逻辑   |
| **初始化逻辑** | `func(){loadConfig()}()`                        | 临时代码块  |
| **函数工厂**  | `func(prefix string) func(string) string {...}` | 生成定制函数 |

### 函数表达式的类型

函数表达式本质上是有类型的。
例如：

```go
var fn func (int, int) int
fn = func (a, b int) int { return a + b }
```

* `func(int, int) int` 是一个函数类型
* 匿名函数赋值时类型必须匹配
* 函数类型可以被显式声明或推断
