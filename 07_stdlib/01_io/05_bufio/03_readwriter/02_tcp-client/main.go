package main

import (
	"bufio"
	"fmt"
	"net"
)

func main() {
	// ReadWriter 在网络通信中最常见，比如 TCP
	conn, _ := net.Dial("tcp", "example.com:80")
	defer conn.Close()

	rw := bufio.NewReadWriter(
		bufio.NewReader(conn),
		bufio.NewWriter(conn),
	)

	// 写入 HTTP 请求
	rw.WriteString("GET / HTTP/1.1\r\n")
	rw.WriteString("Host: example.com\r\n")
	rw.WriteString("\r\n")
	rw.Flush()

	// 读取响应的第一行
	line, _ := rw.ReadString('\n')
	fmt.Println("Response:", line)
}
