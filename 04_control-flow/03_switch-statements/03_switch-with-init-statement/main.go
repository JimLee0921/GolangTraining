package main

import "fmt"

// main switch后还可以添加一个赋值语句（）
func main() {
	/*
		类似于 if 后的赋值语句
		只能在 switch 结构体内使用
	*/
	switch x := 10; {
	case x < 0:
		fmt.Println("y < 0")
	case x == 0:
		fmt.Println("y == 0")
	default:
		fmt.Println("y > 0")
	}
	role := "傻逼"
	switch leaderMessage, staffMessage, message := "欢迎领导", "欢迎员工", "欢迎"; role {
	case "管理员", "主管":
		fmt.Println(leaderMessage)
	case "员工":
		fmt.Println(staffMessage)
	default:
		fmt.Println(message)
	}
}
