package main

import (
	"fmt"
	"os"
)

func main() {
	// MkdirTemp() 创建临时目录 自动随机名
	// 第一个参数指定在哪个目录下创建临时目录，可以传为空，go会自己操作
	// 第二个参数是临时目录名称的前缀，也可以为空
	dir, err := os.MkdirTemp("temp_files", "")
	if err != nil {
		panic(err)
	}
	fmt.Println("temp dir:", dir)
	defer os.RemoveAll(dir) // 使用完成手动删除
}
