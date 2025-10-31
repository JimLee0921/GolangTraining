package main

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

func readShoreLine(text string) {
	r := bufio.NewReader(strings.NewReader(text))
	for {
		line, isPrefix, err := r.ReadLine()

		if err == io.EOF {
			fmt.Println("end reading")
			break
		} else if err != nil {
			fmt.Println("unknown error:", err)
			break
		}
		fmt.Printf("line=%q, isPrefix=%v\n", line, isPrefix)

	}
}

func readLongLine(text string) {
	// 把缓冲区设置很小来模拟一行有很长的状态
	r := bufio.NewReaderSize(strings.NewReader(text), 3)
	// 定义字节切片
	var fullLine []byte
	// 把每次的内容循环拼接
	for {
		line, isPrefix, err := r.ReadLine()
		fullLine = append(fullLine, line...) // 拼接分段
		if err == io.EOF {
			fmt.Println("end reading")
			break
		} else if err != nil {
			fmt.Println("unknown error:", err)
			break
		}
		// isPrefix=true 说明一行还没结束
		if !isPrefix {
			// 说明本行读完
			break
		}
	}
	fmt.Println("complete line: ", string(fullLine))
}

func main() {
	readShoreLine("Hello\nWorld\nGolang\n哈哈")
	readLongLine("this is the first line\n") // 因为缓冲区太小，每次只读到部分；

}
