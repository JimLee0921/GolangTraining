package main

import (
	"fmt"
	"math/rand"
)

// main if支持初始化语句
func main() {
	/*
		Go 提供了 if 初始化语句; 条件 的写法
		初始化语句的作用域只在整个 if-else 语句块内有效（else 中也是可以正常访问到的）
		if 的初始化语句只能写一条，但可以用 多重赋值 初始化多个变量
		if 初始化语句; 条件 {
			条件为 true 执行
		} else {
			条件为 false 执行
		}
	*/
	if adminMessage, userMessage, isAdmin := "欢迎管理员", "欢迎用户", rand.Intn(2) == 0; isAdmin {
		fmt.Println(adminMessage)
	} else {
		fmt.Println(userMessage)
	}
	//fmt.Println(userMessage) // 这里会遍历错误，因为 if 初始化语句的作用域只在 if-else 语句块内

}
