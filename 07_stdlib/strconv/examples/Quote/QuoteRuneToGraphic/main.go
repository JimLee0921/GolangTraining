package main

import (
	"fmt"
	"strconv"
)

func main() {
	s := strconv.QuoteRuneToGraphic('☺')
	fmt.Println(s) // '☺'

	s = strconv.QuoteRuneToGraphic('\u263a')
	fmt.Println(s) // '☺'

	s = strconv.QuoteRuneToGraphic('\u000a')
	fmt.Println(s) // '\n'

	s = strconv.QuoteRuneToGraphic('	') // 一个 tab 空格
	fmt.Println(s)                      // '\t'
}
