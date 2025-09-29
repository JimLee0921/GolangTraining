package main

import (
	"fmt"
)

func main() {
	/*
		从 字符串（string 类型变量） 读取
		fmt.Sscan
			func Sscan(str string, a ...any) (n int, err error)
			从字符串 str 中读取，用空格和换行分隔，按顺序解析并填充到 a 中
		fmt.Sscanln
			func Sscanln(str string, a ...any) (n int, err error)
			类似 Sscan，但遇到换行符就停止，常用于按行解析字符串
		fmt.Sscanf
			func Sscanf(str string, format string, a ...any) (n int, err error)
			从字符串 str 中按格式化字符串 format 解析输入，相当于字符串版本的 Scanf
	*/
	input := "Tom 20"
	var name string
	var age int

	n, err := fmt.Sscan(input, &name, &age)
	if err != nil {
		fmt.Println("读取错误:", err)
		return
	}

	fmt.Printf("成功解析 %d 个字段: name=%s, age=%d\n", n, name, age)

	input1 := "Alice 30\nBob 25"
	var name1 string
	var age1 int

	_, _ = fmt.Sscanln(input1, &name1, &age1)
	fmt.Println(name1, age1) // Alice 30

	input2 := "Name:John Age:28"
	var name2 string
	var age2 int

	_, _ = fmt.Sscanf(input2, "Name:%s Age:%d", &name2, &age2)
	fmt.Println(name2, age2) // John 28
}
