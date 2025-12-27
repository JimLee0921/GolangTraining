package main

import (
	"fmt"
	"maps"
)

func main() {
	m1 := map[string]int{
		"key": 1,
	}
	m2 := maps.Clone(m1)
	// 浅拷贝，不会影响 m1
	m2["key"] = 199
	fmt.Println(m1)
	fmt.Println(m2)

	m3 := map[string][]int{
		"keys": {1, 2, 3},
	}
	m4 := maps.Clone(m3)
	// 引用对象使用的一个，会被修改
	m4["keys"][0] = 100
	fmt.Println(m3)
	fmt.Println(m4)
}
