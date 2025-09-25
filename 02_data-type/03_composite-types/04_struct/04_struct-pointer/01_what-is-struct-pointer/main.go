package main

import "fmt"

type person struct {
	name string
	age  int
}

func main() {
	/*
		在 Go 中，结构体指针（pointer to struct）是很常见的用法，本质上就是存放结构体地址的变量
		pointInstance := &structType{初始化属性}
		pointInstance 是 *structType 类型 指向一个 structType
		Go 做了语法糖：访问字段时 不需要写 (*instance).字段名，直接写 p.字段名 就行。
	*/
	// 普通结构体变量: p1 是一个结构体值，赋值/传参时会 拷贝整个结构体。
	p1 := person{
		name: "AliceFrank",
		age:  18,
	}
	fmt.Println(p1.name)
	// 结构体指针: p2 是 *Person 类型，指向一个 Person
	p2 := &person{
		name: "JimLee",
		age:  0,
	}
	fmt.Println(p2.name)
}
