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

	fmt.Println("database connect successful")

	rows, err := db.Query("SELECT id, name, gender, phone AS telephone FROM student")
	if err != nil {
		return
	}

	defer rows.Close()

	cols, _ := rows.Columns()
	for _, col := range cols {
		fmt.Println(col)
	}
}
