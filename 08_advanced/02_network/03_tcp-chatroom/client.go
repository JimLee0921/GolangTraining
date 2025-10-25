package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
	// 1. 连接服务器
	conn, err := net.Dial("tcp", "127.0.0.1:8080")
	if err != nil {
		fmt.Println("error connecting to server:", err)
		os.Exit(1)
	}
	defer conn.Close()

	go func() {
		// 创建一个带缓冲的 Reader，用来按行读取服务器的消息
		reader := bufio.NewReader(conn)
		for {
			// 一直从服务器读取，直到遇到换行符 '\n'
			message, err := reader.ReadString('\n')
			if err != nil { // 如果出错（例如服务器断开连接），就退出循环
				break
			}
			// 打印服务器发来的消息（如：别人发送的聊天内容）
			fmt.Print(message)
		}
	}()

	// 先读一次服务器发来的“Enter your name”
	reader := bufio.NewReader(conn)
	prompt, _ := reader.ReadString(':') // 或 '\n'，视服务器消息结尾
	fmt.Print(prompt)
	// 创建一个 Scanner，从标准输入读取用户的输入
	scanner := bufio.NewScanner(os.Stdin)
	// 无限循环：不断读取用户输入
	for scanner.Scan() {
		conn.Write([]byte(scanner.Text() + "\n"))
	}
}
