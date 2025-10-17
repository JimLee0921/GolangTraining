package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

// Snap 定义快照结构：Snap 表示某次配置快照，内含 map[string]string，但注意：发布后要当成只读不可改，避免数据竞争
type Snap struct {
	Routes map[string]string // 视为不可变（发布后不再修改）
}

// atomic.Value 可以安全地在多 goroutine 间读写任意类型的值，Store 时是原子操作，Load 时永远拿到某次完整 Store 的对象。
var snapshot atomic.Value // 存 *Snap

// Publish 原子更新当前配置快照
func Publish(s *Snap) { snapshot.Store(s) }

// Current 原子读取，拿到的是某次完整的 *Snap
func Current() *Snap { return snapshot.Load().(*Snap) }

func main() {
	// 初次发布
	Publish(&Snap{Routes: map[string]string{"home": "/", "about": "/about"}})

	var wg sync.WaitGroup
	wg.Add(2)

	// 因为每次都 Load()，所以能看到配置随时变化，但每次读到的都是一份完整的快照
	go func() {
		defer wg.Done()
		for i := 0; i < 5; i++ {
			s := Current()
			fmt.Println("reader sees:", s.Routes)
			time.Sleep(80 * time.Millisecond)
		}
	}()

	// 150ms 后更新配置，把 /about 改成 /help
	// 注意：是新建一个 *Snap，而不是在旧 map 上直接改 这样保证旧快照仍然对其他 goroutine 可见且一致
	go func() {
		defer wg.Done()
		time.Sleep(150 * time.Millisecond)
		Publish(&Snap{Routes: map[string]string{"home": "/", "help": "/help"}})
	}()
	wg.Wait()
}
