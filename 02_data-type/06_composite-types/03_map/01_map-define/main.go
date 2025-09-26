package main

import "fmt"

// main map映射基础创建方式
func main() {
	/*
		map 用来存储 键值对 (key-value)
		相当于其他语言里的字典（Python dict）、哈希表（Java HashMap）、对象字典（JavaScript Object）
		键不能重复值可以重复，注意内部元素是无序的
		Go 的 map 底层是 哈希表，元素的位置由哈希函数决定
		遍历 map (for range) 时，顺序是随机的
		每次运行，甚至同一次运行中的不同遍历，顺序都可能不同
		定于语法：map[键类型]值类型
		1. 使用 var 创建一个 nil map（只能读不能写）
			var m map[KeyType]ValueType
		2. 使用 make 创建，可立即读写但不能初始化
			make(map[KeyType]ValueType, hintSize)
				KeyType：键的类型（必须是可比较的类型，比如 string、int、bool）
				ValueType：值的类型（任意类型）
				hintSize（可选）：初始容量提示，大概需要多少空间
		3. 字面量创建，可立即读写且初始化
			map[KeyType]ValueType{}
			可选择传入初始化值，如果不传入也不是 nil map，而是一个空 map
	*/
	// var 创建一个零值（nil map，只能读不能写）
	var nilMap map[string]string
	fmt.Println(nilMap, nilMap == nil) // map[] true

	var varMap = map[string]string{"name": "Jim", "age": "23"}
	fmt.Println(varMap)

	// make定义：最常用，立即可读写
	makeMap1 := make(map[string]int)
	fmt.Println(makeMap1)

	// 字面量 创建一个空 map，但不是 nil
	shorthandMap1 := map[string]string{}
	fmt.Println(shorthandMap1, shorthandMap1 == nil) // map[] false

	// 字面量 创建一个 map 并初始化
	shorthandMap2 := map[string]int{"a": 1, "b": 2}
	fmt.Println(shorthandMap2)

}
