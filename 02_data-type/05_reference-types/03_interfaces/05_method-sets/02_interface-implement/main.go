package main

import "fmt"

// Editor 定义一个接口
type Editor interface {
	Edit()
}

// Data 定义一个结构体类型
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

	// 普通方法调用阶段：Go 会自动取地址或解引用
	v.Show() // 值接收者方法
	p.Show() // 自动解引用 (*p).Show()

	v.Edit() // 自动取地址 (&v).Edit()
	p.Edit() // 指针方法直接调用

	fmt.Println("------ 接口实现部分 ------")

	// 接口实现检查阶段：Go 不会自动取地址
	//var e Editor = v   // 编译错误：Data 没有 Edit() 方法
	var e Editor = p // *Data 实现了 Edit()

	e.Edit() // 调用接口方法

}
