package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	// 模拟初始化：返回 (值, 错误)。OnceValues 会把两者都缓存住
	loadConfig := sync.OnceValues(func() (string, error) {
		fmt.Println("load config (once)")
		// 模拟从远端拉配置...
		time.Sleep(100 * time.Millisecond)
		// 返回值与错误一起被缓存；若返回 err，下次调用仍会得到同一个 err
		return "cfg:v1", nil
		// return "", errors.New("fetch failed") // 若失败，错误同样被缓存
	})

	cfg, err := loadConfig()
	fmt.Println("1st:", cfg, err)

	cfg, err = loadConfig()
	fmt.Println("2nd:", cfg, err)
}
