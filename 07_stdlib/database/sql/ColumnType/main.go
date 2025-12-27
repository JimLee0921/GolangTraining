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

	rows, err := db.Query(
		"SELECT name, phone FROM student WHERE id > ?",
		3,
	)
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	types, err := rows.ColumnTypes()
	if err != nil {
		return
	}
	for _, ct := range types {
		fmt.Println(ct.Name())
		fmt.Println(ct.Nullable())
		fmt.Println(ct.DatabaseTypeName())
		fmt.Println(ct.Length())
		fmt.Println(ct.DecimalSize())
		fmt.Println(ct.ScanType())
	}

	//for rows.Next() {
	//	var name string
	//	var phone sql.Null[string]
	//	if err := rows.Scan(&name, &phone); err != nil {
	//		panic(err)
	//	}
	//	fmt.Println(name, phone)
	//}

}
