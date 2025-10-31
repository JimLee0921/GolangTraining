package main

import (
	"fmt"
	"os"
)

/*
文件内容
Go 2025
Rust 2023
*/
func main() {
	file, _ := os.Open("temp_files/log.txt")
	defer file.Close()

	var lang string
	var year int

	for {
		n, err := fmt.Fscanf(file, "%s %d", &lang, &year)
		if err != nil || n == 0 {
			fmt.Println(err)
			break
		}
		fmt.Println(lang, year)
	}
}
