package main

import "fmt"

// main if-else_if-else编写多个分支
func main() {
	/*
		if 条件 {
		    // 当条件为 true 时执行
		} else if 条件 {
		    // 当第二个条件为 true 时执行
		} else if 条件 {
			// 当第三个条件为 true 时执行
		} else {
			// 当上面所有条件都为 false 时执行
		}
	*/

	score := 89
	if score < 50 {
		fmt.Println("不及格")
	} else if score >= 50 && score < 80 {
		fmt.Println("及格")
	} else if score >= 80 && score < 90 {
		fmt.Println("优秀")
	} else {
		fmt.Println("NB")
	}
}
