package main

import (
	"fmt"
	"strings"
)

func main() {
	// ToTitle
	fmt.Println(strings.ToTitle("her royal highness"))
	fmt.Println(strings.ToTitle("loud noises"))
	fmt.Println(strings.ToTitle("брат"))

	// ToUpper
	fmt.Println(strings.ToUpper("Gopher"))

	// ToLower
	fmt.Println(strings.ToLower("Gopher"))
}
