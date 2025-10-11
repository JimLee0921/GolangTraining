package main

import "fmt"

func main() {
	fmt.Println(swap("A", "B"))
	fmt.Println(swap(666, 333))
}

// swap 泛型定义函数参数类型和返回值
func swap[T any](a, b T) (T, T) {
	return b, a
}
