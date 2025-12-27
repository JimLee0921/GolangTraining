package main

import (
	"fmt"
	"strconv"
)

func main() {
	// 字符串包含了一个 tab
	s := strconv.QuoteToASCII(`"Fran & Freddie's Diner	☺"`)
	fmt.Println(s) // "\"Fran & Freddie's Diner\t\u263a\""
}
