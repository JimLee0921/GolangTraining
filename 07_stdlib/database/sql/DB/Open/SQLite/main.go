package main

import (
	"database/sql"
	"fmt"

	_ "modernc.org/sqlite" // 常用 sqlit 驱动
)

func main() {
	// 1. 创建 *sql.DB 句柄，sqlit 数据库不存在会在项目根目录创建
	db, err := sql.Open("sqlite", "./test.db")

	if err != nil {
		panic(err)
	}

	// 2. 显式检查是否可用
	if err = db.Ping(); err != nil {
		panic(err)
	}

	fmt.Println("database is ready")

	// 3. 正常使用
	var now string
	err = db.QueryRow("SELECT datetime('now')").Scan(&now)
	if err != nil {
		panic(err)
	}
	fmt.Println("now:", now)
}
