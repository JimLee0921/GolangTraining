# Go 常量

## 基本概念

- 使用 `const` 定义常量。
- 常量的值在 **编译期确定**，运行时不可修改。
- 只能是 **布尔、数值、字符串** 三种类型。
- 常量的零值概念不存在，必须显式赋值。

示例：

```go
const Pi = 3.14159
const Greeting = "Hello, Go!"
const IsFun = true
```

## 使用案例

---

### 常量 vs 变量

* **变量 (var)**：运行时可修改。
* **常量 (const)**：编译期固定，运行时不可修改。

```go
var x = 10
const y = 20

x = 30 // 可以
// y = 40 // 编译错误
```

---

### 有类型 & 无类型常量

* **有类型常量**：只能赋给相同类型变量。

```go
const a int = 10
var x int = a     // 可以
var y float64 = a // 错误
```

* **无类型常量**：可以延迟确定类型，更灵活。

```go
const b = 10
var x int = b
var y float64 = b // 可以自动转换
```

---

### 常量表达式

常量可以由表达式生成，只要编译期能确定：

```go
const (
A = 1 + 2 // 3
B = "Go" + "lang" // "Golang"
C = 2 << 3        // 16
)
```

---

## `iota` 关键字

* `iota` 是一个在 `const` 块里自动递增的计数器。
* 每个 `const` 块的 `iota` 从 0 开始，每行 +1。
* 常用于枚举、位标志。

### 枚举

```go
const (
Sunday = iota // 0
Monday        // 1
Tuesday       // 2
)
```

### 位标志

```go
const (
Read = 1 << iota // 1
Write            // 2
Exec // 4
)
```

组合使用：

```go
perm := Read | Write // 3
if perm&Read != 0 {
fmt.Println("has read permission")
}
```

### 跳过值

```go
const (
_ = iota
KB = 1 << (10 * iota) // 1024
MB                    // 1048576
GB                    // 1073741824
)
```

---

## 特点

1. 常量不能用变量赋值，只能用字面量或常量表达式。
2. 不能是 slice/map/struct 等复合类型。
3. 多个常量在 `const` 块中可以省略表达式，自动继承上一行。

---

## 常见用法

* 定义数学常量：`const Pi = 3.14`
* 定义应用配置：`const AppName = "MyApp"`
* 定义枚举值：配合 `iota`
* 定义位标志：权限/状态管理


