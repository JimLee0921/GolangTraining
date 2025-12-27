package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	dsn := "root:Dayi@516@tcp(192.168.7.236:53306)/test"

	db, err := sql.Open("mysql", dsn)

	if err != nil {
		panic(err)
	}
	if err = db.Ping(); err != nil {
		panic(err)
	}

	fmt.Println("database connect successful")

	// 开启事务
	tx, err := db.Begin()

	if err != nil {
		panic(err)
	}
	// 必须要兜底回滚
	defer tx.Rollback()

	// 事务中执行 sql
	_, err = tx.Exec(
		"INSERT INTO student (name, gender, phone) VALUES (?, ?, ?), (?, ?, ?)",
		"Jim",
		"male",
		"13600000000",
		"Bruce",
		"female",
		"134235235353",
	)

	if err != nil {
		panic(err)
	}

	// 正常提交事务
	if err = tx.Commit(); err != nil {
		panic(err)
	}
	fmt.Println("commit successful")
}
