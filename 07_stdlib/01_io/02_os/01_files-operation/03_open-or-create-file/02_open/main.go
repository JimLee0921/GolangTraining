package main

import (
	"fmt"
	"os"
)

func main() {
	// Open 方法，不能创建只能打开文件读取内容
	file, err := os.Open("temp_files/test.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// 创建一个缓冲区，用来装读到的内容
	buf := make([]byte, 64) // 一次最多读 64 字节

	// 从文件中读取
	n, err := file.Read(buf)
	if err != nil {
		// 注意 EOF 不是严重错误，只表示文件结束
		if err.Error() != "EOF" {
			panic(err)
		}
	}

	// n 是本次实际读取的字节数
	fmt.Println("读取字节数：", n)
	fmt.Println("内容：", string(buf[:n]))
}
