package main

import "fmt"

type Animal interface {
	Speak()
}

type Dog struct{}

// Speak 值接收者
func (d Dog) Speak() {
	fmt.Println("Woof!")
}

func main() {
	/*
		值接收者实现接口，值和指针都实现该接口
			值类型调用（d.Speak()）
			指针类型调用（(&d).Speak()）—— 编译器自动解引用
	*/
	var a Animal
	d := Dog{}
	a = d     // 值类型 Dog 实现接口
	a.Speak() // 正常调用

	a = &d    // 指针类型 *Dog 实现接口
	a.Speak() // 正常调用

}
