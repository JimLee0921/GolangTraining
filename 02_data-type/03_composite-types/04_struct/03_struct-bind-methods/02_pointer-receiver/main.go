package main

import "fmt"

type person struct {
	name string
	age  int
}

func (p *person) Birthday() {
	p.age++
}

// main 指针接收者方法
func main() {
	/*
		如果要定义结构体方法并且希望能修改结构体字段，需要用指针接收者：
		func (p *Person) funcName() {...}
	*/
	// 这里没有取指针因为 Go 编译器的自动取地址机制，等价于 p := &person{"Bob", 20}
	p := person{"Bob", 20}
	fmt.Println(p.age)
	p.Birthday()
	fmt.Println(p.age)
	p.Birthday()
	fmt.Println(p.age)
}
