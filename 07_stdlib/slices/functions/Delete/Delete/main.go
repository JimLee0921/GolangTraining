package main

import (
	"fmt"
	"slices"
)

func main() {
	letters := []string{"a", "b", "c", "d", "e", "f"}
	newLetters := slices.Delete(letters, 1, 4)
	fmt.Println(letters)
	fmt.Println(newLetters)
	letters = newLetters
	fmt.Println(letters)
}
