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

	result, err := db.Exec(`
		CREATE TABLE IF NOT EXISTS student (
			id BIGINT AUTO_INCREMENT PRIMARY KEY,
			name VARCHAR(50) NOT NULL,
			gender ENUM('male', 'female') NOT NULL,
			phone VARCHAR(20)
		);`)

	if err != nil {
		panic(err)
	}

	// Exec 对 DDL 返回的 RESULT 通常没有实际应用

	_, _ = result.RowsAffected()
	_, _ = result.LastInsertId()

	fmt.Println("student table created (or already exists)")

}
