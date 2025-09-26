package main

import "fmt"

// main 查看内存地址
func main() {
	/*
		在 Go 里，可以用 取地址符 & 来得到变量的内存地址
		常量 (const) 不能取地址，因为常量在编译时就已经确定，它们不是存储在一个变量里，而是 直接内嵌在代码里的值
	*/
	a := 10

	fmt.Println("a - ", a)
	fmt.Println("a's memory address - ", &a)
	fmt.Printf("%d\n", &a)

	const PI = 3.14

	fmt.Println("PI - ", PI)
	//fmt.Println("PI's memory address - ", &PI)	编译错误
	//fmt.Printf("%d\n", &PI)	编译错误
}
