package main

import "fmt"

// Ordered 这里接口类型定义不携带 ~
type Ordered interface {
	int | float64 | string
}

type MyInt int

func Max[T Ordered](a, b T) T {
	if a > b {
		return a
	}
	return b
}

func main() {
	/*
		type	只能是 type 本身
		~type	任何底层是 int 的类型（比如 MyInt）
	*/
	fmt.Println(Max(3, 5))     // int ok
	fmt.Println(Max(3.5, 4.8)) // float64 ok
	// fmt.Println(Max(MyInt(1), MyInt(2))) 编译错误，因为 MyInt 是一个新的类型，不等于 int。
}
