package main

import "fmt"

func main() {
	/*
		map 删除元素使用 delete 内置函数，直接修改原 map
		delete(map, key)
			map：需要操作删除的 map
			key：要删除的键
			如果 key 不存在，delete 什么都不会做，也不会报错
			方法没有返回值
	*/
	myMap := map[string]string{"name": "Jim", "age": "23"}
	myMap["sex"] = "male"
	fmt.Println(myMap)
	delete(myMap, "name") // 删除 name 键
	fmt.Println(myMap)
	delete(myMap, "color") // 没有 color 键，不报错，忽略
	fmt.Println(myMap)

	// 存在删除，不存在不进行 delete 操作
	if value, exists := myMap["age"]; exists {
		delete(myMap, "age")
		fmt.Println("age", value, "存在删除成功!")
	} else {
		fmt.Println("age不存在不进行删除操作")
	}
}
