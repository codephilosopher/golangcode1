package concurrency

import "fmt"

// Declaring channel //var ch chan int // <--- unintialized channel
// the default value is nil.

func UnbufferedChannel() {
	ch := make(chan int) // unbuffered because capasity is unspecified
	// ch <- 12 this means value 12 is sending to channel 12
	// fmt.Println(<-ch)
	// <-ch values are recieving from channel
	// result will be fatal error: all goroutines are asleep - deadlock!
	// because the receiver will block the process further because the channel is empty.
	// anonymous gorouting helps the main function to reach the main recieve operation
	go func() { ch <- 12 }()
	fmt.Println(<-ch)
}
