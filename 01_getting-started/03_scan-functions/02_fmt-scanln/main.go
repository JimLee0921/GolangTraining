package main

import (
	"fmt"
)

func main() {
	/*
		标准输入 os.Stdin 读取
		fmt.Scanln
			func Scanln(a ...any) (n int, err error)
			类似 Scan，但以换行结束，不会跨行读取
			如果输入的字段比参数少，会报错
	*/
	var first, last string
	fmt.Println("请输入名字和姓氏，例如: Jim Lee")

	// 输入两个字符串，用空格分隔，遇到换行结束
	n, err := fmt.Scanln(&first, &last)
	if err != nil {
		fmt.Println("输入错误:", err)
		return
	}

	fmt.Printf("成功读取 %d 个字段\n", n)
	fmt.Printf("First: %s, Last: %s\n", first, last)
}
