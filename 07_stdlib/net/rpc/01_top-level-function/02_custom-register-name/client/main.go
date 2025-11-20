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
	// 3. rpc 调用，由于服务端使用 RegisterName 指定了服务名，这里也要保持服务名一致
	client.Call("Cal.Square", args, &reply)        // 服务名不一致，调用失败
	fmt.Println(reply)                             // {0 0}
	client.Call("Calculator.Square", args, &reply) // 调用成功
	fmt.Println(reply)                             // {10 100}

}
