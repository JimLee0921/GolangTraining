package main

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, _ := sql.Open("mysql", "root:Dayi@516@tcp(192.168.7.236:53306)/test")

	// 插入一个没有手机号的学生
	_, _ = db.Exec(
		"INSERT INTO student (name, gender, phone) VALUES (?, ?, ?)",
		"SB",
		"male",
		sql.Null[string]{Valid: false},
	)

	// 插入一个有手机号的学生
	_, _ = db.Exec(
		"INSERT INTO student (name, gender, phone) VALUES (?, ?, ?)",
		"JimWWW",
		"female",
		sql.Null[string]{
			Valid: true,
			V:     "1543534534534",
		},
	)

	// 把手机号更新为 NULL
	_, _ = db.Exec(
		"UPDATE student SET phone = ? WHERE id = ?",
		sql.Null[string]{Valid: false},
		3,
	)
}
