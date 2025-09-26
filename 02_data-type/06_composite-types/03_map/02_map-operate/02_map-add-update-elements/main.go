package main

import "fmt"

// main map添加元素
func main() {
	/*
		map 添加元素直接通过赋值语句即可
		map[key] = value
		如果 key 不存在，就是新增，如果 key 会覆盖原来的值，相当于更新
		注意如果 map 是 nil map，直接赋值会 panic
	*/
	myMap := map[string]string{"name": "Jim", "age": "23"}
	myMap["sex"] = "male"
	myMap["color"] = "orange" // 新增
	myMap["color"] = "yellow" // 更新
	fmt.Println(myMap)
	//var myMap2 map[int]string
	//myMap2[1] = "goodbye"	 会报错
}
