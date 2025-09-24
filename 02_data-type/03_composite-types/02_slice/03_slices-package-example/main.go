package main

import (
	"cmp"
	"fmt"
	"slices"
)

func main() {
	/*
		复制与容量管理
		Clone
			签名：Clone[S ~[]E, E any](s S) S
			作用：拷贝一份新切片（深拷贝元素、独立底层数组）。
			例：s1 := slices.Clone(s)

		Grow
			签名：Grow[S ~[]E, E any](s S, n int) S
			作用：确保有空间可再追加 n 个元素；len 不变，可能重新分配底层数组。
			例：s = slices.Grow(s, 100)

		Clip
			签名：Clip[S ~[]E, E any](s S) S
			作用：把 cap 降到 len，释放多余容量（可能重新分配）。
			例：s = slices.Clip(s)

		排序 / 反转 / 去重

		Sort
			签名：Sort[S ~[]E, E cmp.Ordered](x S)
			作用：按升序原地排序（元素需可比较）。
			例：slices.Sort(s)

		SortFunc
			签名：SortFunc[S ~[]E, E any](x S, cmp func(a, b E) int)
			作用：自定义比较器排序；cmp(a,b) 返回负/0/正。
			例：slices.SortFunc(s, func(a,b T) int { return cmp.Compare(a,b) })

		Reverse
			签名：Reverse[S ~[]E, E any](s S)
			作用：原地反转。
			例：slices.Reverse(s)

		Compact
			签名：Compact[S ~[]E, E comparable](s S) S
			作用：移除相邻重复元素；通常配合 Sort 做“唯一化”。
			例：s = slices.Compact(slices.Sort(s))

		查找 / 比较

		Contains
			签名：Contains[S ~[]E, E comparable](s S, v E) bool
			作用：是否包含某值。
			例：ok := slices.Contains(s, 4)

		Index
			签名：Index[S ~[]E, E comparable](s S, v E) int
			作用：第一次出现的位置，找不到返回 -1。
			例：i := slices.Index(s, 3)

		BinarySearch
			签名：BinarySearch[S ~[]E, E cmp.Ordered](x S, target E) (int, bool)
			作用：已排序切片上二分查找；返回插入位与命中标记。
			例：idx, found := slices.BinarySearch(sorted, 10)

		编辑（改变内容或长度）

		Insert
			签名：Insert[S ~[]E, E any](s S, i int, v ...E) S
			作用：在位置 i 插入若干元素。
			例：s = slices.Insert(s, 1, 7, 8)

		Delete
			签名：Delete[S ~[]E, E any](s S, i, j int) S
			作用：删除区间 [i, j)。
			例：s = slices.Delete(s, 2, 4)

		Replace
			签名：Replace[S ~[]E, E any](s S, i, j int, v ...E) S
			作用：用 v... 替换区间 [i, j)。
			例：s = slices.Replace(s, 1, 2, 99, 100)
	*/
	s := []int{3, 1, 2, 2, 5, 4, 5}

	// 1) 排序 + 去重（唯一化）
	s1 := slices.Clone(s)   // 拷贝一份，避免改原切片
	slices.Sort(s1)         // [1 2 2 3 4 5 5]
	s1 = slices.Compact(s1) // 相邻去重 => [1 2 3 4 5]
	fmt.Println("unique sorted:", s1)

	// 2) 自定义排序（按绝对值）
	t := []int{-3, -1, 2, -4}
	slices.SortFunc(t, func(a, b int) int {
		return cmp.Compare(abs(a), abs(b))
	})
	fmt.Println("abs sort:", t) // [-1 2 -3 -4]

	// 3) 查找
	fmt.Println("Contains 4?", slices.Contains(s1, 4)) // true
	fmt.Println("Index of 3:", slices.Index(s1, 3))    // 2

	// 4) 二分查找（需要已排序）
	idx, found := slices.BinarySearch(s1, 4)
	fmt.Println("BinarySearch 4:", idx, found) // 3 true

	// 5) 插入 / 删除 / 替换
	u := []int{10, 20, 30}
	u = slices.Insert(u, 1, 99, 98)   // [10 99 98 20 30]
	u = slices.Delete(u, 2, 4)        // 删除 [2,4) => [10 99 30]
	u = slices.Replace(u, 1, 2, 7, 8) // 用 7,8 替换 [1,2) => [10 7 8 30]
	fmt.Println("after edit:", u)     // [10 7 8 30]

	// 6) 反转
	slices.Reverse(u)
	fmt.Println("reverse:", u) // [30 8 7 10]

	// 7) Clone / Grow / Clip（内存与容量管理）
	v := slices.Clone(s)                      // 独立拷贝
	v = v[:3]                                 // [3 1 2]
	v = slices.Grow(v, 100)                   // 仅扩 cap（len 不变），减少后续 append 扩容
	v = slices.Clip(v)                        // 把 cap 压到 len，利于释放多余内存
	fmt.Println("v len/cap:", len(v), cap(v)) // 3 3

}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
