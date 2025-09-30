package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	balance = 1000
	wg      sync.WaitGroup
	mu      sync.Mutex
)

func main() {
	/*
		加锁解决资源竞争问题
		mu := sync.Mutex
			互斥量 (mutual exclusion)保证同一时刻只有一个 goroutine 能进入临界区（共享资源的访问代码块）
			mu.Lock()   // 上锁
			// 临界区：安全地访问或修改共享变量
			mu.Unlock() // 解锁
		注意事项：
			1. 成对出现：Lock() 和 Unlock() 必须配对，否则容易死锁
			2. 锁的粒度：锁的范围不要太大，否则会降低并发度
				粒度太小：保护不住数据，还是有竞争
				粒度太大：整个程序都被串行化，失去并发意义
			3. 避免重复上锁：同一个 goroutine 如果在解锁前再次 Lock()，会死锁

	*/
	wg.Add(3)
	go withdraw("JimLee", 600)
	go withdraw("BruceLee", 500)
	go withdraw("FrankStan", 300)
	wg.Wait()

	fmt.Printf("最终余额为: %v", balance)
}

func withdraw(name string, amount int) {
	defer wg.Done()

	// 进入临界区: 读余额 -> 判断 -> 扣款 -> 写回
	mu.Lock()
	defer mu.Unlock()
	if balance >= amount {
		fmt.Printf("%s 正在取钱: %d，当前余额: %d\n", name, amount, balance)
		time.Sleep(10 * time.Millisecond) // 模拟处理耗时
		balance -= amount
		fmt.Printf("%s 取钱成功，取出: %d，剩余余额: %d\n", name, amount, balance)
	} else {
		fmt.Printf("%s 取钱失败，余额不足（尝试取: %d，余额: %d）\n", name, amount, balance)
	}

}

/*
输出结果：
FrankStan 正在取钱: 300，当前余额: 1000
FrankStan 取钱成功，取出: 300，剩余余额: 700
BruceLee 正在取钱: 500，当前余额: 700
BruceLee 取钱成功，取出: 500，剩余余额: 200
JimLee 取钱失败，余额不足（尝试取: 600，余额: 200）
最终余额为: 200

PS C:\demo\GolangTraining> go run -race .\07_goroutines\02_race-condition\02_mutex-to-solve-race-condition\main.go
FrankStan 正在取钱: 300，当前余额: 1000
FrankStan 取钱成功，取出: 300，剩余余额: 700
BruceLee 正在取钱: 500，当前余额: 700
BruceLee 取钱成功，取出: 500，剩余余额: 200
JimLee 取钱失败，余额不足（尝试取: 600，余额: 200）
最终余额为: 200v
*/
