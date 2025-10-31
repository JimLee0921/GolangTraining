package main

import (
	"fmt"
	"time"
)

const Layout = "2006-01-02 15:04:05"

func main() {
	timeStr := "2024-10-31 08:30:00"
	// Parse
	t1, err := time.Parse(Layout, timeStr)
	if err != nil {
		panic(err)
	}
	fmt.Println("Parse: ", t1)
	// 转为当前时区
	fmt.Println("Parse local: ", t1.Local())

	// ParseInLocation 可以指定时区
	loc, _ := time.LoadLocation("Asia/Shanghai")
	t2, err := time.ParseInLocation(Layout, timeStr, loc)
	if err != nil {
		panic(err)
	}

	fmt.Println("ParseInLocation: ", t2)
	fmt.Println("As UTC: ", t2.UTC())
}

/*
Parse:  2024-10-31 08:30:00 +0000 UTC
Parse local:  2024-10-31 16:30:00 +0800 +08
ParseInLocation:  2024-10-31 08:30:00 +0800 CST
As UTC:  2024-10-31 00:30:00 +0000 UTC
*/
