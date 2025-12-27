package main

import (
	"fmt"
	"strings"
)

// ContainsRune 用于判断字符串中是否包含某个 Unicode 代码点 rune
func main() {
	// 小写字母 a 的 Unicode 码点是 97
	fmt.Println(strings.ContainsRune("aardvark", 97)) // true
	fmt.Println(strings.ContainsRune("timeout", 97))  // false
}
