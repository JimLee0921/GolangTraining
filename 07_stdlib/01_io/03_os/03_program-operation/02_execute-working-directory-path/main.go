package main

import (
	"fmt"
	"os"
)

func main() {
	// 获取当前程序的完整路径
	path, err := os.Executable()
	if err != nil {
		panic(err)
	}
	fmt.Println("executable path:", path)

	// 获取当前工作目录
	path, err = os.Getwd()
	if err != nil {
		panic(err)
	}
	fmt.Println("wd path:", path)

	// 切换当前工作目录，必须存在且为目录
	err = os.Chdir("temp_files")
	if err != nil {
		panic(err)
	}
	path, err = os.Getwd()
	if err != nil {
		panic(err)
	}
	fmt.Println("new wd path:", path) // 输出：/tmp
}
