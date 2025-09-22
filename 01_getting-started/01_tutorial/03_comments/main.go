package main

import "fmt"

// main 是程序的入口点 演示如何在 Go 中使用注释
func main() {
	fmt.Println("Hello Comments!") // 单行行尾注释: 打印一句话并换行
	/*
		这里是多行注释(块注释)
		通常用于临时屏蔽多行代码
		或者写一段较长的文字说明
		官方更推荐 // 风格的注释
	*/
}
