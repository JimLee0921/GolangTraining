package main

import (
	"fmt"
	"net"
)

func main() {
	// 1. 指定服务器地址
	addr, _ := net.ResolveUDPAddr("udp", "127.0.0.1:3000")
	conn, _ := net.DialUDP("udp", nil, addr)
	defer conn.Close()

	// 2. 发送消息
	conn.Write([]byte("hello, server"))

	// 3. 接收回复
	buf := make([]byte, 1024)
	n, _, _ := conn.ReadFromUDP(buf)
	fmt.Println("client get reply:", string(buf[:n]))
}
