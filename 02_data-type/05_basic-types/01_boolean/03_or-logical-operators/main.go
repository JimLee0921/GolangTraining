package main

import "fmt"

func left() bool {
	fmt.Println("left 被调用")
	return true
}

func right() bool {
	fmt.Println("right 被调用")
	return false
}

// main || 为 逻辑或（logical OR）运算符
func main() {
	/*
		二元运算符（作用于两个值）只要有一个为 true，结果就是 true
		有 短路求值：左边已经 true，右边就不会再计算
	*/
	isUser := true
	isAdmin := false

	if isUser || isAdmin {
		fmt.Println("已经认证的用户")
	}

	// 演示短路求值：由于 left() 已经为 true，不会调用 right()
	if left() || right() {
		fmt.Println("条件为真")
	}
}
