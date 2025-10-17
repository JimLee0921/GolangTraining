package main

import (
	"fmt"
	"time"
)

func main() {
	/*
		无缓冲 channel + goroutine 通信
		潜在问题
			channel 没有关闭：消费者 goroutine 会在生产者发送完 10 个数后继续死等（阻塞）
			不优雅退出：这里只是靠 time.Sleep 粗暴地让程序撑一会儿，然后主 goroutine 结束 -> 整个程序退出
	*/
	// 1. 创建一个无缓冲 channel，在 goroutine 之间传递 int
	ch := make(chan int)

	// 2. 启动一个 goroutine，当“生产者”
	// 会依次往 channel c 里发送 0 到 9
	// 因为是无缓冲 channel，所以每次 c <- i 都会阻塞，直到有接收方 <-c 把值取走
	go func() {
		for i := 0; i < 10; i++ {
			ch <- i
		}
	}()

	// 3. 启动另一个 goroutine，当消费者
	// 会无限循环，从 channel c 里取值并打印
	// 每次 <-c 阻塞，直到有发送方 c <- i 提供一个新值
	// 因为这个 for 没有退出条件，等生产者发完 10 个值后，它还会继续等着，试图接收新的值（但没人再发送了，就会阻塞）
	go func() {
		for {
			fmt.Println(<-ch)
		}
	}()
	// 主 goroutine 睡眠 1 秒，让上面两个 goroutine 有时间跑完
	// 如果不写这句，主函数会立刻退出，其他 goroutine 也会被杀掉
	time.Sleep(time.Second)
}

/*
输出结果：
0
1
2
3
4
5
6
7
8
9
*/
