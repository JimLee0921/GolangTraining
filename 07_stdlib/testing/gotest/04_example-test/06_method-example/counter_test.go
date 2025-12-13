package counter

import "fmt"

// ExampleType_Method
func ExampleCounter_Inc() {
	c := &Counter{1}
	c.Inc()
	c.Inc()
	fmt.Println(c.n)
	// Output:
	// 3
}
