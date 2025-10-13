package main

import "fmt"

// makeAdder 函数作为返回值
func makeAdder(base int) func(int) int {
	return func(x int) int {
		return base + x
	}
}

func main() {
	add10 := makeAdder(10)
	fmt.Println(add10(5))  // 15
	fmt.Println(add10(20)) // 30
}
