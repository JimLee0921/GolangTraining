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
