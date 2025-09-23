package main

import "fmt"

// main &&（and） 逻辑与（logical AND）运算符
func main() {
	/*
		二元运算符，只有两个值都为 true，结果才是 true。
		也有 短路求值：左边已经 false，右边就不会再计算。
	*/
	isUser := true
	isAdmin := false
	// && and运算符
	if isUser && isAdmin {
		fmt.Println("已经登录且是管理员")
	} else {
		fmt.Println("未登录或不是管理员")
	}
	// || or运算符
	if left, right := true, true; left || right {
		fmt.Println("有一个条件成立")
	} else {
		fmt.Println("条件都没成立")
	}
	// ! not 运算符
	if isTrue := true; !isTrue {
		fmt.Println("条件为真")
	} else {
		fmt.Println("条件为假")
	}
}
