package main

import (
	"fmt"
	"strings"
)

// Line 按换行符拆分返回迭代器
func main() {
	text := "Hello\nWorld\nGo Programing\n"
	for line := range strings.Lines(text) {
		fmt.Printf("%q\n", line)
	}
}
