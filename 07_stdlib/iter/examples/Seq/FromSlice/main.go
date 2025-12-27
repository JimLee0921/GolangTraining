package main

import (
	"fmt"
	"iter"
)

// FromSlice 把 Slice 容器包装为 Seq
func FromSlice[V any](strSlice []V) iter.Seq[V] {
	return func(yield func(V) bool) {
		for _, v := range strSlice {
			if !yield(v) {
				return
			}
		}
	}
}

func main() {
	for s := range FromSlice([]string{"Go", "Java", "Python"}) {
		fmt.Println(s)
	}

	// 只取三个
	count := 0
	for v := range FromSlice([]int{5, 20, 9, 24, 56, 64, 74, 28, 19}) {
		fmt.Println(v)
		count++
		if count == 3 {
			break // 这里会让 yield 返回 false，立即停止继续迭代
		}
	}

	// 使用 Pull
	next, stop := iter.Pull(FromSlice([]string{"Go", "Java", "Python"}))
	defer stop()

	for {
		v, ok := next()
		if !ok {
			break
		}
		fmt.Println(v)
	}

}
