package main

import (
	"fmt"
	"os"
)

func main() {
	// os.Remove() 删除单个文件或空目录
	//err := os.Remove("temp_files/config.yaml")
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Println("temp_files/config.yaml deleted")

	// 删除目录,必须为空目录
	err := os.Remove("temp_files/temp/temp/haha")
	if err != nil {
		panic(err)
	}
	fmt.Println("temp_files deleted")
}
