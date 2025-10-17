package main

import "fmt"

// main 演示如何从字符串取出单个字节再转为 rune，说明索引得到的是 byte。
func main() {
	/*
		Go 的字符串是 UTF-8 字节序列
		用索引 s[i] 取出来的不是“字符”，而是单个 byte
	*/
	// 取字符串 "A" 的第一个字节并转成 rune，确保后续哈希按 Unicode 处理字符
	runeLetter := rune("A"[0])
	bytesLetter := "A"[0]
	fmt.Printf("letter as int=%d, as char=%q, type=%T\n", runeLetter, runeLetter, runeLetter)         // 得到的是 rune(65)：int(32)
	fmt.Printf("bytesLetter as int=%d, as char=%q, type=%T\n", bytesLetter, bytesLetter, bytesLetter) // 得到的是 byte(65)：unit(8)

	/*
		letter as int=65, as char='A', type=int32
		bytesLetter as int=65, as char='A', type=uint8
	*/

} // 结论：在 Go 里字符串索引得到的是 byte，如果要做哈希运算，必须转成 rune（码点），才能正确处理所有字符
