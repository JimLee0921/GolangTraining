package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

const (
	numFactorials = 100
	randLimit     = 20
)

func main() {
	var wg sync.WaitGroup
	wg.Add(numFactorials) // 外层 WaitGroup：计数 = 100
	factorial(&wg)        // 这里必须传入指针
	wg.Wait()             // 等待 100 次 Done() 才会返回
}

func factorial(wgMain *sync.WaitGroup) {
	// 这里如果把 WaitGroup 按值传递，就会复制一个新的计数器
	var wg sync.WaitGroup                                // 内层 WaitGroup：作为统一起跑的屏障
	r := rand.New(rand.NewSource(time.Now().UnixNano())) // 使用时间戳作为种子生成真正的随机数

	wg.Add(numFactorials + 1) // 内层计数 = 101
	for j := 1; j <= numFactorials; j++ {
		go func() {
			defer wgMain.Done() // 内部 goroutine 结束时，外层计数 -1
			// 生成随机数并计算随机数阶乘
			f := r.Intn(randLimit)
			wg.Wait() // 在屏障前等待，直到内层计数清零
			total := 1
			for i := f; i > 0; i-- {
				total *= i
			}
			fmt.Printf("%d阶乘为: %d\n", f, total)

		}()
		wg.Done() // 循环里：每启动完 1 个 goroutine，就把内层计数 -1
	}
	fmt.Println("All done with initialization")
	wg.Done() // 循环结束后，再 -1（把最开始 +1 的主线程票也还回去）
}
