package main

import (
	"net"
	"net/http"
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
	rpc.HandleHTTP()
	// 2. 监听端口
	l, _ := net.Listen("tcp", ":1234")
	// 3. 开启HTTP后需要使用 http.Serve 改为 HTTP 模式
	err = http.Serve(l, nil)
	if err != nil {
		panic(err)
	}
}
