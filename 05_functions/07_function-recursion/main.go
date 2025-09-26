package main

import "fmt"

// main 函数递归
func main() {
	/*
		函数递归就是 函数自己调用自己
		通常用在：
			问题可以分解成规模更小的同类问题
			有明确的终止条件（否则就会无限循环，导致栈溢出）
		递归的两个关键要素
			基准条件：终止条件，防止无限调用
			递归调用：函数在某些情况下再次调用自己
		最简单的例子：计算阶乘，斐波那契数列等
		递归：代码简洁，逻辑直观，但可能会占用更多栈内存
		循环：效率更高，但有些问题用循环表达不如递归自然
	*/

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
