package main

import "fmt"

// main 最基础的 for 循环
func main() {
	/*
		for 初始化语句; 条件表达式; 后置语句 {
			// 循环体
		}
	*/
	for i := 1; i <= 100; i++ {
		fmt.Printf("第 %d 次\n", i)
	}
}
