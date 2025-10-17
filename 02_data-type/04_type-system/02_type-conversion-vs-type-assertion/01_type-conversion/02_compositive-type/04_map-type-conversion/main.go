package main

import "fmt"

type M1 map[string]int
type M2 map[string]int

func main() {
	/*
		映射类型包括：
		key 类型 + value 类型
		两者必须完全相同
	*/
	m1 := M1{"a": 1}
	m2 := M2(m1)        // 可以（底层类型相同）
	fmt.Println(m1, m2) // map[a:1] map[a:1]
	var m3 map[string]string
	// m3 = m1 // key 类型相同，但值类型不同
	fmt.Println(m3)
}
