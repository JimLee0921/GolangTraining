package main

import (
	"fmt"
	"iter"
	"slices"
)

func main() {
	names := []string{"Alice", "Bob", "Vera"}

	for v := range slices.Values(names) {
		fmt.Println(v)
	}

	next, stop := iter.Pull(slices.Values(names))
	defer stop()
	for {
		v, ok := next()
		if !ok {
			break
		}
		fmt.Println(v)
	}
}
