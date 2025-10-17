# 接口（Interface）

## 概述

接口（Interface）是 Go 中用于定义类型行为规范的机制，定义了一组方法签名（方法的名字、参数、返回值），但不包含具体实现。

它描述的是一组方法的集合（method set），
任何类型只要实现了接口里定义的所有方法，都自动满足该接口，无需显式声明。
在 Go 里没有 implements 或 extends 这样的关键字，接口是 隐式实现 的。


> 接口是一种行为契约（contract），它定义能做什么，而不是是什么。

---

## 接口的定义

```
type InterfaceName interface {
    Method1(paramList) returnType
    Method2(paramList) returnType
}
```

示例：

```
type Reader interface {
    Read(p []byte) (n int, err error)
}
```

> 一个接口就是方法签名的集合。

---

## 接口实现机制

Go 的接口采用隐式实现（implicit implementation）。即只要某个类型定义了接口要求的全部方法，就自动实现该接口。

但是如果一个接口里定义了多个方法，那么类型必须实现接口里所有的方法，才能被认为实现了这个接口，
如果只实现部分方法，Go 编译器会直接报错，不会部分匹配。

示意：

```
type Dog struct{}
func (d Dog) Speak() string { return "Woof" }

type Speaker interface {
    Speak() string
}

var s Speaker = Dog{} // 无需声明 implements 
```

没有 implement 关键字，Go 通过编译器静态检查匹配方法集

| 特性   | 说明                 |
|------|--------------------|
| 抽象行为 | 只描述方法，不关心数据结构      |
| 隐式实现 | 无需显式声明 implements  |
| 动态类型 | 运行时可存储任意实现类型       |
| 多态支持 | 不同类型可通过接口表现统一行为    |
| 值语义  | 接口本身是值类型（包含类型信息和值） |

---

## 接口底层原理

在 Go 的运行时中，每个接口变量（非空接口）可以理解为一个装了两样东西的盒子

```
interface value = {
    dynamic type  // 运行时真实类型信息
    dynamic value // 运行时真实的值（可以是值也可以是指针）
}
```

内存图示例：

```text
┌────────────────────────┐
│ interface (s)          │
│ ├── dynamic type: Dog  │
│ └── dynamic value: { } │
└────────────────────────┘
```

### reflect 反射

反射包 `reflect` 基于接口实现。

* `reflect.TypeOf(i)` -> 动态类型
* `reflect.ValueOf(i)` -> 动态值

这两个函数仅接受接口参数，因此：

> 反射是接口机制的延伸与底层实现。

可以使用 reflect 查看内部结构：

```text
package main
import (
    "fmt"
    "reflect"
)

func main() {
    var x interface{} = (*int)(nil)
    fmt.Println(x == nil)                        // false
    fmt.Printf("Type: %v, Value: %v\n", reflect.TypeOf(x), reflect.ValueOf(x))
}
```

* 动态类型是 *int
* 动态值是 nil
* 所以接口整体 != nil

### 零值与空接口

| 类型                        | 含义                 |
|---------------------------|--------------------|
| `nil` 接口                  | 未包含任何类型和值          |
| 空接口 `interface{}`         | 不包含任何方法的接口，可容纳任意类型 |
| `any` 完全等价于 `interface{}` | Go 1.8 之后空接口的写法    |

* `nil` 是值状态
* `interface{}` 和 `any` 是类型定义

示意：

```
var x interface{}   // 可存放任意类型
var x any   // any 与 interface{} 完全等价
x = 10
x = "Go"
x = []int{1, 2, 3}
```

### 接口的比较与零值

| 情况         | 示例                              | 结果    | 原因         |
|------------|---------------------------------|-------|------------|
| 相同类型 + 相同值 | `10 == 10`                      | true  | 类型和值都相等    |
| 相同类型 + 不同值 | `10 != 20`                      | false | 值不同        |
| 不同类型 + 同值  | `int(10) vs int64(10)`          | false | 类型不同       |
| 接口为 nil    | `var i interface{} = nil`       | true  | 类型和值都为 nil |
| 接口中值为 nil  | `var i interface{} = (*T)(nil)` | false | 类型非空       |
| 不可比较类型     | `[]int{1} == []int{1}`          | panic | slice 不可比较 |

### 设计理念

* **小接口原则（Small Interface Rule）**

    * 每个接口只定义最小必要行为；
    * 比如 `io.Reader`, `io.Writer`；
* **组合优于继承**

    * 通过嵌入接口组合行为；
* **隐式实现**

    * 减少耦合、提升可复用性。

---

## 方法集

方法集（Method Set） 就是指：某个类型能调用的所有方法的集合，它决定了一个类型是否满足某个接口。
对于 值类型 T：方法集只包含接收者是 T 的方法
对于 *指针类型 T：方法集包含接收者是 T 和 *T 的方法
T 的方法集是 *T 的方法集的子集，有些接口，必须用 *T 才能实现

## 接口值接收者和指针接收者

当类型 `T` 赋给接口变量时，编译器会检查 `T` 的方法集是否包含接口要求的所有方法。

- 值接收者定义的方法 -> 值、指针都实现接口：`func (v T) Method()`
- 指针接收者定义的方法 -> 只有指针实现接口： `func (p *T) Method()`

| 情况    | Go 编译器的行为   | 原因                      | 用途                 |
|-------|-------------|-------------------------|--------------------|
| 值接收者  | 调用时自动取地址    | 不会修改原值                  | 不修改对象/想被接口以值/指针都使用 |
| 指针接收者 | 不会自动取地址赋给接口 | 接口持有值副本，取地址可能导致副作用或非法引用 | 需要修改对象             |

---

## 接口动态类型与类型断言

> 更多类型断言见 type-system/type-conversion-vs-type-assertion

接口变量可在运行时存储不同类型的值
要取出原始值，可使用类型断言（type assertion）

```
v, ok := i.(T)
```

或类型分支（type switch）：

```
switch v := i.(type) {
case int:
    fmt.Println("int")
case string:
    fmt.Println("string")
default:
    fmt.Println("unknown")
}
```

**这正是接口实现“多态”的核心机制**。

---

## 多态（Polymorphism）

在 Go 中，多态通过接口自然实现。

定义抽象行为：

```
type Speaker interface {
    Speak() string
}
```

实现不同类型：

```
type Dog struct{}
type Cat struct{}

func (Dog) Speak() string { return "Woof" }
func (Cat) Speak() string { return "Meow" }
```

多态使用：

```
animals := []Speaker{Dog{}, Cat{}}
for _, a := range animals {
    fmt.Println(a.Speak())
}
```

**接口让代码面向行为，而非具体类型。**

---

## 接口的嵌入（Embedding）

接口可通过嵌入其他接口实现组合与扩展

```
type Reader interface {
    Read(p []byte) (n int, err error)
}

type Writer interface {
    Write(p []byte) (n int, err error)
}

type ReadWriter interface {
    Reader
    Writer
}
```

接口嵌入约等于行为组合：这体现了 Go 的组合优于继承哲学。



---

## 接口与泛型（Generics）

Go 1.18+ 引入泛型

* 接口是运行时多态（dynamic polymorphism）
* 泛型是编译期多态（static polymorphism）

区别：

| 特征   | 接口      | 泛型     |
|------|---------|--------|
| 检查时机 | 运行时     | 编译期    |
| 实现机制 | 方法集匹配   | 类型参数约束 |
| 性能   | 有动态分派开销 | 无分派开销  |
| 适用场景 | 行为抽象    | 类型抽象   |


