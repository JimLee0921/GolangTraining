package main

import "fmt"

func main() {
	str1 := "Hello"
	str2 := "Go语言"
	// 拼接字符串
	fmt.Println(str1 + " " + str2)

	// 使用 len 获取其字符串的字节长度 str2 为 8 因为 utf- 编码下汉字占 3 个字节
	fmt.Println(len(str1), len(str2))

	// 字节索引（返回 byte）
	fmt.Println(str1[1])
	// 需要手动转换
	fmt.Println(string(str1[1]))

	// 使用 for range 遍历字符(for index 遍历的是字节而不是字符!)
	//for i, c := range str2 {
	//	fmt.Println(i, string(c))
	//	fmt.Printf("%d: %c\n", i, c)
	//}
}
