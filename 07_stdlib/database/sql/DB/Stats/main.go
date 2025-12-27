package main

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	dsn := "root:Dayi@516@tcp(192.168.7.236:53306)/test"
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// 4 个连接池设置
	db.SetMaxOpenConns(10)                  // 最多同时打开 10 个连接（使用中+空闲）
	db.SetMaxIdleConns(5)                   // 最多保留 5 个空闲连接
	db.SetConnMaxLifetime(30 * time.Minute) // 单个连接最多存活 30 分钟
	db.SetConnMaxIdleTime(5 * time.Minute)  // 单个连接最多空闲 5 分钟

	// 让连接池至少创建/使用一次连接（否则 Stats 可能都是 0）
	if err := db.Ping(); err != nil {
		panic(err)
	}

	// 打印 Stats
	s := db.Stats()
	fmt.Printf("Open=%d InUse=%d Idle=%d MaxOpen=%d WaitCount=%d WaitTime=%s\n",
		s.OpenConnections, s.InUse, s.Idle, s.MaxOpenConnections, s.WaitCount, s.WaitDuration,
	)
}
