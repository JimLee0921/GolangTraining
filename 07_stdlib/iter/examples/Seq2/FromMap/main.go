package main

import (
	"fmt"
	"iter"
)

// FromMap 惰性遍历 map
func FromMap[K comparable, V any](m map[K]V) iter.Seq2[K, V] {
	return func(yield func(K, V) bool) {
		for k, v := range m {
			if !yield(k, v) {
				return
			}
		}
	}
}

func main() {
	m := map[string]int{
		"apple":  3,
		"banana": 5,
		"orange": 2,
		"peach":  20,
	}

	for k, v := range FromMap(m) {
		fmt.Println(k, v)
	}

	next, stop := iter.Pull2(FromMap(m))
	defer stop()
	for {
		k, v, ok := next()
		if !ok {
			break
		}
		fmt.Println(k, v)

		if v >= 5 {
			break // 提前停止消费
		}
	}
}
