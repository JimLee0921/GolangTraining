package main

import (
	"fmt"
	"sync"
	"time"
)

var sharedRsc = false

var (
	wg   sync.WaitGroup
	lock sync.Mutex
	cond = sync.NewCond(&lock)
)

func main() {
	wg.Add(2)

	go func() {
		cond.L.Lock()
		for sharedRsc == false {
			fmt.Println("goroutine1 wait")
			cond.Wait()
		}
		fmt.Println("goroutine1", sharedRsc)
		cond.L.Unlock()
		wg.Done()
	}()

	go func() {
		cond.L.Lock()
		for sharedRsc == false {
			fmt.Println("goroutine2 wait")
			cond.Wait()
		}
		fmt.Println("goroutine2", sharedRsc)
		cond.L.Unlock()
		wg.Done()
	}()

	time.Sleep(2 * time.Second)
	cond.L.Lock()
	fmt.Println("main goroutine ready")
	sharedRsc = true
	cond.Broadcast()
	fmt.Println("main goroutine broadcast")
	cond.L.Unlock()
	wg.Wait()
}

/*

goroutine1 wait
goroutine2 wait
main goroutine ready
main goroutine broadcast
goroutine2 true
goroutine1 true

goroutine1和goroutine2进入Wait状态
在main goroutine在2s后资源满足，发出broadcast信号后
从Wait中恢复并判断条件是否确实已经满足(sharedRsc不为空)
满足则消费条件，并解锁和wg.Done()

https://ieevee.com/tech/2019/06/15/cond.html
*/
