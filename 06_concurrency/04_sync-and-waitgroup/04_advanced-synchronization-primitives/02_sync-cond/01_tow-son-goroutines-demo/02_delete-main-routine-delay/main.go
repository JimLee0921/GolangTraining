package main

import (
	"fmt"
	"sync"
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

	cond.L.Lock()
	fmt.Println("main goroutine ready")
	sharedRsc = true
	cond.Broadcast()
	fmt.Println("main goroutine broadcast")
	cond.L.Unlock()
	wg.Wait()
}

/*

main goroutine ready
main goroutine broadcast
goroutine1 true
goroutine2 true


删除 main goroutine 中的延时后
两个goroutine都没有进入Wait状态
原因是，main goroutine执行的更快，在goroutine1/goroutine2加锁之前就已经获得了锁，并完成了修改sharedRsc、发出Broadcast信号
当子goroutine调用Wait之前检验condition时，条件已经满足，因此就没有必要再去调用Wait了
*/
