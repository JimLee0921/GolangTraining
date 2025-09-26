# Go 变量

## 使用 `var` 定义变量

标准写法：

```go
var name type = value


例子：

```go
var age int = 20
var msg string = "Hello"

```

* 类型可以省略，编译器自动推断：

```go
var x = 42       // 自动推断为 int
var y = "golang" // 自动推断为 string
```

* 如果没有赋值，会给一个 **零值**：

    * 数字类型 → `0`
    * 布尔类型 → `false`
    * 字符串 → `""`

---

## 2. 使用 `:=` 简短声明

在函数内部可以用 `:=` 省略 `var`，自动推断类型：

```go
func main() {
x := 10      // int
y := "hello" // string
z := true // bool
}
```

注意：`:=` **只能在函数内部用**，不能在包级作用域使用。

---

## 3. 多变量定义

```go
var a, b, c = 1, 2, 3

func main() {
x, y := 10, "hi"
fmt.Println(x, y) // 10 hi
}
```

---

**总结：**

* **`var`** → 适合在包级和函数内定义，可以带类型或自动推断。
* **`:=`** → 只能在函数内使用，更简洁，自动推断类型。

