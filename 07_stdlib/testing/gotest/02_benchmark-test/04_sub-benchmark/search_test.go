package search

import (
	"strconv"
	"testing"
)

/*
benchmark 创建子测试，查看两种查找算法性能差异
go test -bench .

最后结果可以看出来在在数组长度较短时两种查找方法耗时差别不大，但是数组长度一旦大起来，二分查找的优势就出来了
*/

func BenchmarkSearch(b *testing.B) {
	// 定义每次测试的数组长度
	sizes := []int{10, 100, 1000}

	for _, size := range sizes {
		size := size // 重新赋值，这样在子测试中不会影响源数据
		// 给子 benchmark 进行命名
		b.Run("size="+strconv.Itoa(size), func(b *testing.B) {
			// 针对每个 size 拆分两种实现
			data := make([]int, size)
			for i := range data {
				data[i] = i * 2
			}
			target := size * 2
			b.Run("linear", func(b *testing.B) {
				for b.Loop() {
					_ = LinearSearch(data, target)
				}
			})
			b.Run("binary", func(b *testing.B) {
				for b.Loop() {
					_ = BinarySearch(data, target)
				}
			})
		})
	}
}
