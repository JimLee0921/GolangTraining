package main

import (
	"net/http"
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
	// 1. 注册服务（使用 RegisterName 自定义服务名）
	err := rpc.RegisterName("Calculator", &Cal{})
	if err != nil {
		panic(err)
	}
	// 2. 开启 HTTP 连接，注册到 DefaultRPCPath: "/_goRPC_"
	rpc.HandleHTTP()
	// 3. 开启监听(改成 HTTP 方式，就不能再用 rpc.Accept(l) 或 server.Accept(l)。因为 Accept() 处理的是 纯 TCP 连接，而 HTTP 方式是要让连接先走 HTTP 协议握手 再升级为 RPC。)
	//l, _ := net.Listen("tcp", ":1234")
	//err = http.Serve(l, nil)
	//if err != nil {
	//	panic(err)
	//}

	// 上面可以简化为 ListenAndServe
	if err = http.ListenAndServe(":1234", nil); err != nil {
		panic(err)
	}
}
