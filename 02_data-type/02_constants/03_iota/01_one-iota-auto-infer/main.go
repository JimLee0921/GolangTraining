package main

import "fmt"

const (
	x = iota
	y
	z
)

// main 在 const 声明块中第一个值为 iota 剩余的值会自动递增
func main() {
	fmt.Println(x) // 0
	fmt.Println(y) // 1
	fmt.Println(z) // 2
}
