package main

import (
	"fmt"
	"strings"
)

func TrimDemo() {
	// Trim 去除字符集合
	res := strings.Trim("abchelloabcccc", "abc")
	fmt.Printf("%q\n", res) // "hello"
}

func TrimSpaceDemo() {
	// TrimSpace 去除前后空格制表符等内容
	res := strings.TrimSpace(" \n hello \t \t")
	fmt.Printf("%q\n", res) // "hello"
}

func TrimPrefixDemo() {
	// TrimPrefix 去除开头字串
	res := strings.TrimPrefix("Hello World", "Hello")
	fmt.Printf("%q\n", res) // " World"
}

func TrimSuffixDemo() {
	// TrimPrefix 去除开头字串
	res := strings.TrimSuffix("Hello World", "ld")
	fmt.Printf("%q\n", res) // "Hello Wor"
}

func main() {
	TrimDemo()
	TrimSpaceDemo()
	TrimPrefixDemo()
	TrimSuffixDemo()
}
