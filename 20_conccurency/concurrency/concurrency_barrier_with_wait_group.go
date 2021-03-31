package concurrency

import (
	"fmt"
	"runtime"
	"sync"
)

const MAX = 10
const TRLEN = 2

func ConcurrenyWaitGroup() {
	var wg sync.WaitGroup
	values := make(chan int, MAX)
	results := make(chan int, TRLEN)
	wg.Add(TRLEN)
	total := 0
	go func() {
		defer close(values)
		for i := 0; i < MAX; i++ {
			if (i%3) == 0 || (i%5) == 0 {
				values <- i
			}
		}
	}()
	work := func() {
		runtime.Gosched()
		defer wg.Done()
		ps := 0
		for v := range values {
			ps += v
		}
		results <- ps
	}
	for i := 0; i < TRLEN; i++ {
		go work()
	}
	wg.Wait()
	for i := 0; i < TRLEN; i++ {
		total += <-results
	}
	fmt.Println(total)
}
