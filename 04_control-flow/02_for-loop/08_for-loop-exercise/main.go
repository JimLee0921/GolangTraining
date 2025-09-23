package main

import "fmt"

// 小练习打印1-100中的所有质数
func main() {
	fmt.Println("1-100中所有质数有: ")
	// 外层循环遍历 2-100
	for i := 2; i <= 100; i++ {
		// 先设置变量 isPrime 初始值为 true
		isPrime := true
		// 内层循环判断 i 是否可以被 2 ~ n-1 整除
		for j := 2; j*j <= i; j++ {
			if i%j == 0 {
				isPrime = false
				break
			}
		}
		if isPrime {
			fmt.Print(i, " ")
		}
	}
	fmt.Println()
}
