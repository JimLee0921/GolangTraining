package main

import "fmt"

// main channel 基本认识，更多使用在 07_goroutines/03_channel
func main() {
	/*
		channel 是一种特殊的数据类型，用来让 不同 goroutine 之间安全地传递数据
		可以把理解为一条管道：一端塞东西（发送），另一端取东西（接收）
		是一个先入先出(FIFO)的队列，接收的数据和发送的数据的顺序是一致的
		在传统并发编程里，多个线程访问同一个共享变量容易出现资源竞争（race condition）
		Go 提供了 channel，让 goroutine 之间 不用共享内存，而是 通过消息传递 来协作，替代手动加锁的一种更自然的并发工具
		Channel 的核心特征
			有类型且严格区分
				chan int -> 只能传 int
				chan string -> 只能传 string
				chan int 和 chan string 是完全不同的类型。
				chan int 和 chan<- int / <-chan int 也不同
				所以 channel 就像有类型的安全队列

			同步 / 异步行为
				无缓冲 channel：发送和接收必须同步配对，像握手
				有缓冲 channel：可以存一定数量的数据，发送不会立刻阻塞，允许发送方和接收方的速度不完全一致

			双向或单向
				默认：chan T -> 双向（既能收又能发）
				限制：chan<- T（只发）、<-chan T（只收）

			并发安全
				Go 的 channel 是线程安全的
				多个 goroutine 可以同时向同一个 channel 发数据，或者取数据，不需要额外加锁
		Channel 基本定义：
			var intChan chan int    // 声明一个传输 int 的 channel，
			var strChan chan string // 声明一个传输 string 的 channel
		Channel 创建初始化：
			make 会返回一个已初始化好的 channel，会分配底层数据结构，让 channel 可以正常收发
				make(chan T)：无缓冲（发送和接收必须同步配对）
				make(chan T, n)：有缓冲（最多能存放 n 个元素，满了才阻塞发送）
		Channel 类型定义：
			在函数参数或变量中，可以限制 channel 的方向
			var sendOnly chan<- int   // 只能发送 int
			var recvOnly <-chan int   // 只能接收 int
				chan<- T：只发送通道，单向通道，只能用来发送（<- 放在 chan 右侧），只能执行 ch <- v
				<-chan T：只接收通道，单向通道，只能用来接收（<- 放在 chan 左侧），只能执行 v := <-ch
				普通 chan T 是双向的，可以传给只写或只读，但反过来不行
		注意事项：
			var 声明的 channel 默认值是 nil
				nil channel 不能收发数据
				如果对 nil channel 执行发送或接收操作，会永久阻塞（deadlock）
				但是，读取 nil channel 本身是安全的（只要不去 <-ch 或 ch <- x 就没问题）

	*/
	// 1. 基本声明（零值是 nil）
	var intChan chan int    // 声明一个传输 int 的 channel，
	var strChan chan string // 声明一个传输 string 的 channel

	fmt.Printf("%v, %T\n", intChan, intChan) // <nil>, chan int
	fmt.Printf("%v, %T\n", strChan, strChan) // <nil>, chan string

	// 2. 使用 make 创建（最常用）
	unbufferedChan := make(chan int)          // 创建一个无缓冲区的 channel
	bufferedChan := make(chan int, 5)         // 创建一个容量为 5 的缓冲 channel
	fmt.Println(unbufferedChan, bufferedChan) // 默认打印的引用地址 0xc00001a0e0 0xc0000200a0

	// 3. channel 类型，在函数参数或变量中，可以限制 channel 的方向
	var sendOnly chan<- int
	var recvOnly <-chan int
	fmt.Println(sendOnly, recvOnly) // <nil> <nil>

}
