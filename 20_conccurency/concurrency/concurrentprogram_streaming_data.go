package concurrency

import (
	"fmt"
	"strings"
)

func ConcurrentStreamData() {
	data := []string{
		"The yellow fish swims slowly in the water",
		"The brown dog barks loudly after a drink ...",
		"The dark bird bird of prey lands on a small ...",
	}
	histogram := make(map[string]int)
	streamChannel := make(chan string)
	go func() {
		defer close(streamChannel)
		for _, line := range data {
			words := strings.Split(line, " ")
			for _, word := range words {
				word := strings.ToLower(word)
				streamChannel <- word
			}
		}
	}()
	// always close the channel before acessing it, that the reciever could get proper signal else will raise the deadlock situation.
	for {
		data, opened := <-streamChannel // or we can use the for range loop with the channel
		if !opened {
			break
		}
		histogram[data]++
	}

	for k, v := range histogram {
		fmt.Printf("%s\t(%d)\n", k, v)
	}
}
