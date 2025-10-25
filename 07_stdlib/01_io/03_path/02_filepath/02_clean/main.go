package main

import (
	"fmt"
	"path/filepath"
)

func main() {
	// Clean
	fmt.Println(filepath.Clean("/a/b/../c/./d//")) // "/a/c/d"
	fmt.Println(filepath.Clean("C:\\a\\..\\b\\"))  // "C:\b" (Windows)
}
