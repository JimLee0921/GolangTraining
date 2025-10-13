package main

import "fmt"

type Box struct{ N int }

func (b *Box) Inc() {
	b.N++
}
func (b *Box) Sub(num int) {
	b.N -= num
}

func main() {
	b := Box{N: 1}
	b.Inc()          // 等价于 (&b).Inc()
	fmt.Println(b.N) // 2

	b.Sub(20)
	fmt.Println(b.N) // -18
}
