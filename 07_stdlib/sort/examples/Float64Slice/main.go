package main

import (
	"fmt"
	"math"
	"sort"
)

func main() {
	values := []float64{
		3.5, math.NaN(), 1.2, 7.8, math.NaN(), 2.0,
	}
	fmt.Println("original data:", values)

	// 使用 Float64Slice 排序（NaN-safe）
	fs := sort.Float64Slice(values)
	fs.Sort()
	fmt.Println("after sort:", values)

	// 排序后进行二分查找
	target := 2.0
	idx := fs.Search(target)
	fmt.Println("search index:", idx)
}
