package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"strings"
)

func spiltLines(str string, delimiter byte) {
	// 传入字符串和分隔符，以指定分隔符进行划分
	r := bufio.NewReader(strings.NewReader(str))
	for {
		line, err := r.ReadSlice(delimiter) // 按照行进行分隔
		// 拷贝一份安全副本，如果不拷贝，下一次 ReadSlice() 调用会让 line 的内容变成别的字符串
		copyLine := append([]byte(nil), line...)
		if err == io.EOF {
			fmt.Println("end reading")
			break
		} else if errors.Is(err, bufio.ErrBufferFull) {
			fmt.Println("error: buffer full")
			break
		} else if err != nil {
			fmt.Println("unknown error: ", err)
		}
		fmt.Printf("read: %q\n", copyLine)
	}
}

func main() {
	//	spiltLines(`
	//	dajdjaf
	//aawdhjda
	//大撒和大家我
	//`, '\n')
	spiltLines("haha, this is a fucking day, yes yes!", ',')
}
