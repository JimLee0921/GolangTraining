package main

import "fmt"

// main 切片中的长度 len 和容量 cap 属性
func main() {
	/*
		切片（slice）在 Go 中是一个轻量级的数据结构，底层包含：
			指针：指向底层数组的起始位置
			长度（len）：切片中当前包含的元素个数，可以用 len(s) 获取，只能访问 [0, len(s)-1] 范围内的元素
			容量（cap）：从起始位置到底层数组末尾的元素个数，可以用 cap(s) 获取，主要决定了 append 是否会触发扩容
		关系：
			总是有 0 ≤ len(s) ≤ cap(s)。
			扩容规则：如果 append 导致长度超过 cap，Go 会分配一个新的底层数组，并复制原有数据
	*/
	s := make([]int, 2, 5)      // 长度=2, 容量=5
	fmt.Println(len(s), cap(s)) // 2 5

	s = append(s, 10, 20, 30)
	fmt.Println(len(s), cap(s)) // 5 5 (刚好填满)

	s = append(s, 40)
	fmt.Println(len(s), cap(s)) // 6 10 (触发扩容 → 新数组)
}
