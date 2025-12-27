package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

type Service struct {
	running atomic.Bool
}

func (s *Service) Start() {
	// 只允许从 false -> true
	if !s.running.CompareAndSwap(false, true) {
		fmt.Println("Start(): service already running")
		return
	}

	fmt.Println("Start(): service started")
	go s.loop()
}

func (s *Service) loop() {
	for {
		if !s.running.Load() {
			fmt.Println("loop(): service stopped")
			return
		}

		fmt.Println("loop(): working...")
		time.Sleep(500 * time.Millisecond)
	}
}

func (s *Service) Stop() {
	old := s.running.Swap(false)
	if old {
		fmt.Println("Stop(): service stopped")
	} else {
		fmt.Println("Stop(): service was not running")
	}
}

func main() {
	var svc Service
	var wg sync.WaitGroup

	// 并发启动
	for i := 0; i < 3; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			fmt.Printf("goroutine %d calling Start()\n", id)
			svc.Start()
		}(i)
	}
	wg.Wait()
	time.Sleep(2 * time.Second)
	svc.Stop()
	time.Sleep(1 * time.Second)
}
