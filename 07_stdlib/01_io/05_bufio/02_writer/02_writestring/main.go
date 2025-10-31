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
	n, err := w.WriteString("Hello! Go 语言")
	fmt.Printf("Wrote %d bytes (err: %v)\n", n, err)
	w.Flush()
}
