package main

import "fmt"

func scanInt(prompt string) (int, error) {
	var n int
	fmt.Println(prompt)
	_, err := fmt.Scan(&n)
	if err != nil {
		return 0, fmt.Errorf("please enter a number")
	}
	return n, nil
}

func main() {
	num1, err := scanInt("Please enter the first number: ")
	if err != nil {
		fmt.Println(err)
		return
	}
	num2, err := scanInt("Please enter the second number: ")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("sum is: ", num1+num2)
}
