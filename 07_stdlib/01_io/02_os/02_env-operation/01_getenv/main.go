package main

import (
	"fmt"
	"os"
)

func main() {
	// os.Getenv() 获取环境变量值
	path := os.Getenv("PATH")
	fmt.Println("PATH =", path)
	fmt.Println("MY_ENV=", os.Getenv("MY_ENV")) // 不存在返回 ""
}
