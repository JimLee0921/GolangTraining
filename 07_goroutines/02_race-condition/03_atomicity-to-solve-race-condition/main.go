package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

var (
	balance int64 = 1000 // 使用 int64 以便原子操作
	wg      sync.WaitGroup
)

func main() {
	/*
		atomic 解决资源竞争问题
		atomic = 原子操作，atomic 提供了 轻量级、锁无关 的并发安全手段
			原子操作就是 不可分割的最小操作单元，要么全部完成，要么完全不做
			在并发环境下，原子操作不会被打断，因此可以避免资源竞争
			Go 里，sync/atomic 包提供了一些常用的原子操作函数，用于在多个 goroutine 并发访问共享变量时保证安全
		以 int32 / int64 为例：
			读取 & 写入
				atomic.LoadInt64(&x)   // 原子读取
				atomic.StoreInt64(&x, 100) // 原子写入

			加减
				atomic.AddInt64(&x, 1)   // 原子地加 1
				atomic.AddInt64(&x, -1)  // 原子地减 1

			交换
				atomic.SwapInt64(&x, 200)  // 把 x 设置为 200，并返回旧值

			CAS（Compare And Swap）
				atomic.CompareAndSwapInt64(&x, old, new)
				如果 x 当前的值等于 old，就把它改成 new，并返回 true；
				否则返回 false（说明有别的 goroutine 抢先修改过了）
				CAS 是实现很多并发安全算法的核心

	*/
	wg.Add(3)
	go withdraw("JimLee", 600)
	go withdraw("BruceLee", 500)
	go withdraw("FrankStan", 300)
	wg.Wait()

	fmt.Printf("最终余额为: %v", balance)
}

func withdraw(name string, amount int64) {
	defer wg.Done()

	for {
		old := atomic.LoadInt64(&balance) // 原子读取
		if old < amount {
			fmt.Printf("%s 读取失败，余额不足（尝试取: %d，余额: %d）", name, amount, old)
			return
		}
		newVal := old - amount
		// 只有当余额仍然是 old 时，才把它原子地改成 newVal
		if atomic.CompareAndSwapInt64(&balance, old, newVal) {
			// 模拟后续处理（非临界逻辑，不影响余额正确性）
			time.Sleep(10 * time.Millisecond)
			fmt.Printf("%s 取钱成功，取出: %d，剩余余额: %d\n", name, amount, newVal)
			return
		}
		// 前面逻辑没走，说明被其它资源修改，循环重试直到 return
	}
}

/*
每次打印顺序可能不同
	BruceLee 读取失败，余额不足（尝试取: 500，余额: 400）FrankStan 取钱成功，取出: 300，剩余余额: 100
	JimLee 取钱成功，取出: 600，剩余余额: 400
	最终余额为: 100
*/
