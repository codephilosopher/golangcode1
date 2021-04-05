package cancel

import (
	"context"
	"fmt"
	"log"
	"time"
)

func ContextCancel() {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	go func() {
		fmt.Println("Hello from go routine")
		cancel()
	}()

	mySleepFunction(ctx, 5*time.Second, "hello")
}

func mySleepFunction(ctx context.Context, d time.Duration, s string) {
	select {
	case <-time.After(d):
		fmt.Println(s)
	case <-ctx.Done():
		log.Fatalf(ctx.Err().Error())
	}
}
