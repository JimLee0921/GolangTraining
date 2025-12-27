package main

import (
	"fmt"
	"iter"
	"maps"
)

func main() {
	m := map[string]int{
		"one": 1,
		"two": 2,
	}

	// All 返回 iter.Seq2[K, V] 可以直接使用 for...range
	for k, v := range maps.All(m) {
		fmt.Printf("key=%s, value=%d\n", k, v)
	}

	// 使用 Pull2
	next, stop := iter.Pull2(maps.All(m))
	defer stop()
	for {
		k, v, ok := next()
		if !ok {
			break
		}
		fmt.Printf("key=%s, value=%d\n", k, v)
	}
}
