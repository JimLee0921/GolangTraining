package main

import (
	"context"
	"fmt"
	"time"
)

func worker(ctx context.Context) {
	select {
	case <-ctx.Done():
		fmt.Println("Worker: Context has canceled")
	case <-time.After(5 * time.Second):
		fmt.Println("Worker: work successful")
	}
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	go worker(ctx)

	fmt.Println("Main: waiting 5 seconds then send cancel Context")
	time.Sleep(3 * time.Second)
	cancel()
	fmt.Println("Main: Context has canceled")
}
