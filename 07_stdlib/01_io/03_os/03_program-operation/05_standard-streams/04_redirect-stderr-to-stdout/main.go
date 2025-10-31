package main

import (
	"fmt"
	"log"
)

func main() {
	// 可以使用管道符 2>&1 或者 &> 把 stderr 错误流也输出到 stdout
	// go run main.go > all.txt 2>&1  等价于 go run main.go &> all.txt   覆盖
	// go run main.go > all.txt 2>&1  等价于 go run main.go &>> all.txt  追加
	fmt.Println("dsb")                // 默认 stdout 标准流
	log.Println("unknow err", "haha") // 默认 stderr 标准流

}
