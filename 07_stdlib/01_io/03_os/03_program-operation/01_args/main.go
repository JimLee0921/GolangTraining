package main

import (
	"fmt"
	"os"
)

func main() {
	// args 获取所有运行参数，第 0 个位置上的参数永远是程序自身路径
	// 使用命令行运行 go run main.go 参数列表: go run main.go hello world jimlee
	args := os.Args
	for index, arg := range args {
		fmt.Println(index, ":", arg)
	}
}
