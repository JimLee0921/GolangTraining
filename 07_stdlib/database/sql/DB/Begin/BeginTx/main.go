package main

import (
	"context"
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
	if err = db.Ping(); err != nil {
		panic(err)
	}

	fmt.Println("database connect successful")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	tx, err := db.BeginTx(ctx, &sql.TxOptions{
		Isolation: sql.LevelReadCommitted, // 事务隔离级别
		ReadOnly:  false,                  // 是否只读
	})
	if err != nil {
		panic(err)
	}

	// 兜底回滚（标准写法）
	defer tx.Rollback()

	_, err = tx.ExecContext(
		ctx,
		"INSERT INTO student (name, gender, phone) VALUES (?, ?, ?)",
		"周八",
		"male",
		"13500000000",
	)
	if err != nil {
		panic(err)
	}

	if err := tx.Commit(); err != nil {
		panic(err)
	}

}
