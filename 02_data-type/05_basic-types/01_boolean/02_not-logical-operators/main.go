package main

import "fmt"

// main 条件判断语句中 not 为逻辑非运算符，把 true 变成 false，false 变成 true
func main() {
	isTrue := true
	isFalse := false

	if !isTrue {
		fmt.Println("这里条件被 ! 转为 false 不会输出")
	}
	if !isFalse {
		fmt.Println("这里条件被 ! 转为 true 会输出")
	}
}
