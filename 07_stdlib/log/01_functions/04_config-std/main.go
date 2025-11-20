package main

import (
	"fmt"
	"log"
	"os"
)

// 默认这些顶层函数就是给 log.std 配置的
func main() {
	log.SetPrefix("[INFO]")
	log.Println("Hello")

	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	log.Println("World")

	file, _ := os.Create("temp_files/app.log")
	log.SetOutput(file)
	log.Println("this will go to file")

	// 获取相关信息
	fmt.Println(log.Flags())
	fmt.Println(log.Prefix())
	fmt.Println(log.Writer())
}
