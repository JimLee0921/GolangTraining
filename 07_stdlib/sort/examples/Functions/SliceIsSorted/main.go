package main

import (
	"fmt"
	"sort"
)

func main() {
	numbers := []int{1, 2, 3, 4, 5, 6}
	isSortedAsc := sort.SliceIsSorted(numbers, func(i, j int) bool {
		return numbers[i] < numbers[j]
	})
	fmt.Printf("%v sorted ascending: %t\n", numbers, isSortedAsc)

	numbersDesc := []int{6, 5, 4, 3, 2, 1}

	isSortedDesc := sort.SliceIsSorted(numbersDesc, func(i, j int) bool {
		return numbersDesc[i] > numbersDesc[j]
	})
	fmt.Printf("%v sorted descending: %t\n", numbersDesc, isSortedDesc)

	unsortedNumbers := []int{1, 3, 2, 4, 5}

	isSortedUnsorted := sort.SliceIsSorted(unsortedNumbers, func(i, j int) bool {
		return unsortedNumbers[i] < unsortedNumbers[j]
	})
	fmt.Printf("%v unsorted slice sorted: %t\n", unsortedNumbers, isSortedUnsorted)
}
