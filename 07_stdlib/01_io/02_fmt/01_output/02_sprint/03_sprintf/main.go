package main

import "fmt"

func main() {
	// 完全等同于 Printf，但不打印，而是返回字符串
	// 常用于：构造日志，动态生成文本，返回格式化信息给函数调用方
	name := "Alice"
	age := 25
	msg := fmt.Sprintf("%s is %d years old", name, age)
	fmt.Println(msg)
}
