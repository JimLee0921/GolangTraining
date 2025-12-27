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

	rows, err := db.Query("SELECT name, phone FROM student")
	if err != nil {
		return
	}

	defer rows.Close()

	for rows.Next() {
		var name string
		var phone sql.Null[string]

		if err := rows.Scan(&name, &phone); err != nil {
			panic(err)
		}

		if phone.Valid {
			fmt.Printf("%s phone: %s\n", name, phone.V)
		} else {
			fmt.Printf("%s have no phone\n", name)
		}
	}
	
	// 检查错误
	if err = rows.Err(); err != nil {
		panic(err)
	}
}
