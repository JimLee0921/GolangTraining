package main

import (
	"fmt"
	"net/rpc"
)

// Args 把参数 Num 传递进去
type Args struct {
	Num int
}

// Result 存储结果
type Result struct {
	Num, Ans int
}

func main() {
	// 1. 连接到服务
	client, _ := rpc.Dial("tcp", "localhost:1234")
	// 2. 定义参数和结果
	args := Args{Num: 10}
	var reply Result
	// 3. rpc 调用方法，默认就是 TypeName.Method ，结果存储必须传入指针，才能接收成功
	client.Call("Cal.Square", args, &reply)
	fmt.Println(reply) // {10 100}
}
