package main

import "fmt"

func main() {
	//fmt.Println(x) // x 在块作用域中变量定义前不能使用
	fmt.Println(y) // y 在包级作用域被定义，可以直接使用
	x := 5
	fmt.Println(x)
}

var y = 666
