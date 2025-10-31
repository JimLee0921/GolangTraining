package main

import "fmt"

func main() {
	// 把所有参数拼接成一个字符串，不打印，返回给变量，常用于日志拼接、字符串构造
	s := fmt.Sprint("Hello", " ", "world", "!")
	fmt.Println(s)
}
