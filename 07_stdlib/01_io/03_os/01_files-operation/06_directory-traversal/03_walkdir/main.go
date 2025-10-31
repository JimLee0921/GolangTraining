package main

import (
	"fmt"
	"io/fs"
	"path/filepath"
)

func main() {
	// 最常用的递归遍历方式（会遍历所有子目录）
	err := filepath.WalkDir("temp_files", func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if d.IsDir() {
			fmt.Println("[DIR] ", path)
		} else {
			fmt.Println("[FILE]", path)
		}
		return nil
	})
	if err != nil {
		return
	}
}
