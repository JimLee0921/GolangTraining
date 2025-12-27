package main

import (
	"fmt"
	"strings"
)

// SplitN 可以控制字符串切割返回的切片数量
func main() {
	fmt.Printf("%q\n", strings.SplitN("a,b,c,d,2,d2,d,f", ",", 3))
	z := strings.SplitN("a,b,c,", ",", 0) // 切割 0 次返回 nil
	fmt.Printf("%q(nil = %v)\n", z, z == nil)
}
