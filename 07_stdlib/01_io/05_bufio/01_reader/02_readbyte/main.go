package main

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

func ReadStringByByte(str string) {
	reader := bufio.NewReader(strings.NewReader(str))
	for {
		// 每次读取一个字节
		b, err := reader.ReadByte()
		if err == io.EOF {
			fmt.Println("end reading")
			break
		} else if err != nil {
			fmt.Println("unknown error:", err)
		}
		fmt.Printf("read byte:%d, char:%c\n", b, b)
	}
}

func WithUnReadByte(str string) {
	r := bufio.NewReader(strings.NewReader(str))

	// 第一次肚脐眼
	b1, _ := r.ReadByte()
	fmt.Printf("first read: %c\n", b1)

	// 把字符放回缓冲区
	err := r.UnreadByte()
	if err != nil {
		return
	}

	// 第二次读取
	b2, _ := r.ReadByte()
	fmt.Printf("second read: %c\n", b2)
}

func LookAhead(str string) {
	// 使用 UnReadByte 实现前瞻读取
	r := bufio.NewReader(strings.NewReader(str))

	b, _ := r.ReadByte()
	fmt.Printf("read byte: %c\n", b)
	if b != '=' {
		// 不符合条件，退回
		err := r.UnreadByte()
		if err != nil {
			return
		}
		fmt.Println("return byte")
	}
	// 再次读取整个内容
	rest, _ := r.ReadString('\n')
	fmt.Println("the rest:", rest)
}

func main() {
	//ReadStringByByte("dsb6666哈哈")
	//WithUnReadByte("ab")
	LookAhead("1=2")
}
