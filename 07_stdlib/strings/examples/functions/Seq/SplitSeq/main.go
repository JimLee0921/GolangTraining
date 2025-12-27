package main

import (
	"fmt"
	"strings"
)

// SplitSeq Split的迭代器版本
func main() {
	s := "a,b,c,d"
	for part := range strings.SplitSeq(s, ",") {
		fmt.Printf("%q\n", part)
	}
}
