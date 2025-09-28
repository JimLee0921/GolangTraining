package main

import (
	"fmt"
	"sort"
)

func main() {
	/*
		使用标准库已定义好的适配器类型对切片进行排序
		StringSlice / IntSlice / Float64Slice 都实现了 sort.Interface
		https://golang.org/pkg/sort/#Sort
	*/
	// 1. 字符串切片
	strSlice := []string{"JimLee", "Bruce", "Django", "James", "Tom"}
	fmt.Println("Before string sort: ", strSlice) // [JimLee Bruce Django James Tom]
	sort.StringSlice(strSlice).Sort()             // 等价于 	sort.Sort(sort.StringSlice(strSlice))
	fmt.Println("After string sort: ", strSlice)  //[Bruce Django James JimLee Tom]

	// 2. 整数切片
	intSlice := []int{42, 7, 19, 3, 100}
	fmt.Println("Before int sort: ", intSlice)
	sort.Sort(sort.IntSlice(intSlice)) // 等价于 sort.IntSlice(intSlice).Sort()

	fmt.Println("After int sort: ", intSlice)

	// 3. 浮点数切片
	floatSlice := []float64{3.14, 2.71, 1.41, 1.73, 0.577}
	fmt.Println("Before float sort: ", floatSlice)
	sort.Sort(sort.Float64Slice(floatSlice)) // 等价于 sort.Float64Slice(floatSlice).Sort()
	fmt.Println("After float sort: ", floatSlice)
}
