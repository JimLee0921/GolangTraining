package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

type Starter struct {
	started atomic.Bool
}

// Start 只能把 false -> true 成功一次，并发安全
func (s *Starter) Start() bool {
	return s.started.CompareAndSwap(false, true)
}

func main() {
	var s Starter
	var wg sync.WaitGroup
	wg.Add(5)

	for i := 0; i < 5; i++ {
		go func(id int) {
			defer wg.Done()
			if s.Start() {
				fmt.Println("worker", id, "did the real start")
			} else {
				fmt.Println("worker", id, "saw already started")
			}
		}(i)
	}
	wg.Wait()
}
