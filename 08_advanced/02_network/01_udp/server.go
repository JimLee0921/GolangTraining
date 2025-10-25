package main

import (
	"fmt"
	"net"
)

func main() {
	// 1. 监听 UDP 地址
	addr, _ := net.ResolveUDPAddr("udp", ":3000")
	conn, _ := net.ListenUDP("udp", addr)
	defer conn.Close()

	buf := make([]byte, 1024)
	fmt.Println("UDP server started on port 3000...")

	// 2. 无限 for 循环阻塞
	for {
		// 3. 读取客户端信息
		n, clientAddr, _ := conn.ReadFromUDP(buf)
		msg := string(buf[:n])
		fmt.Printf("reveive %v message:%s\n", clientAddr, msg)

		// 4. 回复响应
		reply := []byte("server get msg:" + msg)
		conn.WriteToUDP(reply, clientAddr)
	}
}
