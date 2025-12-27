package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()

	hour, minute, second := now.Clock()
	fmt.Println(hour, minute, second)

	hour = now.Hour()
	minute = now.Minute()
	second = now.Second()
	nanoSecond := now.Nanosecond()

	fmt.Println(hour, minute, second, nanoSecond)
}
