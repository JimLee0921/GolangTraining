package main

import (
	"fmt"
	"os"
)

func main() {
	// os.LookupEnv()：比 Getenv 更安全
	value, exists := os.LookupEnv("PATH")
	if exists {
		fmt.Println("PATH ENV exists:", value)
	} else {
		fmt.Println("PATH NOT FOUND")
	}

	value, exists = os.LookupEnv("DSB")
	if exists {
		fmt.Println("DSB ENV exists:", value)
	} else {
		fmt.Println("DSB NOT FOUND")
	}

}
