package main

import "fmt"

func main() {
	/*
		Go 提供两个内置函数
			real(z)：获取复数的实部
			imag(z)：获取复数的虚部
	*/
	c := 3 + 4i
	fmt.Println("实部:", real(c)) // 3
	fmt.Println("虚部:", imag(c)) // 4
}
