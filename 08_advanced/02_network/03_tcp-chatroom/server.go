package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
	"sync"
)

// ---------------------------
// 全局变量区
// ---------------------------

// clients 用于保存当前所有已连接的客户端
// key 是连接对象 net.Conn，value 是客户端的名字
var (
	clients   = make(map[net.Conn]string)
	clientsMu sync.Mutex // 互斥锁，用来防止并发访问 clients 时出错
)

// broadcastMessage 向所有客户端广播一条消息
// 参数 message 是要发送的内容
// sender 是发送者（自己不需要收到自己的消息）
func broadcastMessage(message string, sender net.Conn) {
	clientsMu.Lock()         // 加锁，防止多个 goroutine 同时访问 clients 出问题
	defer clientsMu.Unlock() // 函数结束时自动解锁

	// 遍历所有在线客户端
	for conn := range clients {
		if conn != sender { // 不给发送者自己发
			conn.Write([]byte(message)) // 向客户端发送消息
		}
	}
}

// handleConnection 负责处理一个客户端的整个生命周期
// 包括：登录、接收消息、转发消息、退出
func handleConnection(conn net.Conn) {
	defer conn.Close() // 函数结束时关闭连接

	// 创建一个带缓冲的读取器（从客户端读取消息）
	reader := bufio.NewReader(conn)

	// 1 提示用户输入名字
	conn.Write([]byte("Enter your name: "))

	name, _ := reader.ReadString('\n') // 读取用户输入的名字
	name = strings.TrimSpace(name)     // 去掉换行符等多余字符

	// 2 把这个客户端添加到全局 clients 列表中
	clientsMu.Lock()
	clients[conn] = name
	clientsMu.Unlock()

	// 3 通知所有人某人加入聊天室
	broadcastMessage(fmt.Sprintf("%s has joined the chat\n", name), conn)

	// 4 不断读取该客户端发来的消息
	for {
		message, err := reader.ReadString('\n') // 读一行输入
		if err != nil {                         // 如果出错（例如客户端断开）
			break // 结束循环
		}

		// 转发消息给所有人（除了自己）
		broadcastMessage(fmt.Sprintf("%s: %s", name, message), conn)
	}

	// 5 客户端断开连接后，移出 clients
	clientsMu.Lock()
	delete(clients, conn)
	clientsMu.Unlock()

	// 6 广播“某人离开聊天室”
	broadcastMessage(fmt.Sprintf("%s has left the chat\n", name), conn)
}

func main() {
	// 1 启动 TCP 服务器，监听 8080 端口
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Println("Error starting server:", err)
		os.Exit(1)
	}
	defer listener.Close() // 程序退出时关闭监听

	fmt.Println("Chat server is listening on port 8080...")

	// 2 不断接受新的客户端连接
	for {
		conn, err := listener.Accept() // 等待客户端连接
		if err != nil {
			fmt.Println("Error accepting connection:", err)
			continue // 出错就跳过
		}

		// 3 每来一个客户端，就开一个 goroutine 独立处理
		go handleConnection(conn)
	}
}
