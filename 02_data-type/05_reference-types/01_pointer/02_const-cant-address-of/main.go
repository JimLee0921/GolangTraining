package main

import "fmt"

// main 查看内存地址
func main() {
	/*
		常量 (const) 不能取地址，因为常量在编译时就已经确定
		它们不是存储在一个变量里，而是 直接内嵌在代码里的值
	*/
	const PI = 3.14

	fmt.Println("PI - ", PI)
	//fmt.Println("PI's memory address - ", &PI)	编译错误
	//fmt.Printf("%d\n", &PI)	编译错误
}
