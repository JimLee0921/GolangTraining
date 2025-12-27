package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql" // 常用 mysql 标准驱动
)

func main() {
	// 1. 创建 *sql.DB 句柄，sqlit 数据库不存在会在项目根目录创建
	db, err := sql.Open("mysql", "root:Dayi@516@tcp(192.168.7.236:53306)/dayiec")

	if err != nil {
		panic(err)
	}

	// 2. 显式检查是否可用
	if err = db.Ping(); err != nil {
		panic(err)
	}

	fmt.Println("database is ready")

	// 3. 正常使用
	var count string
	err = db.QueryRow("SELECT COUNT(*) FROM information_schema.tables WHERE table_schema = DATABASE();").Scan(&count)
	if err != nil {
		panic(err)
	}
	fmt.Println("count:", count)
}
