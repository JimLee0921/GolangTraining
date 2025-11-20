package main

import (
	"net"
	"net/rpc"
)

// Cal 不是用来存数据的，而是作为RPC 服务的名字 + 方法命名空间（Method Set）容器
/*
某些示例使用的是 type Cal int 跟 struct{} 没什么区别
但是在使用 Register 注册时如果是 struct{} 可以使用 new(Cal) 或者 &Cal{}
这种写法只适用于 struct 而不适用于基础类型
但是如果是 int 这种类型只能使用 new(Cal)
*/
type Cal struct{}

// Args 把参数 Num 传递进去
type Args struct {
	Num int
}

// Result 存储结果
type Result struct {
	Num, Ans int
}

// Square 需要修改为符合 rpc 规则的方法 func (t *T) MethodName(arg T1, reply *T2) error
func (cal *Cal) Square(args Args, reply *Result) error {
	// 把结果存入 reply
	reply.Num = args.Num
	reply.Ans = args.Num * args.Num
	// 返回空
	return nil
}

func main() {
	// 使用 ServeConn 进行手动 Accept 连接
	// 1. 进行监听注册
	l, _ := net.Listen("tcp", ":1234")
	err := rpc.Register(&Cal{})
	if err != nil {
		panic(err)
	}
	// 2. 下面几行加起来等同于 rpc.Accept(l) 也就是并行处理
	for {
		conn, _ := l.Accept()
		go rpc.ServeConn(conn) // 每条连接独立处理
	}
}
