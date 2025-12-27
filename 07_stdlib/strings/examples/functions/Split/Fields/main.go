package main

import (
	"fmt"
	"strings"
)

// Fields 自动对开头结尾空白字符进行清除并按照空白字符进行切割字符串
func main() {
	fmt.Printf("Fields are: %q\n", strings.Fields("    \n foo \n bar     baz  ")) // ["foo" "bar" "baz"]
}
