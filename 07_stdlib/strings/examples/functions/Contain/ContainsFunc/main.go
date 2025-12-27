package main

import (
	"fmt"
	"strings"
)

// ContainsFunc 是否存在满足自定义规则条件的字符
func main() {
	// 查找字符串是否包含某个元音字母
	f := func(r rune) bool {
		return r == 'a' || r == 'e' || r == 'i' || r == 'o' || r == 'u'
	}
	fmt.Println(strings.ContainsFunc("hello", f))   // true
	fmt.Println(strings.ContainsFunc("rhythms", f)) // false
}
