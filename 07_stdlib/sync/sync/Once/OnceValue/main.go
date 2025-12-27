package main

import (
	"fmt"
	"sync"
)

func main() {
	// 只会执行一次，后续直接返回缓存值
	getConfig := sync.OnceValue(func() map[string]string {
		fmt.Println("loading config...")
		return map[string]string{
			"env":  "prod",
			"port": "8080",
		}
	})

	cfg1 := getConfig()
	cfg2 := getConfig()

	fmt.Println(cfg1)
	fmt.Println(cfg2)
}
