package main

import (
	"fmt"
	"io"
	"os"
)

func LimitReaderDemo() {
	f, _ := os.Open("temp_files/test.txt")
	defer f.Close()

	limited := io.LimitReader(f, 10)

	data, _ := io.ReadAll(limited)
	fmt.Println(string(data))
}

func TeeReaderDemo() {
	src, _ := os.Open("temp_files/input.txt")
	defer src.Close()

	dst, _ := os.Create("temp_files/output.txt")
	defer dst.Close()

	r := io.TeeReader(src, dst)
	io.Copy(os.Stdout, r) // 同时输出到屏幕 + 复制到 copy.txt
}

func main() {
	//LimitReaderDemo()
	TeeReaderDemo()
}
