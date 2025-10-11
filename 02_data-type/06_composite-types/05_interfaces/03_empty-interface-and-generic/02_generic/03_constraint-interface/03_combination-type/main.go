package main

import "fmt"

// Add 使用 | 或运算符连接多个类型进行约束
func Add[T int | float64](a, b T) T {
	return a + b
}
func main() {
	fmt.Println(Add(1, 2))
	fmt.Println(Add(1, 23.2))
}
