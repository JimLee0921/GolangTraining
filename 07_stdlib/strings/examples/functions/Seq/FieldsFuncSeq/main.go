package main

import (
	"fmt"
	"strings"
	"unicode"
)

// FieldsFuncSeq FieldsFunc的迭代器版本
func main() {
	text := "The quick brown fox"
	fmt.Println("Split on whitespace(similar to FieldsSeq):")
	for word := range strings.FieldsFuncSeq(text, unicode.IsSpace) {
		fmt.Printf("%q\n", word)
	}

	mixedText := "abc123def456ghi"
	fmt.Println("\nSplit on digits:")
	for word := range strings.FieldsFuncSeq(mixedText, unicode.IsDigit) {
		fmt.Printf("%q\n", word)
	}
}
