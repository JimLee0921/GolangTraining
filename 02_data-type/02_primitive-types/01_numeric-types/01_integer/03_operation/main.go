package main

import "fmt"

func main() {
	a, b := 15, 4

	fmt.Println("a + b =", a+b) // 加法 -> 19
	fmt.Println("a - b =", a-b) // 减法 -> 11
	fmt.Println("a * b =", a*b) // 乘法 -> 60
	fmt.Println("a / b =", a/b) // 除法（整数除法，自动向零取整）-> 3
	fmt.Println("a % b =", a%b) // 取余，仅适用于整数类型 -> 3

	a++                      // 自增运算符
	b--                      // 自减运算符
	fmt.Println("a ++ =", a) // 16
	fmt.Println("b -- =", b) // 3

	// 更多运算符见 operators 章节
}
