package main

import (
	"fmt"
	"os"
)

func main() {
	/*
		main 函数
	*/

	fmt.Println("func main is the entry point to your program")
	fmt.Println("命令行参数:", os.Args)
	os.Exit(1) // 携带问题退出

}
