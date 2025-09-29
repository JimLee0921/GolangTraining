package main

import (
	"fmt"
)

func main() {
	/*
		标准输入 os.Stdin 读取
		fmt.Scanf
			func Scanf(format string, a ...any) (n int, err error)
			可以指定格式化字符串来解析输入
			类似于 C 语言的 scanf
	*/
	var name string
	var age int
	fmt.Println("请输入姓名和年龄，例如: Tom 20")

	// 指定格式：字符串 + 整数
	n, err := fmt.Scanf("%s %d", &name, &age)
	if err != nil {
		fmt.Println("输入错误:", err)
		return
	}

	fmt.Printf("成功读取 %d 个字段\n", n)
	fmt.Printf("Name: %s, Age: %d\n", name, age)
}
