package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	fmt.Println(now.Format(time.UnixDate))
	fmt.Println(now.Format("2006-01-02T15:04:05 -070000"))
	fmt.Println(now.Format("2006-01-02T15:04:05 -07:00:00"))
}
