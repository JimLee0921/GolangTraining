package main

import (
	"bufio"
	"fmt"
	"os"
)

func autoFlushShow() {
	// 只给 4 字节缓冲
	w := bufio.NewWriterSize(os.Stdout, 4)

	w.WriteByte('1')
	fmt.Println("Buffered:", w.Buffered())

	w.WriteByte('2')
	fmt.Println("Buffered:", w.Buffered())

	w.WriteByte('3')
	w.WriteByte('4')
	fmt.Println("Buffered:", w.Buffered())

	// 缓冲区已满，下一次写会触发自动 Flush 所以这里之前会输出 1234
	w.WriteByte('5')
	fmt.Println("Buffered:", w.Buffered())

	w.Flush()
}

func WriteByteShow() {
	f, _ := os.Create("temp_files/output.txt")
	defer f.Close()

	w := bufio.NewWriter(f)
	w.WriteByte('A')
	w.WriteByte('\n')
	w.WriteByte('1')

	w.Flush()
}

func main() {
	autoFlushShow()
	WriteByteShow()
}
