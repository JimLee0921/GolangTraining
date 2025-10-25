package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
	// 1. 连接服务器
	conn, err := net.Dial("tcp", "127.0.0.1:4000")
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	fmt.Errorf("connected tcp server")

	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("send message: ")
		text, _ := reader.ReadString('\n')

		// 2. 发送数据
		conn.Write([]byte(text))

		// 3, 接收服务器回复
		reply := make([]byte, 1024)
		n, _ := conn.Read(reply)
		fmt.Println("server reply:", string(reply[:n]))
	}
}
