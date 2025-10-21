package main

import (
	"fmt"
	"sync"
)

var loadConfig = sync.OnceValue(func() string {
	fmt.Println("loading config~")
	return "config.json"
})

func main() {
	/*
		loadConfig() 只会执行一次
		后续调用直接返回第一次的结果（线程安全）
	*/
	fmt.Println(loadConfig()) // 第一次调用执行函数
	fmt.Println(loadConfig()) // 第二次调用直接返回缓存值
}
