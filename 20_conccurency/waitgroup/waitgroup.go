package waitgroup

import (
	"fmt"
	"runtime"
	"sync"
)

var ws sync.WaitGroup

func WaitGroup() {
	fmt.Println(runtime.GOARCH)
	fmt.Println(runtime.GOOS)
	ws.Add(1)
	go foo()
	bar()
	fmt.Println(runtime.NumCPU())
	fmt.Println(runtime.NumGoroutine())
	ws.Wait()
}

func foo() {
	for i := 0; i < 10; i++ {
		fmt.Println(i, "Hello from foo")
	}
	ws.Done()
}

func bar() {
	for i := 0; i < 10; i++ {
		fmt.Println(i, "Hello from bar")
	}
}
