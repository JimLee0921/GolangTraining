package main

import (
	"fmt"
	"strconv"
)

func main() {
	s := strconv.FormatBool(true)
	fmt.Printf("%T, %v\n", s, s)
}
