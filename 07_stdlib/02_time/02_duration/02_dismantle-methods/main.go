package main

import (
	"fmt"
	"time"
)

func main() {
	d := 90 * time.Minute
	fmt.Println("d:", d)
	fmt.Println("Nanoseconds:", d.Nanoseconds())
	fmt.Println("Microseconds:", d.Microseconds())
	fmt.Println("Milliseconds:", d.Milliseconds())
	fmt.Println("Seconds:", d.Seconds())
	fmt.Println("Minutes:", d.Minutes())
	fmt.Println("Hours:", d.Hours())
}
