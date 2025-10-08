package main

import "fmt"

// main 多个 incrementor
func main() {
	c1 := incrementor("channel one")
	c2 := incrementor("channel two")
	c3 := incrementor("channel three")

	p1 := puller(c1)
	p2 := puller(c2)
	p3 := puller(c3)

	v1 := <-p1
	v2 := <-p2
	v3 := <-p3
	fmt.Println("Final Counter:", v1+v2+v3)

}

func incrementor(s string) <-chan int {
	out := make(chan int)

	go func() {
		for i := 0; i < 10; i++ {
			fmt.Printf("%s：%d\n", s, i)
			out <- i
		}
		close(out)
	}()
	return out
}

func puller(c <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		sum := 0
		for i := range c {
			sum += i
		}
		out <- sum
		close(out)
	}()
	return out
}
