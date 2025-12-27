package main

import (
	"fmt"
	"strings"
	"unsafe"
)

func main() {
	s := "abc"
	clone := strings.Clone(s)
	fmt.Println(s == clone)                                       // true
	fmt.Println(unsafe.StringData(s) == unsafe.StringData(clone)) // false
}
