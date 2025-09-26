package main

import "fmt"

type Person struct {
	name string
	age  int
}

// main struct传递指针
func main() {
	/*
		struct 是值类型，所以当你把 struct 作为函数参数传递时，默认会 复制整个结构体
	*/
	p := Person{
		name: "JimLee",
		age:  1,
	}
	changePerson(&p)
	fmt.Println(p)       // {Bruce 1}
	fmt.Println(&p.name) // 0xc0000080a8
}

func changePerson(p *Person) {
	fmt.Println(p)       // &{JimLee 1}
	fmt.Println(p.name)  // JimLee
	fmt.Println(&p.name) // 0xc0000080a8
	p.name = "Bruce"
	fmt.Println(p)       // &{Bruce 1}
	fmt.Println(p.name)  // Bruce
	fmt.Println(&p.name) // 0xc0000080a8

}
