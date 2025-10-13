package main

import "fmt"

func main() {
	/*
		整数类型在溢出时会 回绕（wrap around），不会报错
		Go 不会检测整数溢出
		结果直接按二进制回绕
		这是底层 CPU 指令的自然行为
	*/
	var x uint8 = 255
	x++            // 溢出，回到 0
	fmt.Println(x) // 输出 0

	var y int8 = 127
	y++ // 溢出，变成 -128
	fmt.Println(y)
}
