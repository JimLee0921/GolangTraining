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

	// 使用 QueryRow / QueryRowContext 获取 sql.Row
	// 查询到多条，返回一条
	row := db.QueryRow("SELECT name, phone FROM student WHERE name = ?", "张三")
	var name string
	var phone sql.Null[string] // 可能为空
	// 使用 Row.Scan 才获取到结果和错误
	err = row.Scan(&name, &phone)
	if err != nil {
		panic(err)
	}
	if phone.Valid {
		fmt.Println(name, phone)
	} else {
		fmt.Println(name + "have no phone")
	}

	// 查询到 0 条，返回错误
	row = db.QueryRow("SELECT name, phone FROM student WHERE name = ?", "nobody")
	err = row.Scan(&name, &phone)
	if err != nil {
		panic(err) // no rows in result set
	}
}
