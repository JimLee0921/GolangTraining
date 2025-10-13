package main

import (
	"fmt"
	"time"
)

// 装饰器 1：日志
func withLog(fn func()) func() {
	return func() {
		fmt.Println("[LOG] Start")
		fn()
		fmt.Println("[LOG] End")
	}
}

// 装饰器 2：计时
func withTimer(fn func()) func() {
	return func() {
		start := time.Now()
		fn()
		fmt.Println("[TIMER]", time.Since(start))
	}
}

// 装饰器 3：错误恢复
func withRecover(fn func()) func() {
	return func() {
		defer func() {
			if r := recover(); r != nil {
				fmt.Println("[RECOVER] Caught panic:", r)
			}
		}()
		fn()
	}
}

// Compose 从右到左执行：Compose(f3, f2, f1)(target)
func Compose(funcs ...func(func()) func()) func(func()) func() {
	return func(target func()) func() {
		result := target
		for i := len(funcs) - 1; i >= 0; i-- { // 反向包裹
			result = funcs[i](result)
		}
		return result
	}
}

func riskyTask() {
	fmt.Println("Running risky task...")
	panic("something went wrong!") // 故意触发 panic
}

func main() {
	decorated := Compose(withRecover, withTimer, withLog)(riskyTask)
	decorated()
}
