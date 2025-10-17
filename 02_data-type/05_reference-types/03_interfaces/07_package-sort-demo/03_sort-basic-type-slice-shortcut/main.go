package main

import (
	"fmt"
	"sort"
)

func main() {
	/*
		Go 标准库提供了便捷函数对常见切片类型排序：
			- sort.Strings([]string)
			- sort.Ints([]int)
			- sort.Float64s([]float64)
		这些函数内部其实就是把切片转换成对应的适配器类型
		然后调用 sort.Sort(...) 完成排序
	*/

	// 1. 对字符串切片直接排序
	s := []string{"JimLee", "Bruce", "Django", "James", "Tom"}
	fmt.Println("Before string sort:", s)
	sort.Strings(s)                      // 等价于 sort.Sort(sort.StringSlice(s))
	fmt.Println("After string sort:", s) // [Bruce Django James JimLee Tom]

	// 2. 对整数切片直接排序
	ints := []int{42, 7, 19, 3, 100}
	fmt.Println("Before int sort:", ints)
	sort.Ints(ints)                      // 等价于 sort.Sort(sort.IntSlice(ints))
	fmt.Println("After int sort:", ints) // [3 7 19 42 100]

	// 3. 对浮点数切片直接排序
	floats := []float64{3.14, 2.71, 1.41, 1.73, 0.577}
	fmt.Println("Before float sort:", floats)
	sort.Float64s(floats)                    // 等价于 sort.Sort(sort.Float64Slice(floats))
	fmt.Println("After float sort:", floats) // [0.577 1.41 1.73 2.71 3.14]
}
