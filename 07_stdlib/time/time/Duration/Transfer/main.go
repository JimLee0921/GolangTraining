package main

import (
	"fmt"
	"time"
)

func main() {
	d, err := time.ParseDuration("1h15m30.918273645s")
	if err != nil {
		panic(err)
	}
	fmt.Println(d.Hours())
	fmt.Println(d.Minutes())
	fmt.Println(d.Seconds())
	fmt.Println(d.Microseconds())
	fmt.Println(d.Milliseconds())
	fmt.Println(d.Nanoseconds())
}
