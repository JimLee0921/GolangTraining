package main

import (
	"fmt"
	"time"
)

var workerID int
var publisherID int

func main() {
	/*
		这段代码虽然是 fan-in fan-out 模型，但是会存在资源竞争问题
	*/
	input := make(chan string)
	go workerProcess(input)
	go workerProcess(input)
	go workerProcess(input)
	go publisher(input)
	go publisher(input)
	go publisher(input)
	go publisher(input)
	time.Sleep(1 * time.Millisecond) // 模拟延时操作（也不够安全）
}

// publisher 无限向一个管道中推入数据
func publisher(out chan string) {
	publisherID++
	thisID := publisherID
	dataID := 0
	for {
		dataID++
		fmt.Printf("publisher: %d is pushing data\n", thisID)
		data := fmt.Sprintf("Data from publisher %d. Data %d", thisID, dataID)
		out <- data
	}
}

// workerProcess 从管道中消费数据
func workerProcess(in <-chan string) {
	workerID++
	thisID := workerID
	for {
		fmt.Printf("worker: %d: waiting for input...\n", thisID)
		input := <-in
		fmt.Printf("worker: %d: input is: %s\n", thisID, input)
	}
}

/*
1. 全局计数器没有同步
	var workerID int
	var publisherID int
	…
	publisherID++   // 竞争
	workerID++      // 竞争
	多个 goroutine 同时读写 publisherID / workerID，未做任何同步，触发数据竞争
	竞争表现：ID 可能重复、跳号，在 -race 下直接报 data race

2. main 用 time.Sleep 结束，goroutine 强制终止

3. 多生产者 / 多消费者对同一 channel
	多个 publisher 并发向同一无缓冲 channel 写入、多个 worker 并发读取，这本身是安全的（channel 是并发安全的），不会因并发读写 channel 本身产生 data race
	但可能出现阻塞（背压）：当没有匹配的接收者时发送会阻塞，反之亦然——属于调度/吞吐问题，不是数据竞争
运行时加竞态检测：
	go run -race main.go
	go build -race && ./main
*/
