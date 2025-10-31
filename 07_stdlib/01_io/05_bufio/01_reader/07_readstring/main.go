package main

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

func ReadText(text string, delimiter byte) {
	r := bufio.NewReader(strings.NewReader(text))

	for {
		line, err := r.ReadString(delimiter)
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
func main() {
	ReadText("Hello\nWorld\nGoLang", '\n')
	ReadText("Hello,World,FFFFK", ',')
}
