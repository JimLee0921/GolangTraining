package main

import "fmt"

type person struct {
	name string
	age  int
}

// main 创建结构体指针方式
func main() {
	/*
		一共有两种方式用来创建结构体指针
		1. 取地址符 &T{...}
			使用 复合字面量 (composite literal) 创建结构体，并立刻取地址
			得到的类型是 *T（指针）
			可以指定字段
		2. 使用 new 关键字（不常用，但等价）
			new(T) 会分配一块内存，把所有字段置为 零值，返回 *T
			没法像 &T{...} 那样直接指定初始字段值，只能先得到零值，再去赋值
	*/
	// 直接初始化赋值
	p1 := &person{"Bob", 20}
	// 都是零值 需要再手动赋值
	p2 := new(person)
	p2.name = "Jack"
	p2.age = 20
	fmt.Println(p1)
	fmt.Println(p2)
}
