package main

import (
	"encoding/json"
	"fmt"
	"geerpc"
	"geerpc/codec"
	"log"
	"net"
	"time"
)

func startServer(addr chan string) {
	// address 设置为 :0 表示监听一个随机端口
	l, err := net.Listen("tcp", ":0")
	if err != nil {
		log.Fatal("network error: ", err)
	}
	log.Println("start rpc server on: ", l.Addr())
	addr <- l.Addr().String() // l.Addr().String() 写进 addr 这个 chan string 里，这样 main 那边就能拿到服务端真实监听的地址
	geerpc.Accept(l)          // 真正开始处理服务
}

func main() {
	// 启动服务端
	addr := make(chan string)
	go startServer(addr)

	// 1. 启动客户端拨号连接
	conn, _ := net.Dial("tcp", <-addr)
	defer func() { _ = conn.Close() }() // 确保连接关闭
	time.Sleep(time.Second)

	// 2. 发送 options(协议协商)
	_ = json.NewEncoder(conn).Encode(geerpc.DefaultOption)

	// 3. 基于 conn 创建 GobCodec
	cc := codec.NewGobCodec(conn)

	// 4. 循环发 5 个请求 + 收 5 个响应
	for i := 0; i < 5; i++ {
		// 构造 Header
		h := &codec.Header{
			ServiceMethod: "Foo.Sum",
			Seq:           uint64(i),
		}
		// 真正发送请求 Body
		_ = cc.Write(h, fmt.Sprintf("geerpc req %d", h.Seq))
		_ = cc.ReadHeader(h)
		var reply string
		_ = cc.ReadBody(&reply)
		log.Println("reply:", reply)
	}
}
