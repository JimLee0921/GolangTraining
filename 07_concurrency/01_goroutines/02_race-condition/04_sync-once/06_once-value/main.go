package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	// 首次调用时“计算一次”，以后直接返回缓存值
	getTS := sync.OnceValue(func() string {
		fmt.Println("compute timestamp once")
		return time.Now().Format(time.RFC3339Nano)
	})

	fmt.Println("v1:", getTS())
	fmt.Println("v2:", getTS())
	fmt.Println("v3:", getTS())
	// 只有第一次会打印 "compute timestamp once"
}
