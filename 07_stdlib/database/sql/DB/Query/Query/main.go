package main

import (
	"database/sql"
	"fmt"

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

	querySQL := `
	SELECT * FROM student;
	`

	rows, err := db.Query(querySQL)

	if err != nil {
		panic(err)
	}

	defer rows.Close()

	for rows.Next() {
		var (
			id        int
			name      string
			gender    string
			telephone string
		)

		if err := rows.Scan(&id, &name, &gender, &telephone); err != nil {
			panic(err)
		}
		fmt.Printf("id: %d; name: %s; gender: %s; telephone: %s\n", id, name, gender, telephone)
	}
}
