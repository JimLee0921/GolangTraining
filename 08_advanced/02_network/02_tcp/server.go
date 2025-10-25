package main

import (
	"bufio"
	"fmt"
	"net"
)

func handleConnection(conn net.Conn) {
	defer conn.Close()
	fmt.Printf("client %v connected\n", conn.RemoteAddr())

	reader := bufio.NewReader(conn)
	for {
		// 4. 读取客户端发送的数据
		msg, err := reader.ReadString('\n')
		if err != nil {
			fmt.Printf("client %v disconnected\n", conn.RemoteAddr())
			return
		}
		fmt.Println("client msg:", msg)
		// 5. 回复客户端
		reply := "server received:" + msg
		conn.Write([]byte(reply))
	}
}

func main() {
	// 1. 监听端口
	listener, err := net.Listen("tcp", ":4000")
	if err != nil {
		panic(err)
	}
	defer listener.Close()
	fmt.Println("TCP server starting, listening 4000 port")

	for {
		// 等待客户端连接
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("accept error:", err)
			continue
		}

		// 为每个客户端启动一个 goroutine 处理
		go handleConnection(conn)
	}
}
