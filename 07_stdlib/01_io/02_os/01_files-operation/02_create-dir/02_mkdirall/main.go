package main

import (
	"fmt"
	"os"
)

func main() {
	// MkdirAll递归创建，自动创建不存在的父目录，路径已存在也不会报错
	err := os.MkdirAll("temp_files/temp/temp/haha", 0755)
	if err != nil {
		panic(err)
	}
	fmt.Println("create dirs successful")
}
