package main

import "fmt"

// main channel 作为函数参数和返回值传递
func main() {
	/*
		可选的 <- 运算符指定了通道的方向，是发送还是接收。如果没有给出方向，则通道是双向的
		https://go.dev/ref/spec#Channel_types
	*/
	c := incrementor()
	cSum := puller(c)
	for n := range cSum {
		fmt.Println(n)
	}
	// 可以直接 range puller(c)
	//for n := range puller(c) {
	//	fmt.Println(n)
	//}

}

func incrementor() chan int {
	/*
		生产者返回一个 channel 里面是 0-9 十个数字
	*/
	out := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			out <- i
		}
		close(out)
	}()
	return out
}

func puller(c chan int) (out chan int) {
	/*
		从输入 channel c 中读出所有数字计算总和
		把总和写到一个新的 channel out 里并返回
	*/
	out = make(chan int)
	go func() {
		sum := 0
		for n := range c {
			sum += n
		}
		out <- sum
		close(out)
	}()
	return
}
