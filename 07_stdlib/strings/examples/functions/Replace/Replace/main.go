package main

import (
	"fmt"
	"strings"
)

// Replace 对字符串中指定子串替换最多 n 个
func main() {
	fmt.Println(strings.Replace("oink oink oink", "k", "ky", 2))
	fmt.Println(strings.Replace("oink oink oink", "oink", "moo", -1)) // -1 默认全部替换
}
