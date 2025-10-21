package main

import "fmt"

func both(ch chan int) {}

func sendOnlyFunc(ch chan<- int) {}

func main() {
	/*
		方向转换规则（只允许降级）
		可以把一个 双向通道 传给单向通道参数（安全）
		但不能把单向管道转换为双向
		chan T      →  chan<- T      允许
		chan T      →  <-chan T      允许
		chan<- T    →  chan T        不允许
		<-chan T    →  chan T        不允许
	*/
	ch := make(chan int) // 双向通道

	var sendOnly chan<- int = ch
	var receiveOnly <-chan int = ch
	fmt.Println(sendOnly)
	fmt.Println(receiveOnly)
	//ch = sendOnly // 编译错误：cannot use sendOnly (type chan<- int) as type chan int
	//ch = receiveOnly // 编译错误：cannot use receiveOnly (type <-chan int) as type chan int
	//both(sendOnly)	// 编译错误：cannot use sendOnly (type chan<- int) as type chan int
}
