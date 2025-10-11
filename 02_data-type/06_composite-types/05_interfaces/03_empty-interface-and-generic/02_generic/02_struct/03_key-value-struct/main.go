package main

import "fmt"

type Pair[K comparable, V any] struct {
	Key   K
	Value V
}

func main() {
	/*
		多类型参数 [K, V]：第一个是键类型，第二个是值类型
		K 用 comparable 约束，因为常用于比较、map 键等
		这种模式在实现泛型 Map、Dict、Set 时非常常见
	*/
	p1 := Pair[string, int]{Key: "age", Value: 20}
	p2 := Pair[int, string]{Key: 1, Value: "Tom"}

	fmt.Println(p1) // {age 20}
	fmt.Println(p2) // {1 Tom}
}
