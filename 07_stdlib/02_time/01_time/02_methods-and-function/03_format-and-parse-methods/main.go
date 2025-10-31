package main

import (
	"fmt"
	"time"
)

func main() {
	// Format
	t := time.Now()
	fmt.Println(t.Format(time.RFC850)) // Friday, 31-Oct-25 11:01:54 +08

	// AppendFormat 手动指定 layout
	const layout = "2006-01-02" // 等同于 time.DateTime
	b := []byte{}
	b = t.AppendFormat(b, layout)
	fmt.Println(string(b)) // 2025-10-31 11:04:15

	// Parse
	timeStr := "2025-10-31 11:04:50"
	t2, err := time.Parse(time.DateTime, timeStr)
	if err != nil {
		panic(err)
	}
	fmt.Println(t2)
	// ParseInLocation
	dateStr := "2022-12-12"
	loc, _ := time.LoadLocation("Asia/Shanghai")
	t3, err := time.ParseInLocation(time.DateOnly, dateStr, loc)
	if err != nil {
		panic(err)
	}
	fmt.Println(t3)
}
