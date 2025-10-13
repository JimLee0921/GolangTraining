package main

import "fmt"

func main() {
	/*
		rune 是 int32 的别名
		适合表示多字节 Unicode 字符（例如中文、emoji）
		遍历字符串时，如果使用 for range，Go 会自动按 rune 解析
	*/
	var r rune = '你'      // Unicode 字符
	fmt.Println(r)        // 输出: 20320 （'你' 的 Unicode 码点）
	fmt.Printf("%c\n", r) // 输出: 你
	fmt.Printf("%U\n", r) // 输出: U+4F60
	fmt.Printf("%T\n", r) // 输出: int32

	// 字符串转 rune 切片
	s := "你好"
	rs := []rune(s)
	fmt.Println(rs)                     // 输出: [20320 22909]
	fmt.Printf("%c %c\n", rs[0], rs[1]) // 输出: 你 好
}
