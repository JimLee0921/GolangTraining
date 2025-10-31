package main

import (
	"bufio"
	"fmt"
	"os"
)

func WriteDemo() {
	w := bufio.NewWriter(os.Stdout)
	defer w.Flush()
	p := []byte("Hello\n")
	n, err := w.Write(p)
	if err != nil {
		fmt.Println("error: ", err)
		return
	}
	fmt.Printf("write %d bytes\n", n)
}

func WriteToFile(text string) {
	// 使用 defer 的先进后出原则，先登记 Close，后登记 Flush（退出时先 Flush 再 Close）
	f, _ := os.Create("temp_files/output.txt")
	defer func() {
		if err := f.Close(); err != nil {
			fmt.Println("close error:", err)
		}
	}()
	w := bufio.NewWriter(f)
	defer func() {
		if err := w.Flush(); err != nil {
			fmt.Println("flush error:", err)
		}
	}()
	p := []byte(text)
	n, err := w.Write(p)
	if err != nil {
		fmt.Println("error:", err)
		return
	}
	fmt.Printf("Wrote %d/%d bytes (error: %v)\n", n, len(p), err)
}

func main() {
	//WriteDemo()
	WriteToFile("Hello World")
}
