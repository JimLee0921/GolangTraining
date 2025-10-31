package main

import (
	"bufio"
	"bytes"
	"fmt"
)

func main() {
	// Scanner 默认的最大 token 长度是 64 KB 如果一行太长，会返回 Err()

	// 生成一个超过 64K 的超长行
	data := bytes.Repeat([]byte("a"), 70000)
	scanner := bufio.NewScanner(bytes.NewReader(data))

	for scanner.Scan() {
		fmt.Println("scanned:", len(scanner.Text()))
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("error:", err)
	}
	// error: bufio.Scanner: token too long
}
