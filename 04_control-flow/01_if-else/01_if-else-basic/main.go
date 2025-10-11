package main

import (
	"fmt"
	"math/rand/v2"
)

// main if-else 分支基础练习
func main() {
	/*
		Go 中的 if-else 最基础语法，else 块可以不写，就是只有一个 if 判断语句块
			if 条件 {
				// 当条件为 true 时执行
			} else {
				// 当条件为 false 时执行
			}
	*/
	isAdmin := rand.IntN(2) == 0 // 返回 0 或 1，50% 概率为 true
	if isAdmin {
		fmt.Println("用户是管理员")
	} else {
		fmt.Println("用户不是管理员")
	}
}
