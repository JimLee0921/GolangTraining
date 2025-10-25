package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	path := "temp_files/test.txt"

	info, _ := os.Stat(path)
	fmt.Println("before change:", info.Mode(), info.ModTime())

	// 修改权限
	os.Chmod(path, 0644)

	// 修改时间
	os.Chtimes(path, time.Now(), time.Now())

	// ✅ 重新获取最新文件信息
	newInfo, _ := os.Stat(path)
	fmt.Println("after change:", newInfo.Mode(), newInfo.ModTime())
}
