package main

import (
	"fmt"
	"strconv"
)

func main() {
	b := []byte("quote-rune: ")
	b = strconv.AppendQuoteRune(b, '@')
	fmt.Println(string(b)) // quote-rune: '@'
}
