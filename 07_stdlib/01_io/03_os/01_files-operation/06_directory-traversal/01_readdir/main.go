package main

import (
	"fmt"
	"os"
)

func main() {
	// os.ReadDir
	dirs, err := os.ReadDir("temp_files")
	if err != nil {
		panic(err)
	}
	for _, dir := range dirs {
		fileInfo, _ := dir.Info() // 详细信息
		fmt.Println(dir.Name(), dir.Type(), dir.IsDir(), fileInfo.ModTime())
	}
}
