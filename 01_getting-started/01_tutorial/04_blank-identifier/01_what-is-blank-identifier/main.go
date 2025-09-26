package main

import "fmt"

func foo() (int, string) {
	return 42, "hello"
}

// main 空白标识符 _
func main() {
	/*
		blank identifier 就是一个下划线 _
		被称为 空白标识符，本质上是一个占位符，但不能读取它的值
			_ 可以用在任何需要变量名的地方
			编译器会忽略它的值
			主要用途是“丢弃不需要的值
	*/
	x, _ := foo()  // 忽略第二个返回值
	fmt.Println(x) // 42
	//fmt.Println(_)	不能读取 _
}
