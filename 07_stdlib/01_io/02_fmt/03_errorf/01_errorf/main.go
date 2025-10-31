package main

import (
	"errors"
	"fmt"
)

func main() {
	// Errorf 可以格式化，error.New 不行
	err := fmt.Errorf("user %s not fount", "Tom")
	fmt.Println(err)
	err = errors.New("new err")
	fmt.Println(err)
}
