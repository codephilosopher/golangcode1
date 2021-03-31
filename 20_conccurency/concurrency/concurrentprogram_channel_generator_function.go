package concurrency

import (
	"fmt"
	"strings"
)

func ChannelGeneratorFunction() {
	data := []string{
		"The yellow fish swims slowly in the water",
		"The brown dog barks loudly after a drink ...",
		"The dark bird bird of prey lands on a small ...",
	}
	histogram := make(map[string]int)
	words := words(data)
	for word := range words {
		histogram[word]++
	}
	for k, v := range histogram {
		fmt.Printf("%s\t(%d)\n", k, v)
	}
}

func words(data []string) <-chan string {
	ch := make(chan string)
	go func() {
		defer close(ch)
		for _, lines := range data {
			words := strings.Split(lines, " ")
			for _, word := range words {
				word := strings.ToLower(word)
				ch <- word
			}
		}
	}()
	return ch
}
