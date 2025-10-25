package main

import (
	"fmt"
	"path/filepath"
)

func main() {
	// join
	fmt.Println(filepath.Join("dir", "subdir", "file.txt"))
	fmt.Println(filepath.Join("/usr/", "local", "bin"))
	/*
		windows：
			dir\subdir\file.txt
			\usr\local\bin
		linux:
			dir/subdir/file.txt
			/usr/local/bin
	*/

	// Dir 和 Base
	fmt.Println(filepath.Dir("/a/b/c.txt"))
	fmt.Println(filepath.Base("/a/b/c.txt"))
	/*
		windows
			\a\b
			c.txt
		linux
			/a/b
			c.txt
	*/

	// Ext
	fmt.Println(filepath.Ext("/a/b/c.txt"))     // ".txt"
	fmt.Println(filepath.Ext("archive.tar.gz")) // ".gz"

}
