package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Date(1, time.January, 1, 0, 0, 0, 0, time.UTC)
	fmt.Println(t.IsZero())
}
