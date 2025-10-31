package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Now()

	// 1. 获取年月日
	fmt.Println(t.Year())
	fmt.Println(t.Month())
	fmt.Println(t.Day())
	fmt.Println(t.Date())
	// 2. 获取时分秒
	fmt.Println(t.Hour())
	fmt.Println(t.Minute())
	fmt.Println(t.Second())
	fmt.Println(t.Clock())
	// 3. 获取周
	fmt.Println(t.Weekday())
	fmt.Println(t.ISOWeek())
}
