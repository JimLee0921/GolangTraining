package main

import (
	"fmt"
	"strings"
)

// SplitAfterN 可以指定切割字符串返回切片数量
func main() {
	fmt.Printf("%q\n", strings.SplitAfterN("a,b,c,d,e,f,", ",", 2)) // ["a," "b,c,d,e,f,"]
	fmt.Printf("%q\n", strings.SplitAfterN("a,b,c,d,e,f", "", 3))   // sep 为 "" 会在每个 UTF-8 序列后进行切割
	fmt.Printf("%q\n", strings.SplitAfterN("a,b,c,d,e,f", ",", 1))  // sep 为 "" 会在每个 UTF-8 序列后进行切割
	fmt.Printf("%q\n", strings.SplitAfterN("a,b,c,d,e,f", ",", -1)) // n 为 -1 等价于 SplitAfter
}
