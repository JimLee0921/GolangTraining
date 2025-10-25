package main

import (
	"fmt"
	"path"
	"path/filepath"
)

func main() {
	// Unix 下运行
	fmt.Println(path.IsAbs("/a/b"))     // true
	fmt.Println(filepath.IsAbs("/a/b")) // true

	// Windows 下运行
	fmt.Println(path.IsAbs("C:\\a\\b"))     // false
	fmt.Println(filepath.IsAbs("C:\\a\\b")) // true

}
