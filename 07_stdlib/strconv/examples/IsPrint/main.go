package main

import (
	"fmt"
	"strconv"
)

func main() {
	c := strconv.IsPrint('\u263a')
	fmt.Println(string('\u263a'), c)

	bel := strconv.IsPrint('\007')
	fmt.Println(string('\007'), bel)
}
