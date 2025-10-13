package main

import "fmt"

func main() {
	/*
		byte 与 rune 是 uint8 和 int32 的别名
		byte 用于处理原始字节
		rune 用于处理 Unicode 字符
		%c 以字符形式输出整数对应的 Unicode
	*/
	var b byte = 'A' // 实际类型 uint8
	var r rune = '你' // 实际类型 int32

	fmt.Printf("b = %d, r = %d\n", b, r)
	fmt.Printf("b = %c, r = %c\n", b, r)
}
