package concurrency

import "fmt"

func UnidirectionalChannel() {
	ch := make(chan int, 10) // here its bidirectional channel
	makeEvenNumber(4, ch)
	close(ch)
	for v := range ch {
		fmt.Println(v)
	}
}

func makeEvenNumber(count int, in chan<- int) { // <- here compiler converts to unidirectional
	for i := 0; i < count; i++ {
		in <- 2 * i
	}
}
