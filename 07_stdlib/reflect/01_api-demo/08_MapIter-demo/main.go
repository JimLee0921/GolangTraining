package main

import (
	"fmt"
	"reflect"
)

func main() {
	// MapIter 使用示例（每轮遍历都可能出现不同顺序，这符合 Go 的随机迭代行为）
	m := map[string]int{
		"a": 1,
		"b": 2,
		"c": 3,
	}

	// 用 reflect 取到 map 的 value
	v := reflect.ValueOf(m)

	// 创建迭代器
	iter := v.MapRange()

	fmt.Println("--- first time ---")

	for iter.Next() {
		k := iter.Key()
		val := iter.Value()
		fmt.Printf("Key: %v, Value: %v\n", k, val)
	}

	// Reset 重置迭代器，从头再次遍历同一张 map
	fmt.Println("--- Reset ---")
	iter.Reset(v)

	for iter.Next() {
		fmt.Printf("Key: %v, Value: %v\n", iter.Key(), iter.Value())
	}

}
