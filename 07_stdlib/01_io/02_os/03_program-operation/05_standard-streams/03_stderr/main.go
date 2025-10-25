package main

import (
	"fmt"
	"os"
)

func main() {
	/*
		Stderr 标准错误流，默认也是直接输出到控制台
		可以使用管道操作符 2> 或者 2>> 输出到指定文件
		2> 是覆盖， 2>> 是追加
		文件不存在也会自动创建
		go run main.go 2> error.txt
		go run main.go 2>> error.txt
	*/
	fmt.Fprintln(os.Stderr, "error output one")
	fmt.Fprintln(os.Stderr, "error output two")

}
