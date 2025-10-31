package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	f, _ := os.OpenFile("temp_files/test.txt", os.O_CREATE|os.O_RDWR, 0644)
	defer f.Close()

	var rw io.ReadWriter = f // os.File 同时满足 Reader + Writer

	rw.Write([]byte("Hello\n"))

	f.Seek(0, 0) // 移到开头
	buf := make([]byte, 10)
	n, _ := rw.Read(buf)

	fmt.Println(string(buf[:n]))
}
