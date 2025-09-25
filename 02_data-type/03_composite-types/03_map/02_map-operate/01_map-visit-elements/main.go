package main

import "fmt"

func main() {
	/*
		访问 map 元素
		直接通过键访问获取值，如果 key 不存在则返回 value 类型的零值
			value, ok = map[key]
			返回值:
				value：key 对应的值
				ok：可选接收返回值，主要区分零值和key不存在的情况
	*/
	myMap := map[string]int{"go": 1, "rust": 2, "python": 0}
	fmt.Println(myMap["go"])   // 1
	fmt.Println(myMap["java"]) // 0 key 不存在 返回 value 类型的零值

	// 接收 ok（推荐）
	v1, ok1 := myMap["rust"]
	fmt.Println(v1, ok1) // 2 true
	v2, ok2 := myMap["yes"]
	fmt.Println(v2, ok2) // 0 false
	v3, ok3 := myMap["python"]
	fmt.Println(v3, ok3) // 0 true

	if value, exists := myMap["hh"]; exists {
		fmt.Println("找到了元素")
		fmt.Println("val: ", value)
		fmt.Println("exists: ", exists)
	} else {
		fmt.Println("未找到元素")
		fmt.Println("val: ", value)
		fmt.Println("exists: ", exists)
	}
}
