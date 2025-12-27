package main

import (
	"fmt"
	"strings"
)

func main() {
	stringSlice := []string{"foo", "bar", "baz"}
	resultStr := strings.Join(stringSlice, ", ")
	fmt.Println(resultStr)
}
