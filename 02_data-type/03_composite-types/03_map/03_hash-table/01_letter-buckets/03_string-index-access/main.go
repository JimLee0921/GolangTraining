package main

import "fmt"

// main 演示对普通字符串取索引会返回字节，再转成 rune 查看字符的数值编码。
func main() {
	// word[0] 取到首字节，将其转换成 rune 并打印，用于说明哈希入口通常取首字符编码。
	word := "Hello"
	letter := rune(word[0])
	fmt.Println(letter)
	// 结论：哈希表按首字符分桶时，需要把索引结果视为 rune 参与运算。
}
