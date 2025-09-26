package main

import "fmt"

const (
	Read  = 1 << iota // 1 × 2^0 = 1
	Write             // 1 × 2^1 = 2
	Exec              // 1 × 2^2 = 4
)

// main 位运算配合 iota
func main() {
	/*
		这种方式常见于 权限控制、状态标志
	*/
	fmt.Println(Read, Write, Exec) // 1 2 4

}
