package concurrency

import "fmt"

func ChannelExamples() {
	// ch := make(chan int) // incapacitate leaves the goroutine to deadlock
	// ch <- 12
	// fmt.Println(<-ch)
	ch1 := make(chan int, 2) // channel capacity of one allows to hold one value at a time
	ch1 <- 12
	ch1 <- 14
	close(ch1)
	for v := range ch1 {
		fmt.Println(v)
	}
}
