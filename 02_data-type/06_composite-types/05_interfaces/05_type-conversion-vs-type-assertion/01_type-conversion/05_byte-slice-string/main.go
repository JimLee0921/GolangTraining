package main

import "fmt"

func main() {
	/*
		因为一个中文占三个字节 这里遍历打印时与 rune[] 不同
	*/
	// 定义一个字符串
	s := "中国hello"

	// string -> []byte
	bs := []byte(s)
	fmt.Println("string -> []byte:")
	fmt.Println(bs) // 打印 UTF-8 编码后的字节
	for i, b := range bs {
		fmt.Printf("index=%d byte=%d char=%c\n", i, b, b)
	}

	// []byte -> string
	s2 := string(bs)
	fmt.Println("\n[]byte -> string:")
	fmt.Println(s2) // 恢复成原始字符串

}
