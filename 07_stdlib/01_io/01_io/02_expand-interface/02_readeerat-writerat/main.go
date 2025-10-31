package main

import (
	"fmt"
	"os"
)

func ReaderAtDemo(filePath string) {
	// 打开文件，默认只读
	f, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	buf := make([]byte, 5)

	// 从第 10 个字节位置(从 0 开始)读取 5 个字节
	_, err = f.ReadAt(buf, 10)
	if err != nil {
		panic(err)
	}
	fmt.Println("read from 10 bytes:", string(buf)) // KLMNO

	// 再从第 0 个字节读取 5 个字节
	_, err = f.ReadAt(buf, 0)
	if err != nil {
		panic(err)
	}
	fmt.Println("read from beginning:", string(buf)) // ABCDE
}

func WriteAtDemo(filePath string) {
	// 以读写模式打开文件
	f, err := os.OpenFile(filePath, os.O_RDWR, 0644)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	// 在第 5 个字节位置覆盖写入 ----
	_, err = f.WriteAt([]byte("----"), 5)
	if err != nil {
		panic(err)
	}

	// 读取整个文件查看效果
	data, _ := os.ReadFile(filePath)
	fmt.Println("changed:")
	fmt.Println(string(data))
}

func main() {
	filePath := "temp_files/test.txt"
	WriteAtDemo(filePath)
}
