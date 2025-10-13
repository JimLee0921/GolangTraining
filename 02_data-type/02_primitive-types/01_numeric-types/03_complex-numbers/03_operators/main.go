package main

import "fmt"

func main() {
	// 复数支持常见的算术运算符
	a := 2 + 3i
	b := 1 + 4i

	fmt.Println("a + b =", a+b)
	fmt.Println("a - b =", a-b)
	fmt.Println("a * b =", a*b)
	fmt.Println("a / b =", a/b)
}

/*
a + b = (3+7i)
a - b = (1-1i)
a * b = (-10+11i)
a / b = (0.8235294117647058-0.29411764705882354i)
*/
