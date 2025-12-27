package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	year, month, day := now.Date()
	fmt.Println(year, month, day)

	year = now.Year()
	month = now.Month()
	day = now.Day()
	fmt.Println(year, month, day)
}
