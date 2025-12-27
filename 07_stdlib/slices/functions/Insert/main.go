package main

import (
	"fmt"
	"slices"
)

func main() {
	names := []string{"JimLee", "Bruce", "Chris"}
	names = slices.Insert(names, 1, "Bill", "Bond")
	names = slices.Insert(names, len(names), "Zac") // 相当于追加
	fmt.Println(names)
}
