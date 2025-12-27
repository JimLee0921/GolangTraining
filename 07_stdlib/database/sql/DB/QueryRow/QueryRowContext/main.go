package main

import (
	"context"
	"database/sql"
	"errors"
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

	defer db.Close()

	// 带超时的 Context
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	var (
		id     int64
		name   string
		gender string
		phone  sql.NullString
	)

	err = db.QueryRowContext(ctx, `
	SELECT id, name, gender, phone FROM student where id = ?
	`, 123).Scan(&id, &name, &gender, &phone)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			fmt.Println("student not found")
			return
		}
		panic(err)
	}
	fmt.Printf(
		"id=%d, name=%s, gender=%s, phone=%s\n",
		id, name, gender, phone.String,
	)
}
