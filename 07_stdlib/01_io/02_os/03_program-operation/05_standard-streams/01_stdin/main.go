package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// 默认从控制台读取。使用 < 管道符指定从文件读取：go run main.go < input.txt
	fmt.Print("Enter your name: ")
	reader := bufio.NewReader(os.Stdin)
	name, _ := reader.ReadString('\n') // 遇到换行结束读取
	fmt.Println("Hello,", name)
}
