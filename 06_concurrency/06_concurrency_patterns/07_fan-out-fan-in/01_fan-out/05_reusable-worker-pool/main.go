package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

type Task int
type Result struct {
	WorkerID int
	In       Task
	Out      int
}

// main 传入 ctx、并发数、输入通道，返回输出通道
func main() {
	/*
		最终版本
			Fan-Out 本质：多个 goroutine 共享一个输入通道读取
			收尾原则：谁生产谁关闭；启动 worker 的一侧负责在 wg.Wait() 后关闭 out
			可取消：所有阻塞点监听 ctx.Done()
			缓冲：out 适度缓冲（64/128 起测）
	*/
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*2)
	defer cancel() // 保底

	// 生产者
	in := make(chan Task, 100)
	go func() {
		defer close(in)
		for i := 1; i <= 100; i++ {
			select {
			case in <- Task(i):
			case <-ctx.Done():
				fmt.Println("producer get context cancel signal")
				return
			}
		}
	}()
	out := startWorkers(ctx, 10, in)
	for r := range out {
		fmt.Printf("worker=%d task=%d result=%d\n", r.WorkerID, r.In, r.Out)
	}
}

func startWorkers(ctx context.Context, n int, in <-chan Task) <-chan Result {
	// 把 out 和 wg 定义在 workers 中
	out := make(chan Result, 128)
	var wg sync.WaitGroup
	wg.Add(n)
	for id := 1; id <= n; id++ {
		go func(id int) {
			defer wg.Done()
			for {
				select {
				case <-ctx.Done():
					fmt.Println("worker", id, "get context cancel signal")
					return
				case t, ok := <-in:
					if !ok {
						fmt.Println("worker", id, "have no new task")
						return
					}
					time.Sleep(10 * time.Millisecond) // 模拟任务耗时
					res := Result{
						WorkerID: id,
						In:       t,
						Out:      int(t) * int(t),
					}
					select {
					case out <- res:
					case <-ctx.Done():
						fmt.Println("worker", id, "get context cancel signal")
						return
					}
				}
			}
		}(id)
	}
	// 谁管理谁关闭原则
	go func() {
		wg.Wait()
		close(out)
	}()
	return out
}
