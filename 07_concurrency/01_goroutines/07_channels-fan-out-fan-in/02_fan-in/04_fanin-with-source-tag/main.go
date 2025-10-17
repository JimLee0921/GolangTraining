package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

type Item[T any] struct {
	Src int // 源编号
	Val T
}

// main 合并后还可以知道结果来自哪个输入（source id）
func main() {
	/*
		合并后带上 Src，方便日志/统计/排错
		泛型 Item[T] 不影响性能（编译期具体化）
		把 ctx 超时时间设置更短或给 produce 加上更长的超时可以看到 ctx 效果
	*/
	ctx, cancel := context.WithTimeout(context.Background(), 200*time.Millisecond)
	defer cancel()

	a := produce(ctx, "a", 0, 10, 50*time.Millisecond)
	b := produce(ctx, "b", 100, 10, 80*time.Millisecond)
	c := produce(ctx, "c", 1000, 10, 20*time.Millisecond)

	out := fanInWithSource(ctx, a, b, c)

	for it := range out {
		fmt.Printf("src=%d val=%d\n", it.Src, it.Val)
	}
}

func fanInWithSource[T any](ctx context.Context, ins ...<-chan T) <-chan Item[T] {
	out := make(chan Item[T], 128)

	var wg sync.WaitGroup
	wg.Add(len(ins))

	for i, ch := range ins {
		i, ch := i, ch // 重新定义变量
		go func() {
			defer wg.Done()
			for {
				select {
				case <-ctx.Done():
					fmt.Println("fanin", i, "get cancel signal")
					return
				case v, ok := <-ch:
					if !ok {
						fmt.Println("channel", i, "have no more data")
						return
					}
					select {
					case out <- Item[T]{Src: i, Val: v}:
					case <-ctx.Done():
						fmt.Println("fanin", i, "get cancel signal")
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

func produce(ctx context.Context, name string, base int, n int, d time.Duration) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for i := 0; i < n; i++ {
			select {
			case <-ctx.Done():
				fmt.Println(name, "get cancel signal")
				return
			case <-time.After(d):
			}
			select {
			case out <- base + i:
			case <-ctx.Done():
				fmt.Println(name, "get cancel signal")
				return
			}
		}
	}()
	return out
}
