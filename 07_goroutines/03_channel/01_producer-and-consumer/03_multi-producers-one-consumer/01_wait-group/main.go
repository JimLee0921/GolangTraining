package main

import (
	"fmt"
	"sync"
)

// main 每次运行结果顺序可能是不固定的
func main() {
	/*
		多生产者 (N) -> 单消费者 (1) 并发模式，也就是 fan-in
		如何在多个生产者结束之后，安全地关闭共享 channel，让消费者能优雅退出
			1. WaitGroup
				WaitGroup 用来统计还剩多少生产者没完成
				wg.Wait() 会阻塞，直到计数归零
				然后统一关闭 channel
			2. done channel
				新创建一个 done channel 设置为 make(chan bool) 每个生产者完成后往 done 里发一个信号
				协调者 goroutine 用两次 <-done 接收，确保两个生产者都结束
				然后关闭 channel
		对比：
			特点			WaitGroup								done channel
			语义			计数器（还有多少任务没完成）					信号（谁完成了就发一个）
			用法			Add(n) / Done() / Wait()				done <- true / <-done
			多生产者情况	一次 Add(n)，最后由一个 goroutine Wait		每个生产者发送一个信号，协调者按数量接收
			风险点		Add 必须在 Wait 前完成，否则有竞态			必须匹配好发送/接收次数，否则阻塞
			场景			更通用，适合任务数较多、动态分配				更直观，适合小规模固定任务数
	*/
	ch := make(chan int)
	var wg sync.WaitGroup
	// 等待 2 个生产者 goroutine 完成
	wg.Add(2)

	// 两个生产者 goroutine 向 ch 里每个循环发送 0..9 共 10 个数，总共会有 20 个整数被发送进 channel
	go func() {
		defer wg.Done()
		for i := 0; i < 10; i++ {
			ch <- i
		}
	}()

	go func() {
		defer wg.Done()
		for i := 0; i < 10; i++ {
			ch <- i
		}
	}()
	// 启动一个“协调者”协程，等待所有生产者完成 (wg.Wait())，调用 close(ch)，告诉消费者 不会再有新数据了
	// 这样主 goroutine 里的 for range ch 才能优雅退出，而不会无限阻塞
	go func() {
		wg.Wait()
		close(ch)
	}()

	// 单消费者
	for n := range ch {
		fmt.Println(n)
	}
}
