package main

import (
	"fmt"
	"os"
)

func main() {
	// os.Setenv
	err := os.Setenv("MODE", "DEV")
	if err != nil {
		return
	}
	fmt.Println("SET MODE ENV: ", os.Getenv("MODE")) // SET MODE ENV:  DEV 表示设置成功
}
