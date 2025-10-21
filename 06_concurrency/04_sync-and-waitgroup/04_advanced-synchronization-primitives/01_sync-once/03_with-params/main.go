package main

import (
	"fmt"
	"sync"
)

type Config struct {
	DSN string
}
type DB struct {
	dsn string
}

var (
	once   sync.Once
	cfg    Config // 由外部在初始化前设置
	dbInst *DB
)

// main 携带参数的初始化
func main() {
	/*
		sync.Once 的 Do 不能每次带不同参数；如果需要外部传入参数，先把参数存起来，再一次性初始化
		cfg Config → 用值（Config{...}），因为它小而简单，且生命周期跟随程序
		dbInst *DB → 用地址（&DB{...}），因为要全局唯一实例，所有地方都指向同一个对象
	*/
	// 在首次 InitDB 之前设置配置（之后修改不会再生效）
	cfg = Config{DSN: "mysql://user:pass@localhost/app"}
	var wg sync.WaitGroup
	wg.Add(2)
	for i := 0; i < 2; i++ {
		go func() {
			defer wg.Done()
			initDB()
		}()
	}
	wg.Wait()
}

func initDB() {
	once.Do(func() {
		// 使用已写入的 cfg 来初始化
		dbInst = &DB{dsn: cfg.DSN}
		fmt.Println("DB initialized with:", dbInst.dsn)
	})
}
