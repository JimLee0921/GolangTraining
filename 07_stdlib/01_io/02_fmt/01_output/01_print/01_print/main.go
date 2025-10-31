package main

import "fmt"

func main() {
	// fmt.Print() 直接打印内容，不自动换行，返回打印的字节数、错误
	fmt.Print(1, 2, "hello")
	fmt.Print("world")
	// 1 2helloworld

	n, err := fmt.Print(1, 2, 3, 4, 5)

	fmt.Print("write ", n, " bytes", " error ", err)
}
