package main

import (
	"fmt"
	"math/rand"
)

func main() {
	/*
		! 为（logical NOT） 逻辑非运算符，把 true 变成 false，false 变成 true
		|| 为 逻辑或（logical OR）运算符
		&&（and） 逻辑与（logical AND）运算符
	*/
	//  随机生成登录与管理员状态
	var loggedIn bool = rand.Intn(2) == 1 // 随机 true 或 false
	fmt.Println("loggedIn:", loggedIn)

	if loggedIn {
		isAdmin := rand.Intn(2) == 1 // 随机 true 或 false
		fmt.Printf("isAdmin: %v\n", isAdmin)
		fmt.Println("用户已登录")
		if isAdmin {
			fmt.Println("欢迎管理员")
		} else {
			fmt.Println("普通用户访问")
		}
	} else {
		fmt.Println("请先登录")
	}
}
