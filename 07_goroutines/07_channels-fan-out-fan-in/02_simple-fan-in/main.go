package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	/*

	 */
	c := fanIn(boring("JimLee"), boring("JamesBond"))
	for i := 0; i < 10; i++ {
		fmt.Println(<-c)
	}
	fmt.Println("You're both boring; I'm leaving.")
}

func boring(msg string) <-chan string {
	// 每个 boring 函数启动一个 goroutine
	c := make(chan string)
	go func() {
		for i := 0; ; i++ {
			c <- fmt.Sprintf("%s %d", msg, i)
			time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
		}
	}()
	return c
}

func fanIn(input1, input2 <-chan string) <-chan string {
	// 把 多个输入 channel 合并成一个输出 channel
	out := make(chan string)
	go func() {
		for i := range input1 {
			out <- i
		}
	}()
	go func() {
		for i := range input2 {
			out <- i
		}
	}()
	return out
}
