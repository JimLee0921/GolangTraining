package main

import (
	"fmt"
	"strings"
)

// SplitAfter 会保留 seq 切割字符串
func main() {
	fmt.Printf("%q\n", strings.SplitAfter("a,b,c", ",")) // ["a," "b," "c"]
}
