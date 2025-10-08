package main

import "fmt"

func main() {
	// 同步、简单、调用计算阶乘
	result := factorial(5)
	fmt.Println(result)
}

func factorial(n int) int {
	total := 1
	for i := n; i > 0; i-- {
		total *= i
	}
	return total
}
