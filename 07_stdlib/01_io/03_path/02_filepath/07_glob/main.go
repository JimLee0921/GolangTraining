package main

import (
	"fmt"
	"path/filepath"
)

func main() {
	// 使用 filepath.Join 构建模式
	pattern := filepath.Join("temp_files", "*.txt")

	files, err := filepath.Glob(pattern)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	if len(files) == 0 {
		fmt.Println("No matches found.")
		return
	}

	fmt.Println("Matched files:")
	for _, f := range files {
		fmt.Println(" -", f)
	}

}
