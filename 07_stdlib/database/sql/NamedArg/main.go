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

	row := db.QueryRow(
		"SELECT name FROM student WHERE id = :id", // 需要驱动支持
		sql.Named("id", 1),
	)

	var name string
	err = row.Scan(&name)

	if err != nil {
		panic(err)
	} else {
		fmt.Printf(name)

	}

}
