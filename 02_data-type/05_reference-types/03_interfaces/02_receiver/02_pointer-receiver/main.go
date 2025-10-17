package main

import "fmt"

type Animal interface {
	Speak()
}

type Cat struct{}

func (c *Cat) Speak() {
	fmt.Println("Meow!")
}

func main() {
	/*
		指针接收者实现接口 只有指针类型实现该接口
		Speak 的接收者是 *Cat
		只有 指针类型 *Cat 实现了接口
		值类型 Cat 没有实现，因为编译器不会你取地址赋值给接口
	*/
	var a Animal
	c := Cat{}
	//a = c // 报错：Cat 没有实现接口（因为 Speak 是 *Cat 定义的）

	a = &c    // *Cat 实现接口
	a.Speak() // 正常调用
}
