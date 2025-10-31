package main

import (
	"fmt"
	"os"
)

func main() {
	// os.Exit() 程序立即退出，不执行 defer，通常传入非 0 表示错误
	defer fmt.Println("hahaha")
	fmt.Println("haha")
	fmt.Println("unknow error, exit immediately")
	os.Exit(1)
}
