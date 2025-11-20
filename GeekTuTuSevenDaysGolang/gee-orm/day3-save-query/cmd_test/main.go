package main

import (
	geeorm "day1-database-sql"
	"fmt"

	_ "modernc.org/sqlite"
)

func main() {
	engine, _ := geeorm.NewEngine("sqlite", "gee.db")
	defer engine.Close()

	s := engine.NewSession()

	_, _ = s.Raw("DROP TABLE IF EXISTS User;").Exec()
	_, _ = s.Raw("CREATE TABLE User(Name text);").Exec()
	_, _ = s.Raw("CREATE TABLE User(Name text);").Exec() // 两次创建，会报错
	result, _ := s.Raw("INSERT INTO User(`Name`) values (?), (?)", "JimLee", "Bruce").Exec()

	count, _ := result.RowsAffected()
	fmt.Printf("Exec success, %d affected", count)
}
