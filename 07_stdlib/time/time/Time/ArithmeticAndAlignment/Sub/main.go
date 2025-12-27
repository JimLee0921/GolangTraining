package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Date(2000, 1, 1, 0, 0, 0, 0, time.UTC)
	end := time.Date(2000, 1, 1, 12, 0, 0, 0, time.UTC)

	differencePos := end.Sub(start)
	differenceNeg := start.Sub(end)
	fmt.Printf("difference pos = %v\n", differencePos)
	fmt.Printf("difference neg = %v\n", differenceNeg)
}
