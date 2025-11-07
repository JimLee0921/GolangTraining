package main

import (
	"crypto/tls"
	"log"
	"net/rpc"
)

// Cal 服务类型（命名空间）
type Cal struct{}

// Args 参数（客户端传过来的输入）
type Args struct {
	X, Y int
}

// Result 结果（服务端回写输出）
type Result struct {
	Product int
}

// Multiply RPC 方法（必须符合签名要求）
func (cal *Cal) Multiply(args Args, reply *Result) error {
	reply.Product = args.X * args.Y
	return nil
}

func main() {
	// 1. 注册服务，server（简单一些直接用默认）
	err := rpc.Register(&Cal{})
	if err != nil {
		panic(err)
	}
	// 2. 加载生成的整数
	cert, _ := tls.LoadX509KeyPair("08_advanced/02_rpc-and-ssl/02_client-skip-verify/server.crt", "08_advanced/02_rpc-and-ssl/02_client-skip-verify/server.key")
	config := &tls.Config{
		Certificates: []tls.Certificate{cert},
	}
	// 2. 使用 tls.Listen 监听端口
	l, _ := tls.Listen("tcp", ":1234", config)
	log.Printf("Serving RPC server on port %d", 1234)
	// 3. 接收监听
	rpc.Accept(l)
}
