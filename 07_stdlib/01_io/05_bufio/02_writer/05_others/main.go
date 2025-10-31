package main

import (
	"bufio"
	"fmt"
	"os"
)

func AvailableShow() {
	w := bufio.NewWriterSize(os.Stdout, 8)
	fmt.Println(w.Available()) // 写入缓冲前缓冲区可用大小 8
	w.WriteString("1234123131的2")
	fmt.Println(w.Available()) // 数据写入缓冲后缓冲区可用大小 4
	w.Flush()                  // 这里会打印 1234
	fmt.Println(w.Available()) // Flush 后缓冲区可用大小 8
}

func BufferedShow() {
	w := bufio.NewWriterSize(os.Stdout, 8)
	w.WriteString("1234")
	fmt.Println(w.Buffered()) // 写入缓冲区后的字节数 4
	w.Flush()
	fmt.Println(w.Buffered()) // Flush 后的字节数 0
}

func SizeShow() {
	// 缓冲区总容量
	w := bufio.NewWriter(os.Stdout)
	fmt.Println(w.Size()) // 4096

	w2 := bufio.NewWriterSize(os.Stdout, 16)
	fmt.Println(w2.Size()) // 16
}

func ResetShow() {
	// 复用缓冲区
	f1, _ := os.Create("temp_files/output.txt")
	f2, _ := os.Create("temp_files/test.txt")
	defer f1.Close()
	defer f2.Close()

	w := bufio.NewWriter(f1)
	w.WriteString("File A\n")
	w.Flush()

	// 复用同一缓冲区，但换底层文件
	w.Reset(f2)
	w.WriteString("File B\n")
	w.Flush()

}

func main() {

	//AvailableShow()
	//BufferedShow()
	//SizeShow()
	ResetShow()
}
