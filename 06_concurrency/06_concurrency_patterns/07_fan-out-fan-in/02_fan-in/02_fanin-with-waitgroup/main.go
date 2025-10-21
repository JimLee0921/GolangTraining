package main

import (
	"fmt"
	"sync"
	"time"
)

// main 使用 WaitGroup 支持任意条输入通道；当所有输入都关闭时，再关闭输出
func main() {
	/*
		fanIn 自己负责在 wg.Wait() 后 close(out)
		消费端 range out 自然结束
		结果无序：谁先到先发
	*/
	a := produce(0, 5, 50*time.Millisecond)
	b := produce(100, 5, 80*time.Millisecond)
	c := produce(1000, 5, 90*time.Millisecond)

	out := fanIn(a, b, c)

	for v := range out {
		fmt.Println("result:", v)
	}
}

// fanIn 传入多个 <-chan T 并进行合并
func fanIn[T any](ins ...<-chan T) <-chan T {
	// out 和 wg 维护在 fanIn 内部并添加适度缓冲
	out := make(chan T, 128)
	var wg sync.WaitGroup
	wg.Add(len(ins))

	// 每个输入管道一个搬运工
	for _, ch := range ins {
		go func(ch <-chan T) {
			defer wg.Done()
			for v := range ch {
				out <- v
			}
		}(ch)
	}
	// 所有搬运工都完成后关闭 out （同样遵循谁维护谁关闭原则）
	go func() {
		wg.Wait()
		close(out)
	}()
	return out
}

// produce 生产者
func produce(base, n int, d time.Duration) <-chan int {
	ch := make(chan int)
	go func() {
		defer close(ch)
		for i := 0; i < n; i++ {
			time.Sleep(d)
			ch <- base + i
		}
	}()
	return ch
}
