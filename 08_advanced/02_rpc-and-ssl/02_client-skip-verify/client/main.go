package main

import (
	"crypto/tls"
	"fmt"
	"net/rpc"
)

type Args struct {
	X, Y int
}
type Result struct {
	Product int
}

func main() {
	// 1. 跳过验证
	config := &tls.Config{
		InsecureSkipVerify: true,
	}
	// 2. 使用 tls.Dial 进行连接服务器
	conn, _ := tls.Dial("tcp", "localhost:1234", config)
	defer conn.Close()
	client := rpc.NewClient(conn)
	// 2. 创建参数和结果存储地址
	args := Args{
		X: 6,
		Y: 7,
	}
	var reply Result
	// 3. Call 调用还是一样
	err := client.Call("Cal.Multiply", args, &reply)
	if err != nil {
		panic(err)
	}
	// 4. 查询结果
	fmt.Println("result: ", reply.Product) // result:  42
}
