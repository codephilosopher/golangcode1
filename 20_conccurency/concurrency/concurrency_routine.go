package concurrency
import "fmt"
const MAX_LEN = 100000
const THRLEN = 2
func ConcurrencyProgram() {
	evenNumCh := make(chan int, MAX_LEN)
	results := make(chan int, THRLEN)
	go func() {
		defer close(evenNumCh)
		for i := 0; i < MAX_LEN; i++ {
			if (i%2) == 0 {
				evenNumCh <- i
			}
		}
		}()
		work := func(){
			sum := 0
			for val := range evenNumCh {
				sum += val
			}
			results <- sum
		}

		for i := 0; i < THRLEN; i++ {
			go work()
		}

		total := <-results + <-results

		fmt.Println("total sum of even numbers", total)
}
