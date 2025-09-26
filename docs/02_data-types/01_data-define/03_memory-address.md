# Go 取地址与指针基本使用

## 取地址符 `&`

- 使用 `&变量名` 可以获取变量的 **内存地址**。

示例：

```go
package main

import "fmt"

func main() {
	x := 42
	fmt.Println("x 的值:", x)
	fmt.Println("x 的地址:", &x)
}
```

输出：

```

x 的值: 42
x 的地址: 0xc0000140a8

```

---

## 指针

* 指针变量用于保存另一个变量的地址。
* 声明指针：`var p *int` → `p` 是一个指向 `int` 的指针
* 获取地址：`p := &x`
* 解引用：`*p` 访问指针指向的值

示例：

```go
package main

import "fmt"

func main() {
	x := 42
	p := &x // p 是 *int 类型，指向 x

	fmt.Println("x 的值:", x)
	fmt.Println("p (指针地址):", p)
	fmt.Println("通过指针访问 x 的值:", *p)

	*p = 100 // 修改指针指向的值
	fmt.Println("修改后的 x:", x)
}
```

输出：

```
x 的值: 42
p (指针地址): 0xc0000140a8
通过指针访问 x 的值: 42
修改后的 x: 100
```

---

## 常量不能取地址

* 常量在编译期就直接替换为值，不会分配独立存储空间。
* 因此常量不能取地址。

示例：

```go
const Pi = 3.14

func main() {
fmt.Println(&Pi) // 编译错误：cannot take the address of Pi
}
```

---

## 小结

* `&` → 取变量地址
* `*` → 通过指针访问或修改值（解引用）
* 常量没有地址，不能取地址
* 指针常用于在函数间传递和修改数据，避免拷贝大结构体




