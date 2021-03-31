package waitgroup

import (
	"fmt"
	"runtime"
	"sync"
)

func Exercise09() {
	fmt.Println("CpuNo: ", runtime.NumCPU())
	fmt.Println("num of Goroutines", runtime.NumGoroutine())
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		defer wg.Done()
		fmt.Println("Go routine one")
	}()

	go func() {
		defer wg.Done()
		fmt.Println("Go routine two")
	}()

	fmt.Println("CpuNo: ", runtime.NumCPU())
	fmt.Println("num of Goroutines", runtime.NumGoroutine())

	wg.Wait()

	fmt.Println("Existed to main goroutine")
	fmt.Println("CpuNo: ", runtime.NumCPU())
	fmt.Println("num of Goroutines", runtime.NumGoroutine())
}
