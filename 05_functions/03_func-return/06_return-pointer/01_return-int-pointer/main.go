package main

import "fmt"

func main() {
	/*
		x 是在函数内部定义的
		但 Go 编译器会自动逃逸分析（escape analysis），确保返回的指针仍然可用
		所以不会像 C/C++ 那样返回局部变量地址导致悬空指针
	*/
	p := newInt()
	fmt.Println(*p) // 10
	*p = 20
	fmt.Println(*p) // 20
}

// newInt 返回一个局部变量的指针，在函数外可以进行操作
func newInt() *int {
	x := 10
	return &x // 返回局部变量的指针
}
