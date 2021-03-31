package waitgroup

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

func Exercise09_05() {
	var wg sync.WaitGroup
	var count int64
	wg.Add(100)
	for i := 0; i < 100; i++ {
		go func() {
			runtime.Gosched()
			atomic.AddInt64(&count, 1)
			fmt.Println(atomic.LoadInt64(&count))
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println("total count: ", count)
}
