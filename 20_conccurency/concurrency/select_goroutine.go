package concurrency
// import "fmt"
func SelectGoRoutine()  {
  even := make(chan int)
  // odd := make(chan int)

  go send(even, odd)
  // recieve(even, odd)
}

func send(even, odd){
  for i := 0; i < 10; i++ {
    if (i%2) == 0 {
      even <- i
    } else {
      odd <- i
    }
  }
  close(even)
  close(odd)
}
