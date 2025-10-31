package main

import (
	"bufio"
	"bytes"
	"fmt"
)

func main() {
	longLine := bytes.Repeat([]byte("A"), 70000) // 70KB
	scanner := bufio.NewScanner(bytes.NewReader(longLine))

	// 自定义缓冲区：初始 1KB，最大 1MB
	scanner.Buffer(make([]byte, 1024), 1024*1024)

	for scanner.Scan() {
		fmt.Println("read length:", len(scanner.Text()))
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("error:", err)
	}
}
