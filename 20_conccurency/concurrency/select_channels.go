package concurrency

import (
	"fmt"
	"strings"
)

func SelectChannels() {
	data := []string{
		"The yellow fish swims slowly in the water",
		"The brown dog barks loudly after a drink ...",
		"The dark bird bird of prey lands on a small ...",
	}
	stopCh := make(chan struct{})
	histogram := make(map[string]int)
	words := words1(stopCh, data)
	for word := range words {
		if histogram["the"] == 3 {
			close(stopCh)
		}
		histogram[word]++
	}
	for k, v := range histogram {
		fmt.Printf("%s\t(%d)\n", k, v)
	}
}

func words1(stopCh chan struct{}, data []string) <-chan string {
	out := make(chan string)
	go func() {
		defer close(out)
		for _, lines := range data {
			words := strings.Split(lines, " ")
			for _, word := range words {
				word := strings.ToLower(word)
				select {
				case out <- word:
				case <-stopCh:
					return
				}
			}
		}
	}()
	return out
}
