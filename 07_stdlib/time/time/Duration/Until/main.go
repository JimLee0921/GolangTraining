package main

import (
	"fmt"
	"time"
)

func main() {
	futureTime := time.Now().Add(5 * time.Second)

	durationUntil := time.Until(futureTime)

	fmt.Printf("Duration until future time: %v seconds", durationUntil)
}
