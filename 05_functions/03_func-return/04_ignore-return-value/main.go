package main

import "fmt"

func divide(a, b int) (int, error) {
	if b == 0 {
		return 0, fmt.Errorf("divisor can't be zero")
	} else {
		return a / b, nil
	}
}

func main() {
	/*
		如果不需要函数的某个返回值，用 `_` 忽略
	*/
	// 忽略错误
	result, _ := divide(10, 2)
	fmt.Println(result)
}
