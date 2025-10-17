package main

import "fmt"

// Operation 定义了一个函数类型 Operation，可以表示任何接收两个 int 并返回 int 的函数
type Operation func(int, int) int

func main() {
	var add Operation = func(a int, b int) int {
		return a + b
	}

	fmt.Println(add(1, 1))
}
