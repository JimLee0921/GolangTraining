package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func ReadAllDemo() {
	resp, _ := http.Get("https://www.baidu.com")
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	fmt.Println(string(body))
}

func ReadFullDemo() {
	f, _ := os.Open("temp_files/input.txt")
	defer f.Close()

	buf := make([]byte, 16)

	n, err := io.ReadFull(f, buf)
	fmt.Println("read bytes: ", n)
	fmt.Println("data:", buf)
	fmt.Println("error:", err)
}

func main() {
	//ReadAllDemo()
	ReadFullDemo()
}
