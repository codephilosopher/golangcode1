package concurrency

import (
	"fmt"
	"strings"
)

func ConcurrentProgram1() {
	done := make(chan bool)
	data := []string{
		"The yellow fish swims slowly in the water",
		"The brown dog barks loudly after a drink ...",
		"The dark bird bird of prey lands on a small ...",
	}
	histogram := make(map[string]int)
	go func() {
		for _, line := range data {
			words := strings.Split(line, " ")
			for _, word := range words {
				word := strings.ToLower(word)
				histogram[word]++
			}
		}
		done <- true
	}()
	if <-done {
		for k, v := range histogram {
			fmt.Printf("%s\t(%d)\n", k, v)
		}
	}
}
