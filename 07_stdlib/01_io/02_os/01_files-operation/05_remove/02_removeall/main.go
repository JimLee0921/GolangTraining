package main

import (
	"fmt"
	"os"
)

func main() {
	/*
		os.RemoveAll
			删除目标路径及其下的所有文件和子目录
			无论是文件、空目录还是非空目录都能删除；
			不存在也不会报错（直接返回 nil）。
	*/
	err := os.RemoveAll("temp_files/temp")
	if err != nil {
		panic(err)
	}
	fmt.Println("temp_files/temp deleted")
}
