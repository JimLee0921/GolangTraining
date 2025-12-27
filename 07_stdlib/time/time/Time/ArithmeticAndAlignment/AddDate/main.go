package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Date(2023, 03, 25, 12, 0, 0, 0, time.UTC)

	oneDayLater := start.AddDate(0, 0, 1)
	oneDayLaterDuration := oneDayLater.Sub(start)
	oneMonthLater := start.AddDate(0, 1, 0)
	oneYearLater := start.AddDate(1, 0, 0)

	fmt.Println(start)
	fmt.Println(oneDayLater)
	fmt.Println(oneDayLaterDuration)
	fmt.Println(oneMonthLater)
	fmt.Println(oneYearLater)

	oneDayAfter := start.AddDate(0, 0, -1)
	fmt.Println(oneDayAfter)
}
