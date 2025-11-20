package main

import (
	"fmt"
	"strings"
)

func SplitDemo() {
	s := "Hello World Go!!"
	res := strings.Split(s, " ")
	fmt.Printf("%q\n", res) // ["Hello" "World" "Go!!"]
}

func SpiltNDemo() {
	s := "Hello World Go!!"
	res := strings.SplitN(s, " ", 2)
	fmt.Printf("%q\n", res) // ["Hello" "World Go!!"]
}

func FieldDemo() {
	s := "Hello World Go!!"
	res := strings.Fields(s)
	fmt.Printf("%q\n", res) // ["Hello" "World" "Go!!"]
}

func JoinDemo() {
	s := []string{"Hello", "World", "Go!!"}
	res := strings.Join(s, " ")
	fmt.Printf("%q\n", res) // "Hello World Go!!"
}

func SplitAfterDemo() {
	s := "Hello World Go!!"
	res := strings.SplitAfter(s, " ")
	fmt.Printf("%q\n", res) // ["Hello " "World " "Go!!"]
}
func main() {
	SplitDemo()
	SpiltNDemo()
	SplitAfterDemo()
	FieldDemo()
	JoinDemo()
}
