package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

// main 使用 context 使得所有阻塞读/写都可被 ctx.Done() 打断，防止泄露
func main() {
	/*
		这里设置任务生产延时，ctx 的超时时间越长，生产数据越多，打印也就越多
	*/
	ctx, cancel := context.WithTimeout(context.Background(), 500*time.Millisecond)
	defer cancel()

	a := produce(ctx, "a", 0, 100, 50*time.Millisecond)
	b := produce(ctx, "b", 100, 100, 60*time.Millisecond)
	c := produce(ctx, "c", 1000, 100, 40*time.Millisecond)

	out := fanInCtx(ctx, a, b, c)
	for v := range out {
		fmt.Println("merged", v)
	}
}

// fanInCtx 合并管道并可取消
func fanInCtx[T any](ctx context.Context, ins ...<-chan T) <-chan T {
	out := make(chan T, 128)
	var wg sync.WaitGroup
	wg.Add(len(ins))

	for _, ch := range ins {
		go func(ch <-chan T) {
			defer wg.Done()
			// 使用 for 无限循环持续监听，不能使用 for range，因为它没办法同时监听 ctx.Done()
			for {
				select {
				case <-ctx.Done():
					fmt.Println("fanin get cancel signal")
					return
				case v, ok := <-ch:
					if !ok {
						return
					}
					select {
					case out <- v:
					case <-ctx.Done():
						fmt.Println("fanin get cancel signal")
						return
					}
				}
			}
		}(ch)
	}

	go func() {
		wg.Wait()
		close(out)
	}()
	return out
}

func produce(ctx context.Context, name string, base, n int, d time.Duration) <-chan int {
	ch := make(chan int)
	go func() {
		defer close(ch)
		for i := 0; i < n; i++ {
			/*
				第一段 select
					作用：实现可取消的 sleep
					行为：它会阻塞到二选一发生——要么 ctx.Done() 被关闭（取消/超时），要么 time.After(d) 的定时器到点
					不是“会一直阻塞”，而是最多阻塞 d 时长（除非很快被取消）
				第二段 select
					作用：实现可取消的发送
					背景：ch 若是无缓冲或缓冲已满，ch <- x 会阻塞等待消费者读取
					如果这时 ctx 被取消，没有这层 select，goroutine 可能会永远卡在发送上（泄漏）。用这层 select 就能在取消时不被卡住，直接退出
				第一段保证等待阶段可取消，第二段保证发送阶段可取消
			*/
			select {
			case <-ctx.Done():
				fmt.Println("producer", name, "get cancel signal")
				return
			case <-time.After(d): // 模拟生产耗时
			}
			select {
			case ch <- base + i: // 把数据发出去
			case <-ctx.Done(): // 发送过程中如果被取消，立刻退出，不要卡死
				fmt.Println("producer", name, "get cancel signal")
				return
			}
		}
	}()
	return ch
}
