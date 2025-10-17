package main

import "fmt"

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
