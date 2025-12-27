package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()

	tenMinuteAfter := now.Add(time.Minute * 10)
	fiveMinuteAfter := now.Add(time.Minute * 5)

	if fiveMinuteAfter.Before(tenMinuteAfter) && fiveMinuteAfter.After(now) {
		fmt.Println("fiveMinuteAfter is early tenMinuteAfter and after now")
	}
}
