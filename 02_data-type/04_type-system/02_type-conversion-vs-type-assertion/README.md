# Go 类型转换与类型断言（Type Conversion & Type Assertion）

## 概述

Go 是一门**强类型语言（strongly typed language）**，
类型之间不会进行隐式转换，所有转换都必须**显式声明**。

类型系统提供两种机制：

| 机制                        | 阶段  | 作用             |
|---------------------------|-----|----------------|
| **类型转换（type conversion）** | 编译期 | 用于已知类型之间的显式转换  |
| **类型断言（type assertion）**  | 运行期 | 用于接口类型到具体类型的提取 |

> 从 Go 1.18（引入泛型） 开始，很多原来必须依赖 interface{} 的地方，现在都可以用 类型参数 实现真正的静态类型
>
> 因此：类型断言使用越来越少了

---

## 类型转换（Type Conversion）

Go 里，类型转换（type conversion）是指把一个值从某种类型显式地转成另一种类型，
这和某些编程语言中的强制类型转换类似，但 Go 的设计更安全：

- Go 中如果不进行类型转换直接运算比如 int + float 会编译失败
- Go 转换必须显式写出目标类型（Go中没有隐式转换，不会自动转换）
- 只能在兼容或可转换的类型之间进行

> Go 不支持野蛮的指针强转（除非用 unsafe 包）

### 基本语法

```
T(expression)
```

表示将 `expression` 的值转换为类型 `T`。

注意：

* Go 不存在隐式类型提升（如 int -> float64）
* 必须显式声明
* 转换规则基于类型兼容性与底层表示

---

### 基本类型之间的转换

```
var a int = 10
var b float64 = float64(a)
var c uint = uint(a)
```

* 可在数值类型间转换；
* 若超出目标类型表示范围，则截断；
* 字符与整数类型也可相互转换。

---

### 字符与数值间的转换

```
var ch byte = 'A'
var n int = int(ch)
var r rune = rune(ch)
var s string = string(ch)
```

| 转换                 | 结果            |
|--------------------|---------------|
| `byte -> int`      | 数值（ASCII 码）   |
| `int -> rune`      | Unicode 码点    |
| `rune -> string`   | 字符串（UTF-8 编码） |
| `string -> []byte` | 字节切片副本        |

Go 的 `string` 是不可变字节序列，转换时会发生复制

---

### 复合类型间的转换

部分复合类型之间也允许转换，但需满足约束：

#### 结构体转换（相同字段）

```
type A struct{ X int }
type B struct{ X int }

var a A
var b B = B(a) // 字段结构完全一致
```

#### 结构体字段不一致

```
type A struct{ X int }
type B struct{ X int; Y int }

var b B = B(A{}) // 编译错误
```

---

### 切片与数组之间的转换

```
var arr [3]int = [3]int{1, 2, 3}
var s []int = arr[:] // 从数组创建切片
```

但：

```
arr = [3]int(s) // 不允许直接转换
```

数组与切片是不同类型（值类型 vs 引用类型）

---

### 指针与类型转换

Go 不支持 C 风格的任意类型指针转换，
只能在相同底层类型的指针之间转换。

```
type MyInt int
var x MyInt = 5
var p *int = (*int)(&x)
```

> `unsafe.Pointer` 是唯一可实现任意转换的方式，但属于低级不安全操作

---

### 接口之间的类型转换（interface conversion）

当一个接口类型可以包含另一个接口时，可以进行转换：

```
var r io.Reader
var rw io.ReadWriter

r = rw // ReadWriter 包含 Reader 方法集
rw = r // Reader 不一定有 Write 方法
```

**规则：**

> 可将更大方法集转换为更小方法集，但反向不行

| 类型      | 转换条件          | 是否允许隐式 | 示例                  |
|---------|---------------|--------|---------------------|
| Array   | 元素类型 + 长度完全一致 | 否      | `B(A(a))`           |
| Slice   | 元素类型相同        | 否      | `MySlice(s)`        |
| Struct  | 字段名、顺序、类型完全一致 | 否      | `B(a)`              |
| Map     | 键和值类型相同       | 否      | `M2(m1)`            |
| Chan    | 元素类型相同、方向兼容   | 部分     | `<-chan` / `chan<-` |
| Pointer | 指向类型相同        | 否      | `(*int)(p)`         |
| Func    | 参数与返回类型完全一致   | 否      | `BinOp(add)`        |

---

## 类型断言（Type Assertion）

GO 中的类型断言 (type assertion) 用于从接口类型变量中取出它的动态值，
在 Go 里，接口变量可以存储任何实现了接口的值，但是当需要拿出里面的具体类型时，就需要使用类型断言。

```
v, ok := i.(T)
```

- i：必须是接口类型的变量（如 interface{}(简写为 any)、自定义接口）
- T：目标类型
- v：如果断言成功，就是 x 存的那个值，并且类型是 T，如果失败则是该类型的空值
- ok：布尔值，表示断言是否成功

> 带ok：如果断言成功 -> v 是类型 T 的值，ok = true 如果断言失败 -> v 是 T 的零值，ok = false
>
> 不带ok：如果 x 不是接口类型，或者接口里存的动态值不是 T，直接 panic（不安全，除非能百分百确定类型）




### 使用示例

```
var x interface{} = "hello"

v1 := x.(string)       // 成功
v2, ok := x.(string)   // ok = true
v3, ok := x.(int) // ok = false
v4 := x.(int)     // panic: interface conversion
```



### 类型断言的本质

接口变量内部存储两部分：

1. **类型信息**（dynamic type）
2. **值信息**（dynamic value）

断言操作其实是：

> 如果接口当前的动态类型是 T，则取出其中的值。

否则：

* 带 `ok` 版本返回零值和 `false`
* 不带 `ok` 版本触发 `panic`



### 对接口类型的断言（Interface-to-Interface）

可以断言为另一个接口类型，只要当前动态类型实现了目标接口

```
var r io.Reader
var w io.Writer

rw, ok := r.(io.ReadWriter)
```

如果 `r` 底层类型同时实现了 `Read` 和 `Write`，断言成功



### 结合 `switch` 使用（类型分支）

```
switch v := i.(type) {
case int:
fmt.Println("int", v)
case string:
fmt.Println("string", v)
default:
fmt.Println("unknown")
}
```

这是类型断言的语法糖形式：编译器自动执行多重断言匹配。

---

### 常见陷阱与建议

| 场景      | 说明                         |
|---------|----------------------------|
| 未判断 ok  | 断言失败会 panic                |
| 断言为错误类型 | 不会隐式转换（如 `int32` != `int`） |
| 空接口嵌套   | 多层 `interface{}` 需要逐层断言    |
| 指针类型断言  | 必须精确匹配（`*T` != `T`）        |

---

## 类型转换 vs 类型断言

| 对比项       | 类型转换（Conversion） | 类型断言（Assertion）  |
|-----------|------------------|------------------|
| 作用对象      | 已知类型之间           | 接口类型中的动态类型       |
| 检查阶段      | 编译期              | 运行期              |
| 是否会 panic | 不会（编译时检查）        | 会（断言失败时）         |
| 用法形式      | `T(x)`           | `x.(T)`          |
| 是否复制值     | 是（新值）            | 否（取出底层值）         |
| 是否改变类型    | 是                | 否，仅解封装           |
| 是否涉及接口    | 否                | 是（必须是 interface） |
| 是否可链式使用   | 可                | 可（在类型 switch 中）  |

> 转换是静态的，断言是动态的

