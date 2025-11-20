package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strings.ToUpper("Hello World"))
	fmt.Println(strings.ToLower("Hello World"))
	fmt.Println(strings.Title("hello world")) // 已弃用
	fmt.Println(strings.ToTitle("Hello World"))
	/*
		HELLO WORLD
		hello world
		Hello World
		HELLO WORLD
	*/
	fmt.Println(strings.EqualFold("Go", "go"))         // true
	fmt.Println(strings.EqualFold("Gopher", "gopher")) // true
	fmt.Println(strings.EqualFold("GO语言", "go语言"))     // true
}
