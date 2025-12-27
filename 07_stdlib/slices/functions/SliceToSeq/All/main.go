package main

import (
	"fmt"
	"iter"
	"slices"
)

func main() {
	names := []string{"Alice", "Bob", "Vera"}

	for i, v := range slices.All(names) {
		fmt.Println(i, ":", v)
	}

	// Pull2 使用
	next, stop := iter.Pull2(slices.All(names))
	defer stop()
	for {
		i, v, ok := next()
		if !ok {
			break
		}
		fmt.Println(i, ":", v)
	}
}
