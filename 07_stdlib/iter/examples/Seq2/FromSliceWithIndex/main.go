package main

import (
	"fmt"
	"iter"
)

// FromSliceWithSlice 惰性遍历 slice，同时拿到索引和值
func FromSliceWithSlice[V any](xs []V) iter.Seq2[int, V] {
	return func(yield func(int, V) bool) {
		for i, v := range xs {
			if !yield(i, v) {
				return
			}
		}
	}
}

func main() {
	xs := []string{"a", "b", "c", "d", "e", "f", "g", "h"}
	//for i, v := range FromSliceWithSlice(xs) {
	//	fmt.Println(i, v)
	//	// 如果遍历到 c，后面的不再遍历
	//	if v == "c" {
	//		return
	//	}
	//}

	next, stop := iter.Pull2(FromSliceWithSlice(xs))
	defer stop()
	for {
		i, v, ok := next()
		if !ok {
			break
		}
		if i%2 == 0 {
			fmt.Println("even index: ", i, v)
		}
		if i == 5 {
			break // 提前结束
		}
	}
}
