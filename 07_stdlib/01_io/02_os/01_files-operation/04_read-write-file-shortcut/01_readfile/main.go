package main

import (
	"fmt"
	"os"
)

func main() {
	// os.ReadFile 可以直接一次性读取整个文件内容
	data, err := os.ReadFile("temp_files/log.txt")
	if err != nil {
		panic(err)
	}
	// 把字节转为字符
	fmt.Println("file content: ", string(data))
}
