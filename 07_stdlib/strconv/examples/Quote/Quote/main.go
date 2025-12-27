package main

import (
	"fmt"
	"strconv"
)

func main() {
	s := strconv.Quote(`"Fran & Freddie's Diner ☺"`)
	fmt.Println(s) // "\"Fran & Freddie's Diner ☺\""
}
