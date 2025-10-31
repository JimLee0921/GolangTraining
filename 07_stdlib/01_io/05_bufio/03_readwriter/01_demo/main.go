package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// ReadWriter 让输入输出在一个对象上完成，代码更整洁
	rw := bufio.NewReadWriter(
		bufio.NewReader(os.Stdin),
		bufio.NewWriter(os.Stdout),
	)

	fmt.Print("enter text: ")
	text, _ := rw.ReadString('\n')

	_, err := rw.WriteString("your text is:" + text)
	if err != nil {
		return
	}
	rw.Flush()
}
