package main

import (
	"fmt"
	"os"
)

func main() {
	os.Setenv("USER", "Jim")
	os.Setenv("HOME", "/Users/Jim")

	text := "Welcome, $USER! Your home is ${HOME}."
	result := os.ExpandEnv(text)
	fmt.Println(result)
}
