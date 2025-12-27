package main

import (
	"fmt"
	"strconv"
)

// AppendBool 把一个 bool 值写入字节切片，必须重新接收
func main() {
	b := []byte("bool: ")
	b = strconv.AppendBool(b, true)
	fmt.Println(string(b))
}
