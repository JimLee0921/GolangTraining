package main

import "fmt"

// main 函数参数和返回值使用只读 channel(<-chan)
func main() {
	c := incrementor()
	for n := range puller(c) {
		fmt.Println(n)
	}

}

func incrementor() <-chan int {
	/*
		返回一个 只读 channel <-chan int，调用方只能读，不能写
		goroutine 会往 out 写入 0..9 共 10 个数，然后关闭
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

func puller(c <-chan int) <-chan int {
	/*
		参数 c <-chan int：函数只能从 channel 里读，不能写
		返回值 <-chan int：调用者只能从 out 读，不能写
		把总和写到一个新的 channel out 里并返回
	*/
	out := make(chan int)
	go func() {
		sum := 0
		for n := range c {
			sum += n
		}
		out <- sum
		close(out)
	}()
	return out
}
