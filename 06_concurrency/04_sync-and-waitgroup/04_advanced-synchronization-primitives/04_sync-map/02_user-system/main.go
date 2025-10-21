package main

import (
	"fmt"
	"sync"
	"time"
)

// 全局并发安全map
var users sync.Map

// 用户登录：Store
func login(id int, name string) {
	users.Store(id, name)
	fmt.Printf("[login] 用户 %d: %s 登录\n", id, name)
}

// 读取用户信息：Load
func getUser(id int) {
	if v, ok := users.Load(id); ok {
		fmt.Printf("[load] 读取用户 %d: %v\n", id, v)
	} else {
		fmt.Printf("[load] 用户 %d 不在线\n", id)
	}
}

// 只初始化一次的用户缓存：LoadOrStore
func getOrInitUser(id int, name string) {
	actual, loaded := users.LoadOrStore(id, name)
	if loaded {
		fmt.Printf("[loadOrStore] 用户 %d 已存在，值 = %v\n", id, actual)
	} else {
		fmt.Printf("[loadOrStore] 用户 %d 不存在，已存入新值 = %v\n", id, name)
	}
}

// 用户下线：Delete
func logout(id int) {
	users.Delete(id)
	fmt.Printf("[delete] 用户 %d 下线\n", id)
}

// 遍历当前在线用户：Range
func showOnlineUsers() {
	fmt.Print("[range] 当前在线用户: ")
	count := 0
	users.Range(func(k, v any) bool {
		fmt.Printf("(%d→%v) ", k, v)
		count++
		return true // 返回false可中断遍历
	})
	if count == 0 {
		fmt.Println("无")
	} else {
		fmt.Println()
	}
}

func main() {
	// 模拟多个操作
	login(1, "Alice")
	login(2, "Bob")

	getUser(1)
	getUser(3) // 不存在

	getOrInitUser(2, "Bobby")   // 已存在
	getOrInitUser(3, "Charlie") // 新插入

	showOnlineUsers()

	logout(2)
	showOnlineUsers()

	// 并发演示
	var wg sync.WaitGroup
	for i := 4; i <= 6; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			login(id, fmt.Sprintf("User%d", id))
			time.Sleep(100 * time.Millisecond)
			getUser(id)
		}(i)
	}
	wg.Wait()

	showOnlineUsers()
}
