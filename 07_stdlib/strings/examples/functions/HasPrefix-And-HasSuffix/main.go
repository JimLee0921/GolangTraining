package main

import (
	"fmt"
	"strings"
)

// HasePrefix 判断前缀
func HasePrefix() {
	fmt.Println(strings.HasPrefix("Gopher", "Go")) // true
	fmt.Println(strings.HasPrefix("Gopher", "go")) // false
	fmt.Println(strings.HasPrefix("Gopher", ""))   // true
	fmt.Println(strings.HasPrefix("吉米李", "吉米"))    // true
}

// HasSuffix 判断后缀
func HasSuffix() {
	fmt.Println(strings.HasSuffix("Amigo", "go"))  // true
	fmt.Println(strings.HasSuffix("Amigo", "O"))   // false
	fmt.Println(strings.HasSuffix("Amigo", "Ami")) // false
	fmt.Println(strings.HasSuffix("Amigo", ""))    // true
}

func main() {
	HasePrefix()
	HasSuffix()
}
