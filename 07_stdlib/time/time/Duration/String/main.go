package main

import (
	"fmt"
	"time"
)

func main() {
	d := time.Second * 10
	d, _ = time.ParseDuration(d.String())
	fmt.Println(d, d.String())
}
