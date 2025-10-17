package main

import "fmt"

type StringSlice = []string
type StringMap = map[string]string

func main() {
	var s StringSlice = []string{"Go", "Rust", "Python"}
	var m StringMap = map[string]string{"A": "Alpha"}

	fmt.Printf("value: %v - type: %T\n", s, s)
	fmt.Printf("value: %v - type: %T\n", m, m)
}

/*
	value: [Go Rust Python] - type: []string
	value: map[A:Alpha] - type: map[string]string
*/
