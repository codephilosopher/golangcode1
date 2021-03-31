package concurrency

import "fmt"

func CloseChannel() {
	ch := make(chan int, 4)
	ch <- 2
	ch <- 4
	close(ch)
	for i := 0; i < 4; i++ {
		if val, opened := <-ch; opened { // channel returns a value and a boolean
			fmt.Println(val)    // after channel close value is 0
			fmt.Println(opened) // after channel close the boolean return is false else true
		} else {
			fmt.Println("Channel closed!")
		}
	}
}
