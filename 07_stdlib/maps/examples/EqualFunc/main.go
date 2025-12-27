package main

import (
	"fmt"
	"maps"
	"strings"
)

func main() {
	m1 := map[int]string{
		1:    "one",
		10:   "Ten",
		1000: "THOUSAND",
	}

	m2 := map[int][]byte{
		1:    []byte("One"),
		10:   []byte("Ten"),
		1000: []byte("Thousand"),
	}
	// 自定义比较规则，字符串与字节数组进行比较，并且忽略大小写
	eq := maps.EqualFunc(m1, m2, func(s string, bytes []byte) bool {
		return strings.ToLower(s) == strings.ToLower(string(bytes))
	})
	fmt.Println(eq) // true
}
