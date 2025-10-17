package main

import "fmt"

type Number interface {
	int | int64 | float64
}

// Sum 自定义 interface 约束
func Sum[T Number](a, b T) T {
	return a + b
}

func main() {
	fmt.Println("sum:", Sum(1, 32.23))
}
