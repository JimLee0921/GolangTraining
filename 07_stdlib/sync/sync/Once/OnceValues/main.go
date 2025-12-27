package main

import (
	"database/sql"
	"fmt"
	"sync"
)

func main() {
	// 可返回多值，常用于 value+error
	openDB := sync.OnceValues(func() (*sql.DB, error) {
		fmt.Println("opening database...")
		return sql.Open("mysql", "user:pass@dbname")
	})

	db1, err1 := openDB()
	db2, err2 := openDB()

	fmt.Println(db1, err1)
	fmt.Println(db2, err2)
}
