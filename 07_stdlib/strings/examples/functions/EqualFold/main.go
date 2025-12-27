package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strings.EqualFold("Go", "go")) // true
	fmt.Println(strings.EqualFold("AB", "ab")) // true
	fmt.Println(strings.EqualFold("ß", "ss"))  // false

}
