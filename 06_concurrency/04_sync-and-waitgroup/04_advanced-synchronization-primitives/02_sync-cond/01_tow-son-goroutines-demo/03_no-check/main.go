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
		fmt.Println("goroutine2 wait")
		cond.Wait()
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
goroutine2 wait

在子 goroutine2 中不做校验且主 main goroutine 没有延时会直接导致死锁
main goroutine（goroutine 1）先执行，并停留在 wg.Wait()中，等待子goroutine的wg.Done()
而子goroutine（goroutine 6）没有判断条件直接调用了cond.Wait。

我们知道cond.Wait会释放锁并等待其他goroutine调用Broadcast或者Signal来通知其恢复执行，除此之外没有其他的恢复途径
但此时main goroutine已经调用了Broadcast并进入了等待状态，没有任何goroutine会去拯救还在cond.Wait中的子goroutine了，而该子goroutine也没有机会调用wg.Done()去恢复main goroutine，造成了死锁


因此，一定要注意，Broadcast必须要在所有的Wait之后（当然了，可以通过条件判断来决定要不要进Wait）
*/
