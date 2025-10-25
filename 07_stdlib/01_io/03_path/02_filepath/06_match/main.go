package main

import (
	"fmt"
	"path/filepath"
)

func main() {
	// Linux 下运行结果
	fmt.Println(filepath.Match("*.go", "main.go"))  // true <nil>
	fmt.Println(filepath.Match("a/*/c", "a/b/c"))   // true <nil>
	fmt.Println(filepath.Match("a/*/c", "a/x/y/c")) // false <nil>

}
