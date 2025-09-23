package main

import "fmt"

func left() bool {
	fmt.Println("left 被调用")
	return false
}

func right() bool {
	fmt.Println("right 被调用")
	return false
}

// main &&（and） 逻辑与（logical AND）运算符
func main() {
	/*
		二元运算符，只有两个值都为 true，结果才是 true。
		也有 短路求值：左边已经 false，右边就不会再计算。
	*/
	isUser := true
	isAdmin := false

	if isUser && isAdmin {
		fmt.Println("已经登录且是管理员")
	} else {
		fmt.Println("未登录或不是管理员")
	}

	// 演示短路求值
	if left() && right() {
		fmt.Println("条件为真")
	}
}
