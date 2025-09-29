package main

import "fmt"

func main() {
	/*
		int 和 float 之间类型转换
	*/
	var x = 12
	var y = 12.1230123
	//fmt.Println(x + y) 会报错
	// int 转换为 float
	fmt.Printf("result: %v - type:%T\n", y+float64(x), y+float64(x))
	// float 转换为 int 会精度丢失
	fmt.Printf("result: %v - type:%T\n", int(y)+x, int(y)+x)
}
