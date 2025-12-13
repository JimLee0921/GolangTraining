package demo

import "fmt"

func ExampleSet_Values() {
	s := Set{}
	s.Add("x")
	s.Add("z")
	s.Add("y")
	for _, v := range s.Values() {
		fmt.Println(v)
	}

	// Unordered output:
	// y
	// z
	// x
}
