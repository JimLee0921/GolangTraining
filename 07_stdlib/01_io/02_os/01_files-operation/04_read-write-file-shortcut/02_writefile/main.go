package main

import (
	"fmt"
	"os"
)

func main() {
	// os.WriteFile 一次性写入整个文件（会覆盖，第二次写入直接把第一次全部覆盖了）
	content1 := []byte("Hello, 大傻逼")
	content2 := []byte("????????")
	err1 := os.WriteFile("temp_files/log.txt", content1, 0644)
	if err1 != nil {
		panic(err1)
	}
	err2 := os.WriteFile("temp_files/log.txt", content2, 0644)
	if err2 != nil {
		panic(err2)
	}
	fmt.Println("write file successful")
}
