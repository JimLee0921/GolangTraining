package main

import (
	"fmt"
	"os"
)

func main() {
	// mkdir 创建单层目录，父目录不存在或目录已存在会报错。返回错误信息
	// Windows 会忽略这些 Unix perm 权限值，但仍需传入
	err := os.Mkdir("temp_files/temp", 0755)
	if err != nil {
		panic(err)
	}
	fmt.Println("create dir successful")
}
