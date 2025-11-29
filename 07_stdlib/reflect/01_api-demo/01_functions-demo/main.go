package main

import (
	"fmt"
	"reflect"
)

func CopyDemo() {
	// reflect.Copy 函数
	src := []int{1, 2, 3, 4}
	dst := make([]int, 2)
	// 返回实际复制的数量
	n := reflect.Copy(reflect.ValueOf(dst), reflect.ValueOf(src))
	fmt.Println("copied: ", n) // 2
	fmt.Println("dst: ", dst)  // [1 2]
}

func DeepEqual() {
	fmt.Println(reflect.DeepEqual([]int{1, 2}, []int{1, 2}))                       // true
	fmt.Println(reflect.DeepEqual(map[string]int{"a": 1}, map[string]int{"a": 1})) // true
	fmt.Println(reflect.DeepEqual([]int{1}, []int{1, 2}))                          // false
	fmt.Println(reflect.DeepEqual(1, 2))                                           // false
	fmt.Println(reflect.DeepEqual(1, 1))                                           // true
}

func SwapperDemo() {
	// Swapper 排序，交换等操作
	names := []string{"a", "b", "c"}
	swap := reflect.Swapper(names)

	swap(0, 2) // 交换下标为 0，2的两个元素
	fmt.Println(names)
}

func TypeAssertDemo() {
	v := reflect.ValueOf(123)

	val, ok := reflect.TypeAssert[int](v)
	fmt.Println(val, ok) // 123 true

	val2, ok2 := reflect.TypeAssert[string](v)
	fmt.Println(val2, ok2) // "" false (不会panic)

	val3, ok3 := reflect.TypeAssert[[]int](reflect.ValueOf([]string{"123", "234"}))
	fmt.Println(val3, ok3)

	val4, ok4 := reflect.TypeAssert[[]int](reflect.ValueOf([]int{123, 234}))
	fmt.Println(val4, ok4)

}

func main() {
	CopyDemo()
	DeepEqual()
	SwapperDemo()
	TypeAssertDemo()
}
