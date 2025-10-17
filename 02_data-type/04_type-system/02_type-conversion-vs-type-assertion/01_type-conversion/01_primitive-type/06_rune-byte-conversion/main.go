package main

import "fmt"

func main() {

	// []byte → []rune
	bs := []byte("Hello World!")
	rs := []rune(string(bs))
	fmt.Println("[]byte → []rune:", rs)
	fmt.Println(string(rs))

	// []rune → []byte
	rs2 := []rune{'H', 'e', 'l', 'l', 'o', ' ', 'W', 'o', 'r', 'l', 'd', '!'}
	bs2 := []byte(string(rs2))
	fmt.Println("[]rune → []byte:", bs2)
	fmt.Println(string(rs2))
}
