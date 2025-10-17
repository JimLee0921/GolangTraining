package main

import "fmt"

func main() {
	// slice 元素为任意类型
	data := []interface{}{123, "Hello", true, 3.14}

	for _, v := range data {
		fmt.Printf("%v (%T)\n", v, v)
	}

	// map 值为任意类型
	info := map[string]any{
		"name": "Tom",
		"age":  24,
		"vip":  true,
	}
	fmt.Println(info)
}
