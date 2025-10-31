package main

import (
	"bytes"
	"fmt"
)

func main() {
	// 输出到网络连接（或内存缓冲）
	// 创建一个 bytes.Buffer（实现了 io.Writer）
	var buf bytes.Buffer
	fmt.Fprint(&buf, "Hahaha")
	fmt.Fprintf(&buf, "User: %s | Score: %d\n", "Tuo", 99)
	fmt.Fprintln(&buf, "All operations done.")

	// 输出最终内容
	fmt.Println(buf.String())
}
