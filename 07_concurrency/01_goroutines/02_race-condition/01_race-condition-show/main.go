package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup
var balance = 1000 // 银行账户余额

func main() {
	/*
		race condition（资源竞争）
	*/
	// 展示资源竞争
	wg.Add(2)
	go withdraw("JimLee", 800)
	go withdraw("BruceLee", 500)
	wg.Wait()
	fmt.Printf("最终余额为: %d", balance)

}

func withdraw(name string, amount int) {
	defer wg.Done()
	if balance >= amount {
		fmt.Printf("%s 正在取钱: %d. 当前余额: %d\n", name, amount, balance)
		time.Sleep(10 * time.Millisecond) // 模拟取钱耗时
		balance -= amount
		fmt.Printf("%s 取钱成功，取出: %d，剩余余额: %d\n", name, amount, balance)
	} else {
		fmt.Printf("%s 取钱失败，余额不足（尝试取: %d，余额: %d）\n", name, amount, balance)
	}

}

/*
最终输出:
	BruceLee 正在取钱: 500. 当前余额: 1000
	JimLee 正在取钱: 800. 当前余额: 1000
	JimLee 取钱成功，取出: 800，剩余余额: 200
	BruceLee 取钱成功，取出: 500，剩余余额: -300
	最终余额为: -300
两个 goroutine 同时读取 balance，都以为够取，结果都扣钱，最终余额出错，这就是典型的 资源竞争


race 检测输出
	go run -race main.go
*/
