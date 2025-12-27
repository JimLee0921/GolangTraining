package main

import (
	"fmt"
	"strings"
)

// SplitAfterSeq SplitAfter的迭代器版本，保留分隔符
func main() {
	s := "a,b,c,d"
	for part := range strings.SplitAfterSeq(s, ",") {
		fmt.Printf("%q\n", part)
	}
}
