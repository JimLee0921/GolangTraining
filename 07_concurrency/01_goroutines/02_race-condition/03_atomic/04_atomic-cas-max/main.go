package main

import (
	"fmt"
	"math/rand"
	"sync"
	"sync/atomic"
	"time"
)

type MaxGauge struct{ v atomic.Int64 }

func (m *MaxGauge) Observe(x int64) {
	for {
		old := m.v.Load()
		if x <= old {
			return
		}
		if m.v.CompareAndSwap(old, x) {
			return
		}
		// 失败就重试；激烈竞争场景可以加入短暂退避
	}
}

func (m *MaxGauge) Get() int64 { return m.v.Load() }

func main() {
	rand.Seed(time.Now().UnixNano())
	var g MaxGauge
	var wg sync.WaitGroup

	workers := 6
	wg.Add(workers)
	for i := 0; i < workers; i++ {
		go func() {
			defer wg.Done()
			for j := 0; j < 1000; j++ {
				g.Observe(rand.Int63n(1_000_000))
			}
		}()
	}
	wg.Wait()

	fmt.Println("max =", g.Get())
}
