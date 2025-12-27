package main

import (
	"fmt"
	"strconv"
)

func main() {
	s, err := strconv.QuotedPrefix("not a quoted string")
	fmt.Printf("%q, %v\n", s, err)

	s, err = strconv.QuotedPrefix("\"double-quoted string\" with trailing text")
	fmt.Printf("%q, %v\n", s, err)

	s, err = strconv.QuotedPrefix("`or backquoted` with more trailing text")
	fmt.Printf("%q, %v\n", s, err)

	s, err = strconv.QuotedPrefix("'\u263a' is also okay")
	fmt.Printf("%q, %v\n", s, err)

}
