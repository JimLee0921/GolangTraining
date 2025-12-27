package main

import (
	"fmt"
	"strings"
	"unicode"
)

// FieldsFunc 自定义分割规则（和SplitXxxx系列不同，会默认把开头结尾符合要求的 rune 删除）
func main() {
	f := func(c rune) bool {
		return !unicode.IsLetter(c) && !unicode.IsNumber(c)
	}
	fmt.Printf("Fields are %q\n", strings.FieldsFunc(" \n foo; bar2,baz3...", f))
}
