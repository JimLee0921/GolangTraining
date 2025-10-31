package main

import (
	"bufio"
	"bytes"
	"fmt"
	"strings"
)

func ScanByWordsDemo(text string) {
	scanner := bufio.NewScanner(strings.NewReader(text))
	// 分割依据是空格、制表符、换行符等 Unicode 空白字符
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}

func ScanByBytesDemo(text string) {
	scanner := bufio.NewScanner(strings.NewReader(text))
	// 每次读取一个字节，非常适合调试或逐字符处理场景
	scanner.Split(bufio.ScanBytes)
	for scanner.Scan() {
		fmt.Printf("%q\n", scanner.Text())
	}
}

func ScanByRuneDemo(text string) {
	scanner := bufio.NewScanner(strings.NewReader(text))
	// 在处理中文或 emoji 时，要用 ScanRunes 而不是 ScanBytes
	scanner.Split(bufio.ScanRunes)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}

func SplitByCustomDelimiter(data []byte, atEOF bool) (advance int, token []byte, err error) {
	for i, b := range data {
		// 逗号作为分隔符
		if b == ',' {
			return i + 1, data[:i], nil // 截取到逗号前
		}
	}
	// 文件结束时返回剩余部分
	if atEOF && len(data) > 0 {
		return len(data), data, nil
	}
	return 0, nil, nil // 继续读取更多数据
}

func CustomDeliDemo(text string) {
	data := []byte(text)
	scanner := bufio.NewScanner(bytes.NewReader(data))
	scanner.Split(SplitByCustomDelimiter)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}

func main() {
	text := "Go Rust Python, Hello!"
	ScanByWordsDemo(text)
	ScanByBytesDemo(text)
	ScanByRuneDemo("你好哇haha👉")
	CustomDeliDemo("apple,banana,pear")
}
