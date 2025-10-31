package main

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

func readBytesShoreText(text string, delimiter byte) {
	r := bufio.NewReader(strings.NewReader(text))
	for {
		line, err := r.ReadBytes(delimiter)
		fmt.Printf("read: %q\n", line)

		if err == io.EOF {
			fmt.Println("end reading")
			break
		} else if err != nil {
			fmt.Println("unknown error: ", err)
			break
		}
	}
}

func readBytesLoneText() {
	// 把缓冲区大小设置的很小来模拟文本很多的情况，但是 ReadBytes 会自动进行扩容拼接
	r := bufio.NewReaderSize(strings.NewReader("this_is_a_very_long_line\n"), 5)

	line, _ := r.ReadBytes('\n')
	fmt.Printf("complete line: %s\n", line)
}

func main() {
	readBytesShoreText("hello\nworld\nend", '\n')
	readBytesShoreText("hello,world,end", ',')
	readBytesLoneText()
}
