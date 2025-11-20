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
	client, _ := rpc.DialHTTPPath("tcp", "localhost:1234", "/rpc")
	// 2. 定义参数和结果
	args := Args{Num: 10}
	var reply Result
	// 3. 这里创建一个 channel，用于接收异步返回结果
	done := make(chan *rpc.Call, 1)

	// Go() 立即返回，不阻塞
	call := client.Go("Cal.Square", args, &reply, done)

	// 4. 在这里等待异步结果（阻塞等待），也可以 select 做超时控制
	call = <-done

	// call.Error 表示 RPC 是否错误
	if call.Error != nil {
		fmt.Println("RPC Error:", call.Error)
		return
	}

	fmt.Printf("result: Num=%d, Ans=%d\n", reply.Num, reply.Ans) // result: Num=10, Ans=100
}
