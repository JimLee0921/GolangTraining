## 值类型传递（复制数据本身）

典型值类型：bool、数值型（int/float…）、string*、array（数组）、struct。
按值传递 意味着：调用函数时会拷贝一份完整的值给形参，函数内对形参的修改不会影响实参。

* string 语义上是不可变的值类型（实现是只读的指针+长度头部），见下文细节。

1. 基础数值与布尔：*int / *bool / *float64 …
    - 函数形参用指针类型 *T
    - 调用处用取址 &x
    - 函数内写 *p = ... 或 *p += ...
2. 结构体 *Struct
    - Go 语言中的结构体（struct）字段在函数传递或修改时，基础数据类型与引用数据类型的行为是不同的
    - 直接赋值：会复制结构体中的所有字段，基础类型会复制值，引用类型会复制引用
    - 传值调用函数：Go 默认按值传递（复制 struct），基础字段是独立的，引用字段共享底层
    - 传指针调用函数：所有字段（包括值类型）都可以被修改
    - slice 是引用类型，但本身是一个 结构体，所以当底层容量不够时直接传递也不会变
3. 数组：*[N]T （注意切片不同）
    - 数组是值类型，传参会整体拷贝；想改外部数组要传 *[N]T
    - 若想更灵活，通常改用 切片 []T（切片是引用类型，无需指针）
4. 字符串：不可变，不能原地改字节
    - 不能对 string 做 (*p)[0] = 'X' 这种操作
    - 但可以用 *string 来重新赋一个新字符串（重建后写回）

### array（数组）

数组是固定长度的值类型，传参会复制整个数组：

```go
func setFirst(a [3]int) { a[0] = 99 }

func main() {
arr := [3]int{1, 2, 3}
setFirst(arr)
fmt.Println(arr) // [1 2 3]（外部不变）
}
```

要修改外部数组：传指针或改用切片（切片是引用类型）

```go
func setFirstPtr(a *[3]int) { a[0] = 99 }
```

### string（不可变值）

```go
func replaceHead(s string) string {
b := []byte(s) // 拷贝字节
if len(b) > 0 { b[0] = 'X' }
return string(b)
}

func main() {
s := "hello"
t := replaceHead(s)
fmt.Println(s) // "hello"（原字符串未改）
fmt.Println(t) // "Xello"
}
```

string 本身不可变，传参就是拷贝它的“只读头部”，函数内无法原地改实参

## 指针方法接收者

指针方法接收者（pointer receiver）是 Go 面向对象风格里最核心的概念之一，用方法让调用更自然（编译器可自动取址）

### 定义

```go
func (r ReceiverType) MethodName(params) returnType { ... }
```

> r 就是“方法接收者”，表示这个方法属于某个类型（可以是结构体、指针、别名等）

```go
type Box struct{ N int }

func (b *Box) Inc() { b.N++ }

func main() {
b := Box{N: 1}
b.Inc() // 等价于 (&b).Inc()
fmt.Println(b.N) // 2
}

```

### 小结

- 数组是值类型，传参会整体拷贝；要改外部请用 `*[N]T` 或改用切片
- struct 按值传递不会改到外部，但里面的引用字段（切片/map/指针/chan）可影响外部
- string 不可变，想改内容需要转 []byte/[]rune 后返回新值
- 频繁传递大 struct 会拷贝开销大，用指针参数或拆分
- 指针参数需要关注并发安全与 nil 检查