package main

import "fmt"

// main 函数递归
func main() {
	fmt.Println(factorial(10))
	fmt.Println(fibonacci(10))
}

// 计算 n 的阶乘（这里是 int 过大可能整数溢出）
func factorial(n int) int {
	if n == 0 { // 基准条件
		return 1
	}
	return n * factorial(n-1)
}

// 斐波那契数列（也不要太大，时间复杂度过高）
func fibonacci(n int) int {
	if n <= 1 {
		return n
	}
	return fibonacci(n-1) + fibonacci(n-2)
}
