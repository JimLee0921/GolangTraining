package main

import (
	"fmt"
	"strings"
)

// ReplaceAll 等价于 Replace 中 n 传入 -1 替换全部
func main() {
	fmt.Println(strings.ReplaceAll("oink oink oink", "oink", "moo"))
}
