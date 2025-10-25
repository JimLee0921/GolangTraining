package main

import (
	"fmt"
	"os"
)

func main() {
	// 返回 file 对象，里面最常用的是里面的写入，获取信息的方法
	file, err := os.Create("temp_files/test.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	fmt.Println("create ", file.Name(), " successful")
}
