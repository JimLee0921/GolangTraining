package main

import (
	"fmt"
	"strconv"
)

func main() {
	s, err := strconv.Unquote("you can't unquote a string without quotes")
	fmt.Printf("%q, %v\n", s, err)

	s, err = strconv.Unquote("\"the string must be either double-quoted\"")
	fmt.Printf("%q, %v\n", s, err)

	s, err = strconv.Unquote("`or backquoted`")
	fmt.Printf("%q, %v\n", s, err)

	s, err = strconv.Unquote("'\u263a'") // 单引号只能里面是单字符 rune
	fmt.Printf("%q, %v\n", s, err)

	// 单引号中不是单个字符会报错
	s, err = strconv.Unquote("'\u2639\u2639'")
	fmt.Printf("%q, %v\n", s, err)

}
