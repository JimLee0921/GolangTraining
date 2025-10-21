package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	counter int
	rw      sync.RWMutex
	wg      sync.WaitGroup
)

func read(id int) {
	defer wg.Done()
	rw.RLock() // 加读锁
	defer rw.RUnlock()
	fmt.Printf("read %d: counter=%d\n", id, counter)
	time.Sleep(200 * time.Millisecond)
}

func write(id int) {
	defer wg.Done()
	rw.Lock() // 加写锁
	defer rw.Unlock()
	counter++
	fmt.Printf("write %d: counter=%d\n", id, counter)
	time.Sleep(300 * time.Millisecond)
}

func main() {
	// 多个读锁可以并存，但写锁会阻塞其它所有读写操作
	for i := 0; i < 3; i++ {
		wg.Add(1)
		go read(i)
	}
	for i := 0; i < 2; i++ {
		wg.Add(1)
		go write(i)
	}
	wg.Wait()
	fmt.Println("All goroutines done")
}
