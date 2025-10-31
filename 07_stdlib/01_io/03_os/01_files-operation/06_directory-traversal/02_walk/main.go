package main

import (
	"fmt"
	"io/fs"
	"path/filepath"
)

func main() {
	// filepath.Walk 了解即可
	filepath.Walk("temp_files", func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			return err
		}
		fmt.Println(path, info.Size())
		return nil
	})
}
