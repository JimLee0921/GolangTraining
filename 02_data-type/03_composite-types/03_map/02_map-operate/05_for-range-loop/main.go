package main

import "fmt"

func main() {
	myMap := map[string]int{"go": 1, "rust": 2, "python": 0}
	for key, value := range myMap {
		fmt.Println(key, value)
	}
}
