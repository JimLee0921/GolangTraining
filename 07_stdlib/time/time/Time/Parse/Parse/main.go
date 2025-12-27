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
}
