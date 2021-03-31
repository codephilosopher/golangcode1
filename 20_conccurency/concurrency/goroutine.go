package concurrency

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func GoRoutine() {
	starts := []int{10, 40, 70, 100}
	wg.Add(len(starts))
	for _, j := range starts {
		go func(s int) {
			fmt.Println("inside closure", s)
			count(s, s+20, 10)
		}(j)
	}
	wg.Wait()
}
func count(start, stop, delta int) {
	for i := start; i <= stop; i += delta {
		fmt.Println(i)
	}
	wg.Done()
}
