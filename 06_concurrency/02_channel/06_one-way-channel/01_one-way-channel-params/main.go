package main

import "fmt"

// 发送者，只能写入管道
func sendData(ch chan<- int) {
	defer close(ch) // 发送完数据发送关闭信号
	for i := 1; i <= 5; i++ {
		ch <- i
		fmt.Println("send:", i)
	}
}

// 接收者，只能读取管道
func receiveData(ch <-chan int) {
	for v := range ch {
		fmt.Println("received: ", v)
	}
}

func main() {
	/*
		创建管道还是正常双向管道，由函数参数设置进行限制单向行为
	*/
	ch := make(chan int, 3)
	go sendData(ch) // 只能发送
	receiveData(ch) // 只能接收
}
