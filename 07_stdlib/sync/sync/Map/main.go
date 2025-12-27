package main

import (
	"fmt"
	"sync"
)

type User struct {
	ID   int
	Name string
}

func main() {
	var users sync.Map

	// 写入 Store / LoadOrStore
	users.Store(1, User{
		ID:   1,
		Name: "JimLee",
	})

	users.Store(2, User{
		ID:   2,
		Name: "Bond",
	})

	// 并发安全的只初始化一次
	actual, loaded := users.LoadOrStore(3, User{
		ID:   3,
		Name: "Alex",
	})
	fmt.Println("loaded:", loaded, "value:", actual)

	// 并发读取
	var wg sync.WaitGroup

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			if v, ok := users.Load(id); ok {
				user := v.(User) // 需要类型断言
				fmt.Printf("goroutine read user %d: %+v\n", id, user)
			}
		}(i%3 + 1)
	}
	wg.Wait()

	// 原子替换 Swap
	old, existed := users.Swap(2, User{
		ID:   2,
		Name: "James Bond",
	})
	fmt.Println("swap existed:", existed, "old value:", old)

	// 条件更新
	ok := users.CompareAndSwap(
		1,
		User{
			ID:   1,
			Name: "JimLee",
		},
		User{
			ID:   1,
			Name: "Jim",
		},
	)
	fmt.Println("compare and swap success:", ok)

	// 遍历
	fmt.Println("range users:")
	users.Range(func(key, value any) bool {
		fmt.Printf("key=%v value=%v\n", key, value)
		return true
	})

	// 读取并删除
	if v, ok := users.LoadAndDelete(3); ok {
		fmt.Println("load and delete:", v)
	}
	
	// 清空

	users.Clear()
	fmt.Println("users cleared")
}
