package main

import (
	"fmt"
	"iter"
	"maps"
)

func main() {
	m := map[int]int{
		1: 10,
		2: 20,
		3: 23,
		4: 32,
	}

	// 返回的是单值 Seq 直接使用 range 遍历 m 的 keys
	for k := range maps.Keys(m) {
		fmt.Println("key:", k)
	}

	// 使用 iter.Pull
	next, stop := iter.Pull(maps.Keys(m))
	defer stop()

	for {
		k, ok := next()
		if !ok {
			break
		}
		fmt.Println("key:", k)

	}
}
