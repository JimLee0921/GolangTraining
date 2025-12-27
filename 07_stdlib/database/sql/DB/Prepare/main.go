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

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	stmt, err := db.PrepareContext(ctx, "INSERT INTO student (name, gender, phone) VALUES (?, ?, ?)")

	if err != nil {
		panic(err)
	}

	defer stmt.Close()

	res, err := stmt.ExecContext(ctx, "张三", "male", "13800000000")

	if err != nil {
		panic(err)
	}
	id, err := res.LastInsertId()
	if err != nil {
		panic(err)
	}

	fmt.Println("insert id:", id)

}
