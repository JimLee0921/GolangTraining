package main

import (
	"fmt"
	"strconv"
)

func main() {
	// 普通文本 true
	fmt.Println(strconv.CanBackquote("Hello World"))

	// 包含 \t（运行）true
	fmt.Println(strconv.CanBackquote("Line1\tLine2"))

	// false
	fmt.Println(strconv.CanBackquote("Line1\nLine2")) // true

	// false
	fmt.Println(strconv.CanBackquote("Lin`e1\tLine2"))

}
