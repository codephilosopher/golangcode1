package concurrency

import (
	"fmt"
	"strings"
)

func ConcurrentProgram2() {
	data := []string{
		"The yellow fish swims slowly in the water",
		"The brown dog barks loudly after a drink ...",
		"The dark bird bird of prey lands on a small ...",
	}
	done := make(chan struct{})
	histogram := make(map[string]int)
	fmt.Println("got block to primary go routine")
	go func() {
		defer close(done)
		for _, line := range data {
			words := strings.Split(line, " ")
			for _, word := range words {
				word := strings.ToLower(word)
				histogram[word]++
			}
		}
		fmt.Println("Existing from secondary goroutine")
	}()
	<-done
	for k, v := range histogram {
		fmt.Printf("%s\t(%d)\n", k, v)
	}
}
