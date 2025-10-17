package main

import "fmt"

type Ch1 chan int
type Ch2 chan int

func main() {
	/*
		通道类型包括：
			元素类型 + 通道方向（<-chan、chan<-、chan）
			相同底层类型可以转换，单向与双向转换（受限）
	*/
	c1 := make(Ch1)
	c2 := Ch2(c1) // 合法
	fmt.Println(c2)
	var c chan int = make(chan int)
	var send chan<- int = c // 双向 -> 单向
	var recv <-chan int = c // 双向 -> 单向
	// c = send // 反向不行
	fmt.Println(send, recv)

}
