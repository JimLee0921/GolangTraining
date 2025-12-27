package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	tenMinuteAfter := now.Add(10 * time.Minute)
	fmt.Println(now.Compare(tenMinuteAfter))
	fmt.Println(tenMinuteAfter.Compare(now))
}
