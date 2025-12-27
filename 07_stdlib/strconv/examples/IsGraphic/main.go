package main

import (
	"fmt"
	"strconv"
)

func main() {
	shamrock := strconv.IsGraphic('☘')
	fmt.Println(string('☘'), shamrock)

	a := strconv.IsGraphic('a')
	fmt.Println(string('a'), a)

	bel := strconv.IsGraphic('\007') // 转义字符
	fmt.Println(string('\007'), bel)
}
