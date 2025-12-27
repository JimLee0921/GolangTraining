package main

import (
	"fmt"
	"maps"
)

func main() {
	m1 := map[string]int{
		"one": 1,
		"two": 2,
	}
	m2 := map[string]int{
		"one": 100,
	}
	maps.Copy(m2, m1)
	fmt.Println(m2)
	// 浅拷贝，基本数据类型不会被修改
	m2["one"] = 1000
	fmt.Println(m1)
	fmt.Println(m2)

	m3 := map[string][]int{
		"one": {1, 2, 3},
		"two": {4, 5, 6},
	}
	m4 := map[string][]int{
		"one": {8, 9, 0},
	}

	maps.Copy(m4, m3)
	fmt.Println(m4)
	// value 是引用数据类型，用的同一份数据，引用数据中某个值变了会被修改
	m4["one"][0] = 100

	fmt.Println(m3)
	fmt.Println(m4)

	// value 是引用数据类型，但是直接修改整个引用数据，老的 m3 不会被修改
	m4["one"] = []int{0, 0, 0, 0, 0}

	fmt.Println(m3)
	fmt.Println(m4)
}
