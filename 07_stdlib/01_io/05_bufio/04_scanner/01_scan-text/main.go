package main

import (
	"bufio"
	"fmt"
	"os"
)

func ScannerDemo() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Please enter multi text(Ctrl+D end)")
	// 每次 Scan 扫描一行
	for scanner.Scan() {
		line := scanner.Text() // 获取当前行内容
		fmt.Println("read: ", line)
	}
	// 检查是否有错误
	if err := scanner.Err(); err != nil {
		fmt.Println("error:", err)
	}
}

func ReadFileDemo(filePath string) {
	file, _ := os.Open(filePath)
	defer file.Close()

	scanner := bufio.NewScanner(file)

	lineNumber := 1

	for scanner.Scan() {
		line := scanner.Text() // 获取一行内容
		fmt.Printf("%d lines: %s\n", lineNumber, line)
		lineNumber++
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("error:", err)
	}
}

func main() {
	//ScannerDemo()
	ReadFileDemo("temp_files/test.txt")
}
