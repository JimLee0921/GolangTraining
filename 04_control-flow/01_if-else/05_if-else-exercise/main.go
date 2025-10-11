package main

import (
	"fmt"
	"math/rand/v2"
)

// main if-else 模拟登陆系统练习
func main() {
	users := map[string]string{
		"admin": "1234",
		"tom":   "r3s2",
		"rose":  "fek32",
	}
	var username string
	if rand.IntN(2) == 0 {
		username = "tom"
	} else {
		username = "admin"
	}
	fmt.Println(username)
	password := "r3s2"
	// 使用初始化语句查询用户名是否存在
	if realPwd, ok := users[username]; ok {
		// 用户存在，继续判断密码
		if password == realPwd {
			// 密码正确，判断是否为 admin
			if username == "admin" {
				fmt.Println("欢迎管理员")
			} else {
				fmt.Println("欢迎普通用户")
			}
		} else {
			// 密码错误逻辑
			fmt.Println("密码错误")
		}
	} else {
		// 用户不存在，提示结束
		fmt.Printf("%s 用户不存在!\n", username)
	}
}
