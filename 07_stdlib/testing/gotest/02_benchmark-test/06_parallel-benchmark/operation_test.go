package opeartion

import (
	"math/rand"
	"testing"
)

/*
使用 RunParallel + SetParallelism 模拟多 goroutine 同时操作一个并发安全的数据结构
go test -bench .

可以观察 mutex 性能冲突情况等
*/
func BenchmarkNewSafeMapParallel(b *testing.B) {
	m := NewSafeMap()

	// 将并行度设为 4 * GOMAXPROCS
	b.SetParallelism(4)

	// 并行压测，内部使用多个 goroutine
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			k := rand.Intn(1_000)
			// 不存在则 Set
			if _, ok := m.Get(k); !ok {
				m.Set(k, k)
			}
		}
	})
}
