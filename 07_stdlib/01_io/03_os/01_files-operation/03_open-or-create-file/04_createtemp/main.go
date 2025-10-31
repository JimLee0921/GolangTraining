package main

import (
	"fmt"
	"os"
)

func main() {
	// CreateTemp 创建临时文件
	file, err := os.CreateTemp("temp_files", "")
	if err != nil {
		panic(err)
	}
	// 程序结束后自动关闭和删除
	//defer os.Remove(file.Name())
	defer file.Close()
	fmt.Println("temp file:", file.Name())

	// 向文件中写入内容
	file.WriteString("temp data：Hello, Go!\n")

	// 读取回内容验证
	data, _ := os.ReadFile(file.Name())
	fmt.Println("file content:", string(data))
}
