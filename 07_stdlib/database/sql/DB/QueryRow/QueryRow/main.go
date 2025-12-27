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

	var (
		id     int64
		name   string
		gender string
		phone  sql.NullString
	)

	err = db.QueryRow(`
	SELECT id, name, gender, phone FROM student where id = ?
	`, 4).Scan(&id, &name, &gender, &phone)

	if err != nil {
		if err == sql.ErrNoRows {
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
