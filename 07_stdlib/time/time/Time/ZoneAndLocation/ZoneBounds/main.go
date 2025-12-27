package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Now()

	start, end := t.ZoneBounds()
	fmt.Println(start, end)
}
