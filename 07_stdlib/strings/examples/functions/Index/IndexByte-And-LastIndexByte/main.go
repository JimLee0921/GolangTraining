package main

import (
	"fmt"
	"strings"
)

// IndexByteDemo 查找字符专用
func IndexByteDemo() {
	fmt.Println(strings.IndexByte("Hello, world", 'l')) // 2
	fmt.Println(strings.IndexByte("Hello, world", 'o')) // 4
	fmt.Println(strings.IndexByte("Hello, world", 'x')) // -1
}

// LastIndexByteDemo 查找字符专用
func LastIndexByteDemo() {
	fmt.Println(strings.LastIndexByte("Hello, world", 'l')) // 10
	fmt.Println(strings.LastIndexByte("Hello, world", 'o')) // 8
	fmt.Println(strings.LastIndexByte("Hello, world", 'x')) // -1
}

func main() {
	IndexByteDemo()
	LastIndexByteDemo()
}
