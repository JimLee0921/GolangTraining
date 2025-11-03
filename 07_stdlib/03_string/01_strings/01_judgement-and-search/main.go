package main

import (
	"fmt"
	"strings"
)

func ContainsDemo() {
	s := "Hello World!"
	fmt.Println(strings.Contains(s, "Hello")) // true
	fmt.Println(strings.Contains(s, "hello")) // false
}

func HasPrefixDemo() {
	url := "/api/user/list"
	fmt.Println(strings.HasPrefix(url, "/api/")) // true
	fmt.Println(strings.HasPrefix(url, "api"))   // false
}

func HasSuffixDemo() {
	filename := "report.pdf"
	fmt.Println(strings.HasSuffix(filename, ".pdf")) // true
	fmt.Println(strings.HasSuffix(filename, ".csv")) // false
}

func IndexDemo() {
	s := "Hello World!"
	fmt.Println(strings.Index(s, "world")) // 找不到返回 -1
	fmt.Println(strings.Index(s, "World")) // 6
}

func LastIndexDemo() {
	s2 := "hello world world"
	fmt.Println(strings.LastIndex(s2, "world")) // 12
	fmt.Println(strings.LastIndex(s2, "dada"))  // -1
}

func main() {
	ContainsDemo()
	HasPrefixDemo()
	HasSuffixDemo()
	IndexDemo()
	LastIndexDemo()
}
