package main

import (
	"bufio"
	"fmt"
	"strings"
)

func BufferedShow() {
	// Buffered：返回当前缓冲区中还未被读取的字节数

	reader := bufio.NewReader(strings.NewReader("hello world"))
	readByte, err := reader.ReadByte()
	if err != nil {
		return
	} // 消费一个字节
	fmt.Printf("read byte: %q\n", readByte)
	// 缓冲区容量默认为 4 KB ，读取 h 后还剩十个字节
	fmt.Println("Buffered: ", reader.Buffered())
}

func DiscardShow() {
	reader := bufio.NewReader(strings.NewReader("hello world"))
	discard, err := reader.Discard(3)

	// 如果要丢弃的数据多于缓冲区，就会触发底层读
	// 若中途遇到 io.EOF，会返回 err=io.EOF 和实际跳过的字节数
	if err != nil {
		fmt.Println(err)
	}
	b, _ := reader.ReadString('\n')
	fmt.Printf("after discard %d bytes the rest: %v", discard, b)
}

func PeekShow() {
	reader := bufio.NewReader(strings.NewReader("hello world"))
	// Peek 返回接下来 n 个字节，但不前进指针（查看接下来的数据但是不消费数据）
	b, _ := reader.Peek(2) // 查看两个字节
	fmt.Printf("Peek(2): %s\n", b)

	ch, _ := reader.ReadByte() // 再次读取一个字节仍然是第一个 h
	fmt.Printf("ReadByte(): %c\n", ch)
}

func ResetShow() {
	// Reset 将现有的 bufio.Reader 重置为读取另一个数据源
	r1 := bufio.NewReader(strings.NewReader("hello"))
	buf := make([]byte, 5)
	r1.Read(buf)
	fmt.Println("first time:", string(buf))

	// 重置为另一个输入流
	r1.Reset(strings.NewReader("world"))
	r1.Read(buf)
	fmt.Println("second time:", string(buf))
}

func SizeShow() {
	// 返回缓冲器总容量（创建时指定大小）
	reader := bufio.NewReader(strings.NewReader("hello world"))
	fmt.Println(reader.Size()) // NewReader 默认大小就是 4096

	// 指定大小就是指定大小
	reader2 := bufio.NewReaderSize(strings.NewReader("Hello World"), 200)
	fmt.Println(reader2.Size())
}

func main() {
	//BufferedShow()
	//DiscardShow()
	//PeekShow()
	//ResetShow()
	SizeShow()
}
