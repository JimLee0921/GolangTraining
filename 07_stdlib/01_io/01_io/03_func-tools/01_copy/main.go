package main

import (
	"io"
	"os"
)

func CopyDemo(srcPath, dstPath string) {
	// 使用 Copy 直接复制文件不需要自己循环 Read/Write
	src, _ := os.Open(srcPath)
	defer src.Close()
	dst, _ := os.Create(dstPath)
	defer dst.Close()

	io.Copy(dst, src)
}
func CopyNDemo(srcPath, dstPath string) {
	// 只拷贝前10个字节
	src, _ := os.Open(srcPath)
	defer src.Close()
	dst, _ := os.Create(dstPath)
	defer dst.Close()
	io.CopyN(dst, src, 10)
}

func main() {
	//CopyDemo("temp_files/input.txt", "temp_files/output.txt")
	CopyNDemo("temp_files/input.txt", "temp_files/output.txt")
}
