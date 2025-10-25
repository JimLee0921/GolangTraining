package main

import "fmt"

func main() {
	//参数之间自动加空格，末尾自动加换行符，返回字符串，不直接输出
	s := fmt.Sprintln("Hello", "Go", "Lang")
	fmt.Println(s)
}
