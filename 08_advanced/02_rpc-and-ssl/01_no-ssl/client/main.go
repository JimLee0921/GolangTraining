package main

import (
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
	// 1. 连接服务器
	client, _ := rpc.DialHTTP("tcp", "localhost:1234")
	// 2. 创建参数和结果存储地址
	args := Args{
		X: 6,
		Y: 7,
	}
	var reply Result
	// 3. 直接使用 Call 调用
	err := client.Call("Cal.Multiply", args, &reply)
	if err != nil {
		panic(err)
	}
	// 4. 查询结果
	fmt.Println("result: ", reply.Product) // result:  42
}
