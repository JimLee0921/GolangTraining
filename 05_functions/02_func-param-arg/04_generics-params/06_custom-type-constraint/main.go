package main

import "fmt"

type Number interface {
	int | int64 | float64
}

func Sum[T Number](a, b T) T {
	return a + b
}

func main() {
	//Sum(1, "123")	// 第二个参数类型不符合，编译时就会报错
	fmt.Println(Sum(1, 1.22))
}
