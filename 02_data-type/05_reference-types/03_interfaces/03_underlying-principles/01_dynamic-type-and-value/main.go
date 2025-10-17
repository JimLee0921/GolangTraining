package main

import "fmt"

// Speaker 定义接口
type Speaker interface {
	Speak()
}

// 定义实现类型

type Dog struct {
	Name string
}

func (d Dog) Speak() {
	fmt.Println("Woof!")
}

type Cat struct {
	Name string
}

func (c Cat) Speak() {
	fmt.Println("Meow!")
}

func main() {
	/*
		在 Go 的运行时中，每个接口变量（非空接口）可以理解为一个装了两样东西的盒子：
			动态类型（dynamic type）：Dog
			动态值（dynamic value）：Dog{} 的实例
		interface value = {
			dynamic type  // 运行时真实类型信息
			dynamic value // 运行时真实的值（可以是值也可以是指针）
		}
	*/
	var s Speaker
	s = Dog{"JimLee"}
	fmt.Printf("%v: %T\n", s, s) // {JimLee}: main.Dog
	s = Cat{"JamesBond"}
	fmt.Printf("%v: %T\n", s, s) // {JamesBond}: main.Cat
}
