package counter

import "fmt"

func ExampleCounter() {
	c := &Counter{1}
	c.Inc()
	fmt.Println(c.n)
	// Output:
	// 2
}
