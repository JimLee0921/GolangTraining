package main

import "fmt"

func main() {
	sum, diff := addSub(1, 2)
	fmt.Println(sum, diff)
}

// addSub 命名返回值也可以显式返回
func addSub(a, b int) (sum, diff int) {
	sum = a + b
	return sum, a - b // 显式返回
}
