## 方法集（Method Set）

## 概念

方法集（Method Set） 就是指：某个类型能调用的所有方法的集合。它决定了一个类型是否满足某个接口。

### 方法集规则

方法集 = 某个类型所拥有的所有方法的集合。

它决定了：

- 这个类型实现了哪些接口
- 该类型的值（或指针）能否调用某个方法

假设有一个类型 T：

| 接收者         | 方法集内容                         |
|-------------|-------------------------------|
| **T 值类型**   | 所有接收者是 `(t T)` 的方法            |
| ***T 指针类型** | 所有接收者是 `(t T)` 和 `(t *T)` 的方法 |
| **接口类型**    | 由接口定义的全部方法                    |

- 对于 值类型 T：方法集只包含接收者是 T 的方法
- 对于 *指针类型 T：方法集包含接收者是 T 和 *T 的方法
- T 的方法集是 *T 的方法集的子集，有些接口，必须用 *T 才能实现

## 接收者

### 方法接收

方法调用语法的自动取址/解引用:

```
type Data struct{}

func (d Data) Show() {
	fmt.Println("Value receiver")
}

func (d *Data) Edit() {
	fmt.Println("Pointer receiver")
}

func main() {
	var v Data
	var p *Data = &v

	v.Show() // 值调用值接收者
	p.Show() // 指针自动解引用调用值接收者

	v.Edit() // 值类型自动取地址调用
	p.Edit() // 指针类型调用指针接收者方法
}

```

Show 为 值接收者，Edit 为指针接收者，进行方法调用时：

- 值接收者：可以用值调用，也能用指针调用（编译器自动取地址）
    - `p.Show()` -> `(*p).Show()` 自动解引用
- 指针接收者：可以使用指针调用，也可以使用值调用（编译器自解引用）
    - `v.Edit()` -> `(&v).Edit()` 自动取地址

### 接口实现接收

```
package main

import "fmt"

type Editor interface {
	Edit()
}

type Data struct{}

// Show 值接收者方法
func (d Data) Show() {
	fmt.Println("Value receiver: Show()")
}

// Edit 指针接收者方法
func (d *Data) Edit() {
	fmt.Println("Pointer receiver: Edit()")
}

func main() {
	var v Data
	var p *Data = &v


	// 接口实现检查阶段：Go 不会自动取地址
	//var e Editor = v   // 编译错误：Data 没有 Edit() 方法
	var e Editor = p // *Data 实现了 Edit()

	e.Edit() // 调用接口方法

}
```

定义一个接口 Editor 定义 Edit 方法，Data 编写 Edit 方法实现 Editor 接口

此时 `var e Editor = v` 直接编译错误，因为接口实现不会自动解引用和取地址，
需要传递指针 *Date 给 e 才能调用指针接收者方法

- 实现了接收者是值类型的方法，相当于自动实现了接收者是指针类型的方法
- 而实现了接收者是指针类型的方法，不会自动生成对应接收者是值类型的方法
- 官方不推荐结构体同时有值接受器方法和指针接收器方法
- 方法调用时如果是指针接收者但是传递的是值，那么 Go 的语法糖会自动取地址或解引用
- 但是在接口赋值调用时 Go 不会自动取地址或解引用，必须手动取地址或解引用

### 补充

一般不推荐同一个结构体同时混用值接收者和指针接收者。
更稳妥的做法是：一旦有任意方法需要指针接收者，就把这个类型的所有方法都用指针接收者。

* `T` 的方法集只有“值接收者方法”；
* `*T` 的方法集包含“值 + 指针接收者方法”。
  混用会导致：`T` 实现了一部分接口，`*T` 实现了另一部分接口 -> 使用起来条理不够清晰。

```
type R interface{ Read() }
type W interface{ Write() }

type File struct{}
func (File) Read()     {}  // 值接收者
func (*File) Write()   {}  // 指针接收者

var r R = File{}  
var w W = File{}   // 编译错误（Write 只有 *File 才有）
var w2 W = &File{}
```

### 结论

* **规则 1**：只要类型里有“会修改状态”的方法，所有方法统一用指针接收者。
* **规则 2**：仅当类型“小、可复制、不可变”时，全部用值接收者。
* **规则 3**：避免混用；如需混用，务必在注释中明确区分哪些方法只读、哪些会改状态，并在接口设计上谨慎评估方法集影响。

