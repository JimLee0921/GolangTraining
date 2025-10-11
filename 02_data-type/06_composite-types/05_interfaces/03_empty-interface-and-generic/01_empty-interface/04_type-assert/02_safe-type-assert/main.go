package main

import "fmt"

func main() {
	var v any
	v = 199
	// 安全类型断言
	n, ok := v.(string)
	if ok {
		fmt.Println("n is string")
		fmt.Println(n)
	} else {
		fmt.Println("n is not string")
	}
}
