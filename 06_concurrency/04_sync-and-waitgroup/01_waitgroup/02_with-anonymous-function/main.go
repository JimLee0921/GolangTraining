package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup

	for i := 1; i <= 3; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			fmt.Printf("worker %d is running~\n", id)
			time.Sleep(500 * time.Millisecond)
		}(i)
	}
	wg.Wait()
	fmt.Println("All done!")
}
