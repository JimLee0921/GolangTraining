package main

import "fmt"

func main() {
	/*
		byte 实质是 uint8
		常用于处理纯英文文本、二进制数据或网络通信中 1 字节的内容
		打印 %c 表示字符，%d 表示整数值
	*/
	var a byte = 'A'      // ASCII 'A' 的十进制值为 65
	fmt.Println(a)        // 输出: 65
	fmt.Printf("%c\n", a) // 输出: A
	fmt.Printf("%T\n", a) // 输出: uint8

	// 字符串可以被当作字节切片（这里一个中文 utf-8 编码下占三个字节）
	s := "Go语言"
	bs := []byte(s)
	fmt.Println(bs)                     // 输出: [71 111 232 175 173 232 168 128]
	fmt.Printf("%c %c\n", bs[0], bs[1]) // 输出: G o
}
