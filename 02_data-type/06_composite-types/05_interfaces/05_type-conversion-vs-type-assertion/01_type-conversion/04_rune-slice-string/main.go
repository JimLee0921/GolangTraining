package main

import "fmt"

// main rune切片转为string
func main() {
	/*
		string 的底层就是 UTF-8 字节，底层存放的是一段 只读的 byte 数组，所以 []byte <-> string 最直接
		[]rune -> string 也是可以的，底层会先把每个 rune 转成 UTF-8，再拼起来
		英文/ASCII 情况下两者等价；涉及中文/emoji 时，推荐用 rune，因为一个字符可能对应多个字节
		[]byte → 面向字节（UTF-8 编码，适合做网络传输、文件存储）
		[]rune → 面向字符（Unicode 码点，适合做字符串处理，比如截取、统计字符）
		string <-> []byte
			[]byte(s)：直接拷贝底层 UTF-8 字节
			string(bs)：直接把字节解释成 UTF-8 字符串
			无编码/解码开销（只是拷贝）

		string <-> []rune
			[]rune(s)：遍历字符串，把 UTF-8 解码成 Unicode 码点
			string(rs)：把每个 rune 再编码成 UTF-8 拼接
			有编码/解码开销

		[]byte <-> []rune
			不能直接互转，必须经过 string 作为桥梁
			rs := []rune(string(bs))
			bs := []byte(string(rs))
	*/
	// 定义一个字符串
	s := "中国hello"

	// string -> []rune
	rs := []rune(s)
	fmt.Println("string -> []rune:")
	fmt.Println(rs) // 打印 Unicode 码点
	for i, r := range rs {
		fmt.Printf("index=%d rune=%d char=%c\n", i, r, r)
	}

	// []rune -> string
	s2 := string(rs)
	fmt.Println("\n[]rune -> string:")
	fmt.Println(s2) // 恢复成原始字符串
	
}
