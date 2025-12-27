package main

import (
	"fmt"
	"strconv"
)

func main() {
	b := []byte("quote-ascii: ")
	b = strconv.AppendQuoteToASCII(b, `"哈哈Fran & Freddie's Diner'"`)
	fmt.Println(string(b)) // quote-ascii: "\"\u54c8\u54c8Fran & Freddie's Diner'\""
}
