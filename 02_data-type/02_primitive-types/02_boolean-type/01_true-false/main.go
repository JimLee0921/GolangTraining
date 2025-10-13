package main

import "fmt"

// main 布尔值最简单的用法，true false 表达式
func main() {
	isTrue := true
	isFalse := false

	if isTrue {
		fmt.Println("这里条件为 true 会输出")
	}

	if isFalse {
		fmt.Println("这里条件为 false 不会输出")
	}
}
