package main

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

func readStringByRune(str string) {
	r := bufio.NewReader(strings.NewReader(str))
	for {
		ch, size, err := r.ReadRune()
		if err == io.EOF {
			fmt.Println("end reading")
			break
		}
		fmt.Printf("char: %c, Unicode: %U, charNum: %d\n", ch, ch, size)
	}
}

func WithUnReadRune(str string) {
	r := bufio.NewReader(strings.NewReader(str))
	// 第一次读取
	ch, size, _ := r.ReadRune()
	fmt.Printf("first read:%c(%d bytes)\n", ch, size)
	// 撤回字符
	err := r.UnreadRune()
	if err != nil {
		return
	}
	// 第二次读取
	ch2, size2, _ := r.ReadRune()
	fmt.Printf("second read:%c(%d bytes)\n", ch2, size2)

}

func main() {
	// 字符读取，可以自动解析中文
	readStringByRune("Hello, Go语言")
	WithUnReadRune("傻X")
}
