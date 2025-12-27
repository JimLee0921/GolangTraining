package main

import (
	"fmt"
	"strings"
)

// IndexRune 查找字符
func main() {
	fmt.Println(strings.IndexRune("chicken", 'k')) // 4
	fmt.Println(strings.IndexRune("chicken", 'd')) // -1
	fmt.Println(strings.IndexRune("中国Chine", '国')) // 3
}
