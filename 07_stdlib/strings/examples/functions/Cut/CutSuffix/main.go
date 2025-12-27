package main

import (
	"fmt"
	"strings"
)

func main() {
	show := func(s, suffix string) {
		before, found := strings.CutSuffix(s, suffix)
		fmt.Printf("CutSuffix(%q, %q) = %q, %v\n", s, suffix, before, found)
	}
	show("Gopher", "Go")
	show("Gopher", "er")
}
