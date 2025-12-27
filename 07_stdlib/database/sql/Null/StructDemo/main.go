package main

import (
	"database/sql"
	"fmt"
)

import _ "github.com/go-sql-driver/mysql"

// StudentRow 数据层结构 SQL 语义
type StudentRow struct {
	ID     int64
	Name   string
	Gender string
	Phone  sql.Null[string] // Null 只用于数据库允许 NULL 的列
}

// Student 业务层结构
type Student struct {
	ID     int64
	Name   string
	Gender string
	Phone  *string // 业务层不关心 sql.Null
}

// ToDomain 转换函数
func (r StudentRow) ToDomain() Student {
	var phone *string
	if r.Phone.Valid {
		phone = &r.Phone.V
	}

	return Student{
		ID:     r.ID,
		Name:   r.Name,
		Gender: r.Gender,
		Phone:  phone,
	}
}

func GetStudent(db *sql.DB, id int64) (*StudentRow, error) {
	var s StudentRow
	err := db.QueryRow(
		"SELECT id, name, gender, phone FROM student WHERE id = ?",
		id,
	).Scan(
		&s.ID,
		&s.Name,
		&s.Gender,
		&s.Phone, // Scan 会自动调用 Phone.Scan
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return &s, nil
}

func main() {
	db, _ := sql.Open("mysql", "root:Dayi@516@tcp(192.168.7.236:53306)/test")

	// GetStudent 测试
	student, err := GetStudent(db, 20)
	if err != nil {
		panic(err)
	}

	if student == nil {
		fmt.Println("no row")
		return
	}

	// 正确使用 sql.Null[T]
	if student.Phone.Valid {
		fmt.Println("phone", student.Phone.V)
	} else {
		fmt.Println("phone is NULL")
	}
}
