package concurrency

import "fmt"

func BufferedChannel() {
	ch := make(chan int, 3) // buffered channel has capacity of values to hold
	ch <- 12
	ch <- 13
	ch <- 14
	close(ch) // sender will get panic but reciever dont block whether it is unbuffered or buffered
	for v := range ch {
		fmt.Println(v)
	}
	fmt.Println(<-ch) // if channel closed the reciever value is zero
}
