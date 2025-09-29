# 接口 interface

> 具体代码见：[interface接口](../../../02_data-type/06_composite-types/05_interfaces)

## 接口的基本概念

* 接口定义了一组方法签名，**任何类型只要实现了这些方法，就隐式实现了接口**
* 接口变量可以存储 **实现了接口的任意具体类型的值**
* 多态：通过接口，可以用统一的函数操作不同的类型

示例：

```go
type shape interface {
area() float64
}

type square struct {
side float64
}

func (s square) area() float64 {
return s.side * s.side
}

func info(z shape) {
fmt.Println(z.area())
}

func main() {
s := square{10}
info(s) // square 实现了 shape 接口
}
```

---

## 空接口 `interface{}` / `any`

* 空接口没有方法 → **所有类型都实现了它**。
* 常用于表示 **任意类型**：

    * 函数参数
    * 通用集合 (`[]interface{}` / `map[string]interface{}`)
    * 标准库函数（如 `fmt.Println(...interface{})`）

Go 1.18 起有别名：

```go
type any = interface{}
```

推荐用 `any`，语义更直观

---

## 类型断言 (Type Assertion)

* 用来从接口值中取出具体类型

```go
var i interface{} = "hello"

// 不安全断言（失败 panic）
s := i.(string)

// 安全断言
s, ok := i.(string)
if ok {
fmt.Println("string:", s)
}
```

* 类型 switch：

```go
switch v := i.(type) {
case string:
fmt.Println("string:", v)
case int:
fmt.Println("int:", v)
default:
fmt.Println("unknown")
}
```

---

## 类型转换 (Type Conversion)

* 语法：`T(x)`，在**编译期**把表达式 `x` 转成类型 `T`
* 用途：`int <-> float64`，`string <-> []byte`，`rune <-> string`，`strconv` 转换

```go
a := 42
b := float64(a)

s := "hi"
bs := []byte(s) // string -> []byte
rs := []rune(s) // string -> []rune
```

⚠️ 与断言不同，转换不依赖接口，编译期检查

---

## 方法集 (Method Set)

决定一个类型是否实现接口。

* **T (值类型)**：方法集包含接收者是 `T` 的方法
* ***T (指针类型)**：方法集包含接收者是 `T` 和 `*T` 的方法

例子：

```go
type coder interface {
code()
debug()
}

type Geeker struct{}

func (g *Geeker) code()  {}
func (g *Geeker) debug() {}

var _ coder = &Geeker{} // *Geeker 实现接口
// var _ coder = Geeker{}  // Geeker 不行
```

**注意**：调用方法时，编译器会自动取地址或解引用，所以 `g.code()` 可行；
但接口赋值时必须严格满足方法集，`Geeker` 就不能直接赋给 `coder`

---

## 常见应用场景

* **多态**：不同类型实现相同接口，统一处理
* **空接口**：通用容器、工具函数
* **类型断言 / switch**：运行时区分不同实现
* **标准库**：

    * `io.Reader` / `io.Writer` 抽象 I/O
    * `sort.Interface` 实现自定义排序

---

## Conversion vs Assertion 对比表

| 特性     | 类型转换 `T(x)`                 | 类型断言 `x.(T)`        |
|--------|-----------------------------|---------------------|
| 作用对象   | 任意类型值                       | 接口值                 |
| 检查阶段   | 编译期                         | 运行期                 |
| 是否可能失败 | 不会（编译期报错）                   | 可能失败（panic / ok 判断） |
| 常见用途   | int<->float，string<->[]byte | 从接口取出具体类型           |

---

## 总结

* 接口提供多态能力
* 空接口（`any`）能表示任意类型
* 类型转换和断言语法类似，但语义不同：
    * 转换是编译期，适用于具体类型之间
    * 断言是运行期，适用于接口取值
* 方法集决定了类型能否实现接口