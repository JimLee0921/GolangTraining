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
		如果要定义方法并且希望它能修改结构体字段，需要用指针接收者：
		func (p *Person) funcName() {...}
	*/

	p := &person{"Bob", 20}
	fmt.Println(p.age)
	p.Birthday()
	fmt.Println(p.age)
	p.Birthday()
	fmt.Println(p.age)

}
