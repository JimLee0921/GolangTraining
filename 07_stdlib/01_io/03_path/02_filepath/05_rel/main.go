package main

import (
	"fmt"
	"path/filepath"
)

func main() {
	base := "/usr/local"
	target := "/usr/local/bin/go"
	rel, _ := filepath.Rel(base, target)
	fmt.Println(rel) // "bin/go"

	// "../c"
	fmt.Println(filepath.Rel("/a/b", "/a/c"))

}
