package main

import (
	"fmt"
	"time"
)

func main() {
	loc, _ := time.LoadLocation("Asia/Shanghai")
	t := time.Date(2025, 7, 1, 12, 0, 0, 0, loc)

	name, offset := t.Zone()
	fmt.Println(name, offset)
}
