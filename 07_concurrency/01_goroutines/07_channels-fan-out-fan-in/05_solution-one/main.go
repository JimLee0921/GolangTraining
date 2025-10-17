package main

import (
	"fmt"
	"sync/atomic"
	"time"
)

var workerID int64
var publisherID int64

func main() {
	/*
		添加 atomic 来修复竞争条件
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

// publisher 向一个管道中推入数据
func publisher(out chan string) {
	atomic.AddInt64(&publisherID, 1)
	thisID := atomic.LoadInt64(&publisherID)
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
	atomic.AddInt64(&workerID, 1)
	thisID := atomic.LoadInt64(&workerID)
	for {
		fmt.Printf("worker: %d: waiting for input...\n", thisID)
		input := <-in
		fmt.Printf("worker: %d: input is: %s\n", thisID, input)
	}
}
