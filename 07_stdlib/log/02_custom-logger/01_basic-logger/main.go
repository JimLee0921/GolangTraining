package main

import (
	"log"
	"os"
)

func main() {
	// 创建文件用于日志输出
	file, err := os.OpenFile("temp_files/app.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// 自定义 logger
	logger := log.New(
		file,                               // 输出到文件
		"[MyApp]",                          // 日志前缀
		log.Ldate|log.Ltime|log.Lshortfile, // 日志格式
	)
	logger.Println("Application started")
	logger.Printf("User %s logged in", "JimLee")
}
