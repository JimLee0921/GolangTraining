package main

import "fmt"

func main() {
	/*
		中文在 UTF-8 中占 3 个字节
		而使用 for value 遍历本质是按 byte 遍历会出现乱码
		使用 for range 本质是按 rune 遍历才能正确输出
	*/
	s := "Go语言"

	fmt.Println("按 byte 遍历：")
	for i := 0; i < len(s); i++ {
		fmt.Printf("%c ", s[i])
	}

	fmt.Println("\n\n按 rune 遍历：")
	for _, r := range s {
		fmt.Printf("%c ", r)
	}
}
