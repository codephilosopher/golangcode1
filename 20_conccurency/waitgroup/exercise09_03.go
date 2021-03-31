package waitgroup

import (
	"fmt"
	"runtime"
	"sync"
)

func Exercise09_03() {
	var wg sync.WaitGroup
	var syn sync.Mutex
	wg.Add(100)
	count := 0
	for i := 0; i < 100; i++ {
		go func() {
			runtime.Gosched()
			syn.Lock()
			count++
			syn.Unlock()
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println("total count: ", count)
}
