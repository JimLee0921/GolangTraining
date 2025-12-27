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
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	querySQL := `
	SELECT name, phone FROM student where id in (?, ?);
	`

	rows, err := db.QueryContext(ctx, querySQL, 5, 100)

	if err != nil {
		panic(err)
	}

	defer rows.Close()

	for rows.Next() {
		var (
			name      string
			telephone string
		)

		if err := rows.Scan(&name, &telephone); err != nil {
			panic(err)
		}
		fmt.Printf("name: %s; telephone: %s\n", name, telephone)
	}
}
