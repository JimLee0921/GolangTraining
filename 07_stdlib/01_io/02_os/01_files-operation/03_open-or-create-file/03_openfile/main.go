package main

import (
	"fmt"
	"os"
)

func main() {
	// OpenFile 方法，最常用，可以组合标志做到追加，读写，清空等操作

	// 创建新文件
	//flag := os.O_CREATE | os.O_APPEND | os.O_WRONLY
	flag := os.O_RDONLY
	file, err := os.OpenFile("temp_files/app.log", flag, 0644)
	if err != nil {
		panic(err)
	}
	fmt.Println(file.Name())
	defer file.Close()
}
