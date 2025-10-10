// 06_fanout_fanin_with_source_tag.go
package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

type Task int

// Result 每条结果里附带哪个 worker 处理的
type Result struct {
	FromWorker int
	Value      int
}

// Item fan-in 时再额外附带“来自哪个结果通道/池子”
type Item[T any] struct {
	Src int // pool/source index
	Val T
}

// ===== fan-out：启动 n 个 worker 共同从 in 读取，处理后写出到各自池子的 out =====
func startWorkers(ctx context.Context, n int, in <-chan Task) <-chan Result {
	out := make(chan Result, 128)
	var wg sync.WaitGroup
	wg.Add(n)

	for wid := 1; wid <= n; wid++ {
		wid := wid
		go func() {
			defer wg.Done()
			for {
				select {
				case <-ctx.Done():
					return
				case t, ok := <-in:
					if !ok {
						return
					}
					// 模拟处理：平方 + 延迟，展示并发无序
					time.Sleep(60 * time.Millisecond)
					res := Result{FromWorker: wid, Value: int(t) * int(t)}

					select {
					case out <- res:
					case <-ctx.Done():
						return
					}
				}
			}
		}()
	}

	// 统一收尾：所有 worker 退出后关闭 out
	go func() {
		wg.Wait()
		close(out)
	}()
	return out
}

// ===== fan-in：带来源标签（Src）把多路结果通道合并到一个 =====
func fanInWithSource[T any](ctx context.Context, ins ...<-chan T) <-chan Item[T] {
	out := make(chan Item[T], 256)
	var wg sync.WaitGroup
	wg.Add(len(ins))

	for i, ch := range ins {
		i, ch := i, ch
		go func() {
			defer wg.Done()
			for {
				select {
				case <-ctx.Done():
					return
				case v, ok := <-ch:
					if !ok {
						return
					}
					select {
					case out <- Item[T]{Src: i, Val: v}:
					case <-ctx.Done():
						return
					}
				}
			}
		}()
	}

	go func() {
		wg.Wait()
		close(out)
	}()
	return out
}

// ===== 简单的任务生产者 =====
func produce(ctx context.Context, n int) <-chan Task {
	ch := make(chan Task, 128)
	go func() {
		defer close(ch)
		for i := 0; i < n; i++ {
			select {
			case <-ctx.Done():
				return
			case ch <- Task(i):
			}
		}
	}()
	return ch
}

func main() {
	// 整体超时，防止泄露
	ctx, cancel := context.WithTimeout(context.Background(), 1200*time.Millisecond)
	defer cancel()

	// 任务源
	in := produce(ctx, 30)

	// 两个 fan-out 池子：配置不同并发度，代表两条并行处理路径
	poolA := startWorkers(ctx, 2, in) // Src = 0
	poolB := startWorkers(ctx, 3, in) // Src = 1

	// fan-in：合并两路结果，并携带来源标签（0/1）
	merged := fanInWithSource[Result](ctx, poolA, poolB)

	// 下游消费（无序）：打印来源池子 & worker id & 值
	for it := range merged {
		fmt.Printf("pool=%d worker=%d value=%d\n", it.Src, it.Val.FromWorker, it.Val.Value)
	}
}
