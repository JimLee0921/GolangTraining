package main

import "fmt"

// main 块作用域引出函数闭包
func main() {
	x := 42
	fmt.Println(x)
	{
		fmt.Println(x)
		y := "The y variable belongs with the block and outer can't visit it"
		fmt.Println(y)
	}
	//fmt.Println(y) // 这里会报错访问不到块作用域中的 y
}
