package main

import "fmt"

// main map映射基础创建方式
func main() {
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
