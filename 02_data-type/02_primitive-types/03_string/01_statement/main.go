package main

import "fmt"

func main() {
	// 定义普通字符串
	str1 := "Hello"
	str2 := "Go语言"
	fmt.Println(str1, str2)
	// 定义多行字符串（原样输出，空格也会保留）
	text := `Hello
		World
	hhh`
	fmt.Println(text)
}
