package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	f, _ := os.Create("temp_files/output.txt")
	defer f.Close()

	w := bufio.NewWriter(f)

	// 英文，中文，表情都可以写入
	n1, _ := w.WriteRune('A')
	n2, _ := w.WriteRune('中')
	n3, _ := w.WriteRune('😀')
	fmt.Println(n1, n2, n3)
	err := w.Flush()
	if err != nil {
		fmt.Println("error: ", err)
		return
	}
	fmt.Println("write successful")
}
