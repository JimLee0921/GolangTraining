package main

import (
	"fmt"
	"sync"
)

var users sync.Map

func main() {
	// 1. 写入，键已存在则更新
	users.Store("JimLee", 20)
	users.Store("JamesBond", 22)
	users.Store("JimLee", 15)

	// 2. 读取
	age, _ := users.Load("JimLee")
	fmt.Println(age.(int))

	// 3. 遍历
	users.Range(func(key, value interface{}) bool {
		name := key.(string)
		age := value.(int)
		fmt.Println(name, age)
		return true
	})

	// 4. 删除
	users.Delete("JimLee")
	age, ok := users.Load("JimLee")
	fmt.Println(age, ok)

	// 5. 读取或写入
	users.LoadOrStore("JimLee", 100)
	age, _ = users.Load("JimLee")
	fmt.Println(age)
}
