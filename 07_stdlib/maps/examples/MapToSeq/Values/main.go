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

	for v := range maps.Values(m) {
		fmt.Println("value:", v)
	}

	next, stop := iter.Pull(maps.Values(m))
	defer stop()
	for {
		v, ok := next()
		if !ok {
			break
		}
		fmt.Println("value:", v)

	}
}
