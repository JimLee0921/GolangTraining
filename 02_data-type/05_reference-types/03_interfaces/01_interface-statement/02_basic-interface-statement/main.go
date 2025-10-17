package main

import "fmt"

// Speaker 定义接口
type Speaker interface {
	Speak()
}

// 定义实现类型

type Dog struct{}

func (d Dog) Speak() {
	fmt.Println("Woof!")
}

type Cat struct{}

func (c Cat) Speak() {
	fmt.Println("Meow!")
}

func main() {
	/*
		Speaker 是接口类型
		任何实现了 Speak() 方法的类型都自动实现该接口
		Go 中没有 implements 关键字，接口是隐式实现，只看方法签名是否匹配
	*/
	// 使用接口
	var s Speaker
	s = Dog{}
	s.Speak() // Woof!

	s = Cat{}
	s.Speak() // Meow!

}
