package main

import "os"

func main() {
	// 使用 ReadFile + WriteFile 实现追加功能
	path := "temp_files/log.txt"
	data, _ := os.ReadFile(path)
	data = append(data, []byte("new line\n")...)
	os.WriteFile(path, data, 0644)
}
