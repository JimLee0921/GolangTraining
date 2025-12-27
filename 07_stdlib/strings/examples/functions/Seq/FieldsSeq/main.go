package main

import (
	"fmt"
	"strings"
)

// FieldsSeq Fields 的迭代器版本
func main() {
	text := "The quick brown fox"
	fmt.Println("Split string into fields:")
	for word := range strings.FieldsSeq(text) {
		fmt.Printf("%q\n", word)
	}

	textWithSpaces := " \n    lots of      spaces   \t"
	fmt.Println("\nSplit string with multiple spaces:")
	for word := range strings.FieldsSeq(textWithSpaces) {
		fmt.Printf("%q\n", word)
	}
}
