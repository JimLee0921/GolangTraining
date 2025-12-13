package sort

import "math/rand"

// 准备数据阶段，就是简单的复制一份，不影响源数据
func prepareData(base []int) []int {
	cp := make([]int, len(base))
	copy(cp, base)
	return cp
}

// 生成数据
func makeData() []int {
	data := make([]int, 10)
	for i := range data {
		data[i] = rand.Intn(100) // 0-99
	}
	return data
}

// Sort 简单排序实现
func Sort(data []int) {
	n := len(data)
	// n 小于 2 说明空或就一个元素，没有排序必要
	if n < 2 {
		return
	}

	for i := 0; i < n-1; i++ {
		for j := 0; j < n-1-i; j++ {
			if data[j] > data[j+1] {
				data[j], data[j+1] = data[j+1], data[j]
			}
		}
	}
}
