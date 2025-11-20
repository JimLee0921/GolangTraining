package main

import (
	"database/sql"
	"log"

	_ "modernc.org/sqlite" // 纯 GO 实现的 sqlite 驱动，无需使用 GCC
)

func main() {
	// 运行后会自动生成 //gee.db 文件
	db, _ := sql.Open("sqlite", "./gee.db")

	defer func() { _ = db.Close() }()
	// Exec() 用于执行 SQL 语句，如果是查询语句，不会返回相关的记录
	_, _ = db.Exec("DROP TABLE IF EXISTS User;")
	_, _ = db.Exec("CREATE TABLE User(Name text);")
	result, err := db.Exec("INSERT INTO User(`Name`) values (?), (?)", "Tom", "Sam")
	if err == nil {
		affected, _ := result.RowsAffected()
		log.Println(affected)
	}
	// 查询语句通常使用 Query() 和 QueryRow()，前者可以返回多条记录，后者只返回一条记录
	row := db.QueryRow("SELECT NAME FROM User LIMIT 1;")
	// QueryRow() 的返回值类型是 *sql.Row，row.Scan() 接受1或多个指针作为参数，可以获取对应列(column)的值
	// 在这个示例中，只有 Name 一列，因此传入字符串指针 &name 即可获取到查询的结果
	var name string
	if err := row.Scan(&name); err == nil {
		log.Println(name)
	}
}
