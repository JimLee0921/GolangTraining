package main

import (
	"fmt"
	"path"
	"path/filepath"
)

func main() {
	fmt.Println(path.IsAbs("/dev/null"))
	fmt.Println(path.IsAbs("a/b/c"))
	fmt.Println(path.IsAbs("04_path-base"))
	fmt.Println(path.IsAbs("07_stdlib/02_path/01_path/04_path-base"))
	fmt.Println(path.IsAbs("C:\\demo\\GolangTraining\\07_stdlib\\02_path\\01_path\\04_path-base"))
	fmt.Println(filepath.IsAbs("/dev/null"))
	fmt.Println(filepath.IsAbs("a/b/c"))
	fmt.Println(filepath.IsAbs("04_path-base"))
	fmt.Println(filepath.IsAbs("07_stdlib/02_path/01_path/04_path-base"))
	fmt.Println(filepath.IsAbs("C:\\demo\\GolangTraining\\07_stdlib\\02_path\\01_path\\04_path-base"))
	/*
			true
		false
		false
		false
		false
		false
		false
		false
		false
		true
	*/
}
