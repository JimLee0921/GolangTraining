package main

import (
	"fmt"
	"os"
)

func main() {
	// 输出到文件
	// 创建或打开文件 os.Create 如果文件存在，会被清空覆盖
	file, err := os.Create("temp_files/log.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	// 输入到文件（不换行没空格）
	fmt.Fprint(file, "hello, ")
	fmt.Fprintln(file, "JimLee!")
	fmt.Fprintf(file, "\nThis is written using Fprint.\n")
}
