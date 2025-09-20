package main

import "fmt"

// main 打印一个小表格，用来练习宽度和对齐的格式化动词
func main() {
	fmt.Printf("%-10s %6s %8s\n", "Item", "Qty", "Price")
	fmt.Printf("%-10s %6d %8.2f\n", "Apples", 5, 1.23)
	fmt.Printf("%-10s %6d %8.2f\n", "Bananas", 12, 0.55)
	fmt.Printf("%-10s %6d %8.2f\n", "Cherries", 103, 3.15)

	fmt.Printf("\nZero padded: %08d\n", 42)
}
