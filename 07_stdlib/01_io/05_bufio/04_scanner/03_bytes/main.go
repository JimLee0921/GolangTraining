package main

import (
	"bufio"
	"fmt"
	"strings"
)

func main() {
	text := "Go Rust Python"
	scanner := bufio.NewScanner(strings.NewReader(text))
	scanner.Split(bufio.ScanWords)

	for scanner.Scan() {
		b := scanner.Bytes()
		// 需要使用 string() 把字节数组转会字符串
		fmt.Printf("%s: %v\n", string(b), b)
	}
}
