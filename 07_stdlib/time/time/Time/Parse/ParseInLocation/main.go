package main

import (
	"fmt"
	"time"
)

const Layout = "2006-01-02 15:04:05"

func main() {
	timeStr := "2024-10-31 08:30:00"
	// ParseInLocation 可以指定时区
	loc, _ := time.LoadLocation("Asia/Shanghai")
	t, err := time.ParseInLocation(Layout, timeStr, loc)
	if err != nil {
		panic(err)
	}

	fmt.Println("ParseInLocation: ", t)
	fmt.Println("As UTC: ", t.UTC())
}
