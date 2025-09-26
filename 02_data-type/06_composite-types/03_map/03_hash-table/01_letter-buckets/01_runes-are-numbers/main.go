package main

import "fmt"

// main 演示 rune 本质上是整数，强调哈希过程中字符可以直接视为数字。
func main() {
    // 取字符 'A'，输出其整数值与类型，说明底层是 int32。
    letter := 'A'
    fmt.Println(letter)
    fmt.Printf("%T \n", letter)
    // 结论：字符可按数字编码参与取模，为后续分桶打基础。
}