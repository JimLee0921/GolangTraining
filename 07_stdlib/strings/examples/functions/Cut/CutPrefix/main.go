package main

import (
	"fmt"
	"strings"
)

func main() {
	show := func(s, prefix string) {
		after, found := strings.CutPrefix(s, prefix)
		fmt.Printf("CutPrefix(%q, %q) = %q, %v\n", s, prefix, after, found)
	}

	show("Gopher", "Go")
	show("Gopher", "ph")
}
