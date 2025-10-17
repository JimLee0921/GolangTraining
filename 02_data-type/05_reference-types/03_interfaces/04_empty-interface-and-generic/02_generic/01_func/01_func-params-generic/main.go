package main

import "fmt"

func main() {
	PrintAll(123)
	PrintAll("hello")
	PrintAll(true)

}

// PrintAll 泛型定义函数参数
func PrintAll[T any](v T) {
	fmt.Printf("%T: %v\n", v, v)
}
