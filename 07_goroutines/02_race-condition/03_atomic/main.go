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
