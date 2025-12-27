package main

import (
	"fmt"
	"slices"
)

func main() {
	names := []string{"JimLee", "Bruce", "Chris", "Bond"}
	names = slices.Replace(names, 1, 3, "Bill", "Billie", "Cat", "SB")
	fmt.Println(names)
}
