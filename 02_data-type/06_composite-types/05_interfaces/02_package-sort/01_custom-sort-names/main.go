package main

import (
	"fmt"
	"sort"
)

// 自定义类型，本质上还是字符串切片 []string
type people []string

// Len 返回切片长度
func (p people) Len() int {
	return len(p)
}

// Swap 交换两个元素
func (p people) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

// Less 比较两个元素的大小，决定排序顺序
func (p people) Less(i, j int) bool {
	return p[i] < p[j]
}

// main
func main() {
	/*
		为自定义 people 类型绑定 Len, Swap, Less 方法从而实现 sort.Interface 接口
		https://golang.org/pkg/sort/#Sort
		https://golang.org/pkg/sort/#Interface
	*/
	peopleSlice := people{"JimLee", "Bruce", "Django", "James", "Tom"}
	fmt.Println(peopleSlice) // [JimLee Bruce Django James Tom]
	sort.Sort(peopleSlice)   // 调用标准库的排序函数，排序依据就是 Less
	fmt.Println(peopleSlice) // [Bruce Django James JimLee Tom]
}
