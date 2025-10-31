package main

import (
	"fmt"
	"os"
)

func main() {
	// os.Rename 重命名或移动文件或文件夹
	// 把 temp_files 下的 log.txt 重命名为 logs.txt
	err := os.Rename("temp_files/log.txt", "temp_files/logs.txt")
	if err != nil {
		panic(err)
	}
	fmt.Println("rename successful")

	// 把 temp_files temp 文件夹重命名为 temp_dir
	err = os.Rename("temp_files/temp", "temp_files/temp_dir")
	if err != nil {
		panic(err)
	}
	fmt.Println("rename successful")

	// 把 temp_files 下的 log.txt 移动到 temp_files 下的 temp_dir 中并重命名
	err = os.Rename("temp_files/logs.txt", "temp_files/temp_dir/log.txt")
	if err != nil {
		panic(err)
	}
	fmt.Println("rename successful")
}
