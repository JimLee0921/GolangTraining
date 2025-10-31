package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func MultiReaderDemo() {
	r := io.MultiReader(
		strings.NewReader("Hello "),
		strings.NewReader("World"),
		strings.NewReader("!"),
	)

	data, _ := io.ReadAll(r)
	fmt.Println(string(data)) // Hello World!
}

func MultiWriterDemo() {
	// 每个文件中都写一些内容
	f1, _ := os.Open("temp_files/1.txt")
	f2, _ := os.Open("temp_files/2.txt")
	f3, _ := os.Open("temp_files/3.txt")
	defer f1.Close()
	defer f2.Close()
	defer f3.Close()

	r := io.MultiReader(f1, f2, f3)
	io.Copy(os.Stdout, r)
}

func main() {
	//MultiReaderDemo()
	MultiWriterDemo()
}
