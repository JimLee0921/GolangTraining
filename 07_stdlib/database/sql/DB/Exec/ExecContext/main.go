package main

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "root:Dayi@516@tcp(192.168.7.236:53306)/test")

	if err != nil {
		panic(err)
	}

	if err = db.Ping(); err != nil {
		panic(err)
	}

	fmt.Println("database connect successful")

	// 1. 使用带超时的 Context
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)

	defer cancel()

	// 2. 使用 ExecContext
	insertSql := `
		INSERT INTO student (name, gender, phone)
		VALUES (?, ?, ?), (?, ?, ?);
	`

	result, err := db.ExecContext(ctx, insertSql, "张三", "male", "145535325252", "李四", "female", "15342535353")

	if err != nil {
		panic(err)
	}

	// 3. 获取执行查询结果
	id, err := result.LastInsertId() // 插入多条返回的是插入的第一条的id
	if err != nil {
		panic(err)
	}
	affected, err := result.RowsAffected()
	if err != nil {
		panic(err)
	}
	fmt.Printf("insert success: id=%d, rows=%d\n", id, affected)

}
