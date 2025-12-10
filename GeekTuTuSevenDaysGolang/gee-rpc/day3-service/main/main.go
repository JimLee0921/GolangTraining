package main

import (
	"geerpc"
	"log"
	"net"
	"sync"
)

type Foo int
type Args struct {
	Num1, Num2 int
}

func (f Foo) Sum(args Args, reply *int) error {
	*reply = args.Num1 + args.Num2
	return nil
}

func startServer(addr chan string) {
	// 注册服务
	var foo Foo
	if err := geerpc.Register(&foo); err != nil {
		log.Fatal("register error: ", err)
	}
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
	log.SetFlags(0)
	// 使用 channel 创建 addr 可以确保真正等服务端监听好了客户端才开始 Dial 连接
	addr := make(chan string)
	go startServer(addr)
	// 使用 Dial 创建 client
	client, _ := geerpc.Dial("tcp", <-addr)
	defer func() { _ = client.Close() }()
	// 并发 goroutine
	var wg sync.WaitGroup
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			args := &Args{
				Num1: i,
				Num2: i * i,
			}
			var reply int
			if err := client.Call("Foo.Sum", args, &reply); err != nil {
				log.Fatal("call Foo.Sum error", err)
			}
			log.Printf("%d + %d = %d", i, i*i, reply)
		}(i)
	}
	wg.Wait()
}
