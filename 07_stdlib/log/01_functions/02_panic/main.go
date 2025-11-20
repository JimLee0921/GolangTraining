package main

import "log"

// panic 系列输出后触发 panic，会执行 defer，可以配合 recover
func main() {
	//log.Panic("something bad happened")

	//log.Panicln("panic with newline")

	log.Panicf("invalid id: %d", 99)

}
