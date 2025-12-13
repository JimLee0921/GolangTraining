package parse

import "fmt"

func ExampleParse_error() {
	_, err := Parse("abc")
	fmt.Println(err != nil)
	// Output:
	// true
}
