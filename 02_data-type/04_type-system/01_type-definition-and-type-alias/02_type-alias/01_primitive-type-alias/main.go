package main

import "fmt"

type MyInt = int
type MyString = string

func main() {
	var a MyInt = 10
	var b int = a         // 可以直接赋值，无需转换
	fmt.Printf("%T\n", a) // 输出：int
	fmt.Printf("%T\n", b) // 输出：int

}
