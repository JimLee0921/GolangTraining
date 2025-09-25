package main

import "fmt"

// main 演示切片中整数元素的自增，用于说明桶计数如何累加
func main() {
	// 创建一个长度为 1 的整型切片，依次赋值与自增，模拟桶命中次数递增
	buckets := make([]int, 1)
	fmt.Println(buckets[0])
	buckets[0] = 42 // 人为给桶计数赋值 42，相当于说这个桶目前被命中了 42 次
	fmt.Println(buckets[0])
	buckets[0]++ // 相当于又有一次命中计入了桶
	fmt.Println(buckets[0])
	// 结论：桶使用整型切片存储计数时，可直接用自增累积命中次数
}
