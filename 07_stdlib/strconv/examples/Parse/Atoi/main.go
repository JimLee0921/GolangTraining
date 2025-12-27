package main

import (
	"fmt"
	"strconv"
)

func main() {
	i, err := strconv.Atoi("10")
	if err == nil {
		fmt.Printf("%T, %v", i, i)
	}
}
