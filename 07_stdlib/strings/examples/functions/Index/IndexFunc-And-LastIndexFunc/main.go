package main

import (
	"fmt"
	"strings"
	"unicode"
)

// IndexFuncDemo 自定义查找规则
func IndexFuncDemo() {
	// 查找第一个中文
	f := func(c rune) bool {
		return unicode.Is(unicode.Han, c)
	}
	fmt.Println(strings.IndexFunc("Hello, 世界", f))    // 7
	fmt.Println(strings.IndexFunc("Hello, world", f)) // -1
}

func LastIndexFuncDemo() {
	// 查找第一个中文
	f := func(c rune) bool {
		return unicode.Is(unicode.Han, c)
	}
	fmt.Println(strings.LastIndexFunc("Hello, 世界", f))             // 10
	fmt.Println(strings.LastIndexFunc("Hello, world", f))          // -1
	fmt.Println(strings.LastIndexFunc("go 123", unicode.IsNumber)) // 现成规则，5
}

func main() {
	IndexFuncDemo()
	LastIndexFuncDemo()
}
